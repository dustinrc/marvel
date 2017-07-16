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
	resp, err := receiveWrapped(cos.sling, "../comics", wrap, params)
	return wrap, resp, err
}

// All returns all comics that match the query parameters.
func (cos *ComicService) All(params *ComicParams) ([]Comic, error) {
	wrap, _, err := cos.AllWrapped(params)
	return wrap.Data.Results, err
}

// GetWrapped returns the comic associated with the given ID. The comic
// details will be encapsulated by ComicDataContainer and ComicDataWrapper.
func (cos *ComicService) GetWrapped(comicID int) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d", comicID), wrap, nil)
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
	ID                 int            `json:"id,omitempty"`
	DigitalID          int            `json:"digitalId,omitempty"`
	Title              string         `json:"title,omitempty"`
	IssueNumber        int            `json:"issueNumber,omitempty"`
	VariantDescription string         `json:"variantDescription,omitempty"`
	Description        string         `json:"description,omitempty"`
	Modified           Time           `json:"modified,omitempty"`
	ISBN               string         `json:"isbn,omitempty"`
	UPC                string         `json:"upc,omitempty"`
	DiamondCode        string         `json:"diamondCode,omitempty"`
	EAN                string         `json:"ean,omitempty"`
	ISSN               string         `json:"issn,omitempty"`
	Format             string         `json:"format,omitempty"`
	PageCount          int            `json:"pageCount,omitempty"`
	TextObjects        []TextObject   `json:"textObjects,omitempty"`
	ResourceURI        string         `json:"resourceURI,omitempty"`
	URLs               []URL          `json:"urls,omitempty"`
	Series             SeriesSummary  `json:"series,omitempty"`
	Variants           []ComicSummary `json:"variants,omitempty"`
	Collections        []ComicSummary `json:"collections,omitempty"`
	CollectedIssues    []ComicSummary `json:"collectedIssues,omitempty"`
	Dates              []ComicDate    `json:"dates,omitempty"`
	Prices             []ComicPrice   `json:"prices,omitempty"`
	Thumbnail          Image          `json:"thumbnail,omitempty"`
	Images             []Image        `json:"images,omitempty"`
	Creators           CreatorList    `json:"creators,omitempty"`
	Characters         CharacterList  `json:"characters,omitempty"`
	Stories            StoryList      `json:"stories,omitempty"`
	Events             EventList      `json:"events,omitempty"`
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
