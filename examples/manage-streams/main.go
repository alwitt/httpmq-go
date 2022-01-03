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

	// Read back that stream
	reqID, streamParam, err := ctrl.GetStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to read back stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	t, _ := json.MarshalIndent(streamParam, "", "  ")
	fmt.Printf("%s\nRequest ID %s\n", t, reqID)

	// Change the stream's target subject
	reqID, err = ctrl.ChangeStreamSubjects(
		context.Background(), streamName, []string{"subj.2", "subj.3"},
	)
	if err != nil {
		fmt.Printf("Failed to change stream %s subjects: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Read back that stream
	reqID, streamParam, err = ctrl.GetStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to read back stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	t, _ = json.MarshalIndent(streamParam, "", "  ")
	fmt.Printf("%s\nRequest ID %s\n", t, reqID)

	// Change the stream's retention policy
	reqID, err = ctrl.UpdateStreamLimits(
		context.Background(), streamName, api.ManagementJSStreamLimits{
			MaxAge: api.PtrInt64((time.Minute * 15).Nanoseconds()),
		},
	)
	if err != nil {
		fmt.Printf("Failed to update stream %s retention policy: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	// Read back that stream
	reqID, streamParam, err = ctrl.GetStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to read back stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}

	t, _ = json.MarshalIndent(streamParam, "", "  ")
	fmt.Printf("%s\nRequest ID %s\n", t, reqID)

	// Delete the stream
	reqID, err = ctrl.DeleteStream(context.Background(), streamName)
	if err != nil {
		fmt.Printf("Failed to delete stream %s: %s\n", streamName, err)
		panic(fmt.Errorf("request ID %s", reqID))
	}
}
