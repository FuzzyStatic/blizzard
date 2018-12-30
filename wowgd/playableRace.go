package wowgd

// PlayableRacesIndex structure
type PlayableRacesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Races []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
	} `json:"races"`
}

// PlayableRace structure
type PlayableRace struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	GenderName struct {
		MaleName   string `json:"male_name"`
		FemaleName string `json:"female_name"`
	} `json:"gender_name"`
	Faction []struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
}
