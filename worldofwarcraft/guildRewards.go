/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:37:16
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:38:43
 */

package worldofwarcraft

import "encoding/json"

// GuildRewards structure
type GuildRewards struct {
	Rewards []struct {
		MinGuildLevel    int `json:"minGuildLevel"`
		MinGuildRepLevel int `json:"minGuildRepLevel"`
		Item             struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats                []interface{} `json:"stats"`
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
		Achievement struct {
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
					TimewalkerLevel int `json:"timewalkerLevel"`
				} `json:"tooltipParams"`
				Stats                []interface{} `json:"stats"`
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
		Races []int `json:"races,omitempty"`
	} `json:"rewards"`
}

// JSON2Struct creates GuildRewards structure from JSON byte array
func (g *GuildRewards) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, g)
}
