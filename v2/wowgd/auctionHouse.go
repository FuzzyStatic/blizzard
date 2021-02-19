package wowgd

// AuctionHouse structure
type AuctionHouse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Auctions []struct {
		ID   int `json:"id"`
		Item struct {
			ID int `json:"id"`
		} `json:"item"`
		Quantity  int      `json:"quantity"`
		UnitPrice int      `json:"unit_price"`
		TimeLeft  TimeLeft `json:"time_left"`
	} `json:"auctions"`
}
