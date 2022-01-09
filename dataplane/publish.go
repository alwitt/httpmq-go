package dataplane

import (
	"context"
	"encoding/base64"

	"github.com/alwitt/httpmq-go/common"
)

/*
Publish publishes a message under a subject

 @param ctxt context.Context - the caller context
 @param subject string - the subject to publish under
 @param message []byte - the message body
 @return request ID (to reference logs)
         an error message if request failed
*/
func (c *dataAPIWrapperImpl) Publish(
	ctxt context.Context, subject string, message []byte,
) (string, error) {
	baseRequest := c.client.DataplaneApi.V1DataSubjectSubjectNamePost(ctxt, subject)
	encoded := base64.StdEncoding.EncodeToString(message)
	request := baseRequest.Message(encoded)

	response, httpResp, err := c.client.DataplaneApi.V1DataSubjectSubjectNamePostExecute(request)
	if err != nil {
		return "", err
	}

	requestID := httpResp.Header.Get(common.RequestIDHeader)

	if !response.Success {
		errorDetail, _ := response.GetErrorOk()
		return requestID, common.GenerateHttpmqError(requestID, httpResp.StatusCode, errorDetail)
	}

	return requestID, nil
}
