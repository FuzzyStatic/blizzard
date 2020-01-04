package blizzard

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/FuzzyStatic/blizzard/hsgd"
)

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

// HSCardsAll returns an up-to-date list of all cards matching the search criteria.
// For more information about the search parameters, see the Hearthstone Guide.
func (c *Client) HSCardsAll() (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/cards?locale=%s", c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// HSCards returns an up-to-date list of all cards matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Hearthstone Guide.
func (c *Client) HSCards(setSlug, classSlug, raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	collectiblility Collectibility, sort Sort, order Order) (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		url string
		b   []byte
		err error
	)

	url = c.apiURL + fmt.Sprintf("/hearthstone/cards?locale=%s", c.locale)

	if setSlug != "" {
		url = url + "&" + "set=" + setSlug
	}

	if classSlug != "" {
		url = url + "&" + "class=" + classSlug
	}

	if raritySlug != "" {
		url = url + "&" + "rarity=" + raritySlug
	}

	if typeSlug != "" {
		url = url + "&" + "type=" + typeSlug
	}

	if minionTypeSlug != "" {
		url = url + "&" + "minionType=" + minionTypeSlug
	}

	if keywordSlug != "" {
		url = url + "&" + "keyword=" + keywordSlug
	}

	if textFilter != "" {
		url = url + "&" + "textFilter=" + textFilter
	}

	if manaCost != nil {
		url = url + "&" + "manaCost=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]")
	}

	if attack != nil {
		url = url + "&" + "attack=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]")
	}

	if health != nil {
		url = url + "&" + "health=" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]")
	}

	if page != 0 {
		url = url + "&" + "page=" + strconv.Itoa(page)
	}

	if pageSize != 0 {
		url = url + "&" + "pageSize=" + strconv.Itoa(pageSize)
	}

	url = url + "&" + "collectible=" + string(collectiblility)
	url = url + "&" + "sort=" + string(sort)
	url = url + "&" + "order=" + string(order)

	b, err = c.getURLBody(url, "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/cards/%s?locale=%s", idOrSlug, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/deck/%s?locale=%s", deckCode, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// HSMetadata returns information about the categorization of cards.
// Metadata includes the card set, set group (for example, Standard or Year of the Dragon), rarity, class, card type, minion type, and keywords.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadata() (*hsgd.Metadata, []byte, error) {
	var (
		dat hsgd.Metadata
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata?locale=%s", c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeSets, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeSetGroups, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeTypes, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeRarities, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeClasses, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeMinionTypes, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", MetadataTypeKeywords, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
