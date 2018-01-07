/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:50
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:39:50
 */

package diablo3

import (
	"encoding/json"
)

// Artisan structure
type Artisan struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Portrait string `json:"portrait"`
	Training struct {
		Tiers []struct {
			Tier   int `json:"tier"`
			Levels []struct {
				Tier           int `json:"tier"`
				TierLevel      int `json:"tierLevel"`
				Percent        int `json:"percent"`
				TrainedRecipes []struct {
					ID       string `json:"id"`
					Slug     string `json:"slug"`
					Name     string `json:"name"`
					Cost     int    `json:"cost"`
					Reagents []struct {
						Quantity int `json:"quantity"`
						Item     struct {
							ID            string `json:"id"`
							Name          string `json:"name"`
							Icon          string `json:"icon"`
							DisplayColor  string `json:"displayColor"`
							TooltipParams string `json:"tooltipParams"`
						} `json:"item"`
					} `json:"reagents"`
					ItemProduced struct {
						ID            string `json:"id"`
						Name          string `json:"name"`
						Icon          string `json:"icon"`
						DisplayColor  string `json:"displayColor"`
						TooltipParams string `json:"tooltipParams"`
					} `json:"itemProduced"`
				} `json:"trainedRecipes"`
				TaughtRecipes []struct {
					ID       string `json:"id"`
					Slug     string `json:"slug"`
					Name     string `json:"name"`
					Cost     int    `json:"cost"`
					Reagents []struct {
						Quantity int `json:"quantity"`
						Item     struct {
							ID            string `json:"id"`
							Name          string `json:"name"`
							Icon          string `json:"icon"`
							DisplayColor  string `json:"displayColor"`
							TooltipParams string `json:"tooltipParams"`
						} `json:"item"`
					} `json:"reagents"`
					ItemProduced struct {
						ID            string `json:"id"`
						Name          string `json:"name"`
						Icon          string `json:"icon"`
						DisplayColor  string `json:"displayColor"`
						TooltipParams string `json:"tooltipParams"`
					} `json:"itemProduced"`
				} `json:"taughtRecipes"`
				UpgradeCost int `json:"upgradeCost"`
			} `json:"levels"`
		} `json:"tiers"`
	} `json:"training"`
}

// JSON2Struct creates Artisan structure from JSON byte array
func (a *Artisan) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, a)
}
