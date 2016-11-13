package hetzner

import (
	"fmt"
	"net/http"
	"strings"
)

// An ErrorResponse reports the error caused by an API request
type APIError struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`

	// Array of missing input parameters or null
	Missing []string `json:"missing"`
	// Array of invalid input paramaters or null
	Invalid []string `json:"invalid"`

	// Maximum allowed requests
	MaxRequest int `json:"max_request"`
	// Time interval in seconds
	Interval int `json:"interval"`
}

func (r *APIError) Error() string {
	if r.Status == 400 && r.Code == "INVALID_INPUT" {
		return fmt.Sprintf("%v %v: %d %v %v (missing %v) (invalid %v)",
			r.Response.Request.Method, r.Response.Request.URL, r.Status, r.Code, r.Message, strings.Join(r.Missing, ","), strings.Join(r.Invalid, ","))
	} else if r.Status == 403 && r.Code == "RATE_LIMIT_EXCEEDED" {
		return fmt.Sprintf("%v %v: %d %v %v (max_request %v) (internal %v sec)",
			r.Response.Request.Method, r.Response.Request.URL, r.Status, r.Code, r.Message, r.MaxRequest, r.Interval)
	} else {
		return fmt.Sprintf("%v %v: %d %v %v",
			r.Response.Request.Method, r.Response.Request.URL, r.Status, r.Code, r.Message)
	}
}
