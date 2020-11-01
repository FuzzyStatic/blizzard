package wowcgd

// PlayableClassesIndex structure
type PlayableClassesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Classes []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"classes"`
}

// PlayableClass structure
type PlayableClass struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	GenderName struct {
		Male   string `json:"male"`
		Female string `json:"female"`
	} `json:"gender_name"`
	PowerType struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"power_type"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	PvpTalentSlots struct {
		Href string `json:"href"`
	} `json:"pvp_talent_slots"`
}

// PlayableClassMedia structure
type PlayableClassMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}
