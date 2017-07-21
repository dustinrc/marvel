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
