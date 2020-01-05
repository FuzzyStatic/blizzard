package wowgd

// PvPSeasonIndex structure
type PvPSeasonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"seasons"`
	CurrentSeason struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"current_season"`
}

// PvPSeason structure
type PvPSeason struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID           int `json:"id"`
	Leaderboards struct {
		Href string `json:"href"`
	} `json:"leaderboards"`
	Rewards struct {
		Href string `json:"href"`
	} `json:"rewards"`
	SeasonStartTimestamp int64 `json:"season_start_timestamp"`
	SeasonEndTimestamp   int64 `json:"season_end_timestamp"`
}

// PvPLeaderboardsIndex structure
type PvPLeaderboardsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	Leaderboards []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"leaderboards"`
}

// PvPLeaderboard structure
type PvPLeaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	Name    string `json:"name"`
	Bracket struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"bracket"`
	Entries []struct {
		Character struct {
			Name  string `json:"name"`
			ID    int    `json:"id"`
			Realm struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"realm"`
		} `json:"character"`
		Faction struct {
			Type string `json:"type"`
		} `json:"faction"`
		Rank                  int `json:"rank"`
		Rating                int `json:"rating"`
		SeasonMatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"season_match_statistics"`
		Tier struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"tier"`
	} `json:"entries"`
}

// PvPRewardsIndex structure
type PvPRewardsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	Rewards []struct {
		Bracket struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		} `json:"bracket"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		RatingCutoff int `json:"rating_cutoff"`
		Faction      struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"rewards"`
}
