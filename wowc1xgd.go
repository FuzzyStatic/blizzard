package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v3/wow"
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

func (c *Client) ClassicWoWEraAuctions(ctx context.Context, connectedRealmID, auctionHouseID int) (*wowcgd.Auctions, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/%d", connectedRealmID, auctionHouseID),
		c.GetDynamicClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticClassic1xNamespace(),
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

func (c *Client) ClassicWoWEraItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.ItemSubclass{},
	)
	return dat.(*wowcgd.ItemSubclass), header, err
}

func (c *Client) ClassicWoWEraItem(ctx context.Context, itemID int) (*wowcgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.Item{},
	)
	return dat.(*wowcgd.Item), header, err
}

func (c *Client) ClassicWoWEraItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.ItemMedia{},
	)
	return dat.(*wowcgd.ItemMedia), header, err
}

func (c *Client) ClassicWoWEraPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-class/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PlayableClassesIndex{},
	)
	return dat.(*wowcgd.PlayableClassesIndex), header, err
}

func (c *Client) ClassicWoWEraPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PlayableClass{},
	)
	return dat.(*wowcgd.PlayableClass), header, err
}

func (c *Client) ClassicWoWEraPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", playableClassID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PlayableClassMedia{},
	)
	return dat.(*wowcgd.PlayableClassMedia), header, err
}

func (c *Client) ClassicWoWEraPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-race/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PlayableRacesIndex{},
	)
	return dat.(*wowcgd.PlayableRacesIndex), header, err
}

func (c *Client) ClassicWoWEraPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PlayableRace{},
	)
	return dat.(*wowcgd.PlayableRace), header, err
}

func (c *Client) ClassicWoWEraPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/power-type/index",
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PowerTypesIndex{},
	)
	return dat.(*wowcgd.PowerTypesIndex), header, err
}

func (c *Client) ClassicWoWEraPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.PowerType{},
	)
	return dat.(*wowcgd.PowerType), header, err
}

func (c *Client) ClassicWoWEraPvPSeasonsIndex(ctx context.Context) (*wowcgd.PvPSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-season/index",
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPSeasonIndex{},
	)
	return dat.(*wowcgd.PvPSeasonIndex), header, err
}

func (c *Client) ClassicWoWEraPvPSeason(ctx context.Context, pvpSeasonID int) (*wowcgd.PvPSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d", pvpSeasonID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPSeason{},
	)
	return dat.(*wowcgd.PvPSeason), header, err
}

func (c *Client) ClassicWoWEraPvPRegionIndex(ctx context.Context) (*wowcgd.PvPRegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-region/index",
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPRegionIndex{},
	)
	return dat.(*wowcgd.PvPRegionIndex), header, err
}

func (c *Client) ClassicWoWEraPvPRegionalSeasonIndex(ctx context.Context, pvpRegionID int) (*wowcgd.PvPRegionalSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/index", pvpRegionID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPRegionalSeasonIndex{},
	)
	return dat.(*wowcgd.PvPRegionalSeasonIndex), header, err
}

func (c *Client) ClassicWoWEraPvPRegionalSeason(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRegionSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPRegionSeason{},
	)
	return dat.(*wowcgd.PvPRegionSeason), header, err
}

func (c *Client) ClassicWoWEraPvPLeaderboardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPLeaderboardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPLeaderboardsIndex{},
	)
	return dat.(*wowcgd.PvPLeaderboardsIndex), header, err
}

func (c *Client) ClassicWoWEraPvPLeaderboards(ctx context.Context, pvpRegionID, pvpSeasonID int, pvpBracket wow.Bracket) (*wowcgd.PvPLeaderboards, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/%s", pvpRegionID, pvpSeasonID, pvpBracket),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPLeaderboards{},
	)
	return dat.(*wowcgd.PvPLeaderboards), header, err
}

func (c *Client) ClassicWoWEraPvPRewardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRewardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-reward/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.PvPRewardsIndex{},
	)
	return dat.(*wowcgd.PvPRewardsIndex), header, err
}

func (c *Client) ClassicWoWEraRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/realm/index",
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.RealmIndex{},
	)
	return dat.(*wowcgd.RealmIndex), header, err
}

func (c *Client) ClassicWoWEraRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.Realm{},
	)
	return dat.(*wowcgd.Realm), header, err
}

func (c *Client) ClassicWoWEraRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.RealmSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.RealmSearch{},
	)
	return dat.(*wowcgd.RealmSearch), header, err
}

func (c *Client) ClassicWoWEraRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/region/index",
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.RegionIndex{},
	)
	return dat.(*wowcgd.RegionIndex), header, err
}

func (c *Client) ClassicWoWEraRegion(ctx context.Context, regionID int) (*wowcgd.Region, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicClassic1xNamespace(),
		&wowcgd.Region{},
	)
	return dat.(*wowcgd.Region), header, err
}
