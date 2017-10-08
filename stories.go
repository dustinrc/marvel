package marvel

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// StoryService provides methods for querying story information from the API.
type StoryService struct {
	sling *sling.Sling
}

// NewStoryService returns a new StoryService.
func NewStoryService(sling *sling.Sling) *StoryService {
	return &StoryService{
		sling: sling.Path("stories/"),
	}
}

// AllWrapped returns all stories that match the query parameters. The story
// slice will be encapsulated by StoryDataContainer and StoryDataWrapper.
func (sts *StoryService) AllWrapped(params *StoryParams) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(sts.sling, "../stories", wrap, params)
	return wrap, resp, err
}

// All returns all stories that match the query parameters.
func (sts *StoryService) All(params *StoryParams) ([]Story, error) {
	wrap, _, err := sts.AllWrapped(params)
	return wrap.Data.Results, err
}

// GetWrapped returns the story associated with the given ID. The story
// details will be encapsulated by StoryDataContainer and StoryDataWrapper.
func (sts *StoryService) GetWrapped(storyID int) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d", storyID), wrap, nil)
	return wrap, resp, err
}

// Get returns the story associated with the given ID.
func (sts *StoryService) Get(storyID int) (*Story, error) {
	wrap, _, err := sts.GetWrapped(storyID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// CharactersWrapped returns all characters involving the given story and match the
// query parameters. The character slice will be encapsulated by CharacterDataContainer
// and CharacterDataWrapper.
func (sts *StoryService) CharactersWrapped(storyID int, params *CharacterParams) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d/characters", storyID), wrap, params)
	return wrap, resp, err
}

// Characters returns all characters involving the given story and match the query parameters.
func (sts *StoryService) Characters(storyID int, params *CharacterParams) ([]Character, error) {
	wrap, _, err := sts.CharactersWrapped(storyID, params)
	return wrap.Data.Results, err
}

// ComicsWrapped returns all comics involving the given story and match the
// query parameters. The comic slice will be encapsulated by ComicDataContainer
// and ComicDataWrapper.
func (sts *StoryService) ComicsWrapped(storyID int, params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d/comics", storyID), wrap, params)
	return wrap, resp, err
}

// Comics returns all comics involving the given story and match the query parameters.
func (sts *StoryService) Comics(storyID int, params *ComicParams) ([]Comic, error) {
	wrap, _, err := sts.ComicsWrapped(storyID, params)
	return wrap.Data.Results, err
}

// CreatorsWrapped returns all creators involving the given story and match the
// query parameters. The creator slice will be encapsulated by CreatorDataContainer
// and CreatorDataWrapper.
func (sts *StoryService) CreatorsWrapped(storyID int, params *CreatorParams) (*CreatorDataWrapper, *http.Response, error) {
	wrap := &CreatorDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d/creators", storyID), wrap, params)
	return wrap, resp, err
}

// Creators returns all creators involving the given story and match the query parameters.
func (sts *StoryService) Creators(storyID int, params *CreatorParams) ([]Creator, error) {
	wrap, _, err := sts.CreatorsWrapped(storyID, params)
	return wrap.Data.Results, err
}

// EventsWrapped returns all events involving the given story and match the
// query parameters. The event slice will be encapsulated by EventDataContainer
// and EventDataWrapper.
func (sts *StoryService) EventsWrapped(eventID int, params *EventParams) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d/events", eventID), wrap, params)
	return wrap, resp, err
}

// Events returns all events involving the given series and match the query parameters.
func (sts *StoryService) Events(eventID int, params *EventParams) ([]Event, error) {
	wrap, _, err := sts.EventsWrapped(eventID, params)
	return wrap.Data.Results, err
}

// SeriesWrapped returns all series involving the given series and match the
// query parameters. The series slice will be encapsulated by SeriesDataContainer
// and SeriesDataWrapper.
func (sts *StoryService) SeriesWrapped(seriesID int, params *SeriesParams) (*SeriesDataWrapper, *http.Response, error) {
	wrap := &SeriesDataWrapper{}
	resp, err := receiveWrapped(sts.sling, fmt.Sprintf("%d/series", seriesID), wrap, params)
	return wrap, resp, err
}

// Series returns all series involving the given series and match the query parameters.
func (sts *StoryService) Series(seriesID int, params *SeriesParams) ([]Series, error) {
	wrap, _, err := sts.SeriesWrapped(seriesID, params)
	return wrap.Data.Results, err
}

// StoryDataWrapper provides story wrapper information returned by the API.
type StoryDataWrapper struct {
	DataWrapper
	Data StoryDataContainer `json:"data,omitempty"`
}

// StoryDataContainer provides story container information returned by the API.
type StoryDataContainer struct {
	DataContainer
	Results []Story `json:"results,omitempty"`
}

// Story represents a Marvel comic story.
type Story struct {
	ID            int           `json:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	ResourceURI   string        `json:"resourceUri,omitempty"`
	Type          string        `json:"type,omitempty"`
	Modified      Time          `json:"modified,omitempty"`
	Thumbnail     *Image        `json:"thumbnail,omitempty"`
	Comics        ComicList     `json:"comics,omitempty"`
	Series        SeriesList    `json:"series,omitempty"`
	Events        EventList     `json:"events,omitempty"`
	Characters    CharacterList `json:"characters,omitempty"`
	Creators      CreatorList   `json:"creators,omitempty"`
	OriginalIssue *ComicSummary `json:"originalIssue,omitempty"`
}

// StoryParams are optional parameters to narrow the story results returned
// by the API, as well as specify the number and order.
type StoryParams struct {
	ModifiedSince time.Time `url:"modifiedSince,omitempty"`
	Comics        []int     `url:"comics,omitempty"`
	Series        []int     `url:"series,omitempty"`
	Events        []int     `url:"events,omitempty"`
	Creators      []int     `url:"creators,omitempty"`
	Characters    []int     `url:"characters,omitempty"`
	OrderBy       string    `url:"orderBy,omitempty"`
	Limit         int       `url:"limit,omitempty"`
	Offset        int       `url:"offset,omitempty"`
}

// StoryList provides stories related to the parent entity.
type StoryList struct {
	List
	Items []StorySummary `json:"items,omitempty"`
}

// StorySummary provides the summary for a story related to the parent entity.
type StorySummary struct {
	Summary
	Type string `json:"type,omitempty"`
}
