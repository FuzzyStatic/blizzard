/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 21:48:31
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 22:09:26
 */

package worldofwarcraft

import "encoding/json"

// PetStats structure
type PetStats struct {
	SpeciesID    int `json:"speciesId"`
	BreedID      int `json:"breedId"`
	PetQualityID int `json:"petQualityId"`
	Level        int `json:"level"`
	Health       int `json:"health"`
	Power        int `json:"power"`
	Speed        int `json:"speed"`
}

// JSON2Struct creates PetStats structure from JSON byte array
func (p *PetStats) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, p)
}
