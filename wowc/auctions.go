package wowc

// AuctionFiles structure
type AuctionFiles struct {
	Files []struct {
		URL          string `json:"url"`
		LastModified int64  `json:"lastModified"`
	} `json:"files"`
}

// AuctionData structure
type AuctionData struct {
	Realms []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
	Auctions []struct {
		Auc        int    `json:"auc"`
		Item       int    `json:"item"`
		Owner      string `json:"owner"`
		OwnerRealm string `json:"ownerRealm"`
		Bid        int    `json:"bid"`
		Buyout     int    `json:"buyout"`
		Quantity   int    `json:"quantity"`
		TimeLeft   string `json:"timeLeft"`
		Rand       int    `json:"rand"`
		Seed       int    `json:"seed"`
		Context    int    `json:"context"`
		BonusLists []struct {
			BonusListID int `json:"bonusListId"`
		} `json:"bonusLists,omitempty"`
		Modifiers []struct {
			Type  int `json:"type"`
			Value int `json:"value"`
		} `json:"modifiers,omitempty"`
		PetSpeciesID int `json:"petSpeciesId,omitempty"`
		PetBreedID   int `json:"petBreedId,omitempty"`
		PetLevel     int `json:"petLevel,omitempty"`
		PetQualityID int `json:"petQualityId,omitempty"`
	} `json:"auctions"`
}
