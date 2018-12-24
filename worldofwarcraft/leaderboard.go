/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 10:40:32
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 11:26:28
 */

package worldofwarcraft

import "encoding/json"

// Leaderboard structure
type Leaderboard struct {
	Rows []struct {
		Ranking      int    `json:"ranking"`
		Rating       int    `json:"rating"`
		Name         string `json:"name"`
		RealmID      int    `json:"realmId"`
		RealmName    string `json:"realmName"`
		RealmSlug    string `json:"realmSlug"`
		RaceID       int    `json:"raceId"`
		ClassID      int    `json:"classId"`
		SpecID       int    `json:"specId"`
		FactionID    int    `json:"factionId"`
		GenderID     int    `json:"genderId"`
		SeasonWins   int    `json:"seasonWins"`
		SeasonLosses int    `json:"seasonLosses"`
		WeeklyWins   int    `json:"weeklyWins"`
		WeeklyLosses int    `json:"weeklyLosses"`
	} `json:"rows"`
}

// JSON2Struct creates Leaderboard structure from JSON byte array
func (l *Leaderboard) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, l)
}
