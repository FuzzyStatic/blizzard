package wowc

// ChallengeModeRealmLeaderboard structure
type ChallengeModeRealmLeaderboard struct {
	Challenge []struct {
		Map struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Slug             string `json:"slug"`
			HasChallengeMode bool   `json:"hasChallengeMode"`
			BronzeCriteria   struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"bronzeCriteria"`
			SilverCriteria struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"silverCriteria"`
			GoldCriteria struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"goldCriteria"`
		} `json:"map"`
		Groups []interface{} `json:"groups"`
		Realm  struct {
			Name            string   `json:"name"`
			Slug            string   `json:"slug"`
			Battlegroup     string   `json:"battlegroup"`
			Locale          string   `json:"locale"`
			Timezone        string   `json:"timezone"`
			ConnectedRealms []string `json:"connected_realms"`
		} `json:"realm,omitempty"`
	} `json:"challenge"`
}

// ChallengeModeRegionLeaderboard structure
type ChallengeModeRegionLeaderboard struct {
	Challenge []struct {
		Groups []interface{} `json:"groups"`
	} `json:"challenge"`
}
