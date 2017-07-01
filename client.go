package marvel

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	// APIURL is the base URL for all API requests.
	APIURL = "https://gateway.marvel.com/v1/public/"
)

// Client is a Marvel client for making all API requests.
type Client struct {
	auth  Authenticator
	sling *sling.Sling

	Characters *CharacterService
}

// NewClient returns an API Client that will authenticate according to the provided
// authenticator. A custom http client may also be used, otherwise pass nil for the
// default.
func NewClient(authenticator Authenticator, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	base := sling.New().Client(httpClient).Base(APIURL)
	base.QueryStruct(authenticator.Auth())

	c := &Client{
		auth:  authenticator,
		sling: base,

		Characters: NewCharacterService(base.New()),
	}

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
