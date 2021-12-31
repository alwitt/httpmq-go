package main

import (
	"os"

	"github.com/alwitt/httpmq-go/cmd"
	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

var mgmtCLIArgs cmd.ManagementCLIArgs
var dataCLIArgs cmd.DataplaneCLIArgs

var logTags log.Fields

func main() {
	logTags = log.Fields{
		"module":    "main",
		"component": "main",
	}

	app := &cli.App{
		Version:     "v0.1.0",
		Usage:       "HTTP MQ demo application",
		Description: "Demo application for trying out functionalities of httpmq",
		Commands: []*cli.Command{
			{
				Name:        "management",
				Usage:       "management API client",
				Aliases:     []string{"mgmt"},
				Description: "Operate the httpmq management API",
				Flags:       cmd.GetManagementCLIFlags(&mgmtCLIArgs),
				Subcommands: cmd.GetManagementCLISubcmds(&mgmtCLIArgs),
			},
			{
				Name:        "dataplane",
				Usage:       "dataplane API client",
				Aliases:     []string{"data"},
				Description: "Operate the httpmq dataplane API",
				Flags:       cmd.GetDataplaneCLIFlags(&dataCLIArgs),
				Subcommands: cmd.GetDataplaneCLISubcmds(&dataCLIArgs),
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).WithFields(logTags).Fatal("Program shutdown")
	}
}
