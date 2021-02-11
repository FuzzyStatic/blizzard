package wowgd

// RealmIndex structure
type RealmIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Realms []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realms"`
}

// Realm structure
type Realm struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int `json:"id"`
	Region struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"region"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Type     struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	IsTournament bool   `json:"is_tournament"`
	Slug         string `json:"slug"`
}

// RealmSearch structure
type RealmSearch struct {
	Page        int `json:"page"`
	PageSize    int `json:"pageSize"`
	MaxPageSize int `json:"maxPageSize"`
	PageCount   int `json:"pageCount"`
	Results     []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			IsTournament bool   `json:"is_tournament"`
			Timezone     string `json:"timezone"`
			Name         struct {
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
			ID     int `json:"id"`
			Region struct {
				Name struct {
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
			} `json:"region"`
			Category struct {
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
			} `json:"category"`
			Locale string `json:"locale"`
			Type   struct {
				Name struct {
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
				Type string `json:"type"`
			} `json:"type"`
			Slug string `json:"slug"`
		} `json:"data"`
	} `json:"results"`
}
