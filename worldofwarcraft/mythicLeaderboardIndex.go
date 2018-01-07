/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 13:29:53
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 16:17:27
 */

package worldofwarcraft

import "encoding/json"

// MythicLeaderboardIndex structure
type MythicLeaderboardIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CurrentLeaderboards []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"current_leaderboards"`
}

// JSON2Struct creates MythicLeaderboardIndex structure from JSON byte array
func (m *MythicLeaderboardIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, m)
}
