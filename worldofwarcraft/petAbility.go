/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 21:48:15
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 21:57:54
 */

package worldofwarcraft

import "encoding/json"

// PetAbility structure
type PetAbility struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Cooldown  int    `json:"cooldown"`
	Rounds    int    `json:"rounds"`
	PetTypeID int    `json:"petTypeId"`
	IsPassive bool   `json:"isPassive"`
	HideHints bool   `json:"hideHints"`
}

// JSON2Struct creates PetAbility structure from JSON byte array
func (p *PetAbility) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, p)
}
