package main

import (
	"context"
	"fmt"

	"github.com/alwitt/httpmq-go/common"
	"github.com/alwitt/httpmq-go/dataplane"
	"github.com/alwitt/httpmq-go/management"
)

func main() {
	// Define httpmq core API client
	coreMgntClient, err := common.DefineAPIClient("http://127.0.0.1:4000", nil, nil, true)
	if err != nil {
		panic(fmt.Errorf("unable to define httpmq-go client: %s", err.Error()))
	}
	coreDataClient, err := common.DefineAPIClient("http://127.0.0.1:4001", nil, nil, true)
	if err != nil {
		panic(fmt.Errorf("unable to define httpmq-go client: %s", err.Error()))
	}

	// Define httpmq management API client
	ctrl := management.GetMgmtAPIWrapper(coreMgntClient)
	if err := ctrl.Ready(context.Background()); err != nil {
		panic(fmt.Errorf("httpmq management API not ready: %s", err))
	}

	// Define httpmq dataplane API client
	data := dataplane.GetDataAPIWrapper(coreDataClient)
	if err := data.Ready(context.Background()); err != nil {
		panic(fmt.Errorf("httpmq dataplane API not ready: %s", err))
	}

	fmt.Printf("httpmq API ready\n")
}
