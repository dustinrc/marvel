package marvel_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dustinrc/marvel"
	"github.com/stretchr/testify/assert"
)

func TestCharactersAll(t *testing.T) {
	c := newTestClient(t, "characters_all")
	defer c.stopRecorder()

	chars, err := c.Characters.All(nil)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.NotEmpty(t, chars, "Characters.All({}) returned empty character list")
}

func TestCharactersAllName(t *testing.T) {
	c := newTestClient(t, "characters_all_name")
	defer c.stopRecorder()

	params := &marvel.CharacterParams{Name: "Spider-Man"}
	chars, err := c.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, 1009610, chars[0].ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(chars[0].Name), "spider-man", "Incorrect Name")
}

func TestCharactersAllModifiedSince(t *testing.T) {
	c := newTestClient(t, "characters_all_modified_since")
	defer c.stopRecorder()

	modDate := time.Date(2016, time.August, 17, 17, 46, 57, 123, time.UTC)
	params := &marvel.CharacterParams{ModifiedSince: modDate}
	chars, err := c.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, 1009268, chars[2].ID, "Incorrect ID")
	assert.Contains(t, strings.ToLower(chars[2].Name), "deadpool", "Incorrect Name")
}

func TestCharactersAllOnlyInComics(t *testing.T) {
	c := newTestClient(t, "characters_all_only_in_comics")
	defer c.stopRecorder()

	params := &marvel.CharacterParams{Comics: []int{11200, 22222}}
	chars, err := c.Characters.All(params)
	assert.NoError(t, err, "Characters.All({}) returned an error")
	assert.Equal(t, 1009515, chars[0].ID, "Incorrect ID")
	assert.Equal(t, 1010791, chars[1].ID, "Incorrect ID")
}

func TestCharactersAllBadParam(t *testing.T) {
	c := newTestClient(t, "characters_all_bad_param")
	defer c.stopRecorder()

	params := &marvel.CharacterParams{OrderBy: "superpower"}
	chars, err := c.Characters.All(params)
	assert.Error(t, err)
	assert.Empty(t, chars, "Character list should have been empty")
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
	assert.Equal(t, 28, char.Comics.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Comics.CollectionURI), "characters/1017575/comics"), "Incorrect comics CollectionURI")
	assert.Equal(t, 11, char.Series.Available)
	assert.True(t, strings.HasSuffix(strings.ToLower(char.Series.CollectionURI), "characters/1017575/series"), "Incorrect series CollectionURI")
	assert.Equal(t, 28, char.Stories.Available)
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

func TestCharactersComics(t *testing.T) {
	c1 := newTestClient(t, "characters_comics1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_comics2")
	defer c2.stopRecorder()

	comicsFromCharacter, err := c1.Characters.Comics(1009149, nil)
	assert.NoError(t, err, "Characters.Comics() returned an error for 1009149")
	params := &marvel.ComicParams{Characters: []int{1009149}}
	comicsLimitedToCharacter, err := c2.Comics.All(params)
	assert.NoError(t, err, "Comics.All({}) returned an error")
	assert.Equal(t, comicsLimitedToCharacter, comicsFromCharacter, "Comic results do not match")
}

func TestCharactersEvents(t *testing.T) {
	c1 := newTestClient(t, "characters_events1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_events2")
	defer c2.stopRecorder()

	eventsFromCharacter, err := c1.Characters.Events(1010817, nil)
	assert.NoError(t, err, "Characters.Events() returned an error for 1010817")
	params := &marvel.EventParams{Characters: []int{1010817}}
	eventsLimitedToCharacter, err := c2.Events.All(params)
	assert.NoError(t, err, "Events.All({}) returned an error")
	assert.Equal(t, eventsLimitedToCharacter, eventsFromCharacter, "Event results do not match")
}

func TestCharactersSeries(t *testing.T) {
	c1 := newTestClient(t, "characters_series1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_series2")
	defer c2.stopRecorder()

	seriesFromCharacter, err := c1.Characters.Series(1009149, nil)
	assert.NoError(t, err, "Characters.Series() returned an error for 1009149")
	params := &marvel.SeriesParams{Characters: []int{1009149}}
	seriesLimitedToCharacter, err := c2.Series.All(params)
	assert.NoError(t, err, "Series.All({}) returned an error")
	assert.Equal(t, seriesLimitedToCharacter, seriesFromCharacter, "Series results do not match")
}

func TestCharactersStories(t *testing.T) {
	c1 := newTestClient(t, "characters_stories1")
	defer c1.stopRecorder()
	c2 := newTestClient(t, "characters_stories2")
	defer c2.stopRecorder()

	storiesFromCharacter, err := c1.Characters.Stories(1009149, nil)
	assert.NoError(t, err, "Characters.Stories() returned an error for 1009149")
	params := &marvel.StoryParams{Characters: []int{1009149}}
	storiesLimitedToCharacter, err := c2.Stories.All(params)
	assert.NoError(t, err, "Stories.All({}) returned an error")
	assert.Equal(t, storiesLimitedToCharacter, storiesFromCharacter, "Story results do not match")
}
