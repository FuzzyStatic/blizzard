package wowcgd

// PvPSeasonIndex structure
type PvPSeasonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"seasons"`
}

// PvPSeason structure
type PvPSeason struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID                   int   `json:"id"`
	SeasonStartTimestamp int64 `json:"season_start_timestamp"`
	SeasonEndTimestamp   int64 `json:"season_end_timestamp"`
	PvpRegions           []struct {
		Href string `json:"href"`
	} `json:"pvp_regions"`
}

// PvPRegionIndex structure
type PvPRegionIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	PvpRegions []struct {
		Href string `json:"href"`
	} `json:"pvp_regions"`
}

// PvPRegionalSeasonIndex structure
type PvPRegionalSeasonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Seasons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"seasons"`
	CurrentSeason struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"current_season"`
}

// PvPRegionSeason structure
type PvPRegionSeason struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID           int `json:"id"`
	Leaderboards struct {
		Href string `json:"href"`
	} `json:"leaderboards"`
	Rewards struct {
		Href string `json:"href"`
	} `json:"rewards"`
	SeasonStartTimestamp int64 `json:"season_start_timestamp"`
	SeasonEndTimestamp   int64 `json:"season_end_timestamp"`
}

// PvPLeaderboardsIndex structure
type PvPLeaderboardsIndex struct {
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
	Leaderboards []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"leaderboards"`
}

// PvPLeaderboards structure
type PvPLeaderboards struct {
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
	Name    string `json:"name"`
	Bracket struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"bracket"`
	Entries []struct {
		Faction struct {
			Type string `json:"type"`
		} `json:"faction,omitempty"`
		Rank                  int `json:"rank"`
		Rating                int `json:"rating"`
		SeasonMatchStatistics struct {
			Played int `json:"played"`
			Won    int `json:"won"`
			Lost   int `json:"lost"`
		} `json:"season_match_statistics"`
		Team struct {
			Name  string `json:"name"`
			Realm struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
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
						ID int `json:"id"`
					} `json:"media"`
					Color struct {
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
						ID int `json:"id"`
					} `json:"media"`
					Color struct {
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
						Rgba struct {
							R int     `json:"r"`
							G int     `json:"g"`
							B int     `json:"b"`
							A float32 `json:"a"`
						} `json:"rgba"`
					} `json:"color"`
				} `json:"background"`
			} `json:"crest"`
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
				SeasonMatchStatistics struct {
					Played int `json:"played"`
					Won    int `json:"won"`
					Lost   int `json:"lost"`
				} `json:"season_match_statistics"`
				Rating int `json:"rating"`
			} `json:"members"`
			ID int `json:"id"`
		} `json:"team,omitempty"`
	} `json:"entries"`
}

// PvPRewardsIndex structure
type PvPRewardsIndex struct {
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
	Rewards []struct {
		Bracket struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
		} `json:"bracket"`
		Achievement struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		RatingCutoff int `json:"rating_cutoff"`
	} `json:"rewards"`
}
