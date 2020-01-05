package wowgd

// MountIndex structure
type MountIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Mounts []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"mounts"`
}

// Mount structure
type Mount struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
	CreatureDisplays []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"creature_displays"`
	Description string `json:"description"`
	Source      struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"source"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
}
