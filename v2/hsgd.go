package blizzard

import (
	"context"
	"fmt"
	"strings"

	"github.com/FuzzyStatic/blizzard/v2/hsgd"
)

// HSCardsSearch returns an up-to-date list of all cards matching the search criteria.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardsSearch(ctx context.Context) (*hsgd.CardSearch, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/cards"),
		&hsgd.CardSearch{},
	)
	return dat.(*hsgd.CardSearch), b, err
}

// HSDetailedCardsSearch returns an up-to-date list of all cards matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSDetailedCardsSearch(ctx context.Context, setSlug, classSlug, raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	collectiblility hsgd.Collectibility, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardSearch, []byte, error) {
	pathAndQuery := fmt.Sprintf("/hearthstone/cards")
	pathAndQuery = pathAndQuery + fmt.Sprintf("?collectible=%s", collectiblility)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&sort=%s", sort)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&order=%s", order)
	if setSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&set=%s", setSlug)
	}
	if classSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&class=%s", classSlug)
	}
	if raritySlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&rarity=%s", raritySlug)
	}
	if typeSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&type=%s", typeSlug)
	}
	if minionTypeSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&minionType=%s", minionTypeSlug)
	}
	if keywordSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&keyword=%s", keywordSlug)
	}
	if textFilter != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&textFilter=%s", textFilter)
	}
	if manaCost != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&manaCost=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]"))
	}
	if attack != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&attack=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]"))
	}
	if health != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&health=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]"))
	}
	if page != 0 {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&page=%d", page)
	}
	if pageSize != 0 {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&pageSize=%d", pageSize)
	}

	dat, b, err := c.getStructDataNoNamespace(ctx,
		pathAndQuery,
		&hsgd.CardSearch{},
	)
	return dat.(*hsgd.CardSearch), b, err
}

// HSBattlegroundsCardsSearch returns an up-to-date list of all cards matching the search criteria for the specified game mode.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSBattlegroundsCardsSearch(ctx context.Context, raritySlug, typeSlug, minionTypeSlug, keywordSlug, textFilter string,
	manaCost, attack, health []int, page, pageSize int,
	tier []hsgd.Tier, collectiblility hsgd.Collectibility, sort hsgd.Sort, order hsgd.Order) (*hsgd.CardSearch, []byte, error) {
	pathAndQuery := fmt.Sprintf("/hearthstone/cards?gameMode=%s", hsgd.GameModeBattlegrounds)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&collectible=%s", collectiblility)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&sort=%s", sort)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&order=%s", order)
	if raritySlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&rarity=%s", raritySlug)
	}
	if typeSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&type=%s", typeSlug)
	}
	if minionTypeSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&minionType=%s", minionTypeSlug)
	}
	if keywordSlug != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&keyword=%s", keywordSlug)
	}
	if textFilter != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&textFilter=%s", textFilter)
	}
	if manaCost != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&manaCost=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(manaCost)), ","), "[]"))
	}
	if attack != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&attack=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(attack)), ","), "[]"))
	}
	if health != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&health=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(health)), ","), "[]"))
	}
	if page != 0 {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&page=%d", page)
	}
	if pageSize != 0 {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&pageSize=%d", pageSize)
	}
	if health != nil {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&tier=%s", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tier)), ","), "[]"))
	}

	dat, b, err := c.getStructDataNoNamespace(ctx,
		pathAndQuery,
		&hsgd.CardSearch{},
	)
	return dat.(*hsgd.CardSearch), b, err
}

// HSCardByIDOrSlug returns card by ID or slug.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardByIDOrSlug(ctx context.Context, idOrSlug string, gameMode hsgd.GameMode) (*hsgd.Card, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/cards/%s?gameMode=%s", idOrSlug, gameMode),
		&hsgd.Card{},
	)
	return dat.(*hsgd.Card), b, err
}

// HSCardBackSearchAllLocales returns an up-to-date list of all card backs matching the search criteria for all locales.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardBackSearchAllLocales(ctx context.Context, cardBackCategory hsgd.CardBackCategory, textFilter string,
	sort hsgd.Sort, order hsgd.Order) (*hsgd.CardBackSearchAllLocales, []byte, error) {
	pathAndQuery := "/hearthstone/cardbacks"
	pathAndQuery = pathAndQuery + fmt.Sprintf("?sort=%s", sort)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&order=%s", order)
	if cardBackCategory != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&cardBackCategory=%s", cardBackCategory)
	}
	if textFilter != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&type=%s", textFilter)
	}

	dat, b, err := c.getStructDataNoNamespaceNoLocale(ctx,
		pathAndQuery,
		&hsgd.CardBackSearchAllLocales{},
	)
	return dat.(*hsgd.CardBackSearchAllLocales), b, err
}

