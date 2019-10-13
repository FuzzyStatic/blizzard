package wowcgd

// GuildCrestComponentsIndex structure
type GuildCrestComponentsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Emblems []struct {
		ID    int `json:"id"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"media"`
	} `json:"emblems"`
	Borders []struct {
		ID    int `json:"id"`
		Media struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"media"`
	} `json:"borders"`
}

// GuildCrestBorderMdedia structure
type GuildCrestBorderMdedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}

// GuildCrestEmblemMdedia structure
type GuildCrestEmblemMdedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}
