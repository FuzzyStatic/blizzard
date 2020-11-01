package wowgd

// MythicRaidLeaderboard structure
type MythicRaidLeaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Slug         string `json:"slug"`
	CriteriaType string `json:"criteria_type"`
	Zone         struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"zone"`
	Entries []struct {
		Guild struct {
			Name  string `json:"name"`
			ID    int    `json:"id"`
			Realm struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"realm"`
		} `json:"guild"`
		Faction struct {
			Type string `json:"type"`
		} `json:"faction"`
		Timestamp int64  `json:"timestamp"`
		Region    string `json:"region"`
		Rank      int    `json:"rank"`
	} `json:"entries"`
}
