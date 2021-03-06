package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestComicsAll(t *testing.T) {
	c := newTestClient(t, "comics_all")
	defer c.stopRecorder()

	comics, err := c.Comics.All(nil)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.NotEmpty(t, comics, "Comics.All({}) returned empty comic list")
}

func TestComicsAllFormat(t *testing.T) {
	c := newTestClient(t, "comics_all_format")
	defer c.stopRecorder()

	params := &marvel.ComicParams{Format: "graphic novel"}
	comics, err := c.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, 52761, comics[0].ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(comics[0].Title), "thanos", "Incorrect Title")
}

func TestComicsAllNoVariants(t *testing.T) {
	t.Run("NoVariant is false", func(t *testing.T) {
		c := newTestClient(t, "comics_all_no_variant_false")
		defer c.stopRecorder()

		params := &marvel.ComicParams{
			NoVariants: false,
			UPC:        "75960608297101621",
		}
		comics, err := c.Comics.All(params)
		assert.NoError(t, err, "Comics.All({}) returned an error")
		assert.Equal(t, 58584, comics[0].ID, "Incorrect ID")
	})
	t.Run("NoVariant is true", func(t *testing.T) {
		c := newTestClient(t, "comics_all_no_variant_true")
		defer c.stopRecorder()

		params := &marvel.ComicParams{
			NoVariants: true,
			UPC:        "75960608297101621",
		}
		comics, err := c.Comics.All(params)
		assert.NoError(t, err, "Comics.All({}) returned an error")
		assert.Empty(t, comics, "Comic list should have been empty")
	})
}

func TestComicsAllDateRange(t *testing.T) {
	c := newTestClient(t, "comics_all_date_range")
	defer c.stopRecorder()

	sDate := time.Date(2016, time.August, 17, 17, 46, 57, 123, time.UTC)
	eDate := time.Date(2016, time.September, 17, 17, 46, 57, 123, time.UTC)
	params := &marvel.ComicParams{DateRange: []time.Time{sDate, eDate}}
	comics, err := c.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, 57387, comics[0].ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(comics[0].Title), "black panther", "Incorrect Title")
}

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
	assert.Equal(t, *comic.Thumbnail, comic.Images[0], "Incorrect Images")
	assert.Equal(t, 1, comic.Creators.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Creators.CollectionURI), "comics/61292/creators"), "Incorrect creators CollectionURI")
	assert.Equal(t, 2, comic.Characters.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Characters.CollectionURI), "comics/61292/characters"), "Incorrect characters CollectionURI")
	assert.Equal(t, 2, comic.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Stories.CollectionURI), "comics/61292/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 0, comic.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(comic.Events.CollectionURI), "comics/61292/events"), "Incorrect events CollectionURI")
}

func TestComicGetBadID(t *testing.T) {
	c := newTestClient(t, "comics_get_bad_id")
	defer c.stopRecorder()

	comic, err := c.Comics.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, comic, "Comic should have been nil")
}

func TestComicsCharacters(t *testing.T) {
	c1 := newTestClient(t, "comics_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_characters2")
	defer c2.stopRecorder()

	charactersFromComic, err := c1.Comics.Characters(1010817, nil)
	assert.NoError(t, err, "Comics.Characters() returned an error for 1010817")
	params := &marvel.CharacterParams{Comics: []int{1010817}}
	charactersLimitedToComic, err := c2.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, charactersLimitedToComic, charactersFromComic, "Character results do not match")
}

func TestComicsCreators(t *testing.T) {
	c1 := newTestClient(t, "comics_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_creators2")
	defer c2.stopRecorder()

	creatorsFromComic, err := c1.Comics.Creators(1010817, nil)
	assert.NoError(t, err, "Comics.Creators() returned an error for 1010817")
	params := &marvel.CreatorParams{Comics: []int{1010817}}
	creatorsLimitedToComic, err := c2.Creators.All(params)
	assert.NoError(t, err, "Creators.All({}) returned an error")
	assert.Equal(t, creatorsLimitedToComic, creatorsFromComic, "Creator results do not match")
}

func TestComicsEvents(t *testing.T) {
	c1 := newTestClient(t, "comics_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_events2")
	defer c2.stopRecorder()

	eventsFromComic, err := c1.Comics.Events(1010817, nil)
	assert.NoError(t, err, "Comics.Events() returned an error for 1010817")
	params := &marvel.EventParams{Comics: []int{1010817}}
	eventsLimitedToComic, err := c2.Events.All(params)
	assert.NoError(t, err, "Events.All({}) returned an error")
	assert.Equal(t, eventsLimitedToComic, eventsFromComic, "Event results do not match")
}

func TestComicsStories(t *testing.T) {
	c1 := newTestClient(t, "comics_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_stories2")
	defer c2.stopRecorder()

	storiesFromComic, err := c1.Comics.Stories(1009149, nil)
	assert.NoError(t, err, "Comics.Stories() returned an error for 1009149")
	params := &marvel.StoryParams{Comics: []int{1009149}}
	storiesLimitedToComic, err := c2.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, storiesLimitedToComic, storiesFromComic, "Story results do not match")
}
