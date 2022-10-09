package wowp

// Guild structure
type Guild struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	AchievementPoints int `json:"achievement_points"`
	MemberCount       int `json:"member_count"`
	Realm             struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realm"`
	Crest struct {
		Emblem struct {
			ID    int `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float32 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"emblem"`
		Border struct {
			ID    int `json:"id"`
			Media struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
			} `json:"media"`
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float32 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"border"`
		Background struct {
			Color struct {
				ID   int `json:"id"`
				Rgba struct {
					R int     `json:"r"`
					G int     `json:"g"`
					B int     `json:"b"`
					A float32 `json:"a"`
				} `json:"rgba"`
			} `json:"color"`
		} `json:"background"`
	} `json:"crest"`
	Roster struct {
		Href string `json:"href"`
	} `json:"roster"`
	Achievements struct {
		Href string `json:"href"`
	} `json:"achievements"`
	CreatedTimestamp int64 `json:"created_timestamp"`
}

// GuildActivity structure
type GuildActivity struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Guild struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
		Faction struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"faction"`
	} `json:"guild"`
	Activities []struct {
		CharacterAchievement struct {
			Character struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name  string `json:"name"`
				ID    int    `json:"id"`
				Realm struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
					Slug string `json:"slug"`
				} `json:"realm"`
			} `json:"character"`
			Achievement struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"achievement"`
		} `json:"character_achievement,omitempty"`
		Activity struct {
			Type string `json:"type"`
		} `json:"activity"`
		Timestamp          int64 `json:"timestamp"`
		EncounterCompleted struct {
			Encounter struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"encounter"`
			Mode struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"mode"`
		} `json:"encounter_completed,omitempty"`
	} `json:"activities"`
}

// GuildAchievements structure
type GuildAchievements struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Guild struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"guild"`
	TotalQuantity int `json:"total_quantity"`
	TotalPoints   int `json:"total_points"`
	Achievements  []struct {
		ID          int `json:"id"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		Criteria struct {
			ID            int     `json:"id"`
			Amount        float64 `json:"amount"`
			IsCompleted   bool    `json:"is_completed"`
			ChildCriteria []struct {
				ID          int     `json:"id"`
				Amount      float64 `json:"amount"`
				IsCompleted bool    `json:"is_completed"`
			} `json:"child_criteria"`
		} `json:"criteria,omitempty"`
		CompletedTimestamp int64 `json:"completed_timestamp,omitempty"`
	} `json:"achievements"`
	CategoryProgress []struct {
		Category struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"category"`
		Quantity int `json:"quantity"`
		Points   int `json:"points"`
	} `json:"category_progress"`
	RecentEvents []struct {
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		Timestamp int64 `json:"timestamp"`
	} `json:"recent_events"`
}

// GuildRoster structure
type GuildRoster struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Guild struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"guild"`
	Members []struct {
		Character struct {
			Name  string `json:"name"`
			ID    int    `json:"id"`
			Realm struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID   int    `json:"id"`
				Slug string `json:"slug"`
			} `json:"realm"`
			Level         int `json:"level"`
			PlayableClass struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"playable_class"`
			PlayableRace struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"playable_race"`
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
}
