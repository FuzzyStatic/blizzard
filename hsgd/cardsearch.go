package hsgd

type Card struct {
	ID           int    `json:"id"`
	Collectible  int    `json:"collectible"`
	Slug         string `json:"slug"`
	ClassID      int    `json:"classId"`
	CardTypeID   int    `json:"cardTypeId"`
	CardSetID    int    `json:"cardSetId"`
	RarityID     int    `json:"rarityId"`
	ArtistName   string `json:"artistName"`
	ManaCost     int    `json:"manaCost"`
	Name         string `json:"name"`
	Text         string `json:"text"`
	Image        string `json:"image"`
	FlavorText   string `json:"flavorText"`
	ChildIds     []int  `json:"childIds,omitempty"`
	KeywordIds   []int  `json:"keywordIds,omitempty"`
	Health       int    `json:"health,omitempty"`
	Attack       int    `json:"attack,omiteCardsmpty"`
	MinionTypeID int    `json:"minionTypeId,omitempty"`
	Armor        int    `json:"armor,omitempty"`
}

type CardSearch struct {
	Cards     []Card `json:"cards"`
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
}
