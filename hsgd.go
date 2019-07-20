package blizzard

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/FuzzyStatic/blizzard/hsgd"
)

const (
	hsgdPath         = "/hearthstone"
	cardsearchPath   = "/cards"
	deckPath         = "/deck"
	hsMetadataPath   = "/metadata"
	setField         = "set="
	classField       = "class="
	manaCostField    = "manaCost="
	attackField      = "attack="
	healthField      = "health="
	collectibleField = "collectible="
	rarityField      = "rarity="
	typeField        = "type="
	minionTypeField  = "minionType="
	keywordField     = "keyword="
	textFilterField  = "textFilter="
	pageField        = "page="
	pageSizeField    = "pageSize="
	sortField        = "sort="
	orderField       = "order="
)

// Collectiblility type
type Collectiblility string

func (collectiblility Collectiblility) String() string {
	return string(collectiblility)
}

// Sort constants
const (
	CollectiblilityCollectible   = Collectiblility("1")
	CollectiblilityUncollectible = Collectiblility("0")
	CollectiblilityBoth          = Collectiblility("0,1")
)

// Sort
// manaCost, attack, health, or name
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

// Order
// asc or desc
type Order string

func (order Order) String() string {
	return string(order)
}

// Order constants
const (
	OrderAsc  = Order("asc")
	OrderDesc = Order("desc")
)

// MetadataType
// sets, setGroups, types, rarities, classes, minionTypes, and keywords
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

// HSCardsAll returns an up-to-date list of all cards matching the search criteria.
// For more information about the search parameters, see the Hearthstone Guide.
func (c *Client) HSCardsAll() (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+cardsearchPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSCards returns an up-to-date list of all cards matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Hearthstone Guide.
func (c *Client) HSCards(setSlug, classSlug, raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	collectiblility Collectiblility, sort Sort, order Order) (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		url string
		b   []byte
		err error
	)

	url = c.apiURL + hsgdPath + cardsearchPath + "?" + localeQuery + c.locale.String()

	if setSlug != "" {
		url = url + "&" + setField + setSlug
	}

	if classSlug != "" {
		url = url + "&" + classField + classSlug
	}

	if raritySlug != "" {
		url = url + "&" + rarityField + raritySlug
	}

	if typeSlug != "" {
		url = url + "&" + typeField + typeSlug
	}

	if minionTypeSlug != "" {
		url = url + "&" + minionTypeField + minionTypeSlug
	}

	if keywordSlug != "" {
		url = url + "&" + keywordField + keywordSlug
	}

	if textFilter != "" {
		url = url + "&" + textFilterField + textFilter
	}

	if manaCost != nil {
		url = url + "&" + manaCostField + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]")
	}

	if attack != nil {
		url = url + "&" + attackField + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]")
	}

	if health != nil {
		url = url + "&" + healthField + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]")
	}

	if page != 0 {
		url = url + "&" + pageField + strconv.Itoa(page)
	}

	if pageSize != 0 {
		url = url + "&" + pageSizeField + strconv.Itoa(pageSize)
	}

	url = url + "&" + collectibleField + string(collectiblility)

	url = url + "&" + sortField + string(sort)

	url = url + "&" + orderField + string(order)

	b, err = c.getURLBody(url, "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSCardByIDOrSlug returns card by ID or slug.
// For more information about the search parameters, see the Hearthstone Guide.
func (c *Client) HSCardByIDOrSlug(idOrSlug string) (*hsgd.Card, []byte, error) {
	var (
		dat hsgd.Card
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+cardsearchPath+"/"+idOrSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSDeck Finds a deck by its deck code.
// For more information, see the Hearthstone Guide.
func (c *Client) HSDeck(deckCode string) (*hsgd.Deck, []byte, error) {
	var (
		dat hsgd.Deck
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+deckPath+"/"+deckCode+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadata returns information about the categorization of cards. Metadata includes the card set, set group (for example, Standard or Year of the Dragon), rarity, class, card type, minion type, and keywords.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadata() (*hsgd.Metadata, []byte, error) {
	var (
		dat hsgd.Metadata
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataSets returns information about set metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataSets() (*[]hsgd.Set, []byte, error) {
	var (
		dat []hsgd.Set
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeSets)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataSetGroups returns information about set group metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataSetGroups() (*[]hsgd.SetGroup, []byte, error) {
	var (
		dat []hsgd.SetGroup
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeSetGroups)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataTypes returns information about type metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataTypes() (*[]hsgd.Type, []byte, error) {
	var (
		dat []hsgd.Type
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeTypes)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataRarities returns information about rarity metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataRarities() (*[]hsgd.Rarity, []byte, error) {
	var (
		dat []hsgd.Rarity
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeRarities)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataClasses returns information about class metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataClasses() (*[]hsgd.Class, []byte, error) {
	var (
		dat []hsgd.Class
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeClasses)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataMinionTypes returns information about minion type metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataMinionTypes() (*[]hsgd.MinionType, []byte, error) {
	var (
		dat []hsgd.MinionType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeMinionTypes)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// HSMetadataKeywords returns information about keyword metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataKeywords() (*[]hsgd.Keyword, []byte, error) {
	var (
		dat []hsgd.Keyword
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+hsgdPath+hsMetadataPath+"/"+string(MetadataTypeKeywords)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}
