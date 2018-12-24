/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:59:14
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 19:03:01
 */

package worldofwarcraft

import "encoding/json"

// PetTypes structure
type PetTypes struct {
	PetTypes []struct {
		ID              int    `json:"id"`
		Key             string `json:"key"`
		Name            string `json:"name"`
		TypeAbilityID   int    `json:"typeAbilityId"`
		StrongAgainstID int    `json:"strongAgainstId"`
		WeakAgainstID   int    `json:"weakAgainstId"`
	} `json:"petTypes"`
}

// JSON2Struct creates PetTypes structure from JSON byte array
func (p *PetTypes) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, p)
}
