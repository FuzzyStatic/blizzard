/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:37:16
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:40:56
 */

package worldofwarcraft

import "encoding/json"

// GuildAchievements structure
type GuildAchievements struct {
	Achievements []struct {
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
	} `json:"achievements"`
}

// JSON2Struct creates GuildAchievements structure from JSON byte array
func (g *GuildAchievements) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, g)
}
