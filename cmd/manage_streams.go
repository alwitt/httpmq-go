package cmd

import (
	"github.com/urfave/cli/v2"
)

// getMgntStreamsCliSubcmds produces a list of subcommands for managing streams through API
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

func actionListAllStreams(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		return listAllStreams(mgntBaseArgs.ServerURL)
	}
}

// ===========================================================================

func listAllStreams(serverURL string) error {
	return nil
}
