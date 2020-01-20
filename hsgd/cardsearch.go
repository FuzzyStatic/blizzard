package hsgd

// Card structure
type Card struct {
	ID            int    `json:"id"`
	Collectible   int    `json:"collectible"`
	Slug          string `json:"slug"`
	ClassID       int    `json:"classId"`
	MultiClassIds []int  `json:"multiClassIds"`
	CardTypeID    int    `json:"cardTypeId"`
	CardSetID     int    `json:"cardSetId"`
	RarityID      int    `json:"rarityId"`
	ArtistName    string `json:"artistName"`
	Health        int    `json:"health"`
	Attack        int    `json:"attack"`
	ManaCost      int    `json:"manaCost"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	Image         string `json:"image"`
	ImageGold     string `json:"imageGold"`
	FlavorText    string `json:"flavorText"`
	CropImage     string `json:"cropImage"`
	ChildIds      []int  `json:"childIds"`
	KeywordIds    []int  `json:"keywordIds"`
	Battlegrounds struct {
		Tier      int    `json:"tier"`
		Hero      bool   `json:"hero"`
		UpgradeID int    `json:"upgradeId"`
		Image     string `json:"image"`
		ImageGold string `json:"imageGold"`
	} `json:"battlegrounds,omitempty"`
}

// CardSearch structure
type CardSearch struct {
	Cards     []Card `json:"cards"`
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
}
