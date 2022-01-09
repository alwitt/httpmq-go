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
	if useID := common.FetchUserProvidedRequestID(ctxt); useID != nil {
		request = request.HttpmqRequestID(*useID)
	}

	response, httpResp, err :=
		c.client.DataplaneApi.V1DataStreamStreamNameConsumerConsumerNameAckPostExecute(request)
	requestID := httpResp.Header.Get(common.RequestIDHeader)
	if err != nil || !response.Success {
		errorDetail, _ := response.GetErrorOk()
		return requestID, common.GenerateHttpmqError(requestID, httpResp.StatusCode, errorDetail)
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

	requestID := ""
	additionalHeader := map[string]string{}
	if useID := common.FetchUserProvidedRequestID(ctxt); useID != nil {
		requestID = *useID
		additionalHeader[common.RequestIDHeader] = *useID
	}

	if err := c.validate.Struct(&params); err != nil {
		return requestID, err
	}

	localBasePath, err := c.client.GetConfig().ServerURLWithContext(
		ctxt, "DataplaneApiService.V1DataStreamStreamNameConsumerConsumerNameGet",
	)
	if err != nil {
		return requestID, err
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
		ctxt, subscribeURL, http.MethodGet, nil, additionalHeader, queryParams, nil, nil,
	)
	if err != nil {
		return requestID, err
	}

	// Make the request
	response, err := httpClient.Do(request.WithContext(ctxt))
	if err != nil {
		return requestID, err
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
		if requestID == "" {
			requestID = parsed.RequestId
		}
		// Error occurred on the other side
		if !parsed.Success {
			errorDetail, _ := parsed.GetErrorOk()
			readErr = common.GenerateHttpmqError(requestID, response.StatusCode, errorDetail)
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

	return requestID, readErr
}
