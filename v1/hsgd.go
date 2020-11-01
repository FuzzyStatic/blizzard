package blizzard

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FuzzyStatic/blizzard/v1/hsgd"
)

// HSCardsSearch returns an up-to-date list of all cards matching the search criteria.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardsSearch() (*hsgd.CardSearch, []byte, error) {
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

// HSDetailedCardsSearch returns an up-to-date list of all cards matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSDetailedCardsSearch(setSlug, classSlug, raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	collectiblility hsgd.Collectibility, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		url string
		b   []byte
		err error
	)

	url = c.apiURL + fmt.Sprintf("/hearthstone/cards?locale=%s", c.locale)

	if setSlug != "" {
		url = url + fmt.Sprintf("&set=%s", setSlug)
	}

	if classSlug != "" {
		url = url + fmt.Sprintf("&class=%s", classSlug)
	}

	if raritySlug != "" {
		url = url + fmt.Sprintf("&rarity=%s", raritySlug)
	}

	if typeSlug != "" {
		url = url + fmt.Sprintf("&type=%s", typeSlug)
	}

	if minionTypeSlug != "" {
		url = url + fmt.Sprintf("&minionType=%s", minionTypeSlug)
	}

	if keywordSlug != "" {
		url = url + fmt.Sprintf("&keyword=%s", keywordSlug)
	}

	if textFilter != "" {
		url = url + fmt.Sprintf("&textFilter=%s", textFilter)
	}

	if manaCost != nil {
		url = url + fmt.Sprintf("&manaCost=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]"))
	}

	if attack != nil {
		url = url + fmt.Sprintf("&attack=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]"))
	}

	if health != nil {
		url = url + fmt.Sprintf("&health=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]"))
	}

	if page != 0 {
		url = url + fmt.Sprintf("&page=%d", page)
	}

	if pageSize != 0 {
		url = url + fmt.Sprintf("&pageSize=%d", pageSize)
	}

	url = url + fmt.Sprintf("&collectible=%s", collectiblility)
	url = url + fmt.Sprintf("&sort=%s", sort)
	url = url + fmt.Sprintf("&order=%s", order)

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

// HSBattlegroundsCardsSearch returns an up-to-date list of all cards matching the search criteria for the specified game mode.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSBattlegroundsCardsSearch(raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	tier []hsgd.Tier, collectiblility hsgd.Collectibility, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardSearch, []byte, error) {
	var (
		dat hsgd.CardSearch
		url string
		b   []byte
		err error
	)

	url = c.apiURL + fmt.Sprintf("/hearthstone/cards?locale=%s&gameMode=%s", c.locale, hsgd.GameModeBattlegrounds)

	if raritySlug != "" {
		url = url + fmt.Sprintf("&rarity=%s", raritySlug)
	}

	if typeSlug != "" {
		url = url + fmt.Sprintf("&type=%s", typeSlug)
	}

	if minionTypeSlug != "" {
		url = url + fmt.Sprintf("&minionType=%s", minionTypeSlug)
	}

	if keywordSlug != "" {
		url = url + fmt.Sprintf("&keyword=%s", keywordSlug)
	}

	if textFilter != "" {
		url = url + fmt.Sprintf("&textFilter=%s", textFilter)
	}

	if manaCost != nil {
		url = url + fmt.Sprintf("&manaCost=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]"))
	}

	if attack != nil {
		url = url + fmt.Sprintf("&attack=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]"))
	}

	if health != nil {
		url = url + fmt.Sprintf("&health=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]"))
	}

	if page != 0 {
		url = url + fmt.Sprintf("&page=%d", page)
	}

	if pageSize != 0 {
		url = url + fmt.Sprintf("&pageSize=%d", pageSize)
	}

	if health != nil {
		url = url + fmt.Sprintf("&tier=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tier)), ","), "[]"))
	}

	url = url + fmt.Sprintf("&collectible=%s", collectiblility)
	url = url + fmt.Sprintf("&sort=%s", sort)
	url = url + fmt.Sprintf("&order=%s", order)

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
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardByIDOrSlug(idOrSlug string, gameMode hsgd.GameMode) (*hsgd.Card, []byte, error) {
	var (
		dat hsgd.Card
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/cards/%s?locale=%s&gameMode=%s", idOrSlug, c.locale, gameMode), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// HSCardBackSearchAllLocales returns an up-to-date list of all card backs matching the search criteria for all locales.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardBackSearchAllLocales(cardBackCategory hsgd.CardBackCategory, textFilter string, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardBackSearchAllLocales, []byte, error) {
	var (
		dat hsgd.CardBackSearchAllLocales
		url string
		b   []byte
		err error
	)

	url = c.apiURL + "/hearthstone/cardbacks"

	if cardBackCategory != "" {
		url = url + fmt.Sprintf("?cardBackCategory=%s", cardBackCategory)
	}

	if textFilter != "" {
		url = url + fmt.Sprintf("&type=%s", textFilter)
	}

	url = url + fmt.Sprintf("&sort=%s", sort)
	url = url + fmt.Sprintf("&order=%s", order)

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

// HSCardBackSearch returns an up-to-date list of all card backs matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardBackSearch(cardBackCategory hsgd.CardBackCategory, textFilter string, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardBackSearch, []byte, error) {
	var (
		dat hsgd.CardBackSearch
		url string
		b   []byte
		err error
	)

	url = c.apiURL + fmt.Sprintf("/hearthstone/cardbacks?locale=%s", c.locale)

	if cardBackCategory != "" {
		url = url + fmt.Sprintf("?cardBackCategory=%s", cardBackCategory)
	}

	if textFilter != "" {
		url = url + fmt.Sprintf("&type=%s", textFilter)
	}

	url = url + fmt.Sprintf("&sort=%s", sort)
	url = url + fmt.Sprintf("&order=%s", order)

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

// HSCardBackByIDOrSlug returns a specific card back by using card back ID or slug.
func (c *Client) HSCardBackByIDOrSlug(idOrSlug string) (*hsgd.CardBack, []byte, error) {
	var (
		dat hsgd.CardBack
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/cardbacks/%s?locale=%s", idOrSlug, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeSets, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeSetGroups, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeTypes, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeRarities, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeClasses, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeMinionTypes, c.locale), "")
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/hearthstone/metadata/%s?locale=%s", hsgd.MetadataTypeKeywords, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
