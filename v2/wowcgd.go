package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/wowcgd"
)

// ClassicWoWConnectedRealmsIndex returns an index of connected realms.
func (c *Client) ClassicWoWConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/index"),
		c.GetDynamicClassicNamespace(),
		&wowcgd.ConnectedRealmsIndex{},
	)
	return dat.(*wowcgd.ConnectedRealmsIndex), b, err
}

// ClassicWoWConnectedRealm returns a connected realm by ID.
func (c *Client) ClassicWoWConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/connected-realm/%d", connectedRealmID),
		c.GetDynamicClassicNamespace(),
		&wowcgd.ConnectedRealm{},
	)
	return dat.(*wowcgd.ConnectedRealm), b, err
}

// ClassicWoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) ClassicWoWCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamiliesIndex{},
	)
	return dat.(*wowcgd.CreatureFamiliesIndex), b, err
}

// ClassicWoWCreatureFamily returns a creature family by ID.
func (c *Client) ClassicWoWCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamily{},
	)
	return dat.(*wowcgd.CreatureFamily), b, err
}

// ClassicWoWCreatureTypesIndex returns an index of creature types.
func (c *Client) ClassicWoWCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureTypesIndex{},
	)
	return dat.(*wowcgd.CreatureTypesIndex), b, err
}

// ClassicWoWCreatureType returns a creature type by ID.
func (c *Client) ClassicWoWCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature-type/%d", creatureTypeID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureType{},
	)
	return dat.(*wowcgd.CreatureType), b, err
}

// ClassicWoWCreature returns a creature type by ID.
func (c *Client) ClassicWoWCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/creature/%d", creatureID),
		c.GetStaticClassicNamespace(),
		&wowcgd.Creature{},
	)
	return dat.(*wowcgd.Creature), b, err
}

// ClassicWoWCreatureDisplayMedia returns media for a creature display by ID.
func (c *Client) ClassicWoWCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-display/%d", creatureDisplayID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureDisplayMedia{},
	)
	return dat.(*wowcgd.CreatureDisplayMedia), b, err
}

// ClassicWoWCreatureFamilyMedia returns media for a creature family by ID.
func (c *Client) ClassicWoWCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/creature-family/%d", creatureFamilyID),
		c.GetStaticClassicNamespace(),
		&wowcgd.CreatureFamilyMedia{},
	)
	return dat.(*wowcgd.CreatureFamilyMedia), b, err
}

// ClassicWoWGuildCrestComponentsIndex returns an index of guild crest media
func (c *Client) ClassicWoWGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/guild-crest/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestComponentsIndex{},
	)
	return dat.(*wowcgd.GuildCrestComponentsIndex), b, err
}

// ClassicWoWGuildCrestBorderMedia returns media for a guild crest border by ID
func (c *Client) ClassicWoWGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/border/%d", borderID),
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestBorderMedia{},
	)
	return dat.(*wowcgd.GuildCrestBorderMedia), b, err
}

// ClassicWoWGuildCrestEmblemMedia returns media for a guild crest emblem by ID
func (c *Client) ClassicWoWGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d", emblemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.GuildCrestEmblemMedia{},
	)
	return dat.(*wowcgd.GuildCrestEmblemMedia), b, err
}

// ClassicWoWItemClassesIndex returns an index of item classes
func (c *Client) ClassicWoWItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemClassesIndex{},
	)
	return dat.(*wowcgd.ItemClassesIndex), b, err
}

// ClassicWoWItemClass returns an item class by ID
func (c *Client) ClassicWoWItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d", itemClassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemClass{},
	)
	return dat.(*wowcgd.ItemClass), b, err
}

// ClassicWoWItemSubclass returns an item subclass by ID
func (c *Client) ClassicWoWItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d", itemClassID, itemSubclassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemSubclass{},
	)
	return dat.(*wowcgd.ItemSubclass), b, err
}

// ClassicWoWItem returns an item by ID
func (c *Client) ClassicWoWItem(ctx context.Context, itemID int) (*wowcgd.Item, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/item/%d", itemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.Item{},
	)
	return dat.(*wowcgd.Item), b, err
}

// ClassicWoWItemMedia returns media for an item by ID
func (c *Client) ClassicWoWItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/item/%d", itemID),
		c.GetStaticClassicNamespace(),
		&wowcgd.ItemMedia{},
	)
	return dat.(*wowcgd.ItemMedia), b, err
}

// ClassicWoWPlayableClassesIndex returns an index of playable classes
func (c *Client) ClassicWoWPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClassesIndex{},
	)
	return dat.(*wowcgd.PlayableClassesIndex), b, err
}

// ClassicWoWPlayableClass returns a playable class by ID
func (c *Client) ClassicWoWPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-class/%d", classID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClass{},
	)
	return dat.(*wowcgd.PlayableClass), b, err
}

// ClassicWoWPlayableClassMedia returns media for a playable class by ID
func (c *Client) ClassicWoWPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/media/playable-class/%d", playableClassID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableClassMedia{},
	)
	return dat.(*wowcgd.PlayableClassMedia), b, err
}

// ClassicWoWPlayableRacesIndex returns an index of playable races
func (c *Client) ClassicWoWPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableRacesIndex{},
	)
	return dat.(*wowcgd.PlayableRacesIndex), b, err
}

// ClassicWoWPlayableRace returns a playable race by ID
func (c *Client) ClassicWoWPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/playable-race/%d", raceID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PlayableRace{},
	)
	return dat.(*wowcgd.PlayableRace), b, err
}

// ClassicWoWPowerTypesIndex returns an index of power types
func (c *Client) ClassicWoWPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/index"),
		c.GetStaticClassicNamespace(),
		&wowcgd.PowerTypesIndex{},
	)
	return dat.(*wowcgd.PowerTypesIndex), b, err
}

// ClassicWoWPowerType returns a power type by ID
func (c *Client) ClassicWoWPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/power-type/%d", powerTypeID),
		c.GetStaticClassicNamespace(),
		&wowcgd.PowerType{},
	)
	return dat.(*wowcgd.PowerType), b, err
}

// ClassicWoWRealmIndex returns an index of realms.
func (c *Client) ClassicWoWRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/index"),
		c.GetDynamicClassicNamespace(),
		&wowcgd.RealmIndex{},
	)
	return dat.(*wowcgd.RealmIndex), b, err
}

// ClassicWoWRealm returns a single realm by slug or ID.
func (c *Client) ClassicWoWRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/realm/%s", realmSlug),
		c.GetDynamicClassicNamespace(),
		&wowcgd.Realm{},
	)
	return dat.(*wowcgd.Realm), b, err
}

// ClassicWoWRegionIndex returns an index of regions.
func (c *Client) ClassicWoWRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/index"),
		c.GetDynamicClassicNamespace(),
		&wowcgd.RegionIndex{},
	)
	return dat.(*wowcgd.RegionIndex), b, err
}

// ClassicWoWRegion returns a region by ID.
func (c *Client) ClassicWoWRegion(ctx context.Context, regionID int) (*wowcgd.Region, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/region/%d", regionID),
		c.GetDynamicClassicNamespace(),
		&wowcgd.Region{},
	)
	return dat.(*wowcgd.Region), b, err
}
