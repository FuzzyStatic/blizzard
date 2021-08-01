package d3gd

// SeasonIndex structure
type SeasonIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Season []struct {
		Href string `json:"href"`
	} `json:"season"`
	CurrentSeason        int    `json:"current_season"`
	ServiceCurrentSeason int    `json:"service_current_season"`
	ServiceSeasonState   string `json:"service_season_state"`
	LastUpdateTime       string `json:"last_update_time"`
	GeneratedBy          string `json:"generated_by"`
}

// Season structure
type Season struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Leaderboard []struct {
		Ladder struct {
			Href string `json:"href"`
		} `json:"ladder"`
		TeamSize        int    `json:"team_size,omitempty"`
		Hardcore        bool   `json:"hardcore,omitempty"`
		HeroClassString string `json:"hero_class_string,omitempty"`
	} `json:"leaderboard"`
	SeasonID       int    `json:"season_id"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}

// Leaderboard structure
type Leaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Row []struct {
		Player []struct {
			Key       string `json:"key"`
			AccountID int    `json:"accountId"`
			Data      []struct {
				ID     string `json:"id"`
				String string `json:"string,omitempty"`
				Number int    `json:"number,omitempty"`
			} `json:"data"`
		} `json:"player"`
		Order int `json:"order"`
		Data  []struct {
			ID        string `json:"id"`
			Number    int    `json:"number,omitempty"`
			Timestamp int64  `json:"timestamp,omitempty"`
			String    string `json:"string,omitempty"`
		} `json:"data"`
	} `json:"row"`
	Key    string `json:"key"`
	Title  string `json:"title"`
	Column []struct {
		ID     string `json:"id"`
		Hidden bool   `json:"hidden"`
		Order  int    `json:"order,omitempty"`
		Label  string `json:"label"`
		Type   string `json:"type"`
	} `json:"column"`
	LastUpdateTime    string `json:"last_update_time"`
	GeneratedBy       string `json:"generated_by"`
	AchievementPoints bool   `json:"achievement_points"`
	Season            int    `json:"season"`
}

// Title structure
type Title struct {
	EnUS string `json:"en_US"`
	EsMX string `json:"es_MX"`
	PtBR string `json:"pt_BR"`
	DeDE string `json:"de_DE"`
	EnGB string `json:"en_GB"`
	EsES string `json:"es_ES"`
	FrFR string `json:"fr_FR"`
	ItIT string `json:"it_IT"`
	PlPL string `json:"pl_PL"`
	PtPT string `json:"pt_PT"`
	RuRU string `json:"ru_RU"`
	KoKR string `json:"ko_KR"`
	ZhTW string `json:"zh_TW"`
	ZhCN string `json:"zh_CN"`
}

// Label structure
type Label struct {
	EnUS string `json:"en_US"`
	EsMX string `json:"es_MX"`
	PtBR string `json:"pt_BR"`
	DeDE string `json:"de_DE"`
	EnGB string `json:"en_GB"`
	EsES string `json:"es_ES"`
	FrFR string `json:"fr_FR"`
	ItIT string `json:"it_IT"`
	PlPL string `json:"pl_PL"`
	PtPT string `json:"pt_PT"`
	RuRU string `json:"ru_RU"`
	KoKR string `json:"ko_KR"`
	ZhTW string `json:"zh_TW"`
	ZhCN string `json:"zh_CN"`
}

// EraIndex structure
type EraIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Era []struct {
		Href string `json:"href"`
	} `json:"era"`
	CurrentEra     int    `json:"current_era"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}

// Era structure
type Era struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Leaderboard []struct {
		TeamSize int `json:"team_size"`
		Ladder   struct {
			Href string `json:"href"`
		} `json:"ladder"`
		Hardcore        bool   `json:"hardcore,omitempty"`
		HeroClassString string `json:"hero_class_string,omitempty"`
	} `json:"leaderboard"`
	EraID          int    `json:"era_id"`
	EraStartDate   int    `json:"era_start_date"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}