// HSCardBackSearch returns an up-to-date list of all card backs matching the search criteria.
// Input values left blank, 0, or nil will return all values for the type.
// For more information about the search parameters, see the Card Search Guide (https://develop.battle.net/documentation/hearthstone/guides/card-search).
func (c *Client) HSCardBackSearch(ctx context.Context, cardBackCategory hsgd.CardBackCategory, textFilter string,
	sort hsgd.Sort, order hsgd.Order) (*hsgd.CardBackSearch, []byte, error) {
	pathAndQuery := fmt.Sprintf("/hearthstone/cardbacks")
	pathAndQuery = pathAndQuery + fmt.Sprintf("?sort=%s", sort)
	pathAndQuery = pathAndQuery + fmt.Sprintf("&order=%s", order)
	if cardBackCategory != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&cardBackCategory=%s", cardBackCategory)
	}
	if textFilter != "" {
		pathAndQuery = pathAndQuery + fmt.Sprintf("&type=%s", textFilter)
	}

	dat, b, err := c.getStructDataNoNamespace(ctx,
		pathAndQuery,
		&hsgd.CardBackSearch{},
	)
	return dat.(*hsgd.CardBackSearch), b, err
}

// HSCardBackByIDOrSlug returns a specific card back by using card back ID or slug.
func (c *Client) HSCardBackByIDOrSlug(ctx context.Context, idOrSlug string) (*hsgd.CardBack, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/cardbacks/%s", idOrSlug),
		&hsgd.CardBack{},
	)
	return dat.(*hsgd.CardBack), b, err
}

// HSDeck Finds a deck by its deck code.
// For more information, see the Hearthstone Guide.
func (c *Client) HSDeck(ctx context.Context, deckCode string) (*hsgd.Deck, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/deck/%s", deckCode),
		&hsgd.Deck{},
	)
	return dat.(*hsgd.Deck), b, err
}

// HSMetadata returns information about the categorization of cards.
// Metadata includes the card set, set group (for example, Standard or Year of the Dragon), rarity, class, card type, minion type, and keywords.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadata(ctx context.Context) (*hsgd.Metadata, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata"),
		&hsgd.Metadata{},
	)
	return dat.(*hsgd.Metadata), b, err
}

// HSMetadataSets returns information about set metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataSets(ctx context.Context) (*[]hsgd.Set, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeSets),
		&[]hsgd.Set{},
	)
	return dat.(*[]hsgd.Set), b, err
}

// HSMetadataSetGroups returns information about set group metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataSetGroups(ctx context.Context) (*[]hsgd.SetGroup, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeSetGroups),
		&[]hsgd.SetGroup{},
	)
	return dat.(*[]hsgd.SetGroup), b, err
}

// HSMetadataTypes returns information about type metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataTypes(ctx context.Context) (*[]hsgd.Type, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeTypes),
		&[]hsgd.Type{},
	)
	return dat.(*[]hsgd.Type), b, err
}

// HSMetadataRarities returns information about rarity metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataRarities(ctx context.Context) (*[]hsgd.Rarity, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeRarities),
		&[]hsgd.Rarity{},
	)
	return dat.(*[]hsgd.Rarity), b, err
}

// HSMetadataClasses returns information about class metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataClasses(ctx context.Context) (*[]hsgd.Class, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeClasses),
		&[]hsgd.Class{},
	)
	return dat.(*[]hsgd.Class), b, err
}

// HSMetadataMinionTypes returns information about minion type metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataMinionTypes(ctx context.Context) (*[]hsgd.MinionType, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeMinionTypes),
		&[]hsgd.MinionType{},
	)
	return dat.(*[]hsgd.MinionType), b, err
}

// HSMetadataKeywords returns information about keyword metadata.
// For more information, see the Hearthstone Guide.
func (c *Client) HSMetadataKeywords(ctx context.Context) (*[]hsgd.Keyword, []byte, error) {
	dat, b, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/hearthstone/metadata/%s", hsgd.MetadataTypeKeywords),
		&[]hsgd.Keyword{},
	)
	return dat.(*[]hsgd.Keyword), b, err
}
