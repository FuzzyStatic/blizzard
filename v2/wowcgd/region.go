package wowcgd

// RegionIndex structure
type RegionIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Regions []struct {
		Href string `json:"href"`
	} `json:"regions"`
}

// Region structure
type Region struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
