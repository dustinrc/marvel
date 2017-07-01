package marvel

// EventList provides event related to the parent entity.
type EventList struct {
	List
	Items []EventSummary `json:"items,omitempty"`
}

// EventSummary provides the summary for an event related to the parent entity.
type EventSummary struct {
	Summary
}
