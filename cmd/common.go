package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/alwitt/httpmq-go/api"
	"github.com/alwitt/httpmq-go/common"
	"github.com/apex/log"
	apexJSON "github.com/apex/log/handlers/json"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
)

// loggingArgs cli arguments related to logging
type loggingArgs struct {
	// JSONLog whether to produce JSON formated logs
	JSONLog bool
	// LogLevel set the logging level
	LogLevel string `validate:"required,oneof=debug info warn error"`
}

// httpClientArgs cli arguments related to the HTTP client used by the API
type httpClientArgs struct {
	// ServerURL the URL for reaching the httpmq management API
	ServerURL string `validate:"required,url"`
	// CustomCA if provided, the CA to use
	CustomCA string `validate:"omitempty,file"`
}

// CommonCLIArgs cli arguments needed for operating against all APIs
type CommonCLIArgs struct {
	// Logging logging related configuration
	Logging loggingArgs `validate:"required,dive"`
	// HTTP are client related configuration
	HTTP httpClientArgs `validate:"required,dive"`
	// isDebug indicates whether application is operating in debug mode
	isDebug bool
}

/*
initialSetup perform basic application setup

 @return either setup is successful
*/
func (c *CommonCLIArgs) initialSetup(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return err
	}
	if c.Logging.JSONLog {
		log.SetHandler(apexJSON.New(os.Stderr))
	}
	c.isDebug = false
	switch c.Logging.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		c.isDebug = true
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}
	{
		tmp, _ := json.Marshal(c)
		log.Debugf("Starting common params %s", tmp)
	}
	return nil
}

// streamManageCLIArgs cli arguments needed for operation stream management
type streamManageCLIArgs struct {
	// createStream argument needed for defining new stream
	createStream createStreamCLIArgs `validate:"-"`
	// getStream argument needed to fetch info on a stream
	getStream getStreamCLIArgs `validate:"-"`
	// deleteStream argument needed to delete a stream
	deleteStream deleteStreamCLIArgs `validate:"-"`
	// changeSubject argument needed to change a stream's target subjects
	changeSubject createChangeSubjectsCLIArgs `validate:"-"`
	// changeRetention argument needed to change a stream's data retention
	changeRetention createChangeRetentionCLIArgs `validate:"-"`
}

// consumerManageCLIArgs cli arguments needed for operation consumer management
type consumerManageCLIArgs struct {
	stream string `validate:"required"`
	// getConsumer argument needed to get info on a consumer
	getConsumer getConsumerCLIArgs `validate:"-"`
	// createConsumer argument needed to create a new consumer
	createConsumer createConsumerCLIArgs `validate:"-"`
	// deleteConsumer argument needed to delete a consumer
	deleteConsumer deleteConsumerCLIArgs `validate:"-"`
}

// ManagementCLIArgs cli arguments needed for operating against management APIs
type ManagementCLIArgs struct {
	CommonCLIArgs
	// stream argument needed for stream management
	stream streamManageCLIArgs `validate:"-"`
	// consumer argument needed for consumer management
	consumer consumerManageCLIArgs `validate:"-"`
}

/*
getCommonCLIFlags fetch the list of CLI arguments common to both management and dataplane API
subcommands.

 @param args *CommonCLIArgs - where CLI arguments are stored
 @return the list of CLI arguments
*/
func getCommonCLIFlags(args *CommonCLIArgs) []cli.Flag {
	return []cli.Flag{
		// LOGGING
		&cli.BoolFlag{
			Name:        "json-log",
			Usage:       "Whether to log in JSON format",
			Aliases:     []string{"j"},
			EnvVars:     []string{"LOG_AS_JSON"},
			Value:       false,
			DefaultText: "false",
			Destination: &args.Logging.JSONLog,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "Logging level: [debug info warn error]",
			Aliases:     []string{"l"},
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       "info",
			DefaultText: "info",
			Destination: &args.Logging.LogLevel,
			Required:    false,
		},
		// HTTP client
		&cli.StringFlag{
			Name:        "custom-ca-file",
			Usage:       "Custom CA file to use with HTTP client",
			Aliases:     []string{"ccf"},
			EnvVars:     []string{"HTTP_CUSTOM_CA_FILE"},
			Value:       "",
			DefaultText: "",
			Destination: &args.HTTP.CustomCA,
			Required:    false,
		},
	}
}

/*
defineAPIClient define an httpmq API client

 @param config httpClientArgs - HTTP client config
 @param debug bool - whether to operate the API client in debug mode
 @return the httpmq API client
*/
func defineAPIClient(config httpClientArgs, debug bool) (*api.APIClient, error) {
	httpClient := http.Client{}
	// Define the TLS settings if custom CA was provided
	if len(config.CustomCA) > 0 {
		caCert, err := ioutil.ReadFile(config.CustomCA)
		if err != nil {
			log.WithError(err).Errorf("Unable to read %s", config.CustomCA)
			return nil, err
		}
		w, err := os.OpenFile("/tmp/tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.WithError(err).Error("Failed to open /tmp/tls-secrets.txt")
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig := &tls.Config{
			RootCAs:      caCertPool,
			KeyLogWriter: w,
		}
		httpClient.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
	} else {
		httpClient.Transport = &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		}
	}

	return common.DefineAPIClient(config.ServerURL, &httpClient, api.PtrString("httpmq-demo"), debug)
}
