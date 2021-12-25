package main

import (
	"os"

	"github.com/apex/log"
	"github.com/urfave/cli/v2"
)

type cliArgs struct {
	JSONLog  bool
	LogLevel string `validate:"required,oneof=debug info warn error"`
	// For various subcommands
}

var cmdArgs cliArgs

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
		Flags: []cli.Flag{
			// LOGGING
			&cli.BoolFlag{
				Name:        "json-log",
				Usage:       "Whether to log in JSON format",
				Aliases:     []string{"j"},
				EnvVars:     []string{"LOG_AS_JSON"},
				Value:       false,
				DefaultText: "false",
				Destination: &cmdArgs.JSONLog,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "log-level",
				Usage:       "Logging level: [debug info warn error]",
				Aliases:     []string{"l"},
				EnvVars:     []string{"LOG_LEVEL"},
				Value:       "warn",
				DefaultText: "warn",
				Destination: &cmdArgs.LogLevel,
				Required:    false,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).WithFields(logTags).Fatal("Program shutdown")
	}
}
