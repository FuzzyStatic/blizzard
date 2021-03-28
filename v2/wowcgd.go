package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/wowcgd"
)

// ClassicWoWConnectedRealmsIndex returns an index of connected realms.
func (c *Client) ClassicWoWConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/connected-realm/index",
		c.GetDynamicClassicNamespace(),
		&wowcgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowcgd.ConnectedRealmsIndex), header, err
}

// ClassicWoWConnectedRealm returns a connected realm by ID.
func (c *Client) ClassicWoWConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassicNamespace(),
		&wowcgd.ConnectedRealm{},
	)
	return dat.(*wowcgd.ConnectedRealm), header, err
}

// ClassicWoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) ClassicWoWCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-family/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowcgd.CreatureFamiliesIndex), header, err
}

// ClassicWoWCreatureFamily returns a creature family by ID.
func (c *Client) ClassicWoWCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamily{},
	)
	return dat.(*wowcgd.CreatureFamily), header, err
}

// ClassicWoWCreatureTypesIndex returns an index of creature types.
func (c *Client) ClassicWoWCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/creature-type/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureTypesIndex{},
	)
	return dat.(*wowcgd.CreatureTypesIndex), header, err
}

// ClassicWoWCreatureType returns a creature type by ID.
func (c *Client) ClassicWoWCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureType{},
	)
	return dat.(*wowcgd.CreatureType), header, err
}

// ClassicWoWCreature returns a creature type by ID.
func (c *Client) ClassicWoWCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassicNamespace(),
		&wowcgd.Creature{},
	)
	return dat.(*wowcgd.Creature), header, err
}

// ClassicWoWCreatureDisplayMedia returns media for a creature display by ID.
func (c *Client) ClassicWoWCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureDisplayMedia{},
	)
	return dat.(*wowcgd.CreatureDisplayMedia), header, err
}

// ClassicWoWCreatureFamilyMedia returns media for a creature family by ID.
func (c *Client) ClassicWoWCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamilyMedia{},
	)
	return dat.(*wowcgd.CreatureFamilyMedia), header, err
}

// ClassicWoWGuildCrestComponentsIndex returns an index of guild crest media
func (c *Client) ClassicWoWGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/guild-crest/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowcgd.GuildCrestComponentsIndex), header, err
}

// ClassicWoWGuildCrestBorderMedia returns media for a guild crest border by ID
func (c *Client) ClassicWoWGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowcgd.GuildCrestBorderMedia), header, err
}

// ClassicWoWGuildCrestEmblemMedia returns media for a guild crest emblem by ID
func (c *Client) ClassicWoWGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowcgd.GuildCrestEmblemMedia), header, err
}

// ClassicWoWItemClassesIndex returns an index of item classes
func (c *Client) ClassicWoWItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/item-class/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemClassesIndex{},
	)
	return dat.(*wowcgd.ItemClassesIndex), header, err
}

// ClassicWoWItemClass returns an item class by ID
func (c *Client) ClassicWoWItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemClass{},
	)
	return dat.(*wowcgd.ItemClass), header, err
}

// ClassicWoWItemSubclass returns an item subclass by ID
func (c *Client) ClassicWoWItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemSubclass{},
	)
	return dat.(*wowcgd.ItemSubclass), header, err
}

// ClassicWoWItem returns an item by ID
func (c *Client) ClassicWoWItem(ctx context.Context, itemID int) (*wowcgd.Item, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.Item{},
	)
	return dat.(*wowcgd.Item), header, err
}

// ClassicWoWItemMedia returns media for an item by ID
func (c *Client) ClassicWoWItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemMedia{},
	)
	return dat.(*wowcgd.ItemMedia), header, err
}

// ClassicWoWPlayableClassesIndex returns an index of playable classes
func (c *Client) ClassicWoWPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-class/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClassesIndex{},
	)
	return dat.(*wowcgd.PlayableClassesIndex), header, err
}

// ClassicWoWPlayableClass returns a playable class by ID
func (c *Client) ClassicWoWPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClass{},
	)
	return dat.(*wowcgd.PlayableClass), header, err
}

// ClassicWoWPlayableClassMedia returns media for a playable class by ID
func (c *Client) ClassicWoWPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", playableClassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClassMedia{},
	)
	return dat.(*wowcgd.PlayableClassMedia), header, err
}

// ClassicWoWPlayableRacesIndex returns an index of playable races
func (c *Client) ClassicWoWPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/playable-race/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableRacesIndex{},
	)
	return dat.(*wowcgd.PlayableRacesIndex), header, err
}

// ClassicWoWPlayableRace returns a playable race by ID
func (c *Client) ClassicWoWPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableRace{},
	)
	return dat.(*wowcgd.PlayableRace), header, err
}

// ClassicWoWPowerTypesIndex returns an index of power types
func (c *Client) ClassicWoWPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/power-type/index",
		c.GetStaticClassicNamespace(),
		&wowcgd.PowerTypesIndex{},
	)
	return dat.(*wowcgd.PowerTypesIndex), header, err
}

// ClassicWoWPowerType returns a power type by ID
func (c *Client) ClassicWoWPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PowerType{},
	)
	return dat.(*wowcgd.PowerType), header, err
}

// ClassicWoWRealmIndex returns an index of realms.
func (c *Client) ClassicWoWRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/realm/index",
		c.GetDynamicClassicNamespace(),
		&wowcgd.RealmIndex{},
	)
	return dat.(*wowcgd.RealmIndex), header, err
}

// ClassicWoWRealm returns a single realm by slug or ID.
func (c *Client) ClassicWoWRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicClassicNamespace(),
		&wowcgd.Realm{},
	)
	return dat.(*wowcgd.Realm), header, err
}

// ClassicWoWRegionIndex returns an index of regions.
func (c *Client) ClassicWoWRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		"/data/wow/region/index",
		c.GetDynamicClassicNamespace(),
		&wowcgd.RegionIndex{},
	)
	return dat.(*wowcgd.RegionIndex), header, err
}

// ClassicWoWRegion returns a region by ID.
func (c *Client) ClassicWoWRegion(ctx context.Context, regionID int) (*wowcgd.Region, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicClassicNamespace(),
		&wowcgd.Region{},
	)
	return dat.(*wowcgd.Region), header, err
}
