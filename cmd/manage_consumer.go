package cmd

import (
	"encoding/json"

	"github.com/alwitt/httpmq-go/api"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
)

/*
getMgntConsumerCLIFlags fetch the list CLI arguments common to all consumer management

 @param args *consumerManageCLIArgs - where CLI arguments are stored
 @return the list of CLI arguments
*/
func getMgntConsumerCLIFlags(args *consumerManageCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "stream",
			Usage:       "Target stream to operate against",
			Aliases:     []string{"s"},
			EnvVars:     []string{"TARGET_STREAM"},
			Destination: &args.stream,
			Required:    true,
		},
	}
}

/*
getMgntConsumerCLISubcmds fetch the list of subcommands for managing consumer through API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the list of stream CLI subcommands
*/
func getMgntConsumerCLISubcmds(mgntBaseArgs *ManagementCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "list",
			Usage:       "List all consumers",
			Description: "List all consumers of a stream through httpmq management API",
			Action:      actionListConsumers(mgntBaseArgs),
		},
		{
			Name:        "get",
			Usage:       "Fetch one consumer",
			Description: "Read information regarding one consumer through management API",
			Flags:       actionGetConsumerCLIFlags(&mgntBaseArgs.consumer.getConsumer),
			Action:      actionGetConsumer(mgntBaseArgs),
		},
		{
			Name:        "create",
			Usage:       "Define a new consumer",
			Description: "Define a new consumer through httpmq management API",
			Flags:       actionCreateConsumerCLIFags(&mgntBaseArgs.consumer.createConsumer),
			Action:      actionCreateConsumer(mgntBaseArgs),
		},
		{
			Name:        "delete",
			Usage:       "Delete a consumer",
			Description: "Delete a consumer through httpmq management API",
			Flags:       actionDeleteConsumerCLIFlags(&mgntBaseArgs.consumer.deleteConsumer),
			Action:      actionDeleteConsumer(mgntBaseArgs),
		},
	}
}

// ==============================================================================

/*
actionListStreams query the management API for list of all consumer on a stream

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionListConsumers(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, ctxt, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		reqID, consumer, err := client.ListAllConsumerOfStream(
			ctxt, mgntBaseArgs.consumer.stream,
		)
		if err != nil {
			log.WithError(err).Errorf("Failed to list all consumers")
			return err
		}
		type response struct {
			RequestID string                                     `json:"request_id"`
			Stream    string                                     `json:"stream"`
			Consumers map[string]api.ApisAPIRestRespConsumerInfo `json:"consumers"`
		}
		resp := response{RequestID: reqID, Stream: mgntBaseArgs.consumer.stream, Consumers: consumer}
		t, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.WithError(err).Errorf("Failed to JSON format consumer list")
			return err
		}
		log.Infof("Set of known consumers on stream %s\n%s", mgntBaseArgs.consumer.stream, string(t))
		return nil
	}
}

// ==============================================================================

// getConsumerCLIArgs cli arguments needed for query one consumer
type getConsumerCLIArgs struct {
	Name string `validate:"required"`
}

/*
actionGetConsumerCLIFlags fetch the list of CLI arguments needed by fetch consumer info

 @param args *getConsumerCLIArgs - arguments needed for fetch consumer info action
*/
func actionGetConsumerCLIFlags(args *getConsumerCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream consumer name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
	}
}

