# Hello World

Basic example showing how to define client instance for operating the `management` and `dataplane` API.

```golang
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
```

> **IMPORTANT:** We define two separate transport clients for reaching `management` and `dataplane` respectively, as they listen on two different URIs.

```shell
$ ./hello-world.bin
2022/01/02 15:07:16
GET /v1/admin/ready HTTP/1.1
Host: 127.0.0.1:4000
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 15:07:16
HTTP/1.1 200 OK
Content-Length: 16
Content-Type: application/json
Date: Sun, 02 Jan 2022 23:07:16 GMT

{"success":true}
2022/01/02 15:07:16
GET /v1/data/ready HTTP/1.1
Host: 127.0.0.1:4001
User-Agent: OpenAPI-Generator/1.0.0/go
Accept: application/json
Accept-Encoding: gzip


2022/01/02 15:07:16
HTTP/1.1 200 OK
Content-Length: 16
Content-Type: application/json
Date: Sun, 02 Jan 2022 23:07:16 GMT

{"success":true}
httpmq API ready
```
