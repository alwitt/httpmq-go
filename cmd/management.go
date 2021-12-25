package cmd

import "github.com/urfave/cli/v2"

// ManagementCLIArgs cli arguments need for the operating against the management APIs
type ManagementCLIArgs struct {
	ServerURL string `validate:"required,url"`
}

// GetManagementCLIFlags retrieve the set of CMD flags for calling management APIs
func GetManagementCLIFlags(args *ManagementCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "management-server-url",
			Usage:       "Management server URL",
			Aliases:     []string{"s"},
			EnvVars:     []string{"MANAGEMENT_SERVER_URL"},
			Value:       "http://127.0.0.1:3000",
			DefaultText: "http://127.0.0.1:3000",
			Destination: &args.ServerURL,
			Required:    false,
		},
	}
}

// GetManagementCLISubcmds produces a list of subcommands supported by management API
func GetManagementCLISubcmds(args *ManagementCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "stream",
			Usage:       "Manage streams",
			Description: "Manages JetStream streams through httpmq management API",
			Subcommands: getMgntStreamsCliSubcmds(args),
		},
	}
}
