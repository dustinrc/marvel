package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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
