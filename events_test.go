package marvel_test

import (
	"strings"
	"testing"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestEventsList(t *testing.T) {
	c := newTestClient(t, "events_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Events.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestEventGet(t *testing.T) {
	c := newTestClient(t, "events_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Events.Get(314)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	event := wrap.Data.Results[0]
	assert.Equal(t, 314, event.ID)
	assert.Contains(t, strings.ToLower(event.Title), "ultron")
}

func TestEventGetBadID(t *testing.T) {
	c := newTestClient(t, "events_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Events.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestEventsCharacters(t *testing.T) {
	c1 := newTestClient(t, "events_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_characters2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Events.Characters(227, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CharacterParams{Events: []int{227}}
	wc2, _, _ := c2.Characters.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestEventsComics(t *testing.T) {
	c1 := newTestClient(t, "events_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_comics2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Events.Comics(227, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.ComicParams{Events: []int{227}}
	wc2, _, _ := c2.Comics.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestEventsCreators(t *testing.T) {
	c1 := newTestClient(t, "events_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_creators2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Events.Creators(227, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CreatorParams{Events: []int{227}}
	wc2, _, _ := c2.Creators.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestEventsSeries(t *testing.T) {
	c1 := newTestClient(t, "events_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_series2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Events.Series(227, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.SeriesParams{Events: []int{227}}
	wc2, _, _ := c2.Series.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestEventsStories(t *testing.T) {
	c1 := newTestClient(t, "events_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "events_stories2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Events.Stories(227, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.StoryParams{Events: []int{227}}
	wc2, _, _ := c2.Stories.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
