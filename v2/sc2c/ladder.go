package sc2c

// LadderGrandmaster structure
type LadderGrandmaster struct {
	LadderTeams []struct {
		TeamMembers []struct {
			ID           string `json:"id"`
			Realm        int    `json:"realm"`
			Region       int    `json:"region"`
			DisplayName  string `json:"displayName"`
			ClanTag      string `json:"clanTag"`
			FavoriteRace string `json:"favoriteRace"`
		} `json:"teamMembers"`
		PreviousRank  int `json:"previousRank"`
		Points        int `json:"points"`
		Wins          int `json:"wins"`
		Losses        int `json:"losses"`
		Mmr           int `json:"mmr"`
		JoinTimestamp int `json:"joinTimestamp"`
	} `json:"ladderTeams"`
}

// LadderSeason structure
type LadderSeason struct {
	SeasonID  int    `json:"seasonId"`
	Number    int    `json:"number"`
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
