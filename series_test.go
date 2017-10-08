package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestSeriesAll(t *testing.T) {
	c := newTestClient(t, "series_all")
	defer c.stopRecorder()

	series, err := c.Series.All(nil)
	assert.NoError(t, err, "Series.All({}) returned an error")
	assert.NotEmpty(t, series, "Series.All({}) returned emtpy series list")
}

func TestSeriesAllSeriesType(t *testing.T) {
	c := newTestClient(t, "series_all_series_type")
	defer c.stopRecorder()

	params := &marvel.SeriesParams{SeriesType: "one shot"}
	series, err := c.Series.All(params)
	assert.NoError(t, err, "Series.All({}) returned an error")
	assert.Equal(t, 8925, series[11].ID, "Incorrect ID")
	assert.Equal(t, 1, series[11].Comics.Available)
}

func TestSeriesGet(t *testing.T) {
	c := newTestClient(t, "series_get")
	defer c.stopRecorder()

	series, err := c.Series.Get(19244)
	assert.NoError(t, err, "Series.Get() returned an error")

	assert.Equal(t, 19244, series.ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(series.Title), "witch hunter", "Incorrect Title")
	assert.Empty(t, series.Description, "Incorrect Description")
	assert.True(t, strings.HasSuffix(strings.ToLower(series.ResourceURI), "series/19244"), "Incorrect ResourceURI")
	assert.NotEmpty(t, series.URLs, "Incorrect URLs")
	assert.Equal(t, 2015, series.StartYear, "Incorrect StartYear")
	assert.Equal(t, 2099, series.EndYear, "Incorrect EndYear")
	assert.Contains(t, strings.ToLower(series.Rating), "rated", "Incorrect Rating")
	assert.Contains(t, strings.ToLower(series.Type), "limited", "Incorrect Type")
	assert.True(t, time.Now().After(series.Modified.Time), "Incorrect Modified")
	assert.NotEmpty(t, series.Thumbnail, "Incorrect Thumbnail")
	assert.Equal(t, 7, series.Creators.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(series.Creators.CollectionURI), "series/19244/creators"), "Incorrect creators CollectionURI")
	assert.Equal(t, 1, series.Characters.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(series.Characters.CollectionURI), "series/19244/characters"), "Incorrect characters CollectionURI")
	assert.Equal(t, 16, series.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(series.Stories.CollectionURI), "series/19244/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 8, series.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(series.Comics.CollectionURI), "series/19244/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 0, series.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(series.Events.CollectionURI), "series/19244/events"), "Incorrect events CollectionURI")
	assert.Nil(t, series.Next, "Incorrect Next")
	assert.Nil(t, series.Previous, "Incorrect Previous")
}

func TestSeriesGetBadID(t *testing.T) {
	c := newTestClient(t, "series_get_bad_id")
	defer c.stopRecorder()

	series, err := c.Series.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, series, "Series should have been nil")
}

func TestSeriesCharacters(t *testing.T) {
	c1 := newTestClient(t, "series_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_characters2")
	defer c2.stopRecorder()

	charactersFromSeries, err := c1.Series.Characters(12429, nil)
	assert.NoError(t, err, "Series.Characters() returned an error for 12429")
	params := &marvel.CharacterParams{Series: []int{12429}}
	charactersLimitedToSeries, err := c2.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, charactersLimitedToSeries, charactersFromSeries, "Character results do not match")
}

func TestSeriesComics(t *testing.T) {
	c1 := newTestClient(t, "series_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_comics2")
	defer c2.stopRecorder()

	comicsFromSeries, err := c1.Series.Comics(12429, nil)
	assert.NoError(t, err, "Series.Comics() returned an error for 12429")
	params := &marvel.ComicParams{Series: []int{12429}}
	comicsLimitedToSeries, err := c2.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, comicsLimitedToSeries, comicsFromSeries, "Comic results do not match")
}

func TestSeriesCreators(t *testing.T) {
	c1 := newTestClient(t, "series_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_creators2")
	defer c2.stopRecorder()

	creatorsFromSeries, err := c1.Series.Creators(12429, nil)
	assert.NoError(t, err, "Series.Creators() returned an error for 12429")
	params := &marvel.CreatorParams{Series: []int{12429}}
	creatorsLimitedToSeries, err := c2.Creators.All(params)
	assert.NoError(t, err, "Creators.All({}) returned an error")
	assert.Equal(t, creatorsLimitedToSeries, creatorsFromSeries, "Creator results do not match")
}

func TestSeriesEvents(t *testing.T) {
	c1 := newTestClient(t, "series_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_events2")
	defer c2.stopRecorder()

	eventsFromSeries, err := c1.Series.Events(12429, nil)
	assert.NoError(t, err, "Series.Events() returned an error for 12429")
	params := &marvel.EventParams{Series: []int{12429}}
	eventsLimitedToSeries, err := c2.Events.All(params)
	assert.NoError(t, err, "Events.All({}) returned an error")
	assert.Equal(t, eventsLimitedToSeries, eventsFromSeries, "Event results do not match")
}

func TestSeriesStories(t *testing.T) {
	c1 := newTestClient(t, "series_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_stories2")
	defer c2.stopRecorder()

	storiesFromSeries, err := c1.Series.Stories(12429, nil)
	assert.NoError(t, err, "Series.Stories() returned an error for 12429")
	params := &marvel.StoryParams{Series: []int{12429}}
	storiesLimitedToSeries, err := c2.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, storiesLimitedToSeries, storiesFromSeries, "Story results do not match")
}
