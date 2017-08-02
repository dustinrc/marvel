package marvel

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// EventService provides methods for querying event information from the API.
type EventService struct {
	sling *sling.Sling
}

// NewEventService returns a new EventService.
func NewEventService(sling *sling.Sling) *EventService {
	return &EventService{
		sling: sling.Path("events/"),
	}
}

// AllWrapped returns all events that match the query parameters. The event
// slice will be encapsulated by EventDataContainer and EventDataWrapper.
func (evs *EventService) AllWrapped(params *EventParams) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(evs.sling, "../events", wrap, params)
	return wrap, resp, err
}

// All returns all events that match the query parameters.
func (evs *EventService) All(params *EventParams) ([]Event, error) {
	wrap, _, err := evs.AllWrapped(params)
	return wrap.Data.Results, err
}

// GetWrapped returns the event associated with the given ID. The event
// details will be encapsulated by EventDataContainer and EventDataWrapper.
func (evs *EventService) GetWrapped(eventID int) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d", eventID), wrap, nil)
	return wrap, resp, err
}

// Get returns the event associated with the given ID.
func (evs *EventService) Get(eventID int) (*Event, error) {
	wrap, _, err := evs.GetWrapped(eventID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// EventDataWrapper provides event wrapper information returned by the API.
type EventDataWrapper struct {
	DataWrapper
	Data EventDataContainer `json:"data,omitempty"`
}

// EventDataContainer provides event container information returned by the API.
type EventDataContainer struct {
	DataContainer
	Results []Event `json:"results,omitempty"`
}

// Event represents a Marvel comic event.
type Event struct {
	ID          int           `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	ResourceURI string        `json:"resourceURI,omitempty"`
	URLs        []URL         `json:"urls,omitempty"`
	Modified    Time          `json:"modified,omitempty"`
	Start       Time          `json:"start,omitempty"`
	End         Time          `json:"end,omitempty"`
	Thumbnail   Image         `json:"thumbnail,omitempty"`
	Comics      ComicList     `json:"comics,omitempty"`
	Stories     StoryList     `json:"stories,omitempty"`
	Series      SeriesList    `json:"series,omitempty"`
	Characters  CharacterList `json:"characters,omitempty"`
	Creators    CreatorList   `json:"creators,omitempty"`
	Next        *EventSummary `json:"next,omitempty"`
	Previous    *EventSummary `json:"previous,omitempty"`
}

// EventParams are optional parameters to narrow the event results returned
// by the API, as well as specify the number and order.
type EventParams struct {
	Name           string    `url:"name,omitempty"`
	NameStartsWith string    `url:"nameStartsWith,omitempty"`
	ModifiedSince  time.Time `url:"modifiedSince,omitempty"`
	Creators       []int     `url:"creators,omitempty"`
	Characters     []int     `url:"characters,omitempty"`
	Series         []int     `url:"series,omitempty"`
	Comics         []int     `url:"comics,omitempty"`
	Stories        []int     `url:"stories,omitempty"`
	OrderBy        string    `url:"orderBy,omitempty"`
	Limit          int       `url:"limit,omitempty"`
	Offset         int       `url:"offset,omitempty"`
}

// EventList provides event related to the parent entity.
type EventList struct {
	List
	Items []EventSummary `json:"items,omitempty"`
}

// EventSummary provides the summary for an event related to the parent entity.
type EventSummary struct {
	Summary
}
