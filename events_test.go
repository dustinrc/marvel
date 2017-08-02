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
