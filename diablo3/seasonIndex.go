/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:40
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:01:44
 */

package diablo3

import "encoding/json"

// SeasonIndex structure
type SeasonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season []struct {
		Href string `json:"href"`
	} `json:"season"`
	CurrentSeason        int    `json:"current_season"`
	ServiceCurrentSeason int    `json:"service_current_season"`
	ServiceSeasonState   string `json:"service_season_state"`
	LastUpdateTime       string `json:"last_update_time"`
	GeneratedBy          string `json:"generated_by"`
}

// JSON2Struct creates SeasonIndex structure from JSON byte array
func (s *SeasonIndex) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, s)
}
