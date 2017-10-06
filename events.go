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

// CharactersWrapped returns all characters involving the given event and match the
// query parameters. The character slice will be encapsulated by CharacterDataContainer
// and CharacterDataWrapper.
func (evs *EventService) CharactersWrapped(eventID int, params *CharacterParams) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d/characters", eventID), wrap, params)
	return wrap, resp, err
}

// Characters returns all characters involving the given event and match the query parameters.
func (evs *EventService) Characters(eventID int, params *CharacterParams) ([]Character, error) {
	wrap, _, err := evs.CharactersWrapped(eventID, params)
	return wrap.Data.Results, err
}

// ComicsWrapped returns all comics involving the given event and match the
// query parameters. The comic slice will be encapsulated by ComicDataContainer
// and ComicDataWrapper.
func (evs *EventService) ComicsWrapped(eventID int, params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d/comics", eventID), wrap, params)
	return wrap, resp, err
}

// Comics returns all comics involving the given event and match the query parameters.
func (evs *EventService) Comics(eventID int, params *ComicParams) ([]Comic, error) {
	wrap, _, err := evs.ComicsWrapped(eventID, params)
	return wrap.Data.Results, err
}

// CreatorsWrapped returns all creators involving the given event and match the
// query parameters. The creator slice will be encapsulated by CreatorDataContainer
// and CreatorDataWrapper.
func (evs *EventService) CreatorsWrapped(eventID int, params *CreatorParams) (*CreatorDataWrapper, *http.Response, error) {
	wrap := &CreatorDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d/creators", eventID), wrap, params)
	return wrap, resp, err
}

// Creators returns all creators involving the given event and match the query parameters.
func (evs *EventService) Creators(eventID int, params *CreatorParams) ([]Creator, error) {
	wrap, _, err := evs.CreatorsWrapped(eventID, params)
	return wrap.Data.Results, err
}

// SeriesWrapped returns all series involving the given event and match the
// query parameters. The series slice will be encapsulated by SeriesDataContainer
// and SeriesDataWrapper.
func (evs *EventService) SeriesWrapped(eventID int, params *SeriesParams) (*SeriesDataWrapper, *http.Response, error) {
	wrap := &SeriesDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d/series", eventID), wrap, params)
	return wrap, resp, err
}

// Series returns all series involving the given event and match the query parameters.
func (evs *EventService) Series(eventID int, params *SeriesParams) ([]Series, error) {
	wrap, _, err := evs.SeriesWrapped(eventID, params)
	return wrap.Data.Results, err
}

// StoriesWrapped returns all stories involving the given event and match the
// query parameters. The story slice will be encapsulated by StoryDataContainer
// and StoryDataWrapper.
func (evs *EventService) StoriesWrapped(eventID int, params *StoryParams) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(evs.sling, fmt.Sprintf("%d/stories", eventID), wrap, params)
	return wrap, resp, err
}

// Stories returns all stories involving the given event and match the query parameters.
func (evs *EventService) Stories(eventID int, params *StoryParams) ([]Story, error) {
	wrap, _, err := evs.StoriesWrapped(eventID, params)
	return wrap.Data.Results, err
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
	Thumbnail   *Image        `json:"thumbnail,omitempty"`
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
