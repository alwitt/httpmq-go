package common

import (
	"fmt"
	"strings"

	"github.com/alwitt/httpmq-go/api"
)

// HttpmqAPIError is the custom error returned by httpmq
type HttpmqAPIError struct {
	// RequestID is the request ID to match against logs
	RequestID string
	// StatusCode is the response code returned by httpmq
	StatusCode int
	// Message is an optional descriptive message
	Message *string
	// Detail is an optional descriptive message providing additional details on the error
	Detail *string
}

// Error implements the error interface for HttpmqAPIError
func (e HttpmqAPIError) Error() string {
	msgBuilder := strings.Builder{}
	msgBuilder.WriteString(fmt.Sprintf("REQ '%s' ==> RESP %d", e.RequestID, e.StatusCode))
	if e.Message != nil {
		msgBuilder.WriteString(fmt.Sprintf(", %s", *e.Message))
	}
	if e.Detail != nil {
		msgBuilder.WriteString(fmt.Sprintf(": %s", *e.Detail))
	}
	return msgBuilder.String()
}

/*
GenerateHttpmqError generate a HttpmqAPIError object from the error response from httpmq

 @param reqID string - the request ID
 @param err *api.ApisErrorDetail - the error message
 @return a HttpmqAPIError instance
*/
func GenerateHttpmqError(reqID string, statusCode int, err *api.GoutilsErrorDetail) HttpmqAPIError {
	var msg, detail *string
	if err != nil {
		m, _ := err.GetMessageOk()
		msg = m
		d, _ := err.GetDetailOk()
		detail = d
	}
	return HttpmqAPIError{
		RequestID: reqID, StatusCode: statusCode, Message: msg, Detail: detail,
	}
}
