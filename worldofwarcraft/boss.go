/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 21:53:29
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 21:53:50
 */

package worldofwarcraft

import "encoding/json"

// Boss structure
type Boss struct {
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
}

// JSON2Struct creates Boss structure from JSON byte array
func (boss *Boss) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, boss)
}
