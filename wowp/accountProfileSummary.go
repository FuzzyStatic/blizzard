package wowp

// AccountProfileSummary structure
type AccountProfileSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		User struct {
			Href string `json:"href"`
		} `json:"user"`
		Profile struct {
			Href string `json:"href"`
		} `json:"profile"`
	} `json:"_links"`
	ID          int `json:"id"`
	WowAccounts []struct {
		ID         int `json:"id"`
		Characters []struct {
			Character struct {
				Href string `json:"href"`
			} `json:"character"`
			ProtectedCharacter struct {
				Href string `json:"href"`
			} `json:"protected_character"`
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
			PlayableClass struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"playable_class"`
			PlayableRace struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"playable_race"`
			Gender struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"gender"`
			Faction struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"faction"`
			Level int `json:"level"`
		} `json:"characters"`
	} `json:"wow_accounts"`
	Collections struct {
		Href string `json:"href"`
	} `json:"collections"`
}
