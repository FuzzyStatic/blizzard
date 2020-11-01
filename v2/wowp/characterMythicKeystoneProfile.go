package wowp

// CharacterMythicKeystoneProfile structure
type CharacterMythicKeystoneProfile struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Href string `json:"href"`
	} `json:"character"`
	CurrentPeriod struct {
		Period struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"period"`
		BestRuns []struct {
			CompletedTimestamp int64 `json:"completed_timestamp"`
			Duration           int   `json:"duration"`
			KeystoneLevel      int   `json:"keystone_level"`
			KeystoneAffixes    []struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"keystone_affixes"`
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
				} `json:"character"`
				Specialization struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"specialization"`
				Race struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"race"`
				EquippedItemLevel int `json:"equipped_item_level"`
			} `json:"members"`
			Dungeon struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"dungeon"`
			IsCompletedWithinTime bool `json:"is_completed_within_time"`
		} `json:"best_runs"`
	} `json:"current_period"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"seasons"`
}

// CharacterMythicKeystoneProfileSeason structure
type CharacterMythicKeystoneProfileSeason struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"season"`
	BestRuns []struct {
		CompletedTimestamp int64 `json:"completed_timestamp"`
		Duration           int   `json:"duration"`
		KeystoneLevel      int   `json:"keystone_level"`
		KeystoneAffixes    []struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"keystone_affixes"`
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
			} `json:"character"`
			Specialization struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"specialization"`
			Race struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"race"`
			EquippedItemLevel int `json:"equipped_item_level"`
		} `json:"members"`
		Dungeon struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"dungeon"`
		IsCompletedWithinTime bool `json:"is_completed_within_time"`
	} `json:"best_runs"`
	Character struct {
		Href string `json:"href"`
	} `json:"character"`
}
