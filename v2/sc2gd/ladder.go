package sc2gd

type Ladder struct {
	Team []struct {
		ID                  uint64 `json:"id"`
		Rating              int    `json:"rating"`
		Wins                int    `json:"wins"`
		Losses              int    `json:"losses"`
		Ties                int    `json:"ties"`
		Points              int    `json:"points"`
		LongestWinStreak    int    `json:"longest_win_streak"`
		CurrentWinStreak    int    `json:"current_win_streak"`
		CurrentRank         int    `json:"current_rank"`
		HighestRank         int    `json:"highest_rank"`
		PreviousRank        int    `json:"previous_rank"`
		JoinTimeStamp       int    `json:"join_time_stamp"`
		LastPlayedTimeStamp int    `json:"last_played_time_stamp"`
		Member              []struct {
			LegacyLink struct {
				ID    int    `json:"id"`
				Realm int    `json:"realm"`
				Name  string `json:"name"`
				Path  string `json:"path"`
			} `json:"legacy_link"`
			PlayedRaceCount []struct {
				Race  string `json:"race"`
				Count int    `json:"count"`
			} `json:"played_race_count"`
			CharacterLink struct {
				ID        int    `json:"id"`
				BattleTag string `json:"battle_tag"`
				Key       struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"character_link"`
			ClanLink struct {
				ID       int    `json:"id"`
				ClanTag  string `json:"clan_tag"`
				ClanName string `json:"clan_name"`
				IconURL  string `json:"icon_url"`
				DecalURL string `json:"decal_url"`
			} `json:"clan_link"`
		} `json:"member"`
	} `json:"team"`
}
