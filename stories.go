package marvel

// StoryList provides stories related to the parent entity.
type StoryList struct {
	List
	Items []StorySummary `json:"items,omitempty"`
}

// StorySummary provides the summary for a story related to the parent entity.
type StorySummary struct {
	Summary
	Type string `json:"type,omitempty"`
}
