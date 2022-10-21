package wowp

// CharacterProfileSummary structure
type CharacterProfileSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"race"`
	CharacterClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_class"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_spec"`
	Realm struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realm"`
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
	Level             int `json:"level"`
	Experience        int `json:"experience"`
	AchievementPoints int `json:"achievement_points"`
	Achievements      struct {
		Href string `json:"href"`
	} `json:"achievements"`
	Titles struct {
		Href string `json:"href"`
	} `json:"titles"`
	PvpSummary struct {
		Href string `json:"href"`
	} `json:"pvp_summary"`
	RaidProgression struct {
		Href string `json:"href"`
	} `json:"raid_progression"`
	Media struct {
		Href string `json:"href"`
	} `json:"media"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	AverageItemLevel   int   `json:"average_item_level"`
	EquippedItemLevel  int   `json:"equipped_item_level"`
	Specializations    struct {
		Href string `json:"href"`
	} `json:"specializations"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	MythicKeystoneProfile struct {
		Href string `json:"href"`
	} `json:"mythic_keystone_profile"`
	Equipment struct {
		Href string `json:"href"`
	} `json:"equipment"`
	Appearance struct {
		Href string `json:"href"`
	} `json:"appearance"`
	Collections struct {
		Href string `json:"href"`
	} `json:"collections"`
	ActiveTitle ActiveTitle `json:"active_title"`
	Reputations struct {
		Href string `json:"href"`
	} `json:"reputations"`
}

// ActiveTitle structure
type ActiveTitle struct {
	Key struct {
		Href string `json:"href"`
	} `json:"key"`
	Name          string `json:"name"`
	ID            int    `json:"id"`
	DisplayString string `json:"display_string"`
}

// CharacterProfileSummaryPreRev24 structure
type CharacterProfileSummaryPreRev24 struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"gender"`
	Faction struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"faction"`
	Race struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"race"`
	CharacterClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"character_class"`
	ActiveSpec struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"active_spec"`
	Realm struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realm"`
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
	Level             int `json:"level"`
	Experience        int `json:"experience"`
	AchievementPoints int `json:"achievement_points"`
	Achievements      struct {
		Href string `json:"href"`
	} `json:"achievements"`
	Titles struct {
		Href string `json:"href"`
	} `json:"titles"`
	PvpSummary struct {
		Href string `json:"href"`
	} `json:"pvp_summary"`
	RaidProgression struct {
		Href string `json:"href"`
	} `json:"raid_progression"`
	Media struct {
		Href string `json:"href"`
	} `json:"media"`
	LastLoginTimestamp int64 `json:"last_login_timestamp"`
	AverageItemLevel   int   `json:"average_item_level"`
	EquippedItemLevel  int   `json:"equipped_item_level"`
	Specializations    struct {
		Href string `json:"href"`
	} `json:"specializations"`
	Statistics struct {
		Href string `json:"href"`
	} `json:"statistics"`
	MythicKeystoneProfile struct {
		Href string `json:"href"`
	} `json:"mythic_keystone_profile"`
	Equipment struct {
		Href string `json:"href"`
	} `json:"equipment"`
	Appearance struct {
		Href string `json:"href"`
	} `json:"appearance"`
	Collections struct {
		Href string `json:"href"`
	} `json:"collections"`
	Reputations struct {
		Href string `json:"href"`
	} `json:"reputations"`
}

// ActiveTitlePreRev24 structure
type ActiveTitlePreRev24 struct {
	ActiveTitle string `json:"active_title"`
}

// CharacterProfileStatus structure
type CharacterProfileStatus struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID      int  `json:"id"`
	IsValid bool `json:"is_valid"`
}

// ConvertCharacterProfileSummaryPreRev24 converts profile summary to pre revision 24
func ConvertCharacterProfileSummaryPreRev24(activeTitle string,
	cpspr *CharacterProfileSummaryPreRev24, cps *CharacterProfileSummary) {
	cps.Links = cpspr.Links
	cps.ID = cpspr.ID
	cps.Name = cpspr.Name
	cps.Gender = cpspr.Gender
	cps.Faction = cpspr.Faction
	cps.Race = cpspr.Race
	cps.CharacterClass = cpspr.CharacterClass
	cps.ActiveSpec = cpspr.ActiveSpec
	cps.Realm = cpspr.Realm
	cps.Guild = cpspr.Guild
	cps.Level = cpspr.Level
	cps.Experience = cpspr.Experience
	cps.AchievementPoints = cpspr.AchievementPoints
	cps.Achievements = cpspr.Achievements
	cps.Titles = cpspr.Titles
	cps.PvpSummary = cpspr.PvpSummary
	cps.RaidProgression = cpspr.RaidProgression
	cps.Media = cpspr.Media
	cps.LastLoginTimestamp = cpspr.LastLoginTimestamp
	cps.AverageItemLevel = cpspr.AverageItemLevel
	cps.EquippedItemLevel = cpspr.EquippedItemLevel
	cps.Specializations = cpspr.Specializations
	cps.Statistics = cpspr.Statistics
	cps.MythicKeystoneProfile = cpspr.MythicKeystoneProfile
	cps.Equipment = cpspr.Equipment
	cps.Appearance = cpspr.Appearance
	cps.Collections = cpspr.Collections
	cps.ActiveTitle.Name = activeTitle
	cps.Reputations = cpspr.Reputations
}
