package wowgd

// TitlesIndex structure
type TitlesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Titles []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"titles"`
}

// Title structure
type Title struct {
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
}
