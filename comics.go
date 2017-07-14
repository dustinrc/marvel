package marvel

import (
	"fmt"
	"net/http"
	"time"

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

// AllWrapped returns all comics that match the query parameters. The comic
// slice will be encapsulated by ComicDataContainer and ComicDataWrapper.
func (cos *ComicService) AllWrapped(params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	apiErr := &APIError{}
	resp, err := cos.sling.New().Get("../comics").QueryStruct(params).Receive(wrap, apiErr)
	if err == nil && apiErr.Code != nil {
		err = apiErr
	}
	return wrap, resp, err
}

// All returns all comics that match the query parameters.
func (cos *ComicService) All(params *ComicParams) ([]Comic, error) {
	wrap, _, err := cos.AllWrapped(params)
	if err != nil {
		return nil, err
	}
	return wrap.Data.Results, nil
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

// ComicParams are optional parameters to narrow the comic results returned
// by the API, as well as specifiy the number and order.
type ComicParams struct {
	Format            string      `url:"format,omitempty"`
	FormatType        string      `url:"formatType,omitempty"`
	NoVariants        bool        `url:"noVariants,omitempty"`
	DateDescriptor    string      `url:"dateDescriptor,omitempty"`
	DateRange         []time.Time `url:"dateRange,omitempty"`
	Title             string      `url:"title,omitempty"`
	TitleStartsWith   string      `url:"titleStartsWith,omitempty"`
	StartYear         int         `url:"startYear,omitempty"`
	IssueNumber       int         `url:"issueNumber,omitempty"`
	DiamondCode       string      `url:"diamondCode,omitempty"`
	DigitalID         int         `url:"digitalId,omitempty"`
	UPC               string      `url:"upc,omitempty"`
	ISBN              string      `url:"isbn,omitempty"`
	EAN               string      `url:"ean,omitempty"`
	ISSN              string      `url:"issn,omitempty"`
	HasDigitalIssue   bool        `url:"hasDigitalIssue,omitempty"`
	ModifiedSince     time.Time   `url:"modifiedSince,omitempty"`
	Creators          []int       `url:"creators,omitempty"`
	Characters        []int       `url:"characters,omitempty"`
	Series            []int       `url:"series,omitempty"`
	Events            []int       `url:"events,omitempty"`
	Stories           []int       `url:"stories,omitempty"`
	SharedAppearances []int       `url:"shared_appearances,omitempty"`
	Collaborators     []int       `url:"collaborators,omitempty"`
	OrderBy           string      `url:"orderBy,omitempty"`
	Limit             int         `url:"limit,omitempty"`
	Offset            int         `url:"offset,omitempty"`
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
