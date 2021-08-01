package wowp

// CharacterPvPBracketStatistics structure
type CharacterPvPBracketStatistics struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Bracket struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"bracket"`
	Rating int `json:"rating"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	Tier struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"tier"`
	SeasonMatchStatistics struct {
		Played int `json:"played"`
		Won    int `json:"won"`
		Lost   int `json:"lost"`
	} `json:"season_match_statistics"`
	WeeklyMatchStatistics struct {
		Played int `json:"played"`
		Won    int `json:"won"`
		Lost   int `json:"lost"`
	} `json:"weekly_match_statistics"`
}

// CharacterPvPSummary structure
type CharacterPvPSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Brackets []struct {
		Href string `json:"href"`
	} `json:"brackets"`
	HonorLevel       int `json:"honor_level"`
	PvpMapStatistics []struct {
		WorldMap struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"world_map"`
		MatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"match_statistics"`
	} `json:"pvp_map_statistics"`
	HonorableKills int `json:"honorable_kills"`
	Character      struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
}
