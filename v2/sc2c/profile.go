package sc2c

// StaticProfile structure
type StaticProfile struct {
	Achievements []struct {
		CategoryID          string   `json:"categoryId"`
		ChainAchievementIds []string `json:"chainAchievementIds"`
		ChainRewardSize     int      `json:"chainRewardSize"`
		CriteriaIds         []string `json:"criteriaIds,omitempty"`
		Description         string   `json:"description"`
		Flags               int      `json:"flags"`
		ID                  string   `json:"id"`
		ImageURL            string   `json:"imageUrl"`
		IsChained           bool     `json:"isChained"`
		Points              int      `json:"points"`
		Title               string   `json:"title"`
		UIOrderHint         int      `json:"uiOrderHint"`
	} `json:"achievements"`
	Criteria []struct {
		AchievementID     string `json:"achievementId"`
		Description       string `json:"description"`
		EvaluationClass   string `json:"evaluationClass"`
		Flags             int    `json:"flags"`
		ID                string `json:"id"`
		NecessaryQuantity int    `json:"necessaryQuantity"`
		UIOrderHint       int    `json:"uiOrderHint"`
	} `json:"criteria"`
	Categories []struct {
		ChildrenCategoryIds   []interface{} `json:"childrenCategoryIds"`
		FeaturedAchievementID string        `json:"featuredAchievementId"`
		ID                    string        `json:"id"`
		Name                  string        `json:"name"`
		ParentCategoryID      string        `json:"parentCategoryId"`
		Points                int           `json:"points"`
		UIOrderHint           int           `json:"uiOrderHint"`
		MedalTiers            []int         `json:"medalTiers,omitempty"`
	} `json:"categories"`
	Rewards []struct {
		Flags          int    `json:"flags"`
		ID             string `json:"id"`
		AchievementID  string `json:"achievementId,omitempty"`
		Name           string `json:"name"`
		ImageURL       string `json:"imageUrl"`
		UnlockableType string `json:"unlockableType"`
		IsSkin         bool   `json:"isSkin"`
		UIOrderHint    int    `json:"uiOrderHint"`
		Command        string `json:"command,omitempty"`
	} `json:"rewards"`
}

// MetadataProfile structure
type MetadataProfile struct {
	Name       string `json:"name"`
	ProfileURL string `json:"profileUrl"`
	AvatarURL  string `json:"avatarUrl"`
	ProfileID  string `json:"profileId"`
	RegionID   int    `json:"regionId"`
	RealmID    int    `json:"realmId"`
}

