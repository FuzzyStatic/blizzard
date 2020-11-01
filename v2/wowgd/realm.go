package wowgd

// RealmIndex structure
type RealmIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Realms []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realms"`
}

// Realm structure
type Realm struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int `json:"id"`
	Region struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"region"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Type     struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	IsTournament bool   `json:"is_tournament"`
	Slug         string `json:"slug"`
}
