package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/dataplane"
	"github.com/apex/log"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
)

/*
GetDataplaneCLIFlags fetch the list of CLI arguments needed by dataplane API subcommands

 @param args *DataplaneCLIArgs - where CLI arguments are stored
 @return the list CLI arguments
*/
func GetDataplaneCLIFlags(args *DataplaneCLIArgs) []cli.Flag {
	flags := getCommonCLIFlags(&args.CommonCLIArgs)
	// HTTP client
	flags = append(
		flags, &cli.StringFlag{
			Name:        "dataplane-server-url",
			Usage:       "Dataplane server URL",
			Aliases:     []string{"s"},
			EnvVars:     []string{"DATAPLANE_SERVER_URL"},
			Value:       "http://127.0.0.1:4001",
			DefaultText: "http://127.0.0.1:4001",
			Destination: &args.HTTP.ServerURL,
			Required:    false,
		},
	)
	return flags
}

/*
GetDataplaneCLISubcmds fetch the list of subcommands supported for the dataplane API
*/
func GetDataplaneCLISubcmds(args *DataplaneCLIArgs) []*cli.Command {
	return []*cli.Command{
		{
			Name:        "publish",
			Aliases:     []string{"pub"},
			Usage:       "Publish messages",
			Description: "Publish messages on a subject through httpmq dataplane API",
			Flags:       actionPublishMessageCLIFlags(&args.publish),
			Action:      actionPublishMessage(args),
		},
		{
			Name:        "subscribe",
			Aliases:     []string{"sub"},
			Usage:       "Subscribe for messages",
			Description: "Subscribe for messages as a consumer on a stream through httpmq dataplane API",
			Flags:       actionSubscribeCLIFlags(&args.subscribe),
			Action:      actionSubscribe(args),
		},
	}
}

// ==============================================================================

// publishMessageCLIArgs cli arguments needed for publishing a message
type publishMessageCLIArgs struct {
	Subject string `validate:"required"`
	Message string `validate:"required"`
}

/*
actionPublishMessageCLIFlags fetch the list of CLI arguments needed to publish a message

 @param args *publishMessageCLIArgs - arguments needed for publishing a message
*/
func actionPublishMessageCLIFlags(args *publishMessageCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "subject",
			Usage:       "Subject to publish message on",
			Aliases:     []string{"s"},
			Destination: &args.Subject,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "message",
			Usage:       "Message body",
			Aliases:     []string{"m"},
			Destination: &args.Message,
			Required:    true,
		},
	}
}

/*
actionPublishMessage publish a message on a subject

 @param args *DataplaneCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionPublishMessage(args *DataplaneCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientDataplaneAPI(args)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&args.publish); err != nil {
			return err
		}
		reqID, err := client.Publish(
			context.Background(), args.publish.Subject, []byte(args.publish.Message),
		)
		if err != nil {
			log.WithError(err).Errorf("Failed to publish message on subject %s", args.publish.Subject)
			return err
		}
		log.Infof(
			"Published '%s' on subject %s. Request ID %s",
			args.publish.Message,
			args.publish.Subject,
			reqID,
		)
		return nil
	}
}

// ==============================================================================

// subscribeForMessageCLIArgs cli arguments needed for subcribing for messages
type subscribeForMessageCLIArgs struct {
	Stream        string `validate:"required"`
	Consumer      string `validate:"required"`
	Subject       string `validate:"required"`
	MaxInflight   int    `validate:"required,gte=1"`
	DeliveryGroup string
}

/*
actionSubscribeCLIFlags fetch the list of CLI arguments needed to subscribe for messages

 @param args *publishMessageCLIArgs - arguments needed for publishing a message
*/
func actionSubscribeCLIFlags(args *subscribeForMessageCLIArgs) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "stream",
			Usage:       "Stream to operate on",
			Aliases:     []string{"s"},
			Destination: &args.Stream,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "consumer",
			Usage:       "Consumer to operate as",
			Aliases:     []string{"c"},
			Destination: &args.Consumer,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "subject",
			Usage:       "Subject to subscribe for",
			Aliases:     []string{"t"},
			Destination: &args.Subject,
			Required:    true,
		},
		&cli.IntFlag{
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
			Usage:       "The consumer's delivery group",
			Aliases:     []string{"g"},
			Destination: &args.DeliveryGroup,
			Required:    false,
		},
	}
}

/*
actionSubscribe subscribe for mesages

 @param args *DataplaneCLIArgs - where CLI arguments are stored
 @return the CLI action for the subcommand
*/
func actionSubscribe(args *DataplaneCLIArgs) cli.ActionFunc {
	return func(c *cli.Context) error {
		client, err := defineClientDataplaneAPI(args)
		if err != nil {
			return err
		}
		validate := validator.New()
		if err := validate.Struct(&args.subscribe); err != nil {
			return err
		}

		// Subscribe support
		rxContext, cancel := context.WithCancel(context.Background())
		defer cancel()
		wg := sync.WaitGroup{}
		defer wg.Wait()
		msgChan := make(chan api.ApisAPIRestRespDataMessage, args.subscribe.MaxInflight+1)
		params := dataplane.PushSubscribeParam{
			Stream:         args.subscribe.Stream,
			Consumer:       args.subscribe.Consumer,
			SubjectFilter:  args.subscribe.Subject,
			MaxMsgInflight: &args.subscribe.MaxInflight,
			MsgChan:        &msgChan,
		}
		if len(args.subscribe.DeliveryGroup) > 0 {
			params.DeliveryGroup = &args.subscribe.DeliveryGroup
		}

		// Start receiving messages
		wg.Add(2)
		go func() {
			defer wg.Done()
			reqID, err := client.PushSubscribe(rxContext, params)
			log.WithError(err).Infof("Push subscription complete. Request ID %s", reqID)
			cancel()
		}()

		go func() {
			defer wg.Done()
			cc := make(chan os.Signal, 1)
			// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
			// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
			signal.Notify(cc, os.Interrupt)
			select {
			case <-cc:
				cancel()
				break
			case <-rxContext.Done():
				break
			}
		}()

		// Handle the messages
		complete := false
		for !complete {
			select {
			case rxMsg, ok := <-msgChan:
				if ok {
					log.Infof(
						"RX %s@%s/%s [S:%d, C:%d]: %s",
						rxMsg.Consumer,
						rxMsg.Stream,
						rxMsg.Subject,
						rxMsg.Sequence.Stream,
						rxMsg.Sequence.Consumer,
						rxMsg.B64Msg,
					)
					ackMsg := dataplane.MsgACKParam{
						Stream:      rxMsg.Stream,
						StreamSeq:   rxMsg.Sequence.Stream,
						Consumer:    rxMsg.Consumer,
						ConsumerSeq: rxMsg.Sequence.Consumer,
					}
					reqID, err := client.SendACK(rxContext, ackMsg)
					log.WithError(err).Infof(
						"Send ACK %s@%s/%s [S:%d, C:%d]. Request ID %s",
						rxMsg.Consumer,
						rxMsg.Stream,
						rxMsg.Subject,
						rxMsg.Sequence.Stream,
						rxMsg.Sequence.Consumer,
						reqID,
					)
				}
			case <-rxContext.Done():
				complete = true
			}
		}
		return nil
	}
}
