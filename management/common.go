package management

import (
	"context"
	"fmt"

	"github.com/alwitt/httpmq-go/api"
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
	request := c.client.ManagementApi.ReadyGet(ctxt)

	response, _, err := c.client.ManagementApi.ReadyGetExecute(request)
	if err != nil {
		return err
	}

	errorDetail, ok := response.GetErrorOk()
	errorMsg := ""
	if ok {
		msg, ok := errorDetail.GetMessageOk()
		if ok {
			errorMsg = *msg
		}
	}

	if !response.Success {
		return fmt.Errorf("management API not ready: %s", errorMsg)
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
