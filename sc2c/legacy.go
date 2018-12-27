package sc2c

// LegacyProfile structure
type LegacyProfile struct {
	ID          string `json:"id"`
	Realm       int    `json:"realm"`
	DisplayName string `json:"displayName"`
	ClanName    string `json:"clanName"`
	ClanTag     string `json:"clanTag"`
	ProfilePath string `json:"profilePath"`
	Portrait    struct {
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
		Offset int    `json:"offset"`
		URL    string `json:"url"`
	} `json:"portrait"`
	Career struct {
		PrimaryRace      string `json:"primaryRace"`
		TerranWins       int    `json:"terranWins"`
		ProtossWins      int    `json:"protossWins"`
		ZergWins         int    `json:"zergWins"`
		Highest1V1Rank   string `json:"highest1v1Rank"`
		HighestTeamRank  string `json:"highestTeamRank"`
		SeasonTotalGames int    `json:"seasonTotalGames"`
		CareerTotalGames int    `json:"careerTotalGames"`
	} `json:"career"`
	SwarmLevels struct {
		Level  int `json:"level"`
		Terran struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"terran"`
		Zerg struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"zerg"`
		Protoss struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"protoss"`
	} `json:"swarmLevels"`
	Campaign struct {
		Wol  string `json:"wol"`
		Hots string `json:"hots"`
	} `json:"campaign"`
	Season struct {
		SeasonID             int `json:"seasonId"`
		SeasonNumber         int `json:"seasonNumber"`
		SeasonYear           int `json:"seasonYear"`
		TotalGamesThisSeason int `json:"totalGamesThisSeason"`
		Stats                []struct {
			Type  string `json:"type"`
			Wins  int    `json:"wins"`
			Games int    `json:"games"`
		} `json:"stats"`
	} `json:"season"`
	Rewards struct {
		Selected []string `json:"selected"`
		Earned   []string `json:"earned"`
	} `json:"rewards"`
	Achievements struct {
		Points struct {
			TotalPoints    int `json:"totalPoints"`
			CategoryPoints struct {
				Num4325377 int `json:"4325377"`
				Num4325379 int `json:"4325379"`
				Num4325382 int `json:"4325382"`
				Num4325408 int `json:"4325408"`
				Num4325410 int `json:"4325410"`
				Num4330138 int `json:"4330138"`
				Num4364473 int `json:"4364473"`
				Num4386911 int `json:"4386911"`
			} `json:"categoryPoints"`
		} `json:"points"`
		Achievements []struct {
			AchievementID  string `json:"achievementId"`
			CompletionDate int    `json:"completionDate"`
		} `json:"achievements"`
	} `json:"achievements"`
}

// LegacyProfileLadders structure
type LegacyProfileLadders struct {
	CurrentSeason     []interface{} `json:"currentSeason"`
	PreviousSeason    []interface{} `json:"previousSeason"`
	ShowcasePlacement []interface{} `json:"showcasePlacement"`
}

// LegacyProfileMatches structure
type LegacyProfileMatches struct {
	Matches []struct {
		Map      string `json:"map"`
		Type     string `json:"type"`
		Decision string `json:"decision"`
		Speed    string `json:"speed"`
		Date     int    `json:"date"`
	} `json:"matches"`
}

// LegacyLadder structure
type LegacyLadder struct {
	LadderMembers []struct {
		Character struct {
			ID          string `json:"id"`
			Realm       int    `json:"realm"`
			Region      int    `json:"region"`
			DisplayName string `json:"displayName"`
			ClanName    string `json:"clanName"`
			ClanTag     string `json:"clanTag"`
			ProfilePath string `json:"profilePath"`
		} `json:"character"`
		JoinTimestamp  int    `json:"joinTimestamp"`
		Points         int    `json:"points"`
		Wins           int    `json:"wins"`
		Losses         int    `json:"losses"`
		HighestRank    int    `json:"highestRank"`
		PreviousRank   int    `json:"previousRank"`
		FavoriteRaceP1 string `json:"favoriteRaceP1"`
	} `json:"ladderMembers"`
}

// LegacyAchievements structure
type LegacyAchievements struct {
	Achievements []struct {
		Title         string `json:"title"`
		Description   string `json:"description"`
		AchievementID string `json:"achievementId"`
		CategoryID    string `json:"categoryId"`
		Points        int    `json:"points"`
		Icon          struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
	} `json:"achievements"`
	Categories []struct {
		Title                 string `json:"title"`
		CategoryID            string `json:"categoryId"`
		FeaturedAchievementID string `json:"featuredAchievementId"`
		Children              []struct {
			Title                 string `json:"title"`
			CategoryID            string `json:"categoryId"`
			FeaturedAchievementID string `json:"featuredAchievementId"`
		} `json:"children,omitempty"`
	} `json:"categories"`
}

// LegacyRewards structure
type LegacyRewards struct {
	Portraits []struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID string `json:"achievementId"`
	} `json:"portraits"`
	TerranDecals []struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID string `json:"achievementId"`
	} `json:"terranDecals"`
	ZergDecals []struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID string `json:"achievementId"`
	} `json:"zergDecals"`
	ProtossDecals []struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID string `json:"achievementId"`
	} `json:"protossDecals"`
	Skins []struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		Name          string `json:"name,omitempty"`
		AchievementID string `json:"achievementId"`
	} `json:"skins"`
	Animations []struct {
		Title   string `json:"title"`
		Command string `json:"command"`
		ID      string `json:"id"`
		Icon    struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID string `json:"achievementId"`
	} `json:"animations"`
}
