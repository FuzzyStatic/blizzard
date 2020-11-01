package wowgd

// PetIndex structure
type PetIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Pets []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"pets"`
}

// Pet structure
type Pet struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BattlePetType struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"battle_pet_type"`
	Description    string `json:"description"`
	IsCapturable   bool   `json:"is_capturable"`
	IsTradable     bool   `json:"is_tradable"`
	IsBattlepet    bool   `json:"is_battlepet"`
	IsAllianceOnly bool   `json:"is_alliance_only"`
	IsHordeOnly    bool   `json:"is_horde_only"`
	Abilities      []struct {
		Ability struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"ability"`
		Slot          int `json:"slot"`
		RequiredLevel int `json:"required_level"`
	} `json:"abilities"`
	Source struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"source"`
	Icon     string `json:"icon"`
	Creature struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"creature"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	IsRandomCreatureDisplay bool `json:"is_random_creature_display"`
}

// PetAbilityIndex structure
type PetAbilityIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Abilities []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"abilities"`
}

// PetAbility structure
type PetAbility struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BattlePetType struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"battle_pet_type"`
	Rounds int `json:"rounds"`
	Media  struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

// PetMedia structure
type PetMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		FileDataID int    `json:"file_data_id"`
	} `json:"assets"`
	ID int `json:"id"`
}
