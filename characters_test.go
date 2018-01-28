package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestCharactersList(t *testing.T) {
	c := newTestClient(t, "characters_list")
	defer c.stopRecorder()

	wrap, resp, err := c.Characters.List(nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wrap.Data.Results)
}

func TestCharactersListBadParam(t *testing.T) {
	c := newTestClient(t, "characters_list_bad_param")
	defer c.stopRecorder()

	params := &marvel.CharacterParams{OrderBy: "superpower"}
	wrap, resp, err := c.Characters.List(params)
	assert.Error(t, err)
	assert.Equal(t, 409, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestCharactersGet(t *testing.T) {
	c := newTestClient(t, "characters_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Characters.Get(1017575)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, wrap.Data.Results, 1)
	char := wrap.Data.Results[0]
	assert.Equal(t, 1017575, char.ID)
	assert.Contains(t, strings.ToLower(char.Name), "wilson")
	assert.True(t, time.Now().After(char.Modified.Time))
}

func TestCharacterGetBadID(t *testing.T) {
	c := newTestClient(t, "characters_get_bad_id")
	defer c.stopRecorder()

	wrap, resp, err := c.Characters.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Empty(t, wrap.Data.Results)
}

func TestCharactersComics(t *testing.T) {
	c1 := newTestClient(t, "characters_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_comics2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Characters.Comics(1009149, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.ComicParams{Characters: []int{1009149}}
	wc2, _, _ := c2.Comics.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCharactersEvents(t *testing.T) {
	c1 := newTestClient(t, "characters_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_events2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Characters.Events(1010817, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.EventParams{Characters: []int{1010817}}
	wc2, _, _ := c2.Events.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCharactersSeries(t *testing.T) {
	c1 := newTestClient(t, "characters_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_series2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Characters.Series(1009149, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.SeriesParams{Characters: []int{1009149}}
	wc2, _, _ := c2.Series.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}

func TestCharactersStories(t *testing.T) {
	c1 := newTestClient(t, "characters_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_stories2")
	defer c2.stopRecorder()

	wc1, resp, err := c1.Characters.Stories(1009149, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, wc1.Data.Results)
	params := &marvel.StoryParams{Characters: []int{1009149}}
	wc2, _, _ := c2.Stories.List(params)
	assert.Equal(t, wc1.Data.Results, wc2.Data.Results)
}
