package wowgd

// QuestsIndex structure
type QuestsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Categories struct {
		Href string `json:"href"`
	} `json:"categories"`
	Areas struct {
		Href string `json:"href"`
	} `json:"areas"`
	Types struct {
		Href string `json:"href"`
	} `json:"types"`
}

// Quest structure
type Quest struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID    int    `json:"id"`
	Title string `json:"title"`
	Area  struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"area"`
	Description  string `json:"description"`
	Requirements struct {
		MinCharacterLevel int `json:"min_character_level"`
		MaxCharacterLevel int `json:"max_character_level"`
		Faction           struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"requirements"`
	Rewards struct {
		Experience  int `json:"experience"`
		Reputations []struct {
			Reward struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"reward"`
			Value int `json:"value"`
		} `json:"reputations"`
		Money struct {
			Value int `json:"value"`
			Units struct {
				Gold   int `json:"gold"`
				Silver int `json:"silver"`
				Copper int `json:"copper"`
			} `json:"units"`
		} `json:"money"`
	} `json:"rewards"`
}

// QuestCategoriesIndex structure
type QuestCategoriesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Categories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"categories"`
}

// QuestCategory structure
type QuestCategory struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID       int    `json:"id"`
	Category string `json:"category"`
	Quests   []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"quests"`
}

// QuestAreasIndex structure
type QuestAreasIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Areas []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"areas"`
}

// QuestArea structure
type QuestArea struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Area   string `json:"area"`
	Quests []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"quests"`
}

// QuestTypesIndex structure
type QuestTypesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Types []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"types"`
}

// QuestType structure
type QuestType struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Quests []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"quests"`
}
