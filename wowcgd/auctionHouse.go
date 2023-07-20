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
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"auctions"`
}

type Auctions struct {
	Linkds struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Auctions []Auction `json:"auctions"`
	ID       int       `json:"id"`
	Name     string    `json:"name"`
}

type Auction struct {
	ID   int `json:"id"`
	Item struct {
		ID int `json:"id"`
	} `json:"item"`
	Bid      int    `json:"bid"`
	Buyout   int    `json:"buyout"`
	Quantity int    `json:"quantity"`
	TimeLeft string `json:"time_left"`
}
