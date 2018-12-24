/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 21:48:26
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 22:03:57
 */

package worldofwarcraft

import "encoding/json"

// PetSpecies structure
type PetSpecies struct {
	SpeciesID   int    `json:"speciesId"`
	PetTypeID   int    `json:"petTypeId"`
	CreatureID  int    `json:"creatureId"`
	Name        string `json:"name"`
	CanBattle   bool   `json:"canBattle"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Abilities   []struct {
		Slot          int    `json:"slot"`
		Order         int    `json:"order"`
		RequiredLevel int    `json:"requiredLevel"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		Cooldown      int    `json:"cooldown"`
		Rounds        int    `json:"rounds"`
		PetTypeID     int    `json:"petTypeId"`
		IsPassive     bool   `json:"isPassive"`
		HideHints     bool   `json:"hideHints"`
	} `json:"abilities"`
}

// JSON2Struct creates PetSpecies structure from JSON byte array
func (p *PetSpecies) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, p)
}