// Profile structure
type Profile struct {
	Summary struct {
		ID                     string `json:"id"`
		Realm                  int    `json:"realm"`
		DisplayName            string `json:"displayName"`
		Portrait               string `json:"portrait"`
		DecalTerran            string `json:"decalTerran"`
		DecalProtoss           string `json:"decalProtoss"`
		DecalZerg              string `json:"decalZerg"`
		TotalSwarmLevel        int    `json:"totalSwarmLevel"`
		TotalAchievementPoints int    `json:"totalAchievementPoints"`
	} `json:"summary"`
	Snapshot struct {
		SeasonSnapshot struct {
			OneV1 struct {
				Rank       int         `json:"rank"`
				LeagueName interface{} `json:"leagueName"`
				TotalGames int         `json:"totalGames"`
				TotalWins  int         `json:"totalWins"`
			} `json:"1v1"`
			TwoV2 struct {
				Rank       int         `json:"rank"`
				LeagueName interface{} `json:"leagueName"`
				TotalGames int         `json:"totalGames"`
				TotalWins  int         `json:"totalWins"`
			} `json:"2v2"`
			ThreeV3 struct {
				Rank       int         `json:"rank"`
				LeagueName interface{} `json:"leagueName"`
				TotalGames int         `json:"totalGames"`
				TotalWins  int         `json:"totalWins"`
			} `json:"3v3"`
			FourV4 struct {
				Rank       int         `json:"rank"`
				LeagueName interface{} `json:"leagueName"`
				TotalGames int         `json:"totalGames"`
				TotalWins  int         `json:"totalWins"`
			} `json:"4v4"`
			Archon struct {
				Rank       int         `json:"rank"`
				LeagueName interface{} `json:"leagueName"`
				TotalGames int         `json:"totalGames"`
				TotalWins  int         `json:"totalWins"`
			} `json:"Archon"`
		} `json:"seasonSnapshot"`
		TotalRankedSeasonGamesPlayed int `json:"totalRankedSeasonGamesPlayed"`
	} `json:"snapshot"`
	Career struct {
		TerranWins                int         `json:"terranWins"`
		ZergWins                  int         `json:"zergWins"`
		ProtossWins               int         `json:"protossWins"`
		TotalCareerGames          int         `json:"totalCareerGames"`
		TotalGamesThisSeason      int         `json:"totalGamesThisSeason"`
		Current1V1LeagueName      interface{} `json:"current1v1LeagueName"`
		CurrentBestTeamLeagueName interface{} `json:"currentBestTeamLeagueName"`
		Best1V1Finish             struct {
			LeagueName    string `json:"leagueName"`
			TimesAchieved int    `json:"timesAchieved"`
		} `json:"best1v1Finish"`
		BestTeamFinish struct {
			LeagueName    interface{} `json:"leagueName"`
			TimesAchieved int         `json:"timesAchieved"`
		} `json:"bestTeamFinish"`
	} `json:"career"`
	SwarmLevels struct {
		Level  int `json:"level"`
		Terran struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"terran"`
		Zerg struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"zerg"`
		Protoss struct {
			Level              int `json:"level"`
			MaxLevelPoints     int `json:"maxLevelPoints"`
			CurrentLevelPoints int `json:"currentLevelPoints"`
		} `json:"protoss"`
	} `json:"swarmLevels"`
	Campaign struct {
		DifficultyCompleted struct {
			WingsOfLiberty  string `json:"wings-of-liberty"`
			HeartOfTheSwarm string `json:"heart-of-the-swarm"`
		} `json:"difficultyCompleted"`
	} `json:"campaign"`
	CategoryPointProgress []struct {
		CategoryID   string `json:"categoryId"`
		PointsEarned int    `json:"pointsEarned"`
	} `json:"categoryPointProgress"`
	AchievementShowcase []string `json:"achievementShowcase"`
	EarnedRewards       []struct {
		RewardID      string `json:"rewardId"`
		Selected      bool   `json:"selected"`
		AchievementID string `json:"achievementId,omitempty"`
		Category      string `json:"category,omitempty"`
	} `json:"earnedRewards"`
	EarnedAchievements []struct {
		AchievementID                    string  `json:"achievementId"`
		CompletionDate                   float64 `json:"completionDate"`
		NumCompletedAchievementsInSeries int     `json:"numCompletedAchievementsInSeries"`
		TotalAchievementsInSeries        int     `json:"totalAchievementsInSeries"`
		IsComplete                       bool    `json:"isComplete"`
		InProgress                       bool    `json:"inProgress"`
		Criteria                         []struct {
			CriterionID string `json:"criterionId"`
		} `json:"criteria"`
		NextProgressEarnedQuantity   int `json:"nextProgressEarnedQuantity,omitempty"`
		NextProgressRequiredQuantity int `json:"nextProgressRequiredQuantity,omitempty"`
	} `json:"earnedAchievements"`
}

// LadderSummary structure
type LadderSummary struct {
	ShowCaseEntries      []interface{} `json:"showCaseEntries"`
	PlacementMatches     []interface{} `json:"placementMatches"`
	AllLadderMemberships []interface{} `json:"allLadderMemberships"`
}

// Ladder structure
type Ladder struct {
	LadderTeams          []interface{} `json:"ladderTeams"`
	AllLadderMemberships []interface{} `json:"allLadderMemberships"`
	RanksAndPools        []interface{} `json:"ranksAndPools"`
}
