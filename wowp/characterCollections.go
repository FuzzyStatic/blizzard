package wowp

// CharacterCollectionsIndex structure
type CharacterCollectionsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets struct {
		Href string `json:"href"`
	} `json:"pets"`
	Mounts struct {
		Href string `json:"href"`
	} `json:"mounts"`
}

// CharacterMountsCollectionSummary structure
type CharacterMountsCollectionSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Mounts []struct {
		Mount struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"mount"`
		IsFavorite bool `json:"is_favorite,omitempty"`
	} `json:"mounts"`
}

// CharacterPetsCollectionSummary structure
type CharacterPetsCollectionSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets []struct {
		Species struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"species"`
		Level   int `json:"level"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Stats struct {
			BreedID int `json:"breed_id"`
			Health  int `json:"health"`
			Power   int `json:"power"`
			Speed   int `json:"speed"`
		} `json:"stats"`
		IsFavorite bool `json:"is_favorite,omitempty"`
	} `json:"pets"`
	UnlockedBattlePetSlots int `json:"unlocked_battle_pet_slots"`
}
