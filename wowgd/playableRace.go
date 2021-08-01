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
		ID   int    `json:"id"`
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
		Male   string `json:"male"`
		Female string `json:"female"`
	} `json:"gender_name"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	IsSelectable bool `json:"is_selectable"`
	IsAlliedRace bool `json:"is_allied_race"`
}
