package marvel

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// CreatorService provides methods for querying creator information from the API.
type CreatorService struct {
	sling *sling.Sling
}

// NewCreatorService returns a new CreatorService.
func NewCreatorService(sling *sling.Sling) *CreatorService {
	return &CreatorService{
		sling: sling.Path("creators/"),
	}
}

// GetWrapped returns the creator associated with the given ID. The creator
// details will be encapsulated by CreatorDataContainer and CreatorDataWrapper.
func (ctrs *CreatorService) GetWrapped(creatorID int) (*CreatorDataWrapper, *http.Response, error) {
	wrap := &CreatorDataWrapper{}
	resp, err := receiveWrapped(ctrs.sling, fmt.Sprintf("%d", creatorID), wrap, nil)
	return wrap, resp, err
}

// Get returns the creator associated with the given ID.
func (ctrs *CreatorService) Get(creatorID int) (*Creator, error) {
	wrap, _, err := ctrs.GetWrapped(creatorID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// CreatorDataWrapper provides creator wrapper information returned by the API.
type CreatorDataWrapper struct {
	DataWrapper
	Data CreatorDataContainer `json:"data,omitempty"`
}

// CreatorDataContainer provides creator container information returned by the API.
type CreatorDataContainer struct {
	DataContainer
	Results []Creator `json:"results,omitempty"`
}

// Creator represents a Marvel comic creator.
type Creator struct {
	ID          int        `json:"id,omitempty"`
	FirstName   string     `json:"firstName,omitempty"`
	MiddleName  string     `json:"middleName,omitempty"`
	LastName    string     `json:"lastName,omitempty"`
	Suffix      string     `json:"suffix,omitempty"`
	FullName    string     `json:"fullName,omitempty"`
	Modified    Time       `json:"modified,omitempty"`
	ResourceURI string     `json:"resourceURI,omitempty"`
	URLs        []URL      `json:"urls,omitempty"`
	Thumbnail   Image      `json:"thumbnail,omitempty"`
	Series      SeriesList `json:"series,omitempty"`
	Stories     StoryList  `json:"stories,omitempty"`
	Comics      ComicList  `json:"comics,omitempty"`
	Events      EventList  `json:"events,omitempty"`
}

// CreatorList provides creators related to the parent entity.
type CreatorList struct {
	List
	Items []CreatorSummary `json:"items,omitempty"`
}

// CreatorSummary provides the summary for a creator related to the parent entity.
type CreatorSummary struct {
	Summary
	Role string `json:"role,omitempty"`
}
