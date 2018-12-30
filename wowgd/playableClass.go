package wowgd

// PlayableClassesIndex structure
type PlayableClassesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Classes []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"classes"`
}

// PlayableClass structure
type PlayableClass struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	GenderName struct {
		Male   string `json:"male"`
		Female string `json:"female"`
	} `json:"gender_name"`
	PowerType struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"power_type"`
	Specializations []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"specializations"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	PvpTalentSlots struct {
		Href string `json:"href"`
	} `json:"pvp_talent_slots"`
}

// PlayableClassPVPTalentSlots structure
type PlayableClassPVPTalentSlots struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	TalentSlots []struct {
		SlotNumber        int `json:"slot_number"`
		UnlockPlayerLevel int `json:"unlock_player_level"`
	} `json:"talent_slots"`
}
