package marvel

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	// APIURL is the base URL for all API requests.
	APIURL = "https://gateway.marvel.com/v1/public"
)

// Client is a Marvel client for making all API requests.
type Client struct {
	auth  Authenticator
	sling *sling.Sling
}

// NewClient returns a Client that will authenticate according to the provided
// authenticator.
func NewClient(authenticator Authenticator) *Client {
	c := &Client{
		auth:  authenticator,
		sling: sling.New().Base(APIURL),
	}
	c.sling.QueryStruct(c.auth.Auth())

	return c
}

// Request returns the currently prepared HTTP request.
func (c *Client) Request() (*http.Request, error) {
	return c.sling.Request()
}

// APIError is the error, if any, returned by the service. Authentication error
// responses will have Code as a string. For usage errors otherwise, Code will be
// an integer.
type APIError struct {
	Code    interface{}
	Message string
}

// Error implements the Error interface.
func (ae *APIError) Error() string {
	return fmt.Sprintf("marvel: %v %v", ae.Code, ae.Message)
}
