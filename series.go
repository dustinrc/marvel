package marvel

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// SeriesService provides methods for querying series information from the API.
type SeriesService struct {
	sling *sling.Sling
}

// NewSeriesService returns a new SeriesService.
func NewSeriesService(sling *sling.Sling) *SeriesService {
	return &SeriesService{
		sling: sling.Path("series/"),
	}
}

// AllWrapped returns all series that match the query parameters. The series
// slice will be encapsulated by SeriesDataContainer and SeriesDataWrapper.
func (srs *SeriesService) AllWrapped(params *SeriesParams) (*SeriesDataWrapper, *http.Response, error) {
	wrap := &SeriesDataWrapper{}
	resp, err := receiveWrapped(srs.sling, "../series", wrap, params)
	return wrap, resp, err
}

// All returns all series that match the query parameters.
func (srs *SeriesService) All(params *SeriesParams) ([]Series, error) {
	wrap, _, err := srs.AllWrapped(params)
	return wrap.Data.Results, err
}

// GetWrapped returns the series associated with the given ID. The series
// details will be encapsulated by SeriesDataContainer and SeriesDataWrapper.
func (srs *SeriesService) GetWrapped(seriesID int) (*SeriesDataWrapper, *http.Response, error) {
	wrap := &SeriesDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d", seriesID), wrap, nil)
	return wrap, resp, err
}

// Get returns the series associated with the given ID.
func (srs *SeriesService) Get(seriesID int) (*Series, error) {
	wrap, _, err := srs.GetWrapped(seriesID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// CharactersWrapped returns all characters involving the given series and match the
// query parameters. The character slice will be encapsulated by CharacterDataContainer
// and CharacterDataWrapper.
func (srs *SeriesService) CharactersWrapped(seriesID int, params *CharacterParams) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d/characters", seriesID), wrap, params)
	return wrap, resp, err
}

// Characters returns all characters involving the given series and match the query parameters.
func (srs *SeriesService) Characters(seriesID int, params *CharacterParams) ([]Character, error) {
	wrap, _, err := srs.CharactersWrapped(seriesID, params)
	return wrap.Data.Results, err
}

// ComicsWrapped returns all comics involving the given series and match the
// query parameters. The comic slice will be encapsulated by ComicDataContainer
// and ComicDataWrapper.
func (srs *SeriesService) ComicsWrapped(seriesID int, params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d/comics", seriesID), wrap, params)
	return wrap, resp, err
}

// Comics returns all comics involving the given series and match the query parameters.
func (srs *SeriesService) Comics(seriesID int, params *ComicParams) ([]Comic, error) {
	wrap, _, err := srs.ComicsWrapped(seriesID, params)
	return wrap.Data.Results, err
}

// CreatorsWrapped returns all creators involving the given series and match the
// query parameters. The creator slice will be encapsulated by CreatorDataContainer
// and CreatorDataWrapper.
func (srs *SeriesService) CreatorsWrapped(seriesID int, params *CreatorParams) (*CreatorDataWrapper, *http.Response, error) {
	wrap := &CreatorDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d/creators", seriesID), wrap, params)
	return wrap, resp, err
}

// Creators returns all creators involving the given series and match the query parameters.
func (srs *SeriesService) Creators(seriesID int, params *CreatorParams) ([]Creator, error) {
	wrap, _, err := srs.CreatorsWrapped(seriesID, params)
	return wrap.Data.Results, err
}

// EventsWrapped returns all events involving the given series and match the
// query parameters. The event slice will be encapsulated by EventDataContainer
// and EventDataWrapper.
func (srs *SeriesService) EventsWrapped(eventID int, params *EventParams) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d/events", eventID), wrap, params)
	return wrap, resp, err
}

// Events returns all events involving the given series and match the query parameters.
func (srs *SeriesService) Events(eventID int, params *EventParams) ([]Event, error) {
	wrap, _, err := srs.EventsWrapped(eventID, params)
	return wrap.Data.Results, err
}

// StoriesWrapped returns all stories involving the given series and match the
// query parameters. The story slice will be encapsulated by StoryDataContainer
// and StoryDataWrapper.
func (srs *SeriesService) StoriesWrapped(seriesID int, params *StoryParams) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(srs.sling, fmt.Sprintf("%d/stories", seriesID), wrap, params)
	return wrap, resp, err
}

// Stories returns all stories involving the given series and match the query parameters.
func (srs *SeriesService) Stories(seriesID int, params *StoryParams) ([]Story, error) {
	wrap, _, err := srs.StoriesWrapped(seriesID, params)
	return wrap.Data.Results, err
}

// SeriesDataWrapper provides series wrapper information returned by the API.
type SeriesDataWrapper struct {
	DataWrapper
	Data SeriesDataContainer `json:"data,omitempty"`
}

// SeriesDataContainer provides series container information returned by the API.
type SeriesDataContainer struct {
	DataContainer
	Results []Series `json:"results,omitempty"`
}

// Series represents a Marvel comic series.
type Series struct {
	ID          int            `json:"id,omitempty"`
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	ResourceURI string         `json:"resourceURI,omitempty"`
	URLs        []URL          `json:"urls,omitempty"`
	StartYear   int            `json:"startYear,omitempty"`
	EndYear     int            `json:"endYear,omitempty"`
	Rating      string         `json:"rating,omitempty"`
	Type        string         `json:"type,omitempty"`
	Modified    Time           `json:"modified,omitempty"`
	Thumbnail   *Image         `json:"thumbnail,omitempty"`
	Comics      ComicList      `json:"comics,omitempty"`
	Stories     StoryList      `json:"stories,omitempty"`
	Events      EventList      `json:"events,omitempty"`
	Characters  CharacterList  `json:"characters,omitempty"`
	Creators    CreatorList    `json:"creators,omitempty"`
	Next        *SeriesSummary `json:"next,omitempty"`
	Previous    *SeriesSummary `json:"previous,omitempty"`
}

// SeriesParams are optional parameters to narrow the series results returned
// by the API, as well as specify the number and order.
type SeriesParams struct {
	Title           string    `url:"title,omitempty"`
	TitleStartsWith string    `url:"titleStartsWith,omitempty"`
	StartYear       int       `url:"startYear,omitempty"`
	ModifiedSince   time.Time `url:"modifiedSince,omitempty"`
	Comics          []int     `url:"comics,omitempty"`
	Stories         []int     `url:"stories,omitempty"`
	Events          []int     `url:"events,omitempty"`
	Creators        []int     `url:"creators,omitempty"`
	Characters      []int     `url:"characters,omitempty"`
	SeriesType      string    `url:"seriesType,omitempty"`
	Contains        string    `url:"contains,omitempty"`
	OrderBy         string    `url:"orderBy,omitempty"`
	Limit           int       `url:"limit,omitempty"`
	Offset          int       `url:"offset,omitempty"`
}

// SeriesList provides series related to the parent entity.
type SeriesList struct {
	List
	Items []SeriesSummary `json:"items,omitempty"`
}

// SeriesSummary provides summary for a series related to the parent entity.
type SeriesSummary struct {
	Summary
}
