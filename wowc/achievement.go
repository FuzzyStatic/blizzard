package wowc

// Achievement structure
type Achievement struct {
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
		Stats                []interface{} `json:"stats"`
		Armor                int           `json:"armor"`
		Context              string        `json:"context"`
		BonusLists           []interface{} `json:"bonusLists"`
		ArtifactID           int           `json:"artifactId"`
		DisplayInfoID        int           `json:"displayInfoId"`
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
}
