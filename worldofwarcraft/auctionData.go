/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 20:53:48
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 21:11:58
 */

package worldofwarcraft

import "encoding/json"

// AuctionData structure
type AuctionData struct {
	Files []struct {
		URL          string `json:"url"`
		LastModified int64  `json:"lastModified"`
	} `json:"files"`
}

// JSON2Struct creates AuctionData structure from JSON byte array
func (a *AuctionData) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, a)
}

// Auctions structure
type Auctions struct {
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
		Modifiers  []struct {
			Type  int `json:"type"`
			Value int `json:"value"`
		} `json:"modifiers,omitempty"`
		PetSpeciesID int `json:"petSpeciesId,omitempty"`
		PetBreedID   int `json:"petBreedId,omitempty"`
		PetLevel     int `json:"petLevel,omitempty"`
		PetQualityID int `json:"petQualityId,omitempty"`
		BonusLists   []struct {
			BonusListID int `json:"bonusListId"`
		} `json:"bonusLists,omitempty"`
	} `json:"auctions"`
}

// JSON2Struct creates Auctions structure from JSON byte array
func (a *Auctions) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, a)
}
