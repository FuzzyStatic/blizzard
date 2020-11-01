package wowp

// CharacterHunterPetsSummary structure
type CharacterHunterPetsSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	HunterPets []struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		Creature struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"creature"`
		Slot            int  `json:"slot"`
		IsActive        bool `json:"is_active"`
		CreatureDisplay struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"creature_display"`
	} `json:"hunter_pets"`
}
