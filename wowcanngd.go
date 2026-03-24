package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v3/wow"
	"github.com/FuzzyStatic/blizzard/v3/wowcgd"
	"github.com/FuzzyStatic/blizzard/v3/wowsearch"
)

func (c *Client) ClassicWoWAnniversaryItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemSubclass{},
	)
	return dat.(*wowcgd.ItemSubclass), header, err
}

func (c *Client) ClassicWoWAnniversaryItem(ctx context.Context, itemID int) (*wowcgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.Item{},
	)
	return dat.(*wowcgd.Item), header, err
}

func (c *Client) ClassicWoWAnniversaryItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemMedia{},
	)
	return dat.(*wowcgd.ItemMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-class/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PlayableClassesIndex{},
	)
	return dat.(*wowcgd.PlayableClassesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PlayableClass{},
	)
	return dat.(*wowcgd.PlayableClass), header, err
}

func (c *Client) ClassicWoWAnniversaryPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", playableClassID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PlayableClassMedia{},
	)
	return dat.(*wowcgd.PlayableClassMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-race/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PlayableRacesIndex{},
	)
	return dat.(*wowcgd.PlayableRacesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PlayableRace{},
	)
	return dat.(*wowcgd.PlayableRace), header, err
}

func (c *Client) ClassicWoWAnniversaryPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/power-type/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PowerTypesIndex{},
	)
	return dat.(*wowcgd.PowerTypesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.PowerType{},
	)
	return dat.(*wowcgd.PowerType), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPSeasonsIndex(ctx context.Context) (*wowcgd.PvPSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-season/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPSeasonIndex{},
	)
	return dat.(*wowcgd.PvPSeasonIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPSeason(ctx context.Context, pvpSeasonID int) (*wowcgd.PvPSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d", pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPSeason{},
	)
	return dat.(*wowcgd.PvPSeason), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPRegionIndex(ctx context.Context) (*wowcgd.PvPRegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-region/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPRegionIndex{},
	)
	return dat.(*wowcgd.PvPRegionIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPRegionalSeasonIndex(ctx context.Context, pvpRegionID int) (*wowcgd.PvPRegionalSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/index", pvpRegionID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPRegionalSeasonIndex{},
	)
	return dat.(*wowcgd.PvPRegionalSeasonIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPRegionalSeason(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRegionSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPRegionSeason{},
	)
	return dat.(*wowcgd.PvPRegionSeason), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPLeaderboardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPLeaderboardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPLeaderboardsIndex{},
	)
	return dat.(*wowcgd.PvPLeaderboardsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPLeaderboards(ctx context.Context, pvpRegionID, pvpSeasonID int, pvpBracket wow.Bracket) (*wowcgd.PvPLeaderboards, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/%s", pvpRegionID, pvpSeasonID, pvpBracket),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPLeaderboards{},
	)
	return dat.(*wowcgd.PvPLeaderboards), header, err
}

func (c *Client) ClassicWoWAnniversaryPvPRewardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRewardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-reward/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.PvPRewardsIndex{},
	)
	return dat.(*wowcgd.PvPRewardsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/realm/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.RealmIndex{},
	)
	return dat.(*wowcgd.RealmIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.Realm{},
	)
	return dat.(*wowcgd.Realm), header, err
}

func (c *Client) ClassicWoWAnniversaryRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.RealmSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.RealmSearch{},
	)
	return dat.(*wowcgd.RealmSearch), header, err
}

func (c *Client) ClassicWoWAnniversaryRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/region/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.RegionIndex{},
	)
	return dat.(*wowcgd.RegionIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryRegion(ctx context.Context, regionID int) (*wowcgd.Region, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.Region{},
	)
	return dat.(*wowcgd.Region), header, err
}

func (c *Client) ClassicWoWAnniversaryAuctionHouseIndex(ctx context.Context, connectedRealmID int) (*wowcgd.AuctionHouseIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/index", connectedRealmID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.AuctionHouseIndex{},
	)
	return dat.(*wowcgd.AuctionHouseIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryAuctions(ctx context.Context, connectedRealmID, auctionHouseID int) (*wowcgd.Auctions, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d/auctions/%d", connectedRealmID, auctionHouseID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.Auctions{},
	)
	return dat.(*wowcgd.Auctions), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/connected-realm/index",
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowcgd.ConnectedRealmsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealm{},
	)
	return dat.(*wowcgd.ConnectedRealm), header, err
}

func (c *Client) ClassicWoWAnniversaryConnectedRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.ConnectedRealmsSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/connected-realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.ConnectedRealmsSearch{},
	)
	return dat.(*wowcgd.ConnectedRealmsSearch), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowcgd.CreatureFamiliesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamily{},
	)
	return dat.(*wowcgd.CreatureFamily), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureTypesIndex{},
	)
	return dat.(*wowcgd.CreatureTypesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureType{},
	)
	return dat.(*wowcgd.CreatureType), header, err
}

func (c *Client) ClassicWoWAnniversaryCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.Creature{},
	)
	return dat.(*wowcgd.Creature), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureDisplayMedia{},
	)
	return dat.(*wowcgd.CreatureDisplayMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.CreatureFamilyMedia{},
	)
	return dat.(*wowcgd.CreatureFamilyMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/guild-crest/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowcgd.GuildCrestComponentsIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowcgd.GuildCrestBorderMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowcgd.GuildCrestEmblemMedia), header, err
}

func (c *Client) ClassicWoWAnniversaryItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemClassesIndex{},
	)
	return dat.(*wowcgd.ItemClassesIndex), header, err
}

func (c *Client) ClassicWoWAnniversaryItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticClassicAnnNamespace(),
		&wowcgd.ItemClass{},
	)
	return dat.(*wowcgd.ItemClass), header, err
}
