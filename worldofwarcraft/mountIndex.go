/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 21:13:58
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 21:16:36
 */

package worldofwarcraft

import "encoding/json"

// MountIndex structure
type MountIndex struct {
	Mounts []struct {
		Name       string `json:"name"`
		SpellID    int    `json:"spellId"`
		CreatureID int    `json:"creatureId"`
		ItemID     int    `json:"itemId"`
		QualityID  int    `json:"qualityId"`
		Icon       string `json:"icon"`
		IsGround   bool   `json:"isGround"`
		IsFlying   bool   `json:"isFlying"`
		IsAquatic  bool   `json:"isAquatic"`
		IsJumping  bool   `json:"isJumping"`
	} `json:"mounts"`
}

// JSON2Struct creates MountIndex structure from JSON byte array
func (m *MountIndex) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, m)
}