/*
actionGetConsumer fetch consumer info through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionGetConsumer(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, ctxt, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.consumer.getConsumer); err != nil {
			return err
		}
		reqID, info, err := client.GetConsumerOfStream(
			ctxt, mgntBaseArgs.consumer.stream, mgntBaseArgs.consumer.getConsumer.Name,
		)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to read consumer %s info", mgntBaseArgs.consumer.getConsumer.Name,
			)
			return err
		}
		type response struct {
			RequestID string                           `json:"request_id"`
			Stream    string                           `json:"stream"`
			Consumer  *api.ApisAPIRestRespConsumerInfo `json:"consumer"`
		}
		resp := response{RequestID: reqID, Stream: mgntBaseArgs.consumer.stream, Consumer: info}
		t, err := json.MarshalIndent(&resp, "", "  ")
		if err != nil {
			log.WithError(err).Errorf("Failed to JSON format consumer info")
			return err
		}
		log.Infof("Consumer %s\n%s", mgntBaseArgs.consumer.getConsumer.Name, string(t))
		return nil
	}
}

// ==============================================================================

// createConsumerCLIArgs cli arguments needed for create new consumer
type createConsumerCLIArgs struct {
	Name          string `validate:"required"`
	SubjectFilter string `validate:"required"`
	MaxInflight   int64  `validate:"required,gte=1"`
	DeliveryGroup string
}

/*
actionCreateConsumerCLIFags fetch the list of CLI arguments needed by create consumer

 @param args *createConsumerCLIArgs - arguments needed for create consumer
*/
func actionCreateConsumerCLIFags(args *createConsumerCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream consumer name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "subject-filter",
			Usage:       "Target subject filter",
			Aliases:     []string{"s"},
			Destination: &args.SubjectFilter,
			Required:    true,
		},
		&cli.Int64Flag{
			Name:        "max-inflight",
			Usage:       "Max number of inflight / unACKed messages allowed at once",
			Aliases:     []string{"m"},
			Value:       1,
			DefaultText: "1",
			Destination: &args.MaxInflight,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "delivery-group",
			Usage:       "Consumer delivery group",
			Aliases:     []string{"g"},
			Destination: &args.DeliveryGroup,
			Required:    false,
		},
	}
}

/*
actionCreateConsumer create new consumer through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionCreateConsumer(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, ctxt, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.consumer.createConsumer); err != nil {
			return err
		}
		params := api.ManagementJetStreamConsumerParam{
			Name:          mgntBaseArgs.consumer.createConsumer.Name,
			FilterSubject: api.PtrString(mgntBaseArgs.consumer.createConsumer.SubjectFilter),
			MaxInflight:   mgntBaseArgs.consumer.createConsumer.MaxInflight,
			Mode:          "push",
		}
		if len(mgntBaseArgs.consumer.createConsumer.DeliveryGroup) > 0 {
			params.DeliveryGroup = api.PtrString(mgntBaseArgs.consumer.createConsumer.DeliveryGroup)
		}
		reqID, err := client.CreateConsumerForStream(
			ctxt, mgntBaseArgs.consumer.stream, params,
		)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to define consumer %s", mgntBaseArgs.consumer.createConsumer.Name,
			)
			return err
		}
		log.Infof(
			"Defined new consumer %s on stream %s. Request ID %s",
			mgntBaseArgs.consumer.createConsumer.Name,
			mgntBaseArgs.consumer.stream,
			reqID,
		)
		return nil
	}
}

// ==============================================================================

// deleteConsumerCLIArgs cli arguments needed for delete one consumer
type deleteConsumerCLIArgs struct {
	Name string `validate:"required"`
}

/*
actionDeleteConsumerCLIFlags fetch the list of CLI arguments needed by delete consumer

 @param args *deleteConsumerCLIArgs - arguments needed for delete consumer action
*/
func actionDeleteConsumerCLIFlags(args *deleteConsumerCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Usage:       "JetStream consumer name",
			Aliases:     []string{"n"},
			Destination: &args.Name,
			Required:    true,
		},
	}
}

/*
actionDeleteConsumer delete consumer info through management API

 @param mgntBaseArgs *ManagementCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionDeleteConsumer(mgntBaseArgs *ManagementCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, ctxt, err := defineClientManagementAPI(mgntBaseArgs)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&mgntBaseArgs.consumer.deleteConsumer); err != nil {
			return err
		}
		reqID, err := client.DeleteConsumerOnStream(
			ctxt, mgntBaseArgs.consumer.stream, mgntBaseArgs.consumer.deleteConsumer.Name,
		)
		if err != nil {
			log.WithError(err).Errorf(
				"Failed to delete consumer %s", mgntBaseArgs.consumer.deleteConsumer.Name,
			)
			return err
		}
		log.Infof(
			"Delete consumer %s from stream %s. Request ID %s",
			mgntBaseArgs.consumer.deleteConsumer.Name,
			mgntBaseArgs.consumer.stream,
			reqID,
		)
		return nil
	}
}
