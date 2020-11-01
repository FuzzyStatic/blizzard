package wowgd

// JournalExpansionsIndex structure
type JournalExpansionsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Tiers []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"tiers"`
}

// JournalExpansion structure
type JournalExpansion struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Dungeons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dungeons"`
	Raids []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"raids"`
}

// JournalEncountersIndex structure
type JournalEncountersIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Encounters []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"encounters"`
}

// JournalEncounter structure
type JournalEncounter struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creatures   []struct {
		ID              int    `json:"id"`
		Name            string `json:"name"`
		CreatureDisplay struct {
			ID  int `json:"id"`
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
		} `json:"creature_display"`
	} `json:"creatures"`
	Items []struct {
		ID   int `json:"id"`
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"item"`
	} `json:"items"`
	Sections []struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		BodyText string `json:"body_text"`
		Sections []struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Sections []struct {
				ID       int    `json:"id"`
				Title    string `json:"title"`
				BodyText string `json:"body_text"`
			} `json:"sections"`
			CreatureDisplay struct {
				ID  int `json:"id"`
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"creature_display"`
		} `json:"sections"`
	} `json:"sections"`
	Instance struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"instance"`
	Category struct {
		Type string `json:"type"`
	} `json:"category"`
	Modes []struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"modes"`
}

// JournalInstancesIndex structure
type JournalInstancesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Instances []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"instances"`
}

// JournalInstance structure
type JournalInstance struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MinimumLevel int    `json:"minimum_level"`
	Map          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"map"`
	Area struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"area"`
	Location struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	Encounters []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"encounters"`
	Expansion struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"expansion"`
	Modes []struct {
		Mode struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"mode"`
		Players   int  `json:"players"`
		IsTracked bool `json:"is_tracked"`
	} `json:"modes"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	Category struct {
		Type string `json:"type"`
	} `json:"category"`
}
