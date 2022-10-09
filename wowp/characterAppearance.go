package wowp

// CharacterAppearanceSummary structure
type CharacterAppearanceSummary struct {
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
	PlayableRace struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"playable_race"`
	PlayableClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"playable_class"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_spec"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	GuildCrest struct {
		Emblem struct {
			ID    int `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float32 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"emblem"`
		Border struct {
			ID    int `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float64 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"border"`
		Background struct {
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float64 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"background"`
	} `json:"guild_crest"`
	Appearance struct {
		FaceVariation        int   `json:"face_variation"`
		SkinColor            int   `json:"skin_color"`
		HairVariation        int   `json:"hair_variation"`
		HairColor            int   `json:"hair_color"`
		FeatureVariation     int   `json:"feature_variation"`
		CustomDisplayOptions []int `json:"custom_display_options"`
	} `json:"appearance"`
	Items []struct {
		ID   int `json:"id"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Enchant                  int `json:"enchant"`
		ItemAppearanceModifierID int `json:"item_appearance_modifier_id"`
		InternalSlotID           int `json:"internal_slot_id"`
		Subclass                 int `json:"subclass"`
	} `json:"items"`
}
