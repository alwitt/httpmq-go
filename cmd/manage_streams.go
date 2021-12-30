package cmd

import (
	"context"
	"encoding/json"

	"github.com/alwitt/httpmq-go/api"
	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

/*
getMgntStreamsCliSubcmds fetch the list of subcommands for managing streams through API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the list of stream CLI subcommands
*/
func getMgntStreamsCliSubcmds(mgntBaseArgs *ManagementCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "list-all",
			Usage:       "List all streams",
			Description: "List all JetStream streams through httpmq management API",
			Action:      actionListAllStreams(mgntBaseArgs),
		},
	}
}

/*
actionListAllStreams query the management API for list of all streams

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionListAllStreams(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := mgntBaseArgs.InitialSetup(); err != nil {
			return err
		}
		client, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		reqID, streams, err := client.ListAllStreams(context.Background())
		if err != nil {
			log.WithError(err).Errorf("Failed to list all streams")
			return err
		}
		type response struct {
			RequestID string                                   `json:"request_id"`
			Streams   map[string]api.ApisAPIRestRespStreamInfo `json:"streams"`
		}
		resp := response{RequestID: reqID, Streams: streams}
		t, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.WithError(err).Errorf("Failed to JSON format stream list")
			return err
		}
		log.Infof("Set of known streams\n%s", string(t))
		return nil
	}
}
