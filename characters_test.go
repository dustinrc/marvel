package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCharactersAll(t *testing.T) {
	c := newTestClient(t, "characters_all")
	defer c.stopRecorder()

	chars, err := c.Characters.All()
	assert.NoError(t, err, "Characters.All() returned an error")
	assert.NotEmpty(t, chars, "Characters.All() returned empty character list")
}

func TestCharactersGetWrapped(t *testing.T) {
	c := newTestClient(t, "characters_get")
	defer c.stopRecorder()

	wrap, resp, err := c.Characters.GetWrapped(1017575)
	assert.NoError(t, err, "Characters.GetWrapped() returned an error")
	assert.NotNil(t, resp)

	t.Run("CharacterDataWrapper is correct", func(t *testing.T) {
		assert.Equal(t, 200, wrap.Code, "Incorrect Code")
		assert.Equal(t, "ok", strings.ToLower(wrap.Status), "Incorrect Status")
		assert.Contains(t, strings.ToLower(wrap.Copyright), "marvel", "Incorrect Copyright")
		assert.Contains(t, strings.ToLower(wrap.AttributionText), "marvel", "Incorrect AttributionText")
		assert.Contains(t, strings.ToLower(wrap.AttributionHTML), "marvel", "Incorrect AttributionHTML")
		assert.NotEmpty(t, wrap.ETag, "Empty ETag")
	})
	t.Run("CharacterDataContainer is correct", func(t *testing.T) {
		assert.Equal(t, 0, wrap.Data.Offset)
		assert.Equal(t, 20, wrap.Data.Limit)
		assert.Equal(t, 1, wrap.Data.Total)
		assert.Equal(t, 1, wrap.Data.Count)
	})
}

func TestCharactersGet(t *testing.T) {
	c := newTestClient(t, "characters_get")
	defer c.stopRecorder()

	char, err := c.Characters.Get(1017575)
	assert.NoError(t, err, "Characters.Get() returned an error for 1017575")

	assert.Equal(t, 1017575, char.ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(char.Name), "wilson", "Incorrect Name")
	assert.Contains(t, strings.ToLower(char.Description), "falcon", "Incorrect Description")
	assert.True(t, time.Now().After(char.Modified.Time), "Incorrect Modified time")
	assert.NotEmpty(t, char.Thumbnail, "Incorrect Thumbnail")
	assert.True(t, strings.HasSuffix(strings.ToLower(char.ResourceURI), "characters/1017575"), "Incorrect ResourceURI")
	assert.Equal(t, 25, char.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Comics.CollectionURI), "characters/1017575/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 9, char.Series.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Series.CollectionURI), "characters/1017575/series"), "Incorrect series CollectionURI")
	assert.Equal(t, 25, char.Stories.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Stories.CollectionURI), "characters/1017575/stories"), "Incorrect stories CollectionURI")
	assert.Equal(t, 1, char.Events.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Events.CollectionURI), "characters/1017575/events"), "Incorrect events CollectionURI")
	assert.NotEmpty(t, char.URLs, "Incorrect URLs")
}

func TestCharacterGetBadID(t *testing.T) {
	c := newTestClient(t, "characters_get_bad_id")
	defer c.stopRecorder()

	char, err := c.Characters.Get(-1)
	assert.Error(t, err)
	assert.Nil(t, char, "Character should have been nil")
}
