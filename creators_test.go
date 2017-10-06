package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestCreatorsAll(t *testing.T) {
	c := newTestClient(t, "creators_all")
	defer c.stopRecorder()

	creators, err := c.Creators.All(nil)
	assert.NoError(t, err, "Creators.All({}) returned an error")
	assert.NotEmpty(t, creators, "Creators.All({}) return empty creator list")
}

func TestCreatorsAllMiddleNameStartsWith(t *testing.T) {
	c := newTestClient(t, "creators_all_middle_starts_with")
	defer c.stopRecorder()

	params := &marvel.CreatorParams{MiddleNameStartsWith: "manu"}
	creators, err := c.Creators.All(params)
	assert.NoError(t, err, "Creator.All({}) returned and error")
	assert.Equal(t, 12533, creators[1].ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(creators[0].MiddleName), "manu", "Incorrect MiddleName")
}

func TestCreatorsGet(t *testing.T) {
	c := newTestClient(t, "creators_get")
	defer c.stopRecorder()

	creator, err := c.Creators.Get(4545)
	assert.NoError(t, err, "Creators.Get() returned an error for 4545")

	assert.Equal(t, 4545, creator.ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(creator.FirstName), "wayne", "Incorrect FirstName")
	assert.Empty(t, creator.MiddleName, "Incorrect MiddleName")
	assert.Contains(t, strings.ToLower(creator.LastName), "robinson", "Incorrect LastName")
	assert.Empty(t, creator.Suffix, "Incorrect Suffix")
	assert.Contains(t, strings.ToLower(creator.FullName), "wayne", "Incorrect LastName")
	assert.Contains(t, strings.ToLower(creator.FullName), "robinson", "Incorrect LastName")
	assert.True(t, time.Now().After(creator.Modified.Time), "Incorrect ModifiedTime")
	assert.NotEmpty(t, creator.Thumbnail, "Incorrect Thumbnail")
	assert.True(t, strings.HasSuffix(strings.ToLower(creator.ResourceURI), "creators/4545"), "Incorrect ResourceURI")
	assert.Equal(t, 3, creator.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(creator.Comics.CollectionURI), "creators/4545/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 3, creator.Series.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(creator.Series.CollectionURI), "creators/4545/series"), "Incorrect series CollectionURI")
	assert.Equal(t, 2, creator.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(creator.Stories.CollectionURI), "creators/4545/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 0, creator.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(creator.Events.CollectionURI), "creators/4545/events"), "Incorrect events CollectionURI")
	assert.NotEmpty(t, creator.URLs, "Incorrect URLs")
}

func TestCreatorGetBadID(t *testing.T) {
	c := newTestClient(t, "creators_get_bad_id")
	defer c.stopRecorder()

	creator, err := c.Creators.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, creator, "Creator should have been nil")
}

func TestCreatorsComics(t *testing.T) {
	c1 := newTestClient(t, "creators_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_comics2")
	defer c2.stopRecorder()

	comicsFromCreator, err := c1.Creators.Comics(2935, nil)
	assert.NoError(t, err, "Creators.Comics() returned an error for 2935")
	params := &marvel.ComicParams{Creators: []int{2935}}
	comicsLimitedToCreator, err := c2.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, comicsLimitedToCreator, comicsFromCreator, "Comic results do not match")
}

func TestCreatorsEvents(t *testing.T) {
	c1 := newTestClient(t, "creators_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_events2")
	defer c2.stopRecorder()

	eventsFromCreator, err := c1.Creators.Events(2935, nil)
	assert.NoError(t, err, "Creators.Events() returned an error for 2935")
	params := &marvel.EventParams{Creators: []int{2935}}
	eventsLimitedToCreator, err := c2.Events.All(params)
	assert.NoError(t, err, "Events.All({}) returned an error")
	assert.Equal(t, eventsLimitedToCreator, eventsFromCreator, "Event results do not match")
}

func TestCreatorsSeries(t *testing.T) {
	c1 := newTestClient(t, "creators_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_series2")
	defer c2.stopRecorder()

	seriesFromCreator, err := c1.Creators.Series(2935, nil)
	assert.NoError(t, err, "Creators.Series() returned an error for 2935")
	params := &marvel.SeriesParams{Creators: []int{2935}}
	seriesLimitedToCreator, err := c2.Series.All(params)
	assert.NoError(t, err, "Series.All({}) returned an error")
	assert.Equal(t, seriesLimitedToCreator, seriesFromCreator, "Series results do not match")
}

func TestCreatorsStories(t *testing.T) {
	c1 := newTestClient(t, "creators_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_stories2")
	defer c2.stopRecorder()

	storiesFromCreator, err := c1.Creators.Stories(2935, nil)
	assert.NoError(t, err, "Creators.Stories() returned an error for 2935")
	params := &marvel.StoryParams{Creators: []int{2935}}
	storiesLimitedToCreator, err := c2.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, storiesLimitedToCreator, storiesFromCreator, "Story results do not match")
}
