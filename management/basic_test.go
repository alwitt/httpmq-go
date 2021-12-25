package management

import (
	"context"
	"testing"

	"github.com/alwitt/httpmq-go/common"
	"github.com/apex/log"
	"github.com/stretchr/testify/assert"
)

func TestClientReadyCheck(t *testing.T) {
	assert := assert.New(t)
	log.SetLevel(log.DebugLevel)

	svrURL := common.GetUnitTestHttpmqMgmtAPIURL()
	coreClient, err := common.DefineAPIClient(svrURL, true)
	assert.Nil(err)

	uut := GetMgmtAPIWrapper(coreClient)

	utCtxt, utCtxtCancel := context.WithCancel(context.Background())
	defer utCtxtCancel()

	err = uut.Ready(utCtxt)
	assert.Nil(err)
}
