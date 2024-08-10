package wowp

// CharacterCollectionsIndex structure
type CharacterCollectionsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Heirlooms struct {
		Href string `json:"href"`
	} `json:"heirlooms"`
	Mounts struct {
		Href string `json:"href"`
	} `json:"mounts"`
	Pets struct {
		Href string `json:"href"`
	} `json:"pets"`
	Toys struct {
		Href string `json:"href"`
	} `json:"toys"`
	Transmog struct {
		Href string `json:"href"`
	} `json:"transmog"`
}

// CharacterHeirloomsCollectionSummary structure
type CharacterHeirloomsCollectionSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Heirlooms []struct {
		Heirloom struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"heirloom"`
		Upgrade struct {
			Level int `json:"level"`
		} `json:"upgrade"`
	} `json:"Heirlooms"`
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

// CharacterToysCollectionSummary structure
type CharacterToysCollectionSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Toys []struct {
		Toy struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"toy"`
	} `json:"toys"`
}

// CharacterTransmogCollectionSummary structure
type CharacterTransmogsCollectionSummary struct {
	AppearanceSets []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"appearance_sets"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Slots []CharacterTransmogsCollectionSummarySlots `json:"slots"`
}

type CharacterTransmogsCollectionSummarySlots struct {
	Slot struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"slot"`
	Appearances []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"appearances"`
}
