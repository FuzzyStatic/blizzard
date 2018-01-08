/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 21:45:21
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 21:46:57
 */

package worldofwarcraft

import "encoding/json"

// BossIndex structure
type BossIndex struct {
	Bosses []struct {
		ID                    int    `json:"id"`
		Name                  string `json:"name"`
		URLSlug               string `json:"urlSlug"`
		Description           string `json:"description,omitempty"`
		ZoneID                int    `json:"zoneId"`
		AvailableInNormalMode bool   `json:"availableInNormalMode"`
		AvailableInHeroicMode bool   `json:"availableInHeroicMode"`
		Health                int    `json:"health"`
		HeroicHealth          int    `json:"heroicHealth"`
		Level                 int    `json:"level"`
		HeroicLevel           int    `json:"heroicLevel"`
		JournalID             int    `json:"journalId,omitempty"`
		Npcs                  []struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			URLSlug           string `json:"urlSlug"`
			CreatureDisplayID int    `json:"creatureDisplayId"`
		} `json:"npcs"`
		Location struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"location,omitempty"`
	} `json:"bosses"`
}

// JSON2Struct creates BossIndex structure from JSON byte array
func (boss *BossIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, boss)
}
