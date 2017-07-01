package marvel

import (
	"time"
)

// DataWrapper provides the common wrapper attributes to unmarshal the API's response.
// It is used to compose more specific wrappers, e.g., CharacterDataWrapper.
type DataWrapper struct {
	Code            int    `json:"code,omitempty"`
	Status          string `json:"status,omitempty"`
	Copyright       string `json:"copyright,omitempty"`
	AttributionText string `json:"attributionText,omitempty"`
	AttributionHTML string `json:"attributionHTML,omitempty"`
	ETag            string `json:"etag,omitempty"`
}

// DataContainer provides the common container attributes to unmarshal the API's response.
// It is used to compose more specific containers, e.g., CharacterDataContainer.
type DataContainer struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
	Total  int `json:"total,omitempty"`
	Count  int `json:"count,omitempty"`
}

// URL represents a public website related to the parent entity.
type URL struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Image represents the available graphics related to the parent entity. This information
// is the basis for constructing paths to specific image variants. See
// https://developer.marvel.com/documentation/images for details.
type Image struct {
	Path      string `json:"path,omitempty"`
	Extension string `json:"extension,omitempty"`
}

// Time allows unique parsing of the time format given in the API's responses.
type Time struct{ time.Time }

// UnmarshalJSON implements the json.Unmarshaler interface. The timezone format
// returned by the API does not parse using any of the default formats in the time
// package.
func (tm *Time) UnmarshalJSON(b []byte) (err error) {
	if tempTime, err := time.Parse(`"2006-01-02T15:04:05Z0700"`, string(b)); err == nil {
		tm.Time = tempTime
	}
	return
}

// List provides the common list attributes to unmarshal the API's response.
// It is used to compose more specific containers, e.g., CharacterList.
type List struct {
	Available     int    `json:"available,omitempty"`
	Returned      int    `json:"returned,omitempty"`
	CollectionURI string `json:"collectionURI,omitempty"`
}

// Summary provides the common summary attributes to unmarshal the API's response.
// It is used to compose more specific containers, e.g., CharacterSummary.
type Summary struct {
	ResourceURI string `json:"resourceURI,omitempty"`
	Name        string `json:"name,omitempty"`
}
