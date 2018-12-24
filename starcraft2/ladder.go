/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-18 16:35:58
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 16:36:48
 */

package starcraft2

import "encoding/json"

// Ladder structure
type Ladder struct {
	LadderMembers []struct {
		Character struct {
			ID          int    `json:"id"`
			Realm       int    `json:"realm"`
			DisplayName string `json:"displayName"`
			ClanName    string `json:"clanName"`
			ClanTag     string `json:"clanTag"`
			ProfilePath string `json:"profilePath"`
		} `json:"character"`
		JoinTimestamp  int     `json:"joinTimestamp"`
		Points         float64 `json:"points"`
		Wins           int     `json:"wins"`
		Losses         int     `json:"losses"`
		HighestRank    int     `json:"highestRank"`
		PreviousRank   int     `json:"previousRank"`
		FavoriteRaceP1 string  `json:"favoriteRaceP1"`
	} `json:"ladderMembers"`
}

// JSON2Struct creates Ladder structure from JSON byte array
func (l *Ladder) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, l)
}
