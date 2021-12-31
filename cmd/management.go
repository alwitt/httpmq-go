package cmd

import (
	"github.com/alwitt/httpmq-go/dataplane"
	"github.com/alwitt/httpmq-go/management"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
)

/*
GetManagementCLIFlags fetch the list of CLI arguments needed by management API subcommands

 @param args *ManagementCLIArgs - where CLI arguments are stored
 @return the list CLI arguments
*/
func GetManagementCLIFlags(args *ManagementCLIArgs) []cli.Flag {
	flags := getCommonCLIFlags(&args.CommonCLIArgs)
	// HTTP client
	flags = append(
		flags, &cli.StringFlag{
			Name:        "management-server-url",
			Usage:       "Management server URL",
			Aliases:     []string{"s"},
			EnvVars:     []string{"MANAGEMENT_SERVER_URL"},
			Value:       "http://127.0.0.1:4000",
			DefaultText: "http://127.0.0.1:4000",
			Destination: &args.HTTP.ServerURL,
			Required:    false,
		},
	)
	return flags
}

/*
GetManagementCLISubcmds fetch the list of subcommands supported for the management API

 @param args *ManagementCLIArgs - the structure where the CLI arguments are stored
 @return the list of CLI subcommands
*/
func GetManagementCLISubcmds(args *ManagementCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "stream",
			Usage:       "Manage streams",
			Description: "Manages streams through httpmq management API",
			Subcommands: getMgntStreamsCLISubcmds(args),
		},
		{
			Name:        "consumer",
			Usage:       "Manage consumers",
			Description: "Manages consumers through httpmq management API",
			Flags:       getMgntConsumerCLIFlags(&args.consumer),
			Subcommands: getMgntConsumerCLISubcmds(args),
		},
	}
}

/*
defineClientManagementAPI creates a httpmq client for management API

 @param args *ManagementCLIArgs - where CLI arguments are stored
 @return the httpmq client
*/
func defineClientManagementAPI(args *ManagementCLIArgs) (management.MgmtAPIWrapper, error) {
	validate := validator.New()
	if err := args.initialSetup(validate); err != nil {
		log.WithError(err).Errorf("Failed to parse command line arguments")
		return nil, err
	}
	if err := validate.Struct(args); err != nil {
		log.WithError(err).Errorf("Failed to parse command line arguments")
		return nil, err
	}

	// Define the client
	client, err := defineAPIClient(args.HTTP, args.isDebug)
	if err != nil {
		log.WithError(err).Errorf("Faild to define httpmq API client")
	}

	wrapper := management.GetMgmtAPIWrapper(client)
	return wrapper, err
}

/*
defineClientDataplaneAPI creates a httpmq client for dataplane API

 @param args *DataplaneCLIArgs - where CLI arguments are stored
 @return the httpmq client
*/
func defineClientDataplaneAPI(args *DataplaneCLIArgs) (dataplane.DataAPIWrapper, error) {
	validate := validator.New()
	if err := args.initialSetup(validate); err != nil {
		log.WithError(err).Errorf("Failed to parse command line arguments")
		return nil, err
	}
	if err := validate.Struct(args); err != nil {
		log.WithError(err).Errorf("Failed to parse command line arguments")
		return nil, err
	}

	// Define the client
	client, err := defineAPIClient(args.HTTP, args.isDebug)
	if err != nil {
		log.WithError(err).Errorf("Faild to define httpmq API client")
	}

	wrapper := dataplane.GetDataAPIWrapper(client)
	return wrapper, err
}
