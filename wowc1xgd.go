package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v3/wowcgd"
	"github.com/FuzzyStatic/blizzard/v3/wowsearch"
)

func (c *Client) ClassicWoWEraAuctionHouseIndex(ctx context.Context, connectedRealmID int) (*wowcgd.AuctionHouseIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/index", connectedRealmID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.AuctionHouseIndex{},
	)
	return dat.(*wowcgd.AuctionHouseIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryAuctionHouseIndex(ctx context.Context, connectedRealmID int) (*wowcgd.AuctionHouseIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/index", connectedRealmID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.AuctionHouseIndex{},
	)
	return dat.(*wowcgd.AuctionHouseIndex), header, err
}

func (c *Client) ClassicWoWEraAuctions(ctx context.Context, connectedRealmID, auctionHouseID int) (*wowcgd.Auctions, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/%d", connectedRealmID, auctionHouseID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.Auctions{},
	)
	return dat.(*wowcgd.Auctions), header, err
}

func (c *Client) ClassicWoWAnniversaryAuctions(ctx context.Context, connectedRealmID, auctionHouseID int) (*wowcgd.Auctions, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/%d", connectedRealmID, auctionHouseID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.Auctions{},
	)
	return dat.(*wowcgd.Auctions), header, err
}

func (c *Client) ClassicWoWEraConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/connected-realm/index",
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowcgd.ConnectedRealmsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/connected-realm/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowcgd.ConnectedRealmsIndex), header, err
}

func (c *Client) ClassicWoWEraConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.ConnectedRealm{},
	)
	return dat.(*wowcgd.ConnectedRealm), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealm{},
	)
	return dat.(*wowcgd.ConnectedRealm), header, err
}

func (c *Client) ClassicWoWEraConnectedRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.ConnectedRealmsSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/connected-realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.ConnectedRealmsSearch{},
	)
	return dat.(*wowcgd.ConnectedRealmsSearch), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.ConnectedRealmsSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/connected-realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealmsSearch{},
	)
	return dat.(*wowcgd.ConnectedRealmsSearch), header, err
}

func (c *Client) ClassicWoWEraCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowcgd.CreatureFamiliesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowcgd.CreatureFamiliesIndex), header, err
}

func (c *Client) ClassicWoWEraCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureFamily{},
	)
	return dat.(*wowcgd.CreatureFamily), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamily{},
	)
	return dat.(*wowcgd.CreatureFamily), header, err
}

func (c *Client) ClassicWoWEraCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureTypesIndex{},
	)
	return dat.(*wowcgd.CreatureTypesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureTypesIndex{},
	)
	return dat.(*wowcgd.CreatureTypesIndex), header, err
}

func (c *Client) ClassicWoWEraCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureType{},
	)
	return dat.(*wowcgd.CreatureType), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureType{},
	)
	return dat.(*wowcgd.CreatureType), header, err
}

func (c *Client) ClassicWoWEraCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.Creature{},
	)
	return dat.(*wowcgd.Creature), header, err
}

func (c *Client) ClassicWoWAnniversaryCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.Creature{},
	)
	return dat.(*wowcgd.Creature), header, err
}

func (c *Client) ClassicWoWEraCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureDisplayMedia{},
	)
	return dat.(*wowcgd.CreatureDisplayMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureDisplayMedia{},
	)
	return dat.(*wowcgd.CreatureDisplayMedia), header, err
}

func (c *Client) ClassicWoWEraCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.CreatureFamilyMedia{},
	)
	return dat.(*wowcgd.CreatureFamilyMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamilyMedia{},
	)
	return dat.(*wowcgd.CreatureFamilyMedia), header, err
}

func (c *Client) ClassicWoWEraGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/guild-crest/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowcgd.GuildCrestComponentsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/guild-crest/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowcgd.GuildCrestComponentsIndex), header, err
}

func (c *Client) ClassicWoWEraGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowcgd.GuildCrestBorderMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowcgd.GuildCrestBorderMedia), header, err
}

func (c *Client) ClassicWoWEraGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowcgd.GuildCrestEmblemMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowcgd.GuildCrestEmblemMedia), header, err
}

func (c *Client) ClassicWoWEraItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.ItemClassesIndex{},
	)
	return dat.(*wowcgd.ItemClassesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemClassesIndex{},
	)
	return dat.(*wowcgd.ItemClassesIndex), header, err
}

func (c *Client) ClassicWoWEraItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.ItemClass{},
	)
	return dat.(*wowcgd.ItemClass), header, err
}

func (c *Client) ClassicWoWAnniversaryItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemClass{},
	)
	return dat.(*wowcgd.ItemClass), header, err
}

func (c *Client) ClassicWoWEraItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.ItemSubclass{},
	)
	return dat.(*wowcgd.ItemSubclass), header, err
}
