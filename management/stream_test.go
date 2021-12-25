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
	coreClient, err := common.DefineAPIClient(svrURL, true)
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
		param := api.ManagementJSStreamParam{Name: stream0, Subjects: &subjects0}
		_, err := uut.CreateStream(utCtxt, param)
		assert.Nil(err)
	}

	// Case 1: create the same stream again
	{
		param := api.ManagementJSStreamParam{Name: stream0, Subjects: &subjects0}
		_, err := uut.CreateStream(utCtxt, param)
		assert.Nil(err)
	}

	// Case 2: create the another stream but with same subject
	{
		param := api.ManagementJSStreamParam{Name: uuid.New().String(), Subjects: &subjects0}
		_, err := uut.CreateStream(utCtxt, param)
		assert.NotNil(err)
	}

	// Case 3: read back the stream info
	{
		_, streamInfo, err := uut.GetStream(utCtxt, stream0)
		assert.Nil(err)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects0, *streamInfo.Config.Subjects)
	}
	{
		_, allStreamInfo, err := uut.ListAllStreams(utCtxt)
		assert.Nil(err)
		streamInfo, ok := allStreamInfo[stream0]
		assert.True(ok)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects0, *streamInfo.Config.Subjects)
	}

	// Case 4: alter subjects for stream
	subjects4 := []string{uuid.New().String(), uuid.New().String()}
	{
		_, err := uut.ChangeStreamSubjects(utCtxt, stream0, subjects4)
		assert.Nil(err)
	}
	{
		_, streamInfo, err := uut.GetStream(utCtxt, stream0)
		assert.Nil(err)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.EqualValues(subjects4, *streamInfo.Config.Subjects)
	}

	// Case 5: alter stream data retention
	{
		msgAge := time.Hour
		newLimit := api.ManagementJSStreamLimits{
			MaxAge: api.PtrInt64(msgAge.Nanoseconds()),
		}
		_, err := uut.UpdateStreamLimits(utCtxt, stream0, newLimit)
		assert.Nil(err)
		_, streamInfo, err := uut.GetStream(utCtxt, stream0)
		assert.Nil(err)
		assert.Equal(stream0, streamInfo.Config.Name)
		assert.Equal(msgAge.Nanoseconds(), streamInfo.Config.MaxAge)
	}

	// Case 6: delete the stream
	{
		_, err := uut.DeleteStream(utCtxt, stream0)
		assert.Nil(err)
	}
	{
		_, _, err := uut.GetStream(utCtxt, stream0)
		assert.NotNil(err)
	}
}
