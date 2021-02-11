package wowgd

// ProfessionsIndex structure
type ProfessionsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Professions []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"professions"`
}

// Profession structure
type Profession struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	SkillTiers []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"skill_tiers"`
}

// ProfessionMedia structure
type ProfessionMedia struct {
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

// ProfessionSkillTier structure
type ProfessionSkillTier struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	MinimumSkillLevel int    `json:"minimum_skill_level"`
	MaximumSkillLevel int    `json:"maximum_skill_level"`
	Categories        []struct {
		Name    string `json:"name"`
		Recipes []struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"recipes"`
	} `json:"categories"`
}

// Recipe structure
type Recipe struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	CraftedItem struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"crafted_item"`
	Reagents []struct {
		Reagent struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"reagent"`
		Quantity int `json:"quantity"`
	} `json:"reagents"`
	CraftedQuantity struct {
		Minimum int `json:"minimum"`
		Maximum int `json:"maximum"`
	} `json:"crafted_quantity"`
}

// RecipeMedia structure
type RecipeMedia struct {
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
