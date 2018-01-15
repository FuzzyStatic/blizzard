/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 21:44:17
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 21:51:41
 */

package worldofwarcraft

import "encoding/json"

// Pet structure
type Pet struct {
	Pets []struct {
		CanBattle  bool   `json:"canBattle"`
		CreatureID int    `json:"creatureId"`
		Name       string `json:"name"`
		Family     string `json:"family"`
		Icon       string `json:"icon"`
		QualityID  int    `json:"qualityId"`
		Stats      struct {
			SpeciesID    int `json:"speciesId"`
			BreedID      int `json:"breedId"`
			PetQualityID int `json:"petQualityId"`
			Level        int `json:"level"`
			Health       int `json:"health"`
			Power        int `json:"power"`
			Speed        int `json:"speed"`
		} `json:"stats"`
		StrongAgainst []string `json:"strongAgainst"`
		TypeID        int      `json:"typeId"`
		WeakAgainst   []string `json:"weakAgainst"`
	} `json:"pets"`
}

// JSON2Struct creates Pet structure from JSON byte array
func (p *Pet) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, p)
}
