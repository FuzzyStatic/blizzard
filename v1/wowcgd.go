package blizzard

import (
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/wowcgd"
	"github.com/FuzzyStatic/blizzard/wowgd"
)

// ClassicWoWCreatureFamiliesIndex returns an index of creature families.
func (c *Client) ClassicWoWCreatureFamiliesIndex() (*wowcgd.CreatureFamiliesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureFamiliesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/creature-family/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreatureFamily(creatureFamilyID int) (*wowcgd.CreatureFamily, []byte, error) {
	var (
		dat wowcgd.CreatureFamily
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreatureTypesIndex() (*wowcgd.CreatureTypesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/creature-type/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreatureType(creatureTypeID int) (*wowcgd.CreatureType, []byte, error) {
	var (
		dat wowcgd.CreatureType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/creature-type/%d?locale=%s", creatureTypeID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreature(creatureID int) (*wowcgd.Creature, []byte, error) {
	var (
		dat wowcgd.Creature
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/creature/%d?locale=%s", creatureID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreatureDisplayMedia(creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, []byte, error) {
	var (
		dat wowcgd.CreatureDisplayMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/creature-display/%d?locale=%s", creatureDisplayID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWCreatureFamilyMedia(creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, []byte, error) {
	var (
		dat wowcgd.CreatureFamilyMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/creature-family/%d?locale=%s", creatureFamilyID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWGuildCrestComponentsIndex() (*wowcgd.GuildCrestComponentsIndex, []byte, error) {
	var (
		dat wowcgd.GuildCrestComponentsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/guild-crest/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWGuildCrestBorderMedia(borderID int) (*wowcgd.GuildCrestBorderMdedia, []byte, error) {
	var (
		dat wowcgd.GuildCrestBorderMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/border/%d?locale=%s", borderID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWGuildCrestEmblemMedia(emblemID int) (*wowcgd.GuildCrestEmblemMdedia, []byte, error) {
	var (
		dat wowcgd.GuildCrestEmblemMdedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/guild-crest/emblem/%d?locale=%s", emblemID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWItemClassesIndex() (*wowcgd.ItemClassesIndex, []byte, error) {
	var (
		dat wowcgd.ItemClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/item-class/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWItemClass(itemClassID int) (*wowcgd.ItemClass, []byte, error) {
	var (
		dat wowcgd.ItemClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/item-class/%d?locale=%s", itemClassID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWItemSubclass(itemClassID, itemSubclassID int) (*wowcgd.ItemSubclass, []byte, error) {
	var (
		dat wowcgd.ItemSubclass
		b   []byte
		err error
	)

	b, err = c.getURLBody(
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
func (c *Client) ClassicWoWItem(itemID int) (*wowcgd.Item, []byte, error) {
	var (
		dat wowcgd.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/item/%d?locale=%s", itemID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWItemMedia(itemID int) (*wowcgd.ItemMedia, []byte, error) {
	var (
		dat wowcgd.ItemMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/item/%d?locale=%s", itemID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPlayableClassesIndex() (*wowcgd.PlayableClassesIndex, []byte, error) {
	var (
		dat wowcgd.PlayableClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-class/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPlayableClass(classID int) (*wowcgd.PlayableClass, []byte, error) {
	var (
		dat wowcgd.PlayableClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d?locale=%s", classID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPlayableClassMedia(playableClassID int) (*wowcgd.PlayableClassMedia, []byte, error) {
	var (
		dat wowcgd.PlayableClassMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/playable-class/%d?locale=%s", playableClassID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPlayableRacesIndex() (*wowcgd.PlayableRacesIndex, []byte, error) {
	var (
		dat wowcgd.PlayableRacesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-race/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPlayableRace(raceID int) (*wowcgd.PlayableRace, []byte, error) {
	var (
		dat wowcgd.PlayableRace
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-race/%d?locale=%s", raceID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPowerTypesIndex() (*wowcgd.PowerTypesIndex, []byte, error) {
	var (
		dat wowcgd.PowerTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/power-type/index?locale=%s", c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWPowerType(powerTypeID int) (*wowcgd.PowerType, []byte, error) {
	var (
		dat wowcgd.PowerType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/power-type/%d?locale=%s", powerTypeID, c.locale), c.staticClassicNamespace)
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
func (c *Client) ClassicWoWRealmIndex() (*wowgd.RealmIndex, []byte, error) {
	var (
		dat wowgd.RealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/realm/index?locale=%s", c.locale), c.dynamicNamespace)
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
func (c *Client) ClassicWoWRealm(realmSlug string) (*wowgd.Realm, []byte, error) {
	var (
		dat wowgd.Realm
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/realm/%s?locale=%s", realmSlug, c.locale), c.dynamicNamespace)
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
func (c *Client) ClassicWoWRegionIndex() (*wowgd.RegionIndex, []byte, error) {
	var (
		dat wowgd.RegionIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/region/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
