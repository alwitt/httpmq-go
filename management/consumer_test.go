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
	coreClient, err := common.DefineAPIClient(svrURL, true)
	assert.Nil(err)

	uut := GetMgmtAPIWrapper(coreClient)

	utCtxt, utCtxtCancel := context.WithCancel(context.Background())
	defer utCtxtCancel()

	err = uut.Ready(utCtxt)
	assert.Nil(err)

	// Case 0: create consumer with unknown stream
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          uuid.New().String(),
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(uuid.New().String()),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, uuid.New().String(), param)
		assert.NotNil(err)
	}

	// Case 1: create a stream
	stream1 := uuid.New().String()
	subjectBase := uuid.New().String()
	subjects1 := []string{fmt.Sprintf("%s.a", subjectBase), fmt.Sprintf("%s.b", subjectBase)}
	{
		param := api.ManagementJSStreamParam{Name: stream1, Subjects: &subjects1}
		_, err := uut.CreateStream(utCtxt, param)
		assert.Nil(err)
	}
	{
		_, streamInfo, err := uut.GetStream(utCtxt, stream1)
		assert.Nil(err)
		assert.Equal(stream1, streamInfo.Config.Name)
		assert.EqualValues(subjects1, *streamInfo.Config.Subjects)
	}

	// Case 2: create consumer on stream
	consumer2 := uuid.New().String()
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[0]),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, stream1, param)
		assert.Nil(err)
	}

	// Case 3: create same consumer again on stream
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[0]),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, stream1, param)
		assert.Nil(err)
	}

	// Case 4: create same consumer with different subject
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer2,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(subjects1[1]),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, stream1, param)
		assert.NotNil(err)
	}

	// Case 5: create consumer against unknown subject
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          uuid.New().String(),
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString(uuid.New().String()),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, stream1, param)
		assert.NotNil(err)
	}

	// Case 6: create consumer against wildcard subject
	consumer6 := uuid.New().String()
	subject6 := fmt.Sprintf("%s.*", subjectBase)
	{
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer6,
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(subject6),
		}
		_, err := uut.CreateConsumerForStream(utCtxt, stream1, param)
		assert.Nil(err)
	}

	// Case 7: read back consumer info
	{
		_, info, err := uut.GetConsumerOfStream(utCtxt, stream1, consumer2)
		assert.Nil(err)
		assert.Equal(consumer2, info.Name)
		assert.Equal(int64(1), info.Config.GetMaxAckPending())
		assert.Equal(subjects1[0], info.Config.GetFilterSubject())
	}
	{
		_, allInfo, err := uut.ListAllConsumerOfStream(utCtxt, stream1)
		assert.Nil(err)
		assert.Len(allInfo, 2)
		info, ok := allInfo[consumer6]
		assert.True(ok)
		assert.Equal(consumer6, info.Name)
		assert.Equal(int64(2), info.Config.GetMaxAckPending())
		assert.Equal(subject6, info.Config.GetFilterSubject())
	}

	// Case #: delete the consumer
	{
		_, err := uut.DeleteConsumerOnStream(utCtxt, stream1, consumer2)
		assert.Nil(err)
		_, err = uut.DeleteConsumerOnStream(utCtxt, stream1, consumer6)
		assert.Nil(err)
		_, allInfo, err := uut.ListAllConsumerOfStream(utCtxt, stream1)
		assert.Nil(err)
		assert.Empty(allInfo)
	}

	// Case #: delete the stream
	{
		_, err := uut.DeleteStream(utCtxt, stream1)
		assert.Nil(err)
		_, _, err = uut.GetStream(utCtxt, stream1)
		assert.NotNil(err)
	}
}
