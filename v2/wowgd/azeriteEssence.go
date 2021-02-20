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

// AzeriteEssenceSearch structure
type AzeriteEssenceSearch struct {
	Page        int `json:"page"`
	PageSize    int `json:"pageSize"`
	MaxPageSize int `json:"maxPageSize"`
	PageCount   int `json:"pageCount"`
	Results     []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			AllowedSpecializations []struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"allowed_specializations"`
			Name struct {
				ItIT string `json:"it_IT"`
				RuRU string `json:"ru_RU"`
				EnGB string `json:"en_GB"`
				ZhTW string `json:"zh_TW"`
				KoKR string `json:"ko_KR"`
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				EsES string `json:"es_ES"`
				ZhCN string `json:"zh_CN"`
				FrFR string `json:"fr_FR"`
				DeDE string `json:"de_DE"`
			} `json:"name"`
		} `json:"data"`
	} `json:"results"`
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
