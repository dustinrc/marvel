package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
