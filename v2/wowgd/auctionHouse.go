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
			ID         int   `json:"id"`
			Context    int   `json:"context"`
			BonusLists []int `json:"bonus_lists"`
			Modifiers  []struct {
				Type  int `json:"type"`
				Value int `json:"value"`
			} `json:"modifiers"`
			PetBreedID   int `json:"pet_breed_id"`
			PetLevel     int `json:"pet_level"`
			PetQualityID int `json:"pet_quality_id"`
			PetSpeciesID int `json:"pet_species_id"`
		} `json:"item"`
		Buyout    int      `json:"buyout"`
		Quantity  int      `json:"quantity"`
		UnitPrice int      `json:"unit_price"`
		TimeLeft  TimeLeft `json:"time_left"`
	} `json:"auctions"`
}
