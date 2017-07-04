package marvel

// CreatorList provides creators related to the parent entity.
type CreatorList struct {
	List
	Items []CreatorSummary `json:"items,omitempty"`
}

// CreatorSummary provides the summary for a creator related to the parent entity.
type CreatorSummary struct {
	Summary
	Role string `json:"role,omitempty"`
}
