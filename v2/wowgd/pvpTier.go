package wowgd

// PvPTierMedia structure
type PvPTierMedia struct {
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

// PvPTiersIndex structure
type PvPTiersIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Tiers []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"tiers"`
}

// PvPTier structure
type PvPTier struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MinRating int    `json:"min_rating"`
	MaxRating int    `json:"max_rating"`
	Media     struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"media"`
	Bracket struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"bracket"`
	RatingType int `json:"rating_type"`
}
