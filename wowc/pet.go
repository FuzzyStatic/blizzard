package wowc

// PetMasterList structure
type PetMasterList struct {
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
