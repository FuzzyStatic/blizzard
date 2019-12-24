package wowp

// CharacterReputationsSummary structure
type CharacterReputationsSummary struct {
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
	Reputations []struct {
		Faction struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"faction"`
		Standing struct {
			Raw   int    `json:"raw"`
			Value int    `json:"value"`
			Max   int    `json:"max"`
			Tier  int    `json:"tier"`
			Name  string `json:"name"`
		} `json:"standing"`
		Paragon struct {
			Raw   int `json:"raw"`
			Value int `json:"value"`
			Max   int `json:"max"`
		} `json:"paragon,omitempty"`
	} `json:"reputations"`
}
