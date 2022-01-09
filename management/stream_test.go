package management

import (
	"context"
	"testing"
	"time"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStreamManagement(t *testing.T) {
	assert := assert.New(t)

	svrURL := common.GetUnitTestHttpmqMgmtAPIURL()
	coreClient, err := common.DefineAPIClient(svrURL, nil, nil, true)
	assert.Nil(err)

	uut := GetMgmtAPIWrapper(coreClient)

	utCtxt, utCtxtCancel := context.WithCancel(context.Background())
	defer utCtxtCancel()

	err = uut.Ready(utCtxt)
	assert.Nil(err)

	// Case 0: create a stream
	stream0 := uuid.New().String()
	subjects0 := []string{uuid.New().String()}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJSStreamParam{Name: stream0, Subjects: &subjects0}
		rid, err := uut.CreateStream(reqCtxt, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 1: create the same stream again
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJSStreamParam{Name: stream0, Subjects: &subjects0}
		rid, err := uut.CreateStream(reqCtxt, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 2: create the another stream but with same subject
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJSStreamParam{Name: uuid.New().String(), Subjects: &subjects0}
		rid, err := uut.CreateStream(reqCtxt, param)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}

	// Case 3: read back the stream info
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, streamInfo, err := uut.GetStream(reqCtxt, stream0)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects0, *streamInfo.Config.Subjects)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, allStreamInfo, err := uut.ListAllStreams(reqCtxt)
		assert.Nil(err)
		assert.Equal(callID, rid)
		streamInfo, ok := allStreamInfo[stream0]
		assert.True(ok)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects0, *streamInfo.Config.Subjects)
	}

	// Case 4: alter subjects for stream
	subjects4 := []string{uuid.New().String(), uuid.New().String()}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := uut.ChangeStreamSubjects(reqCtxt, stream0, subjects4)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, streamInfo, err := uut.GetStream(reqCtxt, stream0)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects4, *streamInfo.Config.Subjects)
	}

	// Case 5: alter stream data retention
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		msgAge := time.Hour
		newLimit := api.ManagementJSStreamLimits{
			MaxAge: api.PtrInt64(msgAge.Nanoseconds()),
		}
		rid, err := uut.UpdateStreamLimits(reqCtxt, stream0, newLimit)
		assert.Nil(err)
		assert.Equal(callID, rid)
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, streamInfo, err := uut.GetStream(reqCtxt, stream0)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.Equal(msgAge.Nanoseconds(), streamInfo.Config.MaxAge)
	}

	// Case 6: delete the stream
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := uut.DeleteStream(reqCtxt, stream0)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, _, err := uut.GetStream(reqCtxt, stream0)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}
}
