/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 20:37:11
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 20:37:42
 */

package worldofwarcraft

import "encoding/json"

// Item structure
type Item struct {
	ID                     int    `json:"id"`
	DisenchantingSkillRank int    `json:"disenchantingSkillRank"`
	Description            string `json:"description"`
	Name                   string `json:"name"`
	Icon                   string `json:"icon"`
	Stackable              int    `json:"stackable"`
	ItemBind               int    `json:"itemBind"`
	BonusStats             []struct {
		Stat   int `json:"stat"`
		Amount int `json:"amount"`
	} `json:"bonusStats"`
	ItemSpells     []interface{} `json:"itemSpells"`
	BuyPrice       int           `json:"buyPrice"`
	ItemClass      int           `json:"itemClass"`
	ItemSubClass   int           `json:"itemSubClass"`
	ContainerSlots int           `json:"containerSlots"`
	WeaponInfo     struct {
		Damage struct {
			Min      int     `json:"min"`
			Max      int     `json:"max"`
			ExactMin float64 `json:"exactMin"`
			ExactMax float64 `json:"exactMax"`
		} `json:"damage"`
		WeaponSpeed float64 `json:"weaponSpeed"`
		Dps         float64 `json:"dps"`
	} `json:"weaponInfo"`
	InventoryType     int  `json:"inventoryType"`
	Equippable        bool `json:"equippable"`
	ItemLevel         int  `json:"itemLevel"`
	MaxCount          int  `json:"maxCount"`
	MaxDurability     int  `json:"maxDurability"`
	MinFactionID      int  `json:"minFactionId"`
	MinReputation     int  `json:"minReputation"`
	Quality           int  `json:"quality"`
	SellPrice         int  `json:"sellPrice"`
	RequiredSkill     int  `json:"requiredSkill"`
	RequiredLevel     int  `json:"requiredLevel"`
	RequiredSkillRank int  `json:"requiredSkillRank"`
	ItemSource        struct {
		SourceID   int    `json:"sourceId"`
		SourceType string `json:"sourceType"`
	} `json:"itemSource"`
	BaseArmor            int           `json:"baseArmor"`
	HasSockets           bool          `json:"hasSockets"`
	IsAuctionable        bool          `json:"isAuctionable"`
	Armor                int           `json:"armor"`
	DisplayInfoID        int           `json:"displayInfoId"`
	NameDescription      string        `json:"nameDescription"`
	NameDescriptionColor string        `json:"nameDescriptionColor"`
	Upgradable           bool          `json:"upgradable"`
	HeroicTooltip        bool          `json:"heroicTooltip"`
	Context              string        `json:"context"`
	BonusLists           []interface{} `json:"bonusLists"`
	AvailableContexts    []string      `json:"availableContexts"`
	BonusSummary         struct {
		DefaultBonusLists []interface{} `json:"defaultBonusLists"`
		ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
		BonusChances      []interface{} `json:"bonusChances"`
	} `json:"bonusSummary"`
	ArtifactID int `json:"artifactId"`
}

// JSON2Struct creates Item structure from JSON byte array
func (i *Item) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, i)
}
