package wowp

// CharacterSpecializationsSummary structure
type CharacterSpecializationsSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Specializations []struct {
		Specialization struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"specialization"`
		Talents []struct {
			Talent struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"talent"`
			SpellTooltip struct {
				Spell struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"spell"`
				Description string      `json:"description"`
				CastTime    string      `json:"cast_time"`
				PowerCost   interface{} `json:"power_cost"`
				Range       string      `json:"range"`
				Cooldown    string      `json:"cooldown"`
			} `json:"spell_tooltip,omitempty"`
		} `json:"talents"`
		Glyphs []struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"glyphs,omitempty"`
		PvpTalentSlots []struct {
			Selected struct {
				Talent struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"talent"`
				SpellTooltip struct {
					Spell struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"spell"`
					Description string `json:"description"`
					CastTime    string `json:"cast_time"`
					Cooldown    string `json:"cooldown"`
				} `json:"spell_tooltip"`
			} `json:"selected"`
			SlotNumber int `json:"slot_number"`
		} `json:"pvp_talent_slots,omitempty"`
	} `json:"specializations"`
	ActiveSpecialization struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_specialization"`
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
}
