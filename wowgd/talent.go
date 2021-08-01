package wowgd

// TalentsIndex structure
type TalentsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Talents []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"talents"`
}

// Talent structure
type Talent struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID          int    `json:"id"`
	TierIndex   int    `json:"tier_index"`
	ColumnIndex int    `json:"column_index"`
	Level       int    `json:"level"`
	Description string `json:"description"`
	Spell       struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"spell"`
	PlayableClass struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"playable_class"`
}

// PvPTalentsIndex structure
type PvPTalentsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	PvpTalents []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"pvp_talents"`
}

// PvPTalent structure
type PvPTalent struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID    int `json:"id"`
	Spell struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"spell"`
	PlayableSpecialization struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"playable_specialization"`
	Description       string `json:"description"`
	UnlockPlayerLevel int    `json:"unlock_player_level"`
	CompatibleSlots   []int  `json:"compatible_slots"`
}
