package wowgd

// MythicKeystoneDungeonIndex structure
type MythicKeystoneDungeonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Dungeons []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"dungeons"`
}

// MythicKeystoneDungeon structure
type MythicKeystoneDungeon struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Map  struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"map"`
	Zone struct {
		Slug string `json:"slug"`
	} `json:"zone"`
	KeystoneUpgrades []struct {
		UpgradeLevel       int `json:"upgrade_level"`
		QualifyingDuration int `json:"qualifying_duration"`
	} `json:"keystone_upgrades"`
}

// MythicKeystoneIndex structure
type MythicKeystoneIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Seasons struct {
		Href string `json:"href"`
	} `json:"seasons"`
	Dungeons struct {
		Href string `json:"href"`
	} `json:"dungeons"`
}

// MythicKeystonePeriodIndex structure
type MythicKeystonePeriodIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Periods []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"periods"`
	CurrentPeriod struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"current_period"`
}

// MythicKeystonePeriod structure
type MythicKeystonePeriod struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID             int   `json:"id"`
	StartTimestamp int64 `json:"start_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`
}

// MythicKeystoneSeasonIndex structure
type MythicKeystoneSeasonIndex struct {
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

// MythicKeystoneSeason structure
type MythicKeystoneSeason struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID             int   `json:"id"`
	StartTimestamp int64 `json:"start_timestamp"`
	Periods        []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"periods"`
}
