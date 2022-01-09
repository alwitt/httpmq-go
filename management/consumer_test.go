package management

import (
	"context"
	"fmt"
	"testing"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConsumerManagement(t *testing.T) {
	assert := assert.New(t)

	svrURL := common.GetUnitTestHttpmqMgmtAPIURL()
	coreClient, err := common.DefineAPIClient(svrURL, nil, nil, true)
	assert.Nil(err)

	uut := GetMgmtAPIWrapper(coreClient)

	utCtxt, utCtxtCancel := context.WithCancel(context.Background())
	defer utCtxtCancel()

	err = uut.Ready(utCtxt)
	assert.Nil(err)

	// Case 0: create consumer with unknown stream
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          uuid.New().String(),
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(uuid.New().String()),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, uuid.New().String(), param)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}

	// Case 1: create a stream
	stream1 := uuid.New().String()
	subjectBase := uuid.New().String()
	subjects1 := []string{fmt.Sprintf("%s.a", subjectBase), fmt.Sprintf("%s.b", subjectBase)}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJSStreamParam{Name: stream1, Subjects: &subjects1}
		rid, err := uut.CreateStream(reqCtxt, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, streamInfo, err := uut.GetStream(reqCtxt, stream1)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Equal(stream1, streamInfo.Config.Name)
		assert.EqualValues(subjects1, *streamInfo.Config.Subjects)
	}

	// Case 2: create consumer on stream
	consumer2 := uuid.New().String()
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[0]),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, stream1, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 3: create same consumer again on stream
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[0]),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, stream1, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 4: create same consumer with different subject
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[1]),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, stream1, param)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}

	// Case 5: create consumer against unknown subject
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          uuid.New().String(),
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(uuid.New().String()),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, stream1, param)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}

	// Case 6: create consumer against wildcard subject
	consumer6 := uuid.New().String()
	subject6 := fmt.Sprintf("%s.*", subjectBase)
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer6,
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(subject6),
		}
		rid, err := uut.CreateConsumerForStream(reqCtxt, stream1, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 7: read back consumer info
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, info, err := uut.GetConsumerOfStream(reqCtxt, stream1, consumer2)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Equal(consumer2, info.Name)
		assert.Equal(int64(1), info.Config.GetMaxAckPending())
		assert.Equal(subjects1[0], info.Config.GetFilterSubject())
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, allInfo, err := uut.ListAllConsumerOfStream(reqCtxt, stream1)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Len(allInfo, 2)
		info, ok := allInfo[consumer6]
		assert.True(ok)
		assert.Equal(consumer6, info.Name)
		assert.Equal(int64(2), info.Config.GetMaxAckPending())
		assert.Equal(subject6, info.Config.GetFilterSubject())
	}

	// Case 8: delete the consumer
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := uut.DeleteConsumerOnStream(reqCtxt, stream1, consumer2)
		assert.Nil(err)
		assert.Equal(callID, rid)
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err = uut.DeleteConsumerOnStream(reqCtxt, stream1, consumer6)
		assert.Nil(err)
		assert.Equal(callID, rid)
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, allInfo, err := uut.ListAllConsumerOfStream(reqCtxt, stream1)
		assert.Nil(err)
		assert.Equal(callID, rid)
		assert.Empty(allInfo)
	}

	// Case 9: delete the stream
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := uut.DeleteStream(reqCtxt, stream1)
		assert.Nil(err)
		assert.Equal(callID, rid)
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, _, err = uut.GetStream(reqCtxt, stream1)
		assert.NotNil(err)
		assert.Equal(callID, rid)
	}
}
