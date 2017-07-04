package marvel

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// ComicService provides methods for querying comic information from the API.
type ComicService struct {
	sling *sling.Sling
}

// NewComicService returns a new ComicService.
func NewComicService(sling *sling.Sling) *ComicService {
	return &ComicService{
		sling: sling.Path("comics/"),
	}
}

// GetWrapped returns the comic associated with the given ID. The comic
// details will be encapsulated by ComicDataContainer and ComicDataWrapper.
func (cos *ComicService) GetWrapped(comicID int) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	apiErr := &APIError{}
	resp, err := cos.sling.New().Get(fmt.Sprintf("%d", comicID)).Receive(wrap, apiErr)
	if err == nil && apiErr.Code != nil {
		err = apiErr
	}
	return wrap, resp, err
}

// Get returns the comic associated with the given ID.
func (cos *ComicService) Get(comicID int) (*Comic, error) {
	wrap, _, err := cos.GetWrapped(comicID)
	if err != nil {
		return nil, err
	}
	return &wrap.Data.Results[0], nil
}

// ComicDataWrapper provides comic wrapper information returned by the API.
type ComicDataWrapper struct {
	DataWrapper
	Data ComicDataContainer `json:"data,omitempty"`
}

// ComicDataContainer provides comic container information returned by the API.
type ComicDataContainer struct {
	DataContainer
	Results []Comic `json:"results,omitempty"`
}

// Comic represents a Marvel comic.
type Comic struct {
	ID                 int
	DigitalID          int
	Title              string
	IssueNumber        int
	VariantDescription string
	Description        string
	Modified           Time
	ISBN               string
	UPC                string
	DiamondCode        string
	EAN                string
	ISSN               string
	Format             string
	PageCount          int
	TextObjects        []TextObject
	ResourceURI        string
	URLs               []URL
	Series             SeriesSummary
	Variants           []ComicSummary
	Collections        []ComicSummary
	CollectedIssues    []ComicSummary
	Dates              []ComicDate
	Prices             []ComicPrice
	Thumbnail          Image
	Images             []Image
	Creators           CreatorList
	Characters         CharacterList
	Stories            StoryList
	Events             EventList
}

// ComicDate represents a moment of importance for the comic.
type ComicDate struct {
	Type string `json:"type,omitempty"`
	Date Time   `json:"date,omitempty"`
}

// ComicPrice represents the price of a comic in a certain medium.
type ComicPrice struct {
	Type  string  `json:"type,omitempty"`
	Price float64 `json:"price,omitempty"`
}

// ComicList provides comics related to the parent entity.
type ComicList struct {
	List
	Items []ComicSummary `json:"items,omitempty"`
}

// ComicSummary provides the summary for a comic related to the parent entity.
type ComicSummary struct {
	Summary
}
