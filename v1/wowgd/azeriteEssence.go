package wowgd

// AzeriteEssenceIndex structure
type AzeriteEssenceIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	AzeriteEssences []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"azerite_essences"`
}

// AzeriteEssence structure
type AzeriteEssence struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	AllowedSpecializations []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"allowed_specializations"`
	Powers []struct {
		ID             int `json:"id"`
		Rank           int `json:"rank"`
		MainPowerSpell struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"main_power_spell"`
		PassivePowerSpell struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"passive_power_spell"`
	} `json:"powers"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

// AzeriteEssenceMedia structure
type AzeriteEssenceMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}
