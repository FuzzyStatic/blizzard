/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-17 19:28:37
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:50:45
 */

package starcraft2

import "encoding/json"

// Profile structure
type Profile struct {
	ID          int    `json:"id"`
	Realm       int    `json:"realm"`
	DisplayName string `json:"displayName"`
	ClanName    string `json:"clanName"`
	ClanTag     string `json:"clanTag"`
	ProfilePath string `json:"profilePath"`
	Portrait    struct {
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
		Offset int    `json:"offset"`
		URL    string `json:"url"`
	} `json:"portrait"`
	Career struct {
		PrimaryRace      string `json:"primaryRace"`
		TerranWins       int    `json:"terranWins"`
		ProtossWins      int    `json:"protossWins"`
		ZergWins         int    `json:"zergWins"`
		HighestTeamRank  string `json:"highestTeamRank"`
		SeasonTotalGames int    `json:"seasonTotalGames"`
		CareerTotalGames int    `json:"careerTotalGames"`
	} `json:"career"`
	SwarmLevels struct {
		Level  int `json:"level"`
		Terran struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"terran"`
		Zerg struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"zerg"`
		Protoss struct {
			Level          int `json:"level"`
			TotalLevelXP   int `json:"totalLevelXP"`
			CurrentLevelXP int `json:"currentLevelXP"`
		} `json:"protoss"`
	} `json:"swarmLevels"`
	Campaign struct {
		Wol string `json:"wol"`
	} `json:"campaign"`
	Season struct {
		SeasonID             int `json:"seasonId"`
		SeasonNumber         int `json:"seasonNumber"`
		SeasonYear           int `json:"seasonYear"`
		TotalGamesThisSeason int `json:"totalGamesThisSeason"`
	} `json:"season"`
	Rewards struct {
		Selected []int64 `json:"selected"`
		Earned   []int64 `json:"earned"`
	} `json:"rewards"`
	Achievements struct {
		Points struct {
			TotalPoints    int `json:"totalPoints"`
			CategoryPoints struct {
				Num4325377 int `json:"4325377"`
				Num4325379 int `json:"4325379"`
				Num4325410 int `json:"4325410"`
				Num4330138 int `json:"4330138"`
				Num4364473 int `json:"4364473"`
				Num4386911 int `json:"4386911"`
			} `json:"categoryPoints"`
		} `json:"points"`
		Achievements []struct {
			AchievementID  int64 `json:"achievementId"`
			CompletionDate int   `json:"completionDate"`
		} `json:"achievements"`
	} `json:"achievements"`
}

// JSON2Struct creates Profile structure from JSON byte array
func (p *Profile) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, p)
}
