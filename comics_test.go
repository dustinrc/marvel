package marvel_test

import (
	"testing"
	"time"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestComicsGet(t *testing.T) {
	c := newTestClient(t, "comics_get")
	defer c.stopRecorder()

	comic, err := c.Comics.Get(61292)
	assert.NoError(t, err, "Comics.Get() returned an error for 61292")

	assert.Equal(t, 61292, comic.ID, "Incorrect ID")
	assert.Equal(t, 0, comic.DigitalID, "Incorrect DigitalID")
	assert.Contains(t, strings.ToLower(comic.Title), "guardians", "Incorrect Title")
	assert.Equal(t, 17, comic.IssueNumber, "Incorrect IssueNumber")
	assert.Empty(t, comic.VariantDescription, "Incorrect VariantDescription")
	assert.Contains(t, strings.ToLower(comic.Description), "thanos", "Incorrect Description")
	assert.True(t, time.Now().After(comic.Modified.Time), "Incorrect ModifiedTime")
	assert.Empty(t, comic.ISBN, "Incorrect ISBN")
	assert.Equal(t, "759606082941001711", comic.UPC, "Incorrect UPC")
	assert.Empty(t, comic.DiamondCode, "Incorrect DiamondCode")
	assert.Empty(t, comic.EAN, "Incorrect EAN")
	assert.Empty(t, comic.ISSN, "Incorrect ISSN")
	assert.Equal(t, "comic", strings.ToLower(comic.Format), "Incorrect Format")
	assert.Equal(t, 32, comic.PageCount, "Incorrect PageCount")
	assert.NotEmpty(t, comic.TextObjects[0].Type, "Incorrect TextObjects")
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.ResourceURI), "comics/61292"), "Incorrect ResourceURI")
	assert.NotEmpty(t, comic.URLs, "Incorrect URLs")
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Series.ResourceURI), "series/20365"), "Incorrect Series")
	assert.Empty(t, comic.Variants, "Incorrect Variants")
	assert.Empty(t, comic.Collections, "Incorrect Collections")
	assert.Empty(t, comic.CollectedIssues, "Incorrect CollectedIssues")
	assert.Equal(t, "onsaledate", strings.ToLower(comic.Dates[0].Type), "Incorrect Dates.Type")
	assert.True(t, time.Now().After(comic.Dates[0].Date.Time), "Incorrect Dates.Date")
	assert.Equal(t, "printprice", strings.ToLower(comic.Prices[0].Type), "Incorrect Prices.Type")
	assert.Equal(t, 2.99, comic.Prices[0].Price, "Incorrect Prices.Price")
	assert.NotEmpty(t, comic.Thumbnail, "Incorrect Thumbnail")
	assert.Equal(t, comic.Thumbnail, comic.Images[0], "Incorrect Images")
	assert.Equal(t, 1, comic.Creators.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Creators.CollectionURI), "comics/61292/creators"), "Incorrect creators CollectionURI")
	assert.Equal(t, 2, comic.Characters.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Characters.CollectionURI), "comics/61292/characters"), "Incorrect characters CollectionURI")
	assert.Equal(t, 2, comic.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Stories.CollectionURI), "comics/61292/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 0, comic.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Events.CollectionURI), "comics/61292/events"), "Incorrect events CollectionURI")
}
