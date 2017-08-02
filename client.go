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
	Comics     *ComicService
	Creators   *CreatorService
	Events     *EventService
	Series     *SeriesService
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
		Comics:     NewComicService(base.New()),
		Creators:   NewCreatorService(base.New()),
		Events:     NewEventService(base.New()),
		Series:     NewSeriesService(base.New()),
	}

	return c
}

// Request returns the currently prepared HTTP request.
func (c *Client) Request() (*http.Request, error) {
	return c.sling.Request()
}

// receiveWrapped prepares a request and unmarshals it into the provided wrapper.
func receiveWrapped(sling *sling.Sling, pathURL string, wrapperV, paramsV interface{}) (*http.Response, error) {
	apiErr := &APIError{}
	resp, err := sling.New().Get(pathURL).QueryStruct(paramsV).Receive(wrapperV, apiErr)
	if err == nil && apiErr.Code != nil {
		err = apiErr
	}
	return resp, err
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
