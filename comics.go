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

// List returns all comics that match the query parameters. The comic
// slice will be encapsulated by ComicDataContainer and ComicDataWrapper.
func (cos *ComicService) List(params *ComicParams) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(cos.sling, "../comics", wrap, params)
	return wrap, resp, err
}

// Get returns the comic associated with the given ID. The comic
// details will be encapsulated by ComicDataContainer and ComicDataWrapper.
func (cos *ComicService) Get(comicID int) (*ComicDataWrapper, *http.Response, error) {
	wrap := &ComicDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d", comicID), wrap, nil)
	return wrap, resp, err
}

// Characters returns all characters involving the given comic and match the
// query parameters. The character slice will be encapsulated by CharacterDataContainer
// and CharacterDataWrapper.
func (cos *ComicService) Characters(comicID int, params *CharacterParams) (*CharacterDataWrapper, *http.Response, error) {
	wrap := &CharacterDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d/characters", comicID), wrap, params)
	return wrap, resp, err
}

// Creators returns all creators involving the given comic and match the
// query parameters. The creator slice will be encapsulated by CreatorDataContainer
// and CreatorDataWrapper.
func (cos *ComicService) Creators(comicID int, params *CreatorParams) (*CreatorDataWrapper, *http.Response, error) {
	wrap := &CreatorDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d/creators", comicID), wrap, params)
	return wrap, resp, err
}

// Events returns all events involving the given comic and match the
// query parameters. The event slice will be encapsulated by EventDataContainer
// and EventDataWrapper.
func (cos *ComicService) Events(comicID int, params *EventParams) (*EventDataWrapper, *http.Response, error) {
	wrap := &EventDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d/events", comicID), wrap, params)
	return wrap, resp, err
}

// Stories returns all stories involving the given comic and match the
// query parameters. The event slice will be encapsulated by StoryDataContainer
// and StoryDataWrapper.
func (cos *ComicService) Stories(comicID int, params *StoryParams) (*StoryDataWrapper, *http.Response, error) {
	wrap := &StoryDataWrapper{}
	resp, err := receiveWrapped(cos.sling, fmt.Sprintf("%d/stories", comicID), wrap, params)
	return wrap, resp, err
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
	Series             *SeriesSummary `json:"series,omitempty"`
	Variants           []ComicSummary `json:"variants,omitempty"`
	Collections        []ComicSummary `json:"collections,omitempty"`
	CollectedIssues    []ComicSummary `json:"collectedIssues,omitempty"`
	Dates              []ComicDate    `json:"dates,omitempty"`
	Prices             []ComicPrice   `json:"prices,omitempty"`
	Thumbnail          *Image         `json:"thumbnail,omitempty"`
	Images             []Image        `json:"images,omitempty"`
	Creators           CreatorList    `json:"creators,omitempty"`
	Characters         CharacterList  `json:"characters,omitempty"`
	Stories            StoryList      `json:"stories,omitempty"`
	Events             EventList      `json:"events,omitempty"`
}

// ComicParams are optional parameters to narrow the comic results returned
// by the API, as well as specify the number and order.
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
	SharedAppearances []int       `url:"sharedAppearances,omitempty"`
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
