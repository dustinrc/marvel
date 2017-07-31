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
	Thumbnail   Image      `json:"thumbnail,omitempty"`
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
