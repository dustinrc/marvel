package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestEventsAll(t *testing.T) {
	c := newTestClient(t, "events_all")
	defer c.stopRecorder()

	events, err := c.Events.All(nil)
	assert.NoError(t, err, "Events.All({}) returned and error")
	assert.NotEmpty(t, events, "Events.All({}) returned empty event list")
}

func TestEventsAllCharacters(t *testing.T) {
	c := newTestClient(t, "events_all_characters")
	defer c.stopRecorder()

	params := &marvel.EventParams{Characters: []int{1010817}}
	events, err := c.Events.All(params)
	assert.NoError(t, err, "Events.All({}) returned an error")
	assert.Equal(t, 318, events[0].ID, "Incorrect ID")
	assert.Equal(t, 277, events[3].ID, "Incorrect ID")
}

func TestEventGet(t *testing.T) {
	c := newTestClient(t, "events_get")
	defer c.stopRecorder()

	event, err := c.Events.Get(314)
	assert.NoError(t, err, "Events.Get() returned an error")

	assert.Equal(t, 314, event.ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(event.Title), "ultron", "Incorrect Title")
	assert.Contains(t, strings.ToLower(event.Description), "avengers", "Incorrect Description")
	assert.True(t, strings.HasSuffix(strings.ToLower(event.ResourceURI), "events/314"), "Incorrect ResourceURI")
	assert.NotEmpty(t, event.URLs, "Incorrect URLs")
	assert.True(t, time.Now().After(event.Modified.Time), "Incorrect ModifiedTime")
	assert.True(t, event.End.Time.After(event.Start.Time), "Incorrect Start / End - ", event.Start.Time, " / ", event.End.Time)
	assert.NotEmpty(t, event.Thumbnail, "Incorrect Thumbnail")
	assert.Equal(t, 30, event.Creators.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(event.Creators.CollectionURI), "events/314/creators"), "Incorrect creators CollectionURI")
	assert.Equal(t, 12, event.Characters.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(event.Characters.CollectionURI), "events/314/characters"), "Incorrect characters CollectionURI")
	assert.Equal(t, 40, event.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(event.Stories.CollectionURI), "events/314/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 20, event.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(event.Comics.CollectionURI), "events/314/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 7, event.Series.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(event.Series.CollectionURI), "events/314/series"), "Incorrect series CollectionURI")
	assert.Contains(t, strings.ToLower(event.Next.Name), "infinity", "Incorrect Next")
	assert.Contains(t, strings.ToLower(event.Previous.Name), "marvel now", "Incorrect Previous")
}

func TestEventGetBadID(t *testing.T) {
	c := newTestClient(t, "events_get_bad_id")
	defer c.stopRecorder()

	event, err := c.Events.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, event, "Event should have been nil")
}

func TestEventsCharacters(t *testing.T) {
	c1 := newTestClient(t, "events_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_characters2")
	defer c2.stopRecorder()

	charactersFromEvent, err := c1.Events.Characters(227, nil)
	assert.NoError(t, err, "Events.Characters() returned an error for 227")
	params := &marvel.CharacterParams{Events: []int{227}}
	charactersLimitedToEvent, err := c2.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, charactersLimitedToEvent, charactersFromEvent, "Character results do not match")
}

func TestEventsComics(t *testing.T) {
	c1 := newTestClient(t, "events_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_comics2")
	defer c2.stopRecorder()

	comicsFromEvent, err := c1.Events.Comics(227, nil)
	assert.NoError(t, err, "Events.Comics() returned an error for 227")
	params := &marvel.ComicParams{Events: []int{227}}
	comicsLimitedToEvent, err := c2.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, comicsLimitedToEvent, comicsFromEvent, "Comic results do not match")
}

func TestEventsCreators(t *testing.T) {
	c1 := newTestClient(t, "events_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_creators2")
	defer c2.stopRecorder()

	creatorsFromEvent, err := c1.Events.Creators(227, nil)
	assert.NoError(t, err, "Events.Creators() returned an error for 227")
	params := &marvel.CreatorParams{Events: []int{227}}
	creatorsLimitedToEvent, err := c2.Creators.All(params)
	assert.NoError(t, err, "Creators.All({}) returned an error")
	assert.Equal(t, creatorsLimitedToEvent, creatorsFromEvent, "Creator results do not match")
}

func TestEventsSeries(t *testing.T) {
	c1 := newTestClient(t, "events_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_series2")
	defer c2.stopRecorder()

	seriesFromEvent, err := c1.Events.Series(227, nil)
	assert.NoError(t, err, "Events.Series() returned an error for 227")
	params := &marvel.SeriesParams{Events: []int{227}}
	seriesLimitedToEvent, err := c2.Series.All(params)
	assert.NoError(t, err, "Series.All({}) returned an error")
	assert.Equal(t, seriesLimitedToEvent, seriesFromEvent, "Series results do not match")
}

func TestEventsStories(t *testing.T) {
	c1 := newTestClient(t, "events_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_stories2")
	defer c2.stopRecorder()

	storiesFromEvent, err := c1.Events.Stories(227, nil)
	assert.NoError(t, err, "Events.Stories() returned an error for 227")
	params := &marvel.StoryParams{Events: []int{227}}
	storiesLimitedToEvent, err := c2.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, storiesLimitedToEvent, storiesFromEvent, "Story results do not match")
}
