package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestComicsList(t *testing.T) {
	c := newTestClient(t, "comics_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Comics.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestComicsListDateRange(t *testing.T) {
	c := newTestClient(t, "comics_list_date_range")
	defer c.stopRecorder()

	sDate := time.Date(2016, time.August, 17, 17, 46, 57, 123, time.UTC)
	eDate := time.Date(2016, time.September, 17, 17, 46, 57, 123, time.UTC)
	params := &marvel.ComicParams{DateRange: []time.Time{sDate, eDate}}
	wrap, resp, err := c.Comics.List(params)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, sDate.Before(wrap.Data.Results[0].Modified.Time))
	assert.True(t, eDate.After(wrap.Data.Results[0].Modified.Time))
}

func TestComicsGet(t *testing.T) {
	c := newTestClient(t, "comics_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Comics.Get(61292)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	comic := wrap.Data.Results[0]
	assert.Equal(t, 61292, comic.ID)
	assert.Contains(t, strings.ToLower(comic.Title), "guardians", "Incorrect Title")
}

func TestComicGetBadID(t *testing.T) {
	c := newTestClient(t, "comics_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Comics.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestComicsCharacters(t *testing.T) {
	c1 := newTestClient(t, "comics_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_characters2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Comics.Characters(61292, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CharacterParams{Comics: []int{61292}}
	wc2, _, _ := c2.Characters.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestComicsCreators(t *testing.T) {
	c1 := newTestClient(t, "comics_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_creators2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Comics.Creators(61292, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CreatorParams{Comics: []int{61292}}
	wc2, _, _ := c2.Creators.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestComicsEvents(t *testing.T) {
	c1 := newTestClient(t, "comics_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_events2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Comics.Events(17701, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.EventParams{Comics: []int{17701}}
	wc2, _, _ := c2.Events.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestComicsStories(t *testing.T) {
	c1 := newTestClient(t, "comics_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "comics_stories2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Comics.Stories(61292, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.StoryParams{Comics: []int{61292}}
	wc2, _, _ := c2.Stories.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
