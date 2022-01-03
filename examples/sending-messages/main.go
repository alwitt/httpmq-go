package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/alwitt/httpmq-go/dataplane"
	"github.com/alwitt/httpmq-go/management"
	"github.com/google/uuid"
	"golang.org/x/net/http2"
)

func main() {
	// Define httpmq core API client
	coreMgntClient, err := common.DefineAPIClient(
		"http://127.0.0.1:4000", &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}, nil, true,
	)
	if err != nil {
		panic(fmt.Errorf("unable to define httpmq-go client: %s", err.Error()))
	}
	coreDataClient, err := common.DefineAPIClient(
		"http://127.0.0.1:4001", &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}, nil, true,
	)
	if err != nil {
		panic(fmt.Errorf("unable to define httpmq-go client: %s", err.Error()))
	}

	// Define httpmq management API client
	ctrl := management.GetMgmtAPIWrapper(coreMgntClient)
	if err := ctrl.Ready(context.Background()); err != nil {
		panic(fmt.Errorf("httpmq management API not ready: %s", err))
	}
	data := dataplane.GetDataAPIWrapper(coreDataClient)
	if err := ctrl.Ready(context.Background()); err != nil {
		panic(fmt.Errorf("httpmq dataplane API not ready: %s", err))
	}

	// Define a stream
	streamName := uuid.New().String()
	reqID, err := ctrl.CreateStream(context.Background(), api.ManagementJSStreamParam{
		Name:     streamName,
		Subjects: &[]string{"subj.1", "subj.2"},
		MaxAge:   api.PtrInt64((time.Minute * 10).Nanoseconds()),
	})
	if err != nil {
		fmt.Printf("Failed to define stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Define a consumer
	consumerName := uuid.New().String()
	reqID, err = ctrl.CreateConsumerForStream(
		context.Background(), streamName, api.ManagementJetStreamConsumerParam{
			Name:          consumerName,
			Mode:          "push",
			MaxInflight:   1,
			FilterSubject: api.PtrString("subj.*"),
		},
	)
	if err != nil {
		fmt.Printf("Failed to define consumer %s: %s\n", consumerName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Subscribe for messages
	rxContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}
	defer wg.Wait()
	msgChan := make(chan api.ApisAPIRestRespDataMessage, 1)
	subParams := dataplane.PushSubscribeParam{
		Stream:        streamName,
		Consumer:      consumerName,
		SubjectFilter: "subj.*",
		MsgChan:       &msgChan,
	}

	// Start receiving messages
	wg.Add(1)
	go func() {
		defer wg.Done()
		reqID, err := data.PushSubscribe(rxContext, subParams)
		fmt.Printf("Push subscription complete. Request ID %s\n", reqID)
		if err != nil {
			fmt.Printf("Subscription errors: %s\n", err.Error())
		}
		cancel()
	}()

	// Publish a message
	message := "Hello world"
	reqID, err = data.Publish(context.Background(), "subj.2", []byte(message))
	if err != nil {
		fmt.Printf("Failed to publish a message to 'subj.2': %s\n", err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Wait for message to be read
	rxMsg, ok := <-msgChan
	if !ok {
		panic("Unable to read message over subscription")
	}
	fmt.Printf("Read: '%s'\n", rxMsg.B64Msg)

	// ACK the message
	ackMsg := dataplane.MsgACKParam{
		Stream:      rxMsg.Stream,
		StreamSeq:   rxMsg.Sequence.Stream,
		Consumer:    rxMsg.Consumer,
		ConsumerSeq: rxMsg.Sequence.Consumer,
	}
	reqID, err = data.SendACK(rxContext, ackMsg)
	if err != nil {
		fmt.Printf("Failed to send ACK for message: %s\n", err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Stop receiving messages
	cancel()

	// Delete the stream
	reqID, err = ctrl.DeleteStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to delete stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}
}
