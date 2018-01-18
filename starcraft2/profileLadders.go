/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-17 19:53:35
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:53:59
 */

package starcraft2

import "encoding/json"

// ProfileLadders structure
type ProfileLadders struct {
	CurrentSeason []struct {
		Ladder []struct {
			LadderName       string `json:"ladderName"`
			LadderID         int    `json:"ladderId"`
			Division         int    `json:"division"`
			Rank             int    `json:"rank"`
			League           string `json:"league"`
			MatchMakingQueue string `json:"matchMakingQueue"`
			Wins             int    `json:"wins"`
			Losses           int    `json:"losses"`
			Showcase         bool   `json:"showcase"`
		} `json:"ladder"`
		Characters []struct {
			ID          int    `json:"id"`
			Realm       int    `json:"realm"`
			DisplayName string `json:"displayName"`
			ClanName    string `json:"clanName"`
			ClanTag     string `json:"clanTag"`
			ProfilePath string `json:"profilePath"`
		} `json:"characters"`
		NonRanked []interface{} `json:"nonRanked"`
	} `json:"currentSeason"`
	PreviousSeason []struct {
		Ladder []struct {
			LadderName       string `json:"ladderName"`
			LadderID         int    `json:"ladderId"`
			Division         int    `json:"division"`
			Rank             int    `json:"rank"`
			League           string `json:"league"`
			MatchMakingQueue string `json:"matchMakingQueue"`
			Wins             int    `json:"wins"`
			Losses           int    `json:"losses"`
			Showcase         bool   `json:"showcase"`
		} `json:"ladder"`
		Characters []struct {
			ID          int    `json:"id"`
			Realm       int    `json:"realm"`
			DisplayName string `json:"displayName"`
			ClanName    string `json:"clanName"`
			ClanTag     string `json:"clanTag"`
			ProfilePath string `json:"profilePath"`
		} `json:"characters"`
		NonRanked []interface{} `json:"nonRanked"`
	} `json:"previousSeason"`
	ShowcasePlacement []interface{} `json:"showcasePlacement"`
}

// JSON2Struct creates ProfileLadders structure from JSON byte array
func (p *ProfileLadders) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, p)
}
