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

func (c *Client) ClassicWoWEraItem(ctx context.Context, itemID int) (*wowcgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassic1xNamespace(),
		&wowcgd.Item{},
	)
	return dat.(*wowcgd.Item), header, err
}

func (c *Client) ClassicWoWAnniversaryItem(ctx context.Context, itemID int) (*wowcgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-class/index",
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", playableClassID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-race/index",
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/power-type/index",
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPSeasonsIndex(ctx context.Context) (*wowcgd.PvPSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-season/index",
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPSeason(ctx context.Context, pvpSeasonID int) (*wowcgd.PvPSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d", pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPRegionIndex(ctx context.Context) (*wowcgd.PvPRegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/pvp-region/index",
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPRegionalSeasonIndex(ctx context.Context, pvpRegionID int) (*wowcgd.PvPRegionalSeasonIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/index", pvpRegionID),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPRegionalSeason(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRegionSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPLeaderboardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPLeaderboardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPLeaderboards(ctx context.Context, pvpRegionID, pvpSeasonID int, pvpBracket wow.Bracket) (*wowcgd.PvPLeaderboards, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-leaderboard/%s", pvpRegionID, pvpSeasonID, pvpBracket),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryPvPRewardsIndex(ctx context.Context, pvpRegionID, pvpSeasonID int) (*wowcgd.PvPRewardsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/pvp-region/%d/pvp-season/%d/pvp-reward/index", pvpRegionID, pvpSeasonID),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/realm/index",
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryRealmSearch(ctx context.Context, opts ...wowsearch.Opt) (*wowcgd.RealmSearch, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/search/realm%s", buildSearchParams(opts...)),
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/region/index",
		c.GetDynamicClassicAnnNamespace(),
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

func (c *Client) ClassicWoWAnniversaryRegion(ctx context.Context, regionID int) (*wowcgd.Region, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicClassicAnnNamespace(),
		&wowcgd.Region{},
	)
	return dat.(*wowcgd.Region), header, err
}
