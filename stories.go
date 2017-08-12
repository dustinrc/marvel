package marvel

import (
	"fmt"
	"net/http"

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
