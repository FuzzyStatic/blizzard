package hsgd

type CardBackSearchAllLocales struct {
	CardBacks []struct {
		ID           int `json:"id"`
		SortCategory int `json:"sortCategory"`
		Text         struct {
			DeDE string `json:"de_DE"`
			EnUS string `json:"en_US"`
			EsES string `json:"es_ES"`
			EsMX string `json:"es_MX"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			JaJP string `json:"ja_JP"`
			KoKR string `json:"ko_KR"`
			PlPL string `json:"pl_PL"`
			PtBR string `json:"pt_BR"`
			RuRU string `json:"ru_RU"`
			ThTH string `json:"th_TH"`
			ZhCN string `json:"zh_CN"`
			ZhTW string `json:"zh_TW"`
		} `json:"text"`
		Name struct {
			DeDE string `json:"de_DE"`
			EnUS string `json:"en_US"`
			EsES string `json:"es_ES"`
			EsMX string `json:"es_MX"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			JaJP string `json:"ja_JP"`
			KoKR string `json:"ko_KR"`
			PlPL string `json:"pl_PL"`
			PtBR string `json:"pt_BR"`
			RuRU string `json:"ru_RU"`
			ThTH string `json:"th_TH"`
			ZhCN string `json:"zh_CN"`
			ZhTW string `json:"zh_TW"`
		} `json:"name"`
		Image string `json:"image"`
		Slug  string `json:"slug"`
	} `json:"cardBacks"`
	CardCount int `json:"cardCount"`
	PageCount int `json:"pageCount"`
	Page      int `json:"page"`
}

type CardBackSearch struct {
	CardBacks []struct {
		ID           int    `json:"id"`
		SortCategory int    `json:"sortCategory"`
		Text         string `json:"text"`
		Name         string `json:"name"`
		Image        string `json:"image"`
		Slug         string `json:"slug"`
	} `json:"cardBacks"`
	CardCount int `json:"cardCount"`
	PageCount int `json:"pageCount"`
	Page      int `json:"page"`
}
