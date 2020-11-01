package hsgd

// Metadata structure
type Metadata struct {
	Sets             []Set        `json:"sets"`
	SetGroups        []SetGroup   `json:"setGroups"`
	Types            []Type       `json:"types"`
	Rarities         []Rarity     `json:"rarities"`
	Classes          []Class      `json:"classes"`
	MinionTypes      []MinionType `json:"minionTypes"`
	Keywords         []Keyword    `json:"keywords"`
	FilterableFields []string     `json:"filterableFields"`
	NumericFields    []string     `json:"numericFields"`
}

// Set structure
type Set struct {
	ID                          int    `json:"id"`
	Slug                        string `json:"slug"`
	ReleaseDate                 string `json:"releaseDate,omitempty"`
	Name                        string `json:"name"`
	CollectibleCount            int    `json:"collectibleCount"`
	CollectibleRevealedCount    int    `json:"collectibleRevealedCount"`
	NonCollectibleCount         int    `json:"nonCollectibleCount"`
	NonCollectibleRevealedCount int    `json:"nonCollectibleRevealedCount"`
	Type                        string `json:"type,omitempty"`
}

// SetGroup structure
type SetGroup struct {
	Slug      string   `json:"slug"`
	Year      int      `json:"year,omitempty"`
	CardSets  []string `json:"cardSets"`
	Name      string   `json:"name"`
	Standard  bool     `json:"standard,omitempty"`
	Icon      string   `json:"icon,omitempty"`
	YearRange string   `json:"yearRange,omitempty"`
}

// Type structure
type Type struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Rarity structure
type Rarity struct {
	ID           int    `json:"id"`
	Slug         string `json:"slug"`
	CraftingCost []int  `json:"craftingCost"`
	DustValue    []int  `json:"dustValue"`
	Name         string `json:"name"`
}

// Class structure
type Class struct {
	ID   interface{} `json:"id"`
	Slug string      `json:"slug"`
	Name string      `json:"name"`
}

// MinionType structure
type MinionType struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Keyword structure
type Keyword struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	RefText string `json:"refText"`
	Text    string `json:"text"`
}
