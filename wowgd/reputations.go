package wowgd

// ReputationFactionsIndex structure
type ReputationFactionsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Factions []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"factions"`
	RootFactions []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"root_factions"`
}

// ReputationFaction structure
type ReputationFaction struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ReputationTiers struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"reputation_tiers"`
}

// ReputationTiersIndex structure
type ReputationTiersIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ReputationTiers []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name,omitempty"`
	} `json:"reputation_tiers"`
}

// ReputationTiers structure
type ReputationTiers struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID    int `json:"id"`
	Tiers []struct {
		Name     string `json:"name"`
		MinValue int    `json:"min_value"`
		MaxValue int    `json:"max_value"`
		ID       int    `json:"id"`
	} `json:"tiers"`
	Faction struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"faction"`
}
