package wowp

// CharacterEncountersSummary structure
type CharacterEncountersSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
	Dungeons struct {
		Href string `json:"href"`
	} `json:"dungeons"`
	Raids struct {
		Href string `json:"href"`
	} `json:"raids"`
}

// CharacterDungeons structure
type CharacterDungeons struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Expansions []struct {
		Expansion struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"expansion"`
		Instances []struct {
			Instance struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"instance"`
			Modes []struct {
				Difficulty struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"difficulty"`
				Status struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"status"`
				Progress struct {
					CompletedCount int `json:"completed_count"`
					TotalCount     int `json:"total_count"`
					Encounters     []struct {
						Encounter struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"encounter"`
						CompletedCount    int   `json:"completed_count"`
						LastKillTimestamp int64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}

// CharacterRaids structure
type CharacterRaids struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
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
	Expansions []struct {
		Expansion struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"expansion"`
		Instances []struct {
			Instance struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"instance"`
			Modes []struct {
				Difficulty struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"difficulty"`
				Status struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"status"`
				Progress struct {
					CompletedCount int `json:"completed_count"`
					TotalCount     int `json:"total_count"`
					Encounters     []struct {
						Encounter struct {
							Key struct {
								Href string `json:"href"`
							} `json:"key"`
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"encounter"`
						CompletedCount    int   `json:"completed_count"`
						LastKillTimestamp int64 `json:"last_kill_timestamp"`
					} `json:"encounters"`
				} `json:"progress"`
			} `json:"modes"`
		} `json:"instances"`
	} `json:"expansions"`
}
