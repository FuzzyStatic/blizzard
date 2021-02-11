package wowgd

// Spell structure
type Spell struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Media       struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

// SpellMedia structure
type SpellMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		FileDataID int    `json:"file_data_id"`
	} `json:"assets"`
	ID int `json:"id"`
}

// SpellSearch structure
type SpellSearch struct {
	Page        int `json:"page"`
	PageSize    int `json:"pageSize"`
	MaxPageSize int `json:"maxPageSize"`
	PageCount   int `json:"pageCount"`
	Results     []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
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
			ID    int `json:"id"`
			Media struct {
				ID int `json:"id"`
			} `json:"media"`
		} `json:"data"`
	} `json:"results"`
}
