package dataplane

import (
	"context"

	"github.com/alwitt/httpmq-go/api"
)

// MsgACKParam parameters for sending a message ACK
type MsgACKParam struct {
	// Stream is the name of the stream
	Stream string
	// StreamSeq is the stream scope message sequence number
	StreamSeq int64
	// Consumer is the name of the consumer
	Consumer string
	// ConsumerSeq is the consumer scope message sequence number
	ConsumerSeq int64
}

// PushSubscribeParam parameters for starting a push subscription
type PushSubscribeParam struct {
	// Stream is the name of the stream
	Stream string
	// Consumer is the name of the consumer
	Consumer string
	// SubjectFilter is the subject filter to subscribe to message for
	SubjectFilter string
	// MaxMsgInflight is the max number of inflight messages if provided
	MaxMsgInflight *int
	// DeliveryGroup is the delivery group the consumer belongs if the consumer uses one
	DeliveryGroup *string
	// MsgChan channel for passing back messages
	MsgChan *chan api.ApisAPIRestRespDataMessage
}

// DataAPIWrapper is a client wrapper object for operating the httpmq dataplane API
type DataAPIWrapper interface {
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
	client *api.APIClient
}

/*
GetDataAPIWrapper gets an instance of DataAPIWrapper

 @param core *api.APIClient - the base APIClient object to use
 @return the DataAPIWrapper object
*/
func GetDataAPIWrapper(core *api.APIClient) DataAPIWrapper {
	return &dataAPIWrapperImpl{client: core}
}
