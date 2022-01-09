package management

import (
	"context"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
)

// MgmtAPIWrapper is a client wrapper object for operating the httpmq management API
type MgmtAPIWrapper interface {
	/*
		Ready check whether the httpmq management API is ready

		 @param ctxt context.Context - the caller context
		 @return whether the management API is ready, or an error message is given
	*/
	Ready(ctxt context.Context) error

	// --------------------------------------------------------------------------------
	// Stream related API functions

	/*
		CreateStream define a new stream

		 @param ctxt context.Context - the caller context
		 @param params api.ManagementJSStreamParam - stream parameters
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	CreateStream(ctxt context.Context, params api.ManagementJSStreamParam) (string, error)

	/*
		ListAllStreams query for list of all known streams

		 @param ctxt context.Context - the caller context
		 @return request ID (to reference logs)
		         list of known streams,
		         or an error message if request failed
	*/
	ListAllStreams(ctxt context.Context) (
		string, map[string]api.ApisAPIRestRespStreamInfo, error,
	)

	/*
		GetStream query for a particular stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @return request ID (to reference logs)
		         information on that stream
		         or an error message if request failed
	*/
	GetStream(ctxt context.Context, stream string) (
		string, *api.ApisAPIRestRespStreamInfo, error,
	)

	/*
		ChangeStreamSubjects change the target subjects of a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @param newSubjects []string - list of new subjects
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	ChangeStreamSubjects(ctxt context.Context, stream string, newSubjects []string) (
		string, error,
	)

	/*
		UpdateStreamLimits update the data retention limits of a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @param limits api.ManagementJSStreamLimits - new data retention limits
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	UpdateStreamLimits(
		ctxt context.Context, stream string, limits api.ManagementJSStreamLimits,
	) (string, error)

	/*
		DeleteStream delete a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	DeleteStream(ctxt context.Context, stream string) (string, error)

	// --------------------------------------------------------------------------------
	// Consumer related API functions

	/*
		CreateConsumerForStream define a new consumer on a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - the stream to create the consumer on
		 @param params api.ManagementJetStreamConsumerParam - consumer parameters
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	CreateConsumerForStream(
		ctxt context.Context, stream string, params api.ManagementJetStreamConsumerParam,
	) (string, error)

	/*
		ListAllConsumerOfStream query for list of all known consumers on a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - the stream to query for
		 @return request ID (to reference logs)
		         list of known consumer of a stream,
		         or an error message if request failed
	*/
	ListAllConsumerOfStream(ctxt context.Context, stream string) (
		string, map[string]api.ApisAPIRestRespConsumerInfo, error,
	)

	/*
		GetConsumerOfStream query for a particular consumer on a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @param consumer string - name of the consumer
		 @return request ID (to reference logs)
		         information on that consumer
		         or an error message if request failed
	*/
	GetConsumerOfStream(ctxt context.Context, stream, consumer string) (
		string, *api.ApisAPIRestRespConsumerInfo, error,
	)

	/*
		DeleteConsumerOnStream delete a consumer on a stream

		 @param ctxt context.Context - the caller context
		 @param stream string - name of the stream
		 @param consumer string - name of the consumer
		 @return request ID (to reference logs)
		         an error message if request failed
	*/
	DeleteConsumerOnStream(ctxt context.Context, stream, consumer string) (string, error)
}

// mgmtAPIWrapperImpl implements MgmtAPIWrapper
type mgmtAPIWrapperImpl struct {
	client *api.APIClient
}

/*
Ready check whether the httpmq management API is ready

 @param ctxt context.Context - the caller context
 @return whether the management API is ready, or an error message is given
*/
func (c *mgmtAPIWrapperImpl) Ready(ctxt context.Context) error {
	request := c.client.ManagementApi.V1AdminReadyGet(ctxt)

	response, httpResp, err := c.client.ManagementApi.V1AdminReadyGetExecute(request)
	requestID := httpResp.Header.Get(common.RequestIDHeader)
	if err != nil || !response.Success {
		errorDetail, _ := response.GetErrorOk()
		return common.GenerateHttpmqError(requestID, httpResp.StatusCode, errorDetail)
	}

	return nil
}

/*
GetMgmtAPIWrapper gets an instance of MgmtAPIWrapper

 @param core *api.APIClient - the base APIClient object to use
 @return the MgmtAPIWrapper object
*/
func GetMgmtAPIWrapper(core *api.APIClient) MgmtAPIWrapper {
	return &mgmtAPIWrapperImpl{client: core}
}
