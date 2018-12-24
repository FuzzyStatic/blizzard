/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:45
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:01:39
 */

package diablo3

import "encoding/json"

// Season structure
type Season struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Leaderboard []struct {
		Ladder struct {
			Href string `json:"href"`
		} `json:"ladder"`
		TeamSize        int    `json:"team_size,omitempty"`
		Hardcore        bool   `json:"hardcore,omitempty"`
		HeroClassString string `json:"hero_class_string,omitempty"`
	} `json:"leaderboard"`
	SeasonID       int    `json:"season_id"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}

// JSON2Struct creates Season structure from JSON byte array
func (s *Season) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, s)
}
