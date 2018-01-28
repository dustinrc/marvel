package marvel_test

import (
	"strings"
	"testing"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestCreatorsList(t *testing.T) {
	c := newTestClient(t, "creators_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Creators.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestCreatorsListMiddleNameStartsWith(t *testing.T) {
	c := newTestClient(t, "creators_list_middle_starts_with")
	defer c.stopRecorder()

	params := &marvel.CreatorParams{MiddleNameStartsWith: "manu"}
	wrap, resp, err := c.Creators.List(params)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, strings.ToLower(wrap.Data.Results[0].MiddleName), "manu")
}

func TestCreatorsGet(t *testing.T) {
	c := newTestClient(t, "creators_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Creators.Get(4545)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	creator := wrap.Data.Results[0]
	assert.Equal(t, 4545, creator.ID)
	assert.Contains(t, strings.ToLower(creator.FirstName), "wayne")
}

func TestCreatorGetBadID(t *testing.T) {
	c := newTestClient(t, "creators_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Creators.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestCreatorsComics(t *testing.T) {
	c1 := newTestClient(t, "creators_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_comics2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Creators.Comics(2935, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.ComicParams{Creators: []int{2935}}
	wc2, _, _ := c2.Comics.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCreatorsEvents(t *testing.T) {
	c1 := newTestClient(t, "creators_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_events2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Creators.Events(24, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.EventParams{Creators: []int{24}}
	wc2, _, _ := c2.Events.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCreatorsSeries(t *testing.T) {
	c1 := newTestClient(t, "creators_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_series2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Creators.Series(2935, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.SeriesParams{Creators: []int{2935}}
	wc2, _, _ := c2.Series.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCreatorsStories(t *testing.T) {
	c1 := newTestClient(t, "creators_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "creators_stories2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Creators.Stories(2935, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.StoryParams{Creators: []int{2935}}
	wc2, _, _ := c2.Stories.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
