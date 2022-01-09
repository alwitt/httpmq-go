package dataplane

import (
	"context"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/go-playground/validator/v10"
)

// MsgACKParam parameters for sending a message ACK
type MsgACKParam struct {
	// Stream is the name of the stream
	Stream string `validate:"required"`
	// StreamSeq is the stream scope message sequence number
	StreamSeq int64 `validate:"required,gte=1"`
	// Consumer is the name of the consumer
	Consumer string `validate:"required"`
	// ConsumerSeq is the consumer scope message sequence number
	ConsumerSeq int64 `validate:"required,gte=1"`
}

// PushSubscribeParam parameters for starting a push subscription
type PushSubscribeParam struct {
	// Stream is the name of the stream
	Stream string `validate:"required"`
	// Consumer is the name of the consumer
	Consumer string `validate:"required"`
	// SubjectFilter is the subject filter to subscribe to message for
	SubjectFilter string `validate:"required"`
	// MaxMsgInflight is the max number of inflight messages if provided
	MaxMsgInflight *int `validate:"omitempty,gte=1"`
	// DeliveryGroup is the delivery group the consumer belongs if the consumer uses one
	DeliveryGroup *string
	// MsgChan channel for passing back messages
	MsgChan *chan api.ApisAPIRestRespDataMessage `validate:"-"`
}

// DataAPIWrapper is a client wrapper object for operating the httpmq dataplane API
type DataAPIWrapper interface {
	/*
		Ready check whether the httpmq dataplane API is ready

		 @param ctxt context.Context - the caller context
		 @return whether the dataplane API is ready, or an error message is given
	*/
	Ready(ctxt context.Context) error

	/*
		Publish publishes a message under a subject

		 @param ctxt context.Context - the caller context
		 @param subject string - the subject to publish under
		 @param message []byte - the message body
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	Publish(ctxt context.Context, subject string, message []byte) (string, error)

	/*
		SendACK send a message ACK for an associated JetStream message

		 @param ctxt context.Context - the caller context
		 @param params MsgACKParam - parameters for the message ACK
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	SendACK(ctxt context.Context, params MsgACKParam) (string, error)

	/*
		PushSubscribe start a push subscription for a consumer on a stream.

		This is a blocking function which only exits when either
		 * The caller context is cancelled
		 * Connection breaks
		 * Server closes the connection

		 @param ctxt context.Context - the caller context
		 @param params PushSubscribeParam - parameters for the push subscription
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	PushSubscribe(ctxt context.Context, params PushSubscribeParam) (string, error)
}

// dataAPIWrapperImpl implements DataAPIWrapper
type dataAPIWrapperImpl struct {
	client   *api.APIClient
	validate *validator.Validate
}

/*
GetDataAPIWrapper gets an instance of DataAPIWrapper

 @param core *api.APIClient - the base APIClient object to use
 @return the DataAPIWrapper object
*/
func GetDataAPIWrapper(core *api.APIClient) DataAPIWrapper {
	return &dataAPIWrapperImpl{client: core, validate: validator.New()}
}

/*
Ready check whether the httpmq dataplane API is ready

 @param ctxt context.Context - the caller context
 @return whether the dataplane API is ready, or an error message is given
*/
func (c *dataAPIWrapperImpl) Ready(ctxt context.Context) error {
	request := c.client.DataplaneApi.V1DataReadyGet(ctxt)

	response, httpResp, err := c.client.DataplaneApi.V1DataReadyGetExecute(request)
	requestID := httpResp.Header.Get(common.RequestIDHeader)
	if err != nil || !response.Success {
		errorDetail, _ := response.GetErrorOk()
		return common.GenerateHttpmqError(requestID, httpResp.StatusCode, errorDetail)
	}

	return nil
}
