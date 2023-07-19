package wowcgd

type AuctionHouseIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Auctions []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"name"`
		ID int `json:"id"`
	} `json:"auctions"`
}
