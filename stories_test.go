package marvel_test

import (
	"strings"
	"testing"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestStoriesList(t *testing.T) {
	c := newTestClient(t, "stories_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Stories.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestStoriesListOrderBy(t *testing.T) {
	c := newTestClient(t, "stories_list_order_by")
	defer c.stopRecorder()

	params := &marvel.StoryParams{OrderBy: "id"}
	wrap, resp, err := c.Stories.List(params)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, wrap.Data.Results[0].ID < wrap.Data.Results[1].ID)
}

func TestStoriesGet(t *testing.T) {
	c := newTestClient(t, "stories_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Stories.Get(16)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	story := wrap.Data.Results[0]
	assert.Equal(t, 16, story.ID)
	assert.Contains(t, strings.ToLower(story.Title), "daredevil")
}

func TestStoryGetBadID(t *testing.T) {
	c := newTestClient(t, "stories_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Stories.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestStoriesCharacters(t *testing.T) {
	c1 := newTestClient(t, "stories_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "stories_characters2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Stories.Characters(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CharacterParams{Stories: []int{12429}}
	wc2, _, _ := c2.Characters.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestStoriesComics(t *testing.T) {
	c1 := newTestClient(t, "stories_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "stories_comics2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Stories.Comics(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.ComicParams{Stories: []int{12429}}
	wc2, _, _ := c2.Comics.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestStoriesCreators(t *testing.T) {
	c1 := newTestClient(t, "stories_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "stories_creators2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Stories.Creators(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CreatorParams{Stories: []int{12429}}
	wc2, _, _ := c2.Creators.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestStoriesEvents(t *testing.T) {
	c1 := newTestClient(t, "stories_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "stories_events2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Stories.Events(26280, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.EventParams{Stories: []int{26280}}
	wc2, _, _ := c2.Events.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestStoriesSeries(t *testing.T) {
	c1 := newTestClient(t, "stories_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "stories_series2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Stories.Series(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.SeriesParams{Stories: []int{12429}}
	wc2, _, _ := c2.Series.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
