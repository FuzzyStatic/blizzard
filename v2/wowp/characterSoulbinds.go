package wowp

// CharacterSoulbinds structure
type CharacterSoulbinds struct {
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
	ChosenCovenant struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"chosen_covenant"`
	RenownLevel int `json:"renown_level"`
	Soulbinds   []struct {
		Soulbind struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"soulbind"`
		Traits []struct {
			Trait struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"trait,omitempty"`
			Tier          int `json:"tier"`
			DisplayOrder  int `json:"display_order"`
			ConduitSocket struct {
				Type struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"type"`
				Socket struct {
					Conduit struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"conduit"`
					Rank int `json:"rank"`
				} `json:"socket"`
			} `json:"conduit_socket,omitempty"`
		} `json:"traits"`
		IsActive bool `json:"is_active,omitempty"`
	} `json:"soulbinds"`
}
