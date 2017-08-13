package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestStoriesAll(t *testing.T) {
	c := newTestClient(t, "stories_all")
	defer c.stopRecorder()

	stories, err := c.Stories.All(nil)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.NotEmpty(t, stories, "Stories.All({}) returned empty stories list")
}

func TestStoriesAllOrderBy(t *testing.T) {
	c := newTestClient(t, "stories_all_order_by")
	defer c.stopRecorder()

	params := &marvel.StoryParams{OrderBy: "id"}
	stories, err := c.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, 7, stories[0].ID, "Incorrect ID")
	assert.Equal(t, 1, stories[0].Comics.Available)
}

func TestStoriesGet(t *testing.T) {
	c := newTestClient(t, "stories_get")
	defer c.stopRecorder()

	story, err := c.Stories.Get(16)
	assert.NoError(t, err, "Stories.Get() returned an error")

	assert.Equal(t, 16, story.ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(story.Title), "daredevil", "Incorrect Title")
	assert.Empty(t, story.Description, "Incorrect Description")
	assert.True(t, strings.HasSuffix(strings.ToLower(story.ResourceURI), "stories/16"), "Incorrect ResourceURI")
	assert.Equal(t, "story", story.Type, "Incorrect Type")
	assert.True(t, time.Now().After(story.Modified.Time), "Incorrect Modified")
	assert.Empty(t, story.Thumbnail, "Incorrect Thumbnail")
	assert.Equal(t, 0, story.Creators.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(story.Creators.CollectionURI), "stories/16/creators"), "Incorrect creators CollectionURI")
	assert.Equal(t, 0, story.Characters.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(story.Characters.CollectionURI), "stories/16/characters"), "Incorrect characters CollectionURI")
	assert.Equal(t, 1, story.Series.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(story.Series.CollectionURI), "stories/16/series"), "Incorrect series CollectionURI")
	assert.Equal(t, 1, story.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(story.Comics.CollectionURI), "stories/16/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 0, story.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(story.Events.CollectionURI), "stories/16/events"), "Incorrect events CollectionURI")
	assert.True(t, strings.HasSuffix(strings.ToLower(story.OriginalIssue.ResourceURI), "comics/950"), "Incorrect OriginalIssue")
}

func TestStoryGetBadID(t *testing.T) {
	c := newTestClient(t, "stories_get_bad_id")
	defer c.stopRecorder()

	story, err := c.Stories.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, story, "Story should have been nil")
}
