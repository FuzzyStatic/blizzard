package sps

// PlayableTitles represents the response returned by the Streaming
// Provider Service when querying which titles a player may launch.
//
// The API returns a list of title records; each record includes the
// integer ID and the human-readable name.
type PlayableTitles struct {
	Titles []Title `json:"titles"`
}

// Title holds information about a single playable title or sub-title.
type Title struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
