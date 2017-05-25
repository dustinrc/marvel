package marvel

import (
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
