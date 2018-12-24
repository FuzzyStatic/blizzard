/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:18
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:00:50
 */

package diablo3

import "encoding/json"

// Era structure
type Era struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Leaderboard []struct {
		TeamSize int `json:"team_size"`
		Ladder   struct {
			Href string `json:"href"`
		} `json:"ladder"`
		Hardcore        bool   `json:"hardcore,omitempty"`
		HeroClassString string `json:"hero_class_string,omitempty"`
	} `json:"leaderboard"`
	EraID          int    `json:"era_id"`
	EraStartDate   int    `json:"era_start_date"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}

// JSON2Struct creates Era structure from JSON byte array
func (e *Era) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, e)
}
