package marvel_test

import (
	"strings"
	"testing"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestSeriesList(t *testing.T) {
	c := newTestClient(t, "series_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Series.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestSeriesListSeriesType(t *testing.T) {
	c := newTestClient(t, "series_list_series_type")
	defer c.stopRecorder()

	params := &marvel.SeriesParams{SeriesType: "one shot"}
	wrap, resp, err := c.Series.List(params)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "one shot", wrap.Data.Results[0].Type)
}

func TestSeriesGet(t *testing.T) {
	c := newTestClient(t, "series_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Series.Get(19244)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	series := wrap.Data.Results[0]
	assert.Equal(t, 19244, series.ID)
	assert.Contains(t, strings.ToLower(series.Title), "witch hunter")
}

func TestSeriesGetBadID(t *testing.T) {
	c := newTestClient(t, "series_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Series.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestSeriesCharacters(t *testing.T) {
	c1 := newTestClient(t, "series_characters1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_characters2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Series.Characters(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CharacterParams{Series: []int{12429}}
	wc2, _, _ := c2.Characters.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestSeriesComics(t *testing.T) {
	c1 := newTestClient(t, "series_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_comics2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Series.Comics(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.ComicParams{Series: []int{12429}}
	wc2, _, _ := c2.Comics.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestSeriesCreators(t *testing.T) {
	c1 := newTestClient(t, "series_creators1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_creators2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Series.Creators(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.CreatorParams{Series: []int{12429}}
	wc2, _, _ := c2.Creators.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestSeriesEvents(t *testing.T) {
	c1 := newTestClient(t, "series_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_events2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Series.Events(2116, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.EventParams{Series: []int{2116}}
	wc2, _, _ := c2.Events.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestSeriesStories(t *testing.T) {
	c1 := newTestClient(t, "series_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "series_stories2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Series.Stories(12429, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.StoryParams{Series: []int{12429}}
	wc2, _, _ := c2.Stories.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
