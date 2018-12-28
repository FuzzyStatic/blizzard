package wowc

// RegionBattlegroups structure
type RegionBattlegroups struct {
	Battlegroups []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"battlegroups"`
}

// CharacterRaces structure
type CharacterRaces struct {
	Races []struct {
		ID   int    `json:"id"`
		Mask int    `json:"mask"`
		Side string `json:"side"`
		Name string `json:"name"`
	} `json:"races"`
}

// CharacterClasses structure
type CharacterClasses struct {
	Classes []struct {
		ID        int    `json:"id"`
		Mask      int    `json:"mask"`
		PowerType string `json:"powerType"`
		Name      string `json:"name"`
	} `json:"classes"`
}

// CharacterAchievements strudture
type CharacterAchievements struct {
	Achievements []Achievements `json:"achievements"`
}

// GuildRewards structure
type GuildRewards struct {
	Rewards []struct {
		MinGuildLevel    int `json:"minGuildLevel"`
		MinGuildRepLevel int `json:"minGuildRepLevel"`
		Achievement      struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Points      int    `json:"points"`
			Description string `json:"description"`
			Reward      string `json:"reward"`
			RewardItems []struct {
				ID            int    `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				Quality       int    `json:"quality"`
				ItemLevel     int    `json:"itemLevel"`
				TooltipParams struct {
					TimewalkerLevel   int `json:"timewalkerLevel"`
					AzeritePower0     int `json:"azeritePower0"`
					AzeritePower1     int `json:"azeritePower1"`
					AzeritePower2     int `json:"azeritePower2"`
					AzeritePower3     int `json:"azeritePower3"`
					AzeritePowerLevel int `json:"azeritePowerLevel"`
					AzeritePower4     int `json:"azeritePower4"`
				} `json:"tooltipParams"`
				Stats []struct {
					Stat   int `json:"stat"`
					Amount int `json:"amount"`
				} `json:"stats"`
				Armor                int           `json:"armor"`
				Context              string        `json:"context"`
				BonusLists           []interface{} `json:"bonusLists"`
				DisplayInfoID        int           `json:"displayInfoId"`
				ArtifactID           int           `json:"artifactId"`
				ArtifactAppearanceID int           `json:"artifactAppearanceId"`
				ArtifactTraits       []interface{} `json:"artifactTraits"`
				Relics               []interface{} `json:"relics"`
				Appearance           struct {
				} `json:"appearance"`
			} `json:"rewardItems"`
			Icon     string `json:"icon"`
			Criteria []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool `json:"accountWide"`
			FactionID   int  `json:"factionId"`
		} `json:"achievement,omitempty"`
		Item struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel   int `json:"timewalkerLevel"`
				AzeritePower0     int `json:"azeritePower0"`
				AzeritePower1     int `json:"azeritePower1"`
				AzeritePower2     int `json:"azeritePower2"`
				AzeritePower3     int `json:"azeritePower3"`
				AzeritePowerLevel int `json:"azeritePowerLevel"`
				AzeritePower4     int `json:"azeritePower4"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"item"`
		Races []int `json:"races,omitempty"`
	} `json:"rewards"`
}

// GuildPerks structure
type GuildPerks struct {
	Perks []struct {
		GuildLevel int `json:"guildLevel"`
		Spell      struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Subtext     string `json:"subtext"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			CastTime    string `json:"castTime"`
		} `json:"spell"`
	} `json:"perks"`
}

// GuildAchievements structure
type GuildAchievements struct {
	Achievements []Achievements `json:"achievements"`
}

// Achievements structure
type Achievements struct {
	ID           int `json:"id"`
	Achievements []struct {
		ID          int           `json:"id"`
		Title       string        `json:"title"`
		Points      int           `json:"points"`
		Description string        `json:"description"`
		RewardItems []interface{} `json:"rewardItems"`
		Icon        string        `json:"icon"`
		Criteria    []struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			OrderIndex  int    `json:"orderIndex"`
			Max         int    `json:"max"`
		} `json:"criteria"`
		AccountWide bool   `json:"accountWide"`
		FactionID   int    `json:"factionId"`
		Reward      string `json:"reward,omitempty"`
	} `json:"achievements"`
	Name       string `json:"name"`
	Categories []struct {
		ID           int `json:"id"`
		Achievements []struct {
			ID          int           `json:"id"`
			Title       string        `json:"title"`
			Points      int           `json:"points"`
			Description string        `json:"description"`
			RewardItems []interface{} `json:"rewardItems"`
			Icon        string        `json:"icon"`
			Criteria    []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool `json:"accountWide"`
			FactionID   int  `json:"factionId"`
		} `json:"achievements"`
		Name string `json:"name"`
	} `json:"categories,omitempty"`
}
