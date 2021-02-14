package blizzard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/wowcgd"
)

// ClassicWoWConnectedRealmsIndex returns an index of connected realms.
func (c *Client) ClassicWoWConnectedRealmsIndex(ctx context.Context) (*wowcgd.ConnectedRealmsIndex, []byte, error) {
	var (
		dat wowcgd.ConnectedRealmsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/index?locale=%s",
		c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWConnectedRealm returns a connected realm by ID.
func (c *Client) ClassicWoWConnectedRealm(ctx context.Context, connectedRealmID int) (*wowcgd.ConnectedRealm, []byte, error) {
	var (
		dat wowcgd.ConnectedRealm
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d?locale=%s",
		connectedRealmID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) ClassicWoWCreatureFamiliesIndex(ctx context.Context) (*wowcgd.CreatureFamiliesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureFamiliesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-family/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureFamily returns a creature family by ID.
func (c *Client) ClassicWoWCreatureFamily(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamily, []byte, error) {
	var (
		dat wowcgd.CreatureFamily
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureTypesIndex returns an index of creature types.
func (c *Client) ClassicWoWCreatureTypesIndex(ctx context.Context) (*wowcgd.CreatureTypesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-type/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureType returns a creature type by ID.
func (c *Client) ClassicWoWCreatureType(ctx context.Context, creatureTypeID int) (*wowcgd.CreatureType, []byte, error) {
	var (
		dat wowcgd.CreatureType
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature-type/%d?locale=%s", creatureTypeID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreature returns a creature type by ID.
func (c *Client) ClassicWoWCreature(ctx context.Context, creatureID int) (*wowcgd.Creature, []byte, error) {
	var (
		dat wowcgd.Creature
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/creature/%d?locale=%s", creatureID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureDisplayMedia returns media for a creature display by ID.
func (c *Client) ClassicWoWCreatureDisplayMedia(ctx context.Context, creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, []byte, error) {
	var (
		dat wowcgd.CreatureDisplayMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/creature-display/%d?locale=%s", creatureDisplayID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureFamilyMedia returns media for a creature family by ID.
func (c *Client) ClassicWoWCreatureFamilyMedia(ctx context.Context, creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, []byte, error) {
	var (
		dat wowcgd.CreatureFamilyMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWGuildCrestComponentsIndex returns an index of guild crest media
func (c *Client) ClassicWoWGuildCrestComponentsIndex(ctx context.Context) (*wowcgd.GuildCrestComponentsIndex, []byte, error) {
	var (
		dat wowcgd.GuildCrestComponentsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/guild-crest/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWGuildCrestBorderMedia returns media for a guild crest border by ID
func (c *Client) ClassicWoWGuildCrestBorderMedia(ctx context.Context, borderID int) (*wowcgd.GuildCrestBorderMdedia, []byte, error) {
	var (
		dat wowcgd.GuildCrestBorderMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/border/%d?locale=%s", borderID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWGuildCrestEmblemMedia returns media for a guild crest emblem by ID
func (c *Client) ClassicWoWGuildCrestEmblemMedia(ctx context.Context, emblemID int) (*wowcgd.GuildCrestEmblemMdedia, []byte, error) {
	var (
		dat wowcgd.GuildCrestEmblemMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d?locale=%s", emblemID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWItemClassesIndex returns an index of item classes
func (c *Client) ClassicWoWItemClassesIndex(ctx context.Context) (*wowcgd.ItemClassesIndex, []byte, error) {
	var (
		dat wowcgd.ItemClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item-class/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWItemClass returns an item class by ID
func (c *Client) ClassicWoWItemClass(ctx context.Context, itemClassID int) (*wowcgd.ItemClass, []byte, error) {
	var (
		dat wowcgd.ItemClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item-class/%d?locale=%s", itemClassID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWItemSubclass returns an item subclass by ID
func (c *Client) ClassicWoWItemSubclass(ctx context.Context, itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, []byte, error) {
	var (
		dat wowcgd.ItemSubclass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx,
		c.apiURL+fmt.Sprintf("/data/wow/item-class/%d/item-subclass/%d?locale=%s", itemClassID, itemSubclassID, c.locale),
		c.staticClassicNamespace,
	)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWItem returns an item by ID
func (c *Client) ClassicWoWItem(ctx context.Context, itemID int) (*wowcgd.Item, []byte, error) {
	var (
		dat wowcgd.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/item/%d?locale=%s", itemID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWItemMedia returns media for an item by ID
func (c *Client) ClassicWoWItemMedia(ctx context.Context, itemID int) (*wowcgd.ItemMedia, []byte, error) {
	var (
		dat wowcgd.ItemMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/item/%d?locale=%s", itemID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPlayableClassesIndex returns an index of playable classes
func (c *Client) ClassicWoWPlayableClassesIndex(ctx context.Context) (*wowcgd.PlayableClassesIndex, []byte, error) {
	var (
		dat wowcgd.PlayableClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-class/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPlayableClass returns a playable class by ID
func (c *Client) ClassicWoWPlayableClass(ctx context.Context, classID int) (*wowcgd.PlayableClass, []byte, error) {
	var (
		dat wowcgd.PlayableClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d?locale=%s", classID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPlayableClassMedia returns media for a playable class by ID
func (c *Client) ClassicWoWPlayableClassMedia(ctx context.Context, playableClassID int) (*wowcgd.PlayableClassMedia, []byte, error) {
	var (
		dat wowcgd.PlayableClassMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/media/playable-class/%d?locale=%s", playableClassID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPlayableRacesIndex returns an index of playable races
func (c *Client) ClassicWoWPlayableRacesIndex(ctx context.Context) (*wowcgd.PlayableRacesIndex, []byte, error) {
	var (
		dat wowcgd.PlayableRacesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-race/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPlayableRace returns a playable race by ID
func (c *Client) ClassicWoWPlayableRace(ctx context.Context, raceID int) (*wowcgd.PlayableRace, []byte, error) {
	var (
		dat wowcgd.PlayableRace
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/playable-race/%d?locale=%s", raceID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPowerTypesIndex returns an index of power types
func (c *Client) ClassicWoWPowerTypesIndex(ctx context.Context) (*wowcgd.PowerTypesIndex, []byte, error) {
	var (
		dat wowcgd.PowerTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/power-type/index?locale=%s", c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWPowerType returns a power type by ID
func (c *Client) ClassicWoWPowerType(ctx context.Context, powerTypeID int) (*wowcgd.PowerType, []byte, error) {
	var (
		dat wowcgd.PowerType
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/power-type/%d?locale=%s", powerTypeID, c.locale), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWRealmIndex returns an index of realms.
func (c *Client) ClassicWoWRealmIndex(ctx context.Context) (*wowcgd.RealmIndex, []byte, error) {
	var (
		dat wowcgd.RealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/realm/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWRealm returns a single realm by slug or ID.
func (c *Client) ClassicWoWRealm(ctx context.Context, realmSlug string) (*wowcgd.Realm, []byte, error) {
	var (
		dat wowcgd.Realm
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/realm/%s?locale=%s", realmSlug, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWRegionIndex returns an index of regions.
func (c *Client) ClassicWoWRegionIndex(ctx context.Context) (*wowcgd.RegionIndex, []byte, error) {
	var (
		dat wowcgd.RegionIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/region/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWRegion returns a region by ID.
func (c *Client) ClassicWoWRegion(ctx context.Context, regionID int) (*wowcgd.Region, []byte, error) {
	var (
		dat wowcgd.Region
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/wow/region/%d?locale=%s",
		regionID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
