package marvel

// SeriesList provides series related to the parent entity.
type SeriesList struct {
	List
	Items []SeriesSummary `json:"items,omitempty"`
}

// SeriesSummary provides summary for a series related to the parent entity.
type SeriesSummary struct {
	Summary
}
