package common

import (
	"os"
)

/*
GetUnitTestHttpmqMgmtAPIURL is a helper function to build the httpmq management API URL for testing

 @return the httpmq management API URL
*/
func GetUnitTestHttpmqMgmtAPIURL() string {
	if url, ok := os.LookupEnv("HTTPMQ_MGMT_API_URL"); ok {
		return url
	}
	return "http://127.0.0.1:4000"
}

/*
GetUnitTestHttpmqDataAPIURL is a helper function to build the httpmq dataplane API URL for testing

 @return the httpmq dataplane API URL
*/
func GetUnitTestHttpmqDataAPIURL() string {
	if url, ok := os.LookupEnv("HTTPMQ_DATA_API_URL"); ok {
		return url
	}
	return "http://127.0.0.1:4001"
}
