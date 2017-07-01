package marvel

import (
	"fmt"
	"net/http"

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
func (chs *CharacterService) AllWrapped() (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	apiErr := &APIError{}
	resp, err := chs.sling.New().Get("../characters").Receive(wrap, apiErr)
	if err == nil && apiErr.Code != nil {
		err = apiErr
	}
	return wrap, resp, err
}

// All returns all characters that match the query parameters.
func (chs *CharacterService) All() ([]Character, error) {
	wrap, _, err := chs.AllWrapped()
	if err != nil {
		return nil, err
	}
	return wrap.Data.Results, nil
}

// GetWrapped returns the character associated with the given ID. The character
// details will be encapsulated by CharacterDataContainer and CharacterDataWrapper.
func (chs *CharacterService) GetWrapped(characterID int) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	apiErr := &APIError{}
	resp, err := chs.sling.New().Get(fmt.Sprintf("%d", characterID)).Receive(wrap, apiErr)
	if err == nil && apiErr.Code != nil {
		err = apiErr
	}
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
