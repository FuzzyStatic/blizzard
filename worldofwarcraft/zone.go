/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:44:06
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:51:50
 */

package worldofwarcraft

import "encoding/json"

// Zone structure
type Zone struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URLSlug     string `json:"urlSlug"`
	Description string `json:"description"`
	Location    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	ExpansionID           int      `json:"expansionId"`
	NumPlayers            string   `json:"numPlayers"`
	IsDungeon             bool     `json:"isDungeon"`
	IsRaid                bool     `json:"isRaid"`
	AdvisedMinLevel       int      `json:"advisedMinLevel"`
	AdvisedMaxLevel       int      `json:"advisedMaxLevel"`
	AdvisedHeroicMinLevel int      `json:"advisedHeroicMinLevel"`
	AdvisedHeroicMaxLevel int      `json:"advisedHeroicMaxLevel"`
	AvailableModes        []string `json:"availableModes"`
	LfgNormalMinGearLevel int      `json:"lfgNormalMinGearLevel"`
	LfgHeroicMinGearLevel int      `json:"lfgHeroicMinGearLevel"`
	Floors                int      `json:"floors"`
	Bosses                []struct {
		ID                    int    `json:"id"`
		Name                  string `json:"name"`
		URLSlug               string `json:"urlSlug"`
		Description           string `json:"description"`
		ZoneID                int    `json:"zoneId"`
		AvailableInNormalMode bool   `json:"availableInNormalMode"`
		AvailableInHeroicMode bool   `json:"availableInHeroicMode"`
		Health                int    `json:"health"`
		HeroicHealth          int    `json:"heroicHealth"`
		Level                 int    `json:"level"`
		HeroicLevel           int    `json:"heroicLevel"`
		JournalID             int    `json:"journalId"`
		Npcs                  []struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			URLSlug           string `json:"urlSlug"`
			CreatureDisplayID int    `json:"creatureDisplayId"`
		} `json:"npcs"`
	} `json:"bosses"`
}

// JSON2Struct creates Zone structure from JSON byte array
func (z *Zone) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, z)
}
