package marvel

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// CharacterService provides methods for querying character information from the API.
type CharacterService struct {
	sling *sling.Sling
}

// NewCharacterService returns a new CharacterService.
func NewCharacterService(sling *sling.Sling) *CharacterService {
	return &CharacterService{
		sling: sling.Path("characters/"),
	}
}

// AllWrapped returns all characters that match the query parameters. The character
// slice will be encapsulated by CharacterDataContainer and CharacterDataWrapper.
func (chs *CharacterService) AllWrapped(params *CharacterParams) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(chs.sling, "../characters", wrap, params)
	return wrap, resp, err
}

// All returns all characters that match the query parameters.
func (chs *CharacterService) All(params *CharacterParams) ([]Character, error) {
	wrap, _, err := chs.AllWrapped(params)
	return wrap.Data.Results, err
}

// GetWrapped returns the character associated with the given ID. The character
// details will be encapsulated by CharacterDataContainer and CharacterDataWrapper.
func (chs *CharacterService) GetWrapped(characterID int) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(chs.sling, fmt.Sprintf("%d", characterID), wrap, nil)
	return wrap, resp, err
}

// Get returns the character associated with the given ID.
func (chs *CharacterService) Get(characterID int) (*Character, error) {
	wrap, _, err := chs.GetWrapped(characterID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// ComicsWrapped returns all comics involving the given character and match the
// query parameters. The comic slice will be encapsulated by ComicDataContainer
// and ComicDataWrapper.
func (chs *CharacterService) ComicsWrapped(characterID int, params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(chs.sling, fmt.Sprintf("%d/comics", characterID), wrap, params)
	return wrap, resp, err
}

// Comics returns all comics involving the given character and match the query parameters.
func (chs *CharacterService) Comics(characterID int, params *ComicParams) ([]Comic, error) {
	wrap, _, err := chs.ComicsWrapped(characterID, params)
	return wrap.Data.Results, err
}

// EventsWrapped returns all events involving the given character and match the
// query parameters. The event slice will be encapsulated by EventDataContainer
// and EventDataWrapper.
func (chs *CharacterService) EventsWrapped(characterID int, params *EventParams) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(chs.sling, fmt.Sprintf("%d/events", characterID), wrap, params)
	return wrap, resp, err
}

// Events returns all events involving the given character and match the query parameters.
func (chs *CharacterService) Events(characterID int, params *EventParams) ([]Event, error) {
	wrap, _, err := chs.EventsWrapped(characterID, params)
	return wrap.Data.Results, err
}

// SeriesWrapped returns all series involving the given character and match the
// query parameters. The series slice will be encapsulated by SeriesDataContainer
// and SeriesDataWrapper.
func (chs *CharacterService) SeriesWrapped(characterID int, params *SeriesParams) (*SeriesDataWrapper, *http.Response, error) {
	wrap := &SeriesDataWrapper{}
	resp, err := receiveWrapped(chs.sling, fmt.Sprintf("%d/series", characterID), wrap, params)
	return wrap, resp, err
}

// Series returns all series involving the given character and match the query parameters.
func (chs *CharacterService) Series(characterID int, params *SeriesParams) ([]Series, error) {
	wrap, _, err := chs.SeriesWrapped(characterID, params)
	return wrap.Data.Results, err
}

// StoriesWrapped returns all stories involving the given character and match the
// query parameters. The story slice will be encapsulated by StoryDataContainer
// and StoryDataWrapper.
func (chs *CharacterService) StoriesWrapped(characterID int, params *StoryParams) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(chs.sling, fmt.Sprintf("%d/stories", characterID), wrap, params)
	return wrap, resp, err
}

// Stories returns all stories involving the given character and match the query parameters.
func (chs *CharacterService) Stories(characterID int, params *StoryParams) ([]Story, error) {
	wrap, _, err := chs.StoriesWrapped(characterID, params)
	return wrap.Data.Results, err
}

// CharacterDataWrapper provides character wrapper information returned by the API.
type CharacterDataWrapper struct {
	DataWrapper
	Data CharacterDataContainer `json:"data,omitempty"`
}

// CharacterDataContainer provides character container information returned by the API.
type CharacterDataContainer struct {
	DataContainer
	Results []Character `json:"results,omitempty"`
}

// Character represents a Marvel comic character.
type Character struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Modified    Time       `json:"modified,omitempty"`
	ResourceURI string     `json:"resourceURI,omitempty"`
	URLs        []URL      `json:"urls,omitempty"`
	Thumbnail   *Image     `json:"thumbnail,omitempty"`
	Comics      ComicList  `json:"comics,omitempty"`
	Stories     StoryList  `json:"stories,omitempty"`
	Events      EventList  `json:"events,omitempty"`
	Series      SeriesList `json:"series,omitempty"`
}

// CharacterParams are optional parameters to narrow the character results returned
// by the API, as well as specify the number and order.
type CharacterParams struct {
	Name           string    `url:"name,omitempty"`
	NameStartsWith string    `url:"nameStartsWith,omitempty"`
	ModifiedSince  time.Time `url:"modifiedSince,omitempty"`
	Comics         []int     `url:"comics,omitempty"`
	Series         []int     `url:"series,omitempty"`
	Events         []int     `url:"events,omitempty"`
	Stories        []int     `url:"stories,omitempty"`
	OrderBy        string    `url:"orderBy,omitempty"`
	Limit          int       `url:"limit,omitempty"`
	Offset         int       `url:"offset,omitempty"`
}

// CharacterList provides characters related to the parent entity.
type CharacterList struct {
	List
	Items []CharacterSummary `json:"items,omitempty"`
}

// CharacterSummary provides the summary for a character related to the parent entity.
type CharacterSummary struct {
	Summary
}
