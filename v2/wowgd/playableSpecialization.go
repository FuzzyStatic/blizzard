package wowgd

// PlayableSpecializationIndex structure
type PlayableSpecializationIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	CharacterSpecializations []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_specializations"`
	PetSpecializations []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"pet_specializations"`
}

// PlayableSpecialization structure
type PlayableSpecialization struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID            int `json:"id"`
	PlayableClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"playable_class"`
	Name              string `json:"name"`
	GenderDescription struct {
		Male   string `json:"male"`
		Female string `json:"female"`
	} `json:"gender_description"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	Role struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"role"`
	TalentTiers []struct {
		Level   int `json:"level"`
		Talents []struct {
			Talent struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"talent"`
			SpellTooltip struct {
				Description string `json:"description"`
				CastTime    string `json:"cast_time"`
			} `json:"spell_tooltip"`
		} `json:"talents"`
	} `json:"talent_tiers"`
	PvpTalents []struct {
		Talent struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"talent"`
		SpellTooltip struct {
			Description string `json:"description"`
			CastTime    string `json:"cast_time"`
		} `json:"spell_tooltip"`
	} `json:"pvp_talents"`
}

// PlayableSpecializationMedia structure
type PlayableSpecializationMedia struct {
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
