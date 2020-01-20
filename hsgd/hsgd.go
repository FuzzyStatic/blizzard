// Package hsgd contains types for the Hearthstone Game Data APIs
package hsgd

// Collectibility type
type Collectibility string

func (collectibility Collectibility) String() string {
	return string(collectibility)
}

// Collectibility constants
const (
	CollectibilityCollectible   = Collectibility("1")
	CollectibilityUncollectible = Collectibility("0")
	CollectibilityBoth          = Collectibility("0,1")
)

// Sort - manaCost, attack, health, or name
type Sort string

func (sort Sort) String() string {
	return string(sort)
}

// Sort constants
const (
	SortManaCost = Sort("manaCost")
	SortAttack   = Sort("attack")
	SortHealth   = Sort("health")
	SortName     = Sort("name")
)

// Order - asc or desc
type Order string

func (order Order) String() string {
	return string(order)
}

// Order constants
const (
	OrderAsc  = Order("asc")
	OrderDesc = Order("desc")
)

// MetadataType - sets, setGroups, types, rarities, classes, minionTypes, and keywords
type MetadataType string

func (metadataType MetadataType) String() string {
	return string(metadataType)
}

// MetadataType constants
const (
	MetadataTypeSets        = MetadataType("sets")
	MetadataTypeSetGroups   = MetadataType("setGroups")
	MetadataTypeTypes       = MetadataType("types")
	MetadataTypeRarities    = MetadataType("rarities")
	MetadataTypeClasses     = MetadataType("classes")
	MetadataTypeMinionTypes = MetadataType("minionTypes")
	MetadataTypeKeywords    = MetadataType("keywords")
)

// GameMode - constructed or battlegrounds
type GameMode string

func (gameMode GameMode) String() string {
	return string(gameMode)
}

// GameMode constants
const (
	GameModeConstructed   = GameMode("constructed")
	GameModeBattlegrounds = GameMode("battlegrounds")
)

// Tier - hero, 1, 2, 3, 4, 5, or 6
type Tier string

func (tier Tier) String() string {
	return string(tier)
}

// GameMode constants
const (
	TierHero = Tier("hero")
	Tier1    = Tier("1")
	Tier2    = Tier("2")
	Tier3    = Tier("3")
	Tier4    = Tier("4")
	Tier5    = Tier("5")
	Tier6    = Tier("6")
)
