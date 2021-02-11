package wowp

// CharacterProfessionsSummary
type CharacterProfessionsSummary struct {
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
	Primaries []struct {
		Profession struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"profession"`
		Tiers []struct {
			SkillPoints    int `json:"skill_points"`
			MaxSkillPoints int `json:"max_skill_points"`
			Tier           struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"tier"`
			KnownRecipes []struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"known_recipes"`
		} `json:"tiers"`
	} `json:"primaries"`
	Secondaries []struct {
		Profession struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"profession"`
		Tiers []struct {
			SkillPoints    int `json:"skill_points"`
			MaxSkillPoints int `json:"max_skill_points"`
			Tier           struct {
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"tier"`
			KnownRecipes []struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"known_recipes"`
		} `json:"tiers"`
	} `json:"secondaries"`
}
