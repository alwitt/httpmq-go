package management

import (
	"context"
	"fmt"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
)

/*
CreateStream define a new stream

 @param ctxt context.Context - the caller context
 @param params api.ManagementJSStreamParam - stream parameters
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) CreateStream(
	ctxt context.Context, params api.ManagementJSStreamParam,
) (string, error) {
	initRequest := c.client.ManagementApi.V1AdminStreamPost(ctxt)
	request := initRequest.Setting(params)

	response, httpResp, err := c.client.ManagementApi.V1AdminStreamPostExecute(request)
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
			"failed to define new stream %s: %s", params.Name, errorMsg,
		)
	}

	return requestID, nil
}

/*
ListAllStreams query for list of all known streams

 @param ctxt context.Context - the caller context
 @return request ID (to reference logs)
         list of known streams,
         or an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) ListAllStreams(ctxt context.Context) (
	string, map[string]api.ApisAPIRestRespStreamInfo, error,
) {
	request := c.client.ManagementApi.V1AdminStreamGet(ctxt)

	response, httpResp, err := c.client.ManagementApi.V1AdminStreamGetExecute(request)
	if err != nil {
		return "", map[string]api.ApisAPIRestRespStreamInfo{}, err
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
		return requestID, map[string]api.ApisAPIRestRespStreamInfo{}, fmt.Errorf(
			"failed to query for all known streams: %s", errorMsg,
		)
	}

	return requestID, response.GetStreams(), nil
}

/*
GetStream query for a particular stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @return request ID (to reference logs)
         information on that stream
         or an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) GetStream(ctxt context.Context, stream string) (
	string, *api.ApisAPIRestRespStreamInfo, error,
) {
	request := c.client.ManagementApi.V1AdminStreamStreamNameGet(ctxt, stream)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameGetExecute(request)
	if err != nil {
		return "", nil, err
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
		return requestID, nil, fmt.Errorf("failed to query for stream %s: %s", stream, errorMsg)
	}

	return requestID, response.Stream, nil
}

/*
ChangeStreamSubjects change the target subjects of a stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @param newSubjects []string - list of new subjects
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) ChangeStreamSubjects(
	ctxt context.Context, stream string, newSubjects []string,
) (string, error) {
	baseRequest := c.client.ManagementApi.V1AdminStreamStreamNameSubjectPut(ctxt, stream)
	request := baseRequest.Subjects(*api.NewApisAPIRestReqStreamSubjects(newSubjects))

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameSubjectPutExecute(request)
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
			"failed to change for stream %s subjects: %s", stream, errorMsg,
		)
	}

	return requestID, nil
}

/*
UpdateStreamLimits update the data retention limits of a stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @param limits api.ManagementJSStreamLimits - new data retention limits
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) UpdateStreamLimits(
	ctxt context.Context, stream string, limits api.ManagementJSStreamLimits,
) (string, error) {
	baseRequest := c.client.ManagementApi.V1AdminStreamStreamNameLimitPut(ctxt, stream)
	request := baseRequest.Limits(limits)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameLimitPutExecute(request)
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
			"failed to change for stream %s retention limits: %s", stream, errorMsg,
		)
	}

	return requestID, nil
}

/*
DeleteStream delete a stream

 @param ctxt context.Context - the caller context
 @param stream string - name of the stream
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *mgmtAPIWrapperImpl) DeleteStream(
	ctxt context.Context, stream string,
) (string, error) {
	request := c.client.ManagementApi.V1AdminStreamStreamNameDelete(ctxt, stream)

	response, httpResp, err :=
		c.client.ManagementApi.V1AdminStreamStreamNameDeleteExecute(request)
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
		return requestID, fmt.Errorf("failed to delete stream %s: %s", stream, errorMsg)
	}

	return requestID, nil
}
