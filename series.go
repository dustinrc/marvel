package marvel

import (
	"fmt"
	"net/http"

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
	Thumbnail   Image          `json:"thumbnail,omitempty"`
	Comics      ComicList      `json:"comics,omitempty"`
	Stories     StoryList      `json:"stories,omitempty"`
	Events      EventList      `json:"events,omitempty"`
	Characters  CharacterList  `json:"characters,omitempty"`
	Creators    CreatorList    `json:"creators,omitempty"`
	Next        *SeriesSummary `json:"next,omitempty"`
	Previous    *SeriesSummary `json:"previous,omitempty"`
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
