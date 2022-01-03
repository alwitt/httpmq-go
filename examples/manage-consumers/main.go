package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
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

	// Define httpmq management API client
	ctrl := management.GetMgmtAPIWrapper(coreMgntClient)
	if err := ctrl.Ready(context.Background()); err != nil {
		panic(fmt.Errorf("httpmq management API not ready: %s", err))
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
			Name:        consumerName,
			Mode:        "push",
			MaxInflight: 1,
		},
	)
	if err != nil {
		fmt.Printf("Failed to define consumer %s: %s\n", consumerName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Read the consumer back
	reqID, consumerParam, err := ctrl.GetConsumerOfStream(
		context.Background(), streamName, consumerName,
	)
	if err != nil {
		fmt.Printf("Failed to read back consumer %s: %s\n", consumerName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	t, _ := json.MarshalIndent(consumerParam, "", "  ")
	fmt.Printf("%s\nRequest ID %s\n", t, reqID)

	// Delete the stream
	reqID, err = ctrl.DeleteStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to delete stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}
}
