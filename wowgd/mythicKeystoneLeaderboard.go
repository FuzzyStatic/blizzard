package wowgd

// MythicKeystoneLeaderboardIndex structure
type MythicKeystoneLeaderboardIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CurrentLeaderboards []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"current_leaderboards"`
}

// MythicKeystoneLeaderboard structure
type MythicKeystoneLeaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Map struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"map"`
	Period               int   `json:"period"`
	PeriodStartTimestamp int64 `json:"period_start_timestamp"`
	PeriodEndTimestamp   int64 `json:"period_end_timestamp"`
	ConnectedRealm       struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	LeadingGroups []struct {
		Ranking            int   `json:"ranking"`
		Duration           int   `json:"duration"`
		CompletedTimestamp int64 `json:"completed_timestamp"`
		KeystoneLevel      int   `json:"keystone_level"`
		Members            []struct {
			Profile struct {
				Name  string `json:"name"`
				ID    int    `json:"id"`
				Realm struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					ID   int    `json:"id"`
					Slug string `json:"slug"`
				} `json:"realm"`
			} `json:"profile"`
			Faction struct {
				Type string `json:"type"`
			} `json:"faction"`
			Specialization struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"specialization"`
		} `json:"members"`
	} `json:"leading_groups"`
	KeystoneAffixes []struct {
		KeystoneAffix struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"keystone_affix"`
		StartingLevel int `json:"starting_level"`
	} `json:"keystone_affixes"`
	MapChallengeModeID int    `json:"map_challenge_mode_id"`
	Name               string `json:"name"`
}
