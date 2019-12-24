package wowp

// CharacterTitlesSummary structure
type CharacterTitlesSummary struct {
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
	ActiveTitle struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name          string `json:"name"`
		ID            int    `json:"id"`
		DisplayString string `json:"display_string"`
	} `json:"active_title"`
	Titles []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"titles"`
}
