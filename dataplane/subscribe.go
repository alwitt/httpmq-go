package dataplane

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
)

/*
SendACK send a message ACK for an associated JetStream message

 @param ctxt context.Context - the caller context
 @param params MsgACKParam - parameters for the message ACK
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *dataAPIWrapperImpl) SendACK(ctxt context.Context, params MsgACKParam) (string, error) {
	if err := c.validate.Struct(&params); err != nil {
		return "", err
	}

	baseRequest := c.client.DataplaneApi.V1DataStreamStreamNameConsumerConsumerNameAckPost(
		ctxt, params.Stream, params.Consumer,
	)
	request := baseRequest.SequenceNum(
		api.DataplaneAckSeqNum{Stream: params.StreamSeq, Consumer: params.ConsumerSeq},
	)

	response, httpResp, err :=
		c.client.DataplaneApi.V1DataStreamStreamNameConsumerConsumerNameAckPostExecute(request)
	if err != nil {
		return "", err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		return requestID, fmt.Errorf(
			"failed to send ACK [S:%d, C:%d] for consumer %s on stream %s: %s",
			params.StreamSeq,
			params.ConsumerSeq,
			params.Consumer,
			params.Stream,
			errorMsg,
		)
	}

	return requestID, nil
}

/*
PushSubscribe start a push subscription for a consumer on a stream

This is a blocking function which only exits when either
 * The caller context is cancelled
 * Connection breaks
 * Server closes the connection

 @param ctxt context.Context - the caller context
 @param params PushSubscribeParam - parameters for the push subscription
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *dataAPIWrapperImpl) PushSubscribe(
	ctxt context.Context, params PushSubscribeParam,
) (string, error) {
	// Because the push subscription is a expected to be a server-send-event stream, will
	// have to directly operate the HTTP client

	if err := c.validate.Struct(&params); err != nil {
		return "", err
	}

	localBasePath, err := c.client.GetConfig().ServerURLWithContext(
		ctxt, "DataplaneApiService.V1DataStreamStreamNameConsumerConsumerNameGet",
	)
	if err != nil {
		return "", err
	}
	clientCfg := c.client.GetConfig()
	httpClient := clientCfg.HTTPClient
	// Assume there is only one target server
	subscribeURL := fmt.Sprintf(
		"%s/v1/data/stream/%s/consumer/%s",
		localBasePath,
		url.PathEscape(params.Stream),
		url.PathEscape(params.Consumer),
	)

	// Build the request
	queryParams := url.Values{}
	queryParams["subject_name"] = []string{params.SubjectFilter}
	if params.MaxMsgInflight != nil {
		queryParams["max_msg_inflight"] = []string{fmt.Sprintf("%d", *params.MaxMsgInflight)}
	}
	if params.DeliveryGroup != nil {
		queryParams["delivery_group"] = []string{*params.DeliveryGroup}
	}
	request, err := c.client.PrepareRequest(
		ctxt, subscribeURL, http.MethodGet, nil, nil, queryParams, nil, nil,
	)
	if err != nil {
		return "", err
	}

	// Make the request
	response, err := httpClient.Do(request.WithContext(ctxt))
	if err != nil {
		return "", err
	}

	// ----------------------------------------------------------------------------------------
	// Start reading the messages

	// response stream scanner
	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines)

	var readErr error
	readErr = nil
	for scanner.Scan() {
		oneMessage := scanner.Text()
		var parsed api.ApisAPIRestRespDataMessage
		if readErr = json.Unmarshal([]byte(oneMessage), &parsed); readErr != nil {
			break
		}
		// Error occurred on the other side
		if !parsed.Success {
			errorDetail, ok := parsed.GetErrorOk()
			errorMsg := ""
			if ok {
				msg, ok := errorDetail.GetMessageOk()
				if ok {
					errorMsg = *msg
				}
			}
			readErr = fmt.Errorf(
				"push-subscription stream for %s@%s/%s stopped: %s",
				params.Consumer,
				params.Stream,
				params.SubjectFilter,
				errorMsg,
			)
			break
		}
		// Return the message
		*params.MsgChan <- parsed
	}

	// End the request
	_ = response.Body.Close()

	// Get scanner level error
	if readErr == nil {
		readErr = scanner.Err()
	}

	// Finally get the request ID from the completed response
	requestID := response.Header.Get(common.RequestIDHeader)

	return requestID, readErr
}
