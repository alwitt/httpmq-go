package dataplane

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/alwitt/httpmq-go/management"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPushSubscribe(t *testing.T) {
	assert := assert.New(t)

	mgmtSvrURL := common.GetUnitTestHttpmqMgmtAPIURL()
	mgmtCoreClient, err := common.DefineAPIClient(mgmtSvrURL, nil, nil, true)
	assert.Nil(err)
	dataSvrURL := common.GetUnitTestHttpmqDataAPIURL()
	dataCoreClient, err := common.DefineAPIClient(dataSvrURL, nil, nil, true)
	assert.Nil(err)

	ctrl := management.GetMgmtAPIWrapper(mgmtCoreClient)
	uut := GetDataAPIWrapper(dataCoreClient)

	utCtxt, utCtxtCancel := context.WithCancel(context.Background())
	defer utCtxtCancel()

	assert.Nil(ctrl.Ready(utCtxt))
	assert.Nil(uut.Ready(utCtxt))

	// Case 0: create a stream and consumer
	stream0 := uuid.New().String()
	consumer0 := uuid.New().String()
	consumerWildCard := uuid.New().String()
	subjectBase := uuid.New().String()
	subjects0 := []string{fmt.Sprintf("%s.a", subjectBase), fmt.Sprintf("%s.b", subjectBase)}
	subjectWildcard := fmt.Sprintf("%s.*", subjectBase)
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJSStreamParam{Name: stream0, Subjects: &subjects0}
		rid, err := ctrl.CreateStream(reqCtxt, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumer0,
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(subjects0[0]),
		}
		rid, err := ctrl.CreateConsumerForStream(reqCtxt, stream0, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		param := api.ManagementJetStreamConsumerParam{
			Name:          consumerWildCard,
			Mode:          "push",
			MaxInflight:   2,
			FilterSubject: api.PtrString(subjectWildcard),
		}
		rid, err := ctrl.CreateConsumerForStream(reqCtxt, stream0, param)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Case 1: subscribe for messages
	msgChan1 := make(chan api.ApisAPIRestRespDataMessage, 2)
	rdContext1, tdCtxtCancel1 := context.WithCancel(utCtxt)
	go func() {
		callID := uuid.New().String()
		reqCtxt := context.WithValue(rdContext1, common.UseGivenRequestID{}, callID)
		rid, err := uut.PushSubscribe(
			reqCtxt, PushSubscribeParam{
				Stream:         stream0,
				Consumer:       consumer0,
				SubjectFilter:  subjects0[0],
				MaxMsgInflight: nil,
				DeliveryGroup:  nil,
				MsgChan:        &msgChan1,
			},
		)
		assert.True(strings.Contains(err.Error(), "context canceled"))
		assert.Equal(callID, rid)
	}()
	// Verify messages are published
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		msg := uuid.New().String()
		rid, err := uut.Publish(reqCtxt, subjects0[0], []byte(msg))
		assert.Nil(err)
		assert.Equal(callID, rid)
		rdTimeout, lclCancel := context.WithCancel(utCtxt)
		defer lclCancel()
		select {
		case rxMsg, ok := <-msgChan1:
			assert.True(ok)
			assert.EqualValues(msg, string(rxMsg.B64Msg))
			// ACK message
			callID := uuid.New().String()
			reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
			rid, err := uut.SendACK(
				reqCtxt, MsgACKParam{
					Stream:      stream0,
					StreamSeq:   rxMsg.Sequence.Stream,
					Consumer:    consumer0,
					ConsumerSeq: rxMsg.Sequence.Consumer,
				},
			)
			assert.Nil(err)
			assert.Equal(callID, rid)
		case <-rdTimeout.Done():
			assert.False(true)
		}
	}
	// Also clear the message for the wild card consumer
	go func() {
		callID := uuid.New().String()
		reqCtxt := context.WithValue(rdContext1, common.UseGivenRequestID{}, callID)
		rid, err := uut.PushSubscribe(
			reqCtxt, PushSubscribeParam{
				Stream:         stream0,
				Consumer:       consumerWildCard,
				SubjectFilter:  subjectWildcard,
				MaxMsgInflight: nil,
				DeliveryGroup:  nil,
				MsgChan:        &msgChan1,
			},
		)
		assert.True(strings.Contains(err.Error(), "context canceled"))
		assert.Equal(callID, rid)
	}()
	{
		rdTimeout, lclCancel := context.WithCancel(utCtxt)
		defer lclCancel()
		select {
		case rxMsg, ok := <-msgChan1:
			assert.True(ok)
			// ACK message
			callID := uuid.New().String()
			reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
			rid, err := uut.SendACK(
				reqCtxt, MsgACKParam{
					Stream:      stream0,
					StreamSeq:   rxMsg.Sequence.Stream,
					Consumer:    consumer0,
					ConsumerSeq: rxMsg.Sequence.Consumer,
				},
			)
			assert.Nil(err)
			assert.Equal(callID, rid)
		case <-rdTimeout.Done():
			assert.False(true)
		}
	}
	tdCtxtCancel1()

	// Case 2: subscribe for wildcard subjects
	msgChan2 := make(chan api.ApisAPIRestRespDataMessage, 2)
	rdContext2, tdCtxtCancel2 := context.WithCancel(utCtxt)
	go func() {
		callID := uuid.New().String()
		reqCtxt := context.WithValue(rdContext2, common.UseGivenRequestID{}, callID)
		rid, err := uut.PushSubscribe(
			reqCtxt, PushSubscribeParam{
				Stream:         stream0,
				Consumer:       consumerWildCard,
				SubjectFilter:  subjectWildcard,
				MaxMsgInflight: nil,
				DeliveryGroup:  nil,
				MsgChan:        &msgChan2,
			},
		)
		assert.True(strings.Contains(err.Error(), "context canceled"))
		assert.Equal(callID, rid)
	}()
	// Verify messages are published
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		msg1 := fmt.Sprintf("test-message-1.%s", uuid.New().String())
		rid, err := uut.Publish(reqCtxt, subjects0[0], []byte(msg1))
		assert.Nil(err)
		assert.Equal(callID, rid)
		rdTimeout, lclCancel := context.WithCancel(utCtxt)
		defer lclCancel()
		select {
		case rxMsg, ok := <-msgChan2:
			assert.True(ok)
			assert.EqualValues(msg1, string(rxMsg.B64Msg))
			// ACK message
			callID := uuid.New().String()
			reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
			rid, err := uut.SendACK(
				reqCtxt, MsgACKParam{
					Stream:      stream0,
					StreamSeq:   rxMsg.Sequence.Stream,
					Consumer:    consumerWildCard,
					ConsumerSeq: rxMsg.Sequence.Consumer,
				},
			)
			assert.Nil(err)
			assert.Equal(callID, rid)
		case <-rdTimeout.Done():
			assert.False(true)
		}
		msg2 := fmt.Sprintf("test-message-2.%s", uuid.New().String())
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err = uut.Publish(reqCtxt, subjects0[1], []byte(msg2))
		assert.Nil(err)
		assert.Equal(callID, rid)
		select {
		case rxMsg, ok := <-msgChan2:
			assert.True(ok)
			assert.EqualValues(msg2, string(rxMsg.B64Msg))
			// ACK message
			callID := uuid.New().String()
			reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
			rid, err := uut.SendACK(
				reqCtxt, MsgACKParam{
					Stream:      stream0,
					StreamSeq:   rxMsg.Sequence.Stream,
					Consumer:    consumerWildCard,
					ConsumerSeq: rxMsg.Sequence.Consumer,
				},
			)
			assert.Nil(err)
			assert.Equal(callID, rid)
		case <-rdTimeout.Done():
			assert.False(true)
		}
	}
	tdCtxtCancel2()

	// Delete the consumers
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := ctrl.DeleteConsumerOnStream(reqCtxt, stream0, consumer0)
		assert.Nil(err)
		assert.Equal(callID, rid)
		callID = uuid.New().String()
		reqCtxt = context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err = ctrl.DeleteConsumerOnStream(reqCtxt, stream0, consumerWildCard)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}

	// Delete the stream
	{
		callID := uuid.New().String()
		reqCtxt := context.WithValue(utCtxt, common.UseGivenRequestID{}, callID)
		rid, err := ctrl.DeleteStream(reqCtxt, stream0)
		assert.Nil(err)
		assert.Equal(callID, rid)
	}
}
