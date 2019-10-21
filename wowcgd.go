package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/wowcgd"
)

const (
	wowCreatureFamilyPath        = dataWowPath + "/creature-family"
	wowCreatureFamiliesIndexPath = wowCreatureFamilyPath + "/index"
	wowCreatureTypePath          = dataWowPath + "/creature-type"
	wowCreatureTypesIndexPath    = wowCreatureTypePath + "/index"
	wowCreaturePath              = dataWowPath + "/creature"
	wowMediaCreatureDisplayPath  = dataWowPath + "/media/creature-display"
	wowMediaCreatureFamilyPath   = dataWowPath + "/media/creature-family"
	wowGuildCrestPath            = dataWowPath + "/guild-crest"
	wowGuildCrestIndexPath       = wowGuildCrestPath + "/index"
	wowMediaGuildCrestBorderPath = dataWowPath + "/media/guild-crest/border"
	wowMediaGuildCrestEmblemPath = dataWowPath + "/media/guild-crest/emblem"
	wowItemClassPath             = dataWowPath + "/item-class"
	wowItemClassIndexPath        = wowItemClassPath + "/index"
	wowItemSubclassPath          = "/item-subclass"
	wowDataItemPath              = dataWowPath + "/item"
	wowMediaItemPath             = dataWowPath + "/media/item"
	wowMediaPlayableClassPath    = dataWowPath + "/media/playable-class"
)

// ClassicWoWCreatureFamiliesIndex returns an index of creature families
func (c *Client) ClassicWoWCreatureFamiliesIndex() (*wowcgd.CreatureFamiliesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureFamiliesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowCreatureFamiliesIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureFamily returns a creature family by ID
func (c *Client) ClassicWoWCreatureFamily(creatureFamilyID int) (*wowcgd.CreatureFamily, []byte, error) {
	var (
		dat wowcgd.CreatureFamily
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowCreatureFamilyPath+"/"+strconv.Itoa(creatureFamilyID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureTypesIndex returns an index of creature types
func (c *Client) ClassicWoWCreatureTypesIndex() (*wowcgd.CreatureTypesIndex, []byte, error) {
	var (
		dat wowcgd.CreatureTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowCreatureTypesIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureType returns a creature type by ID
func (c *Client) ClassicWoWCreatureType(creatureTypeID int) (*wowcgd.CreatureType, []byte, error) {
	var (
		dat wowcgd.CreatureType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowCreatureTypePath+"/"+strconv.Itoa(creatureTypeID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreature returns a creature type by ID
func (c *Client) ClassicWoWCreature(creatureID int) (*wowcgd.Creature, []byte, error) {
	var (
		dat wowcgd.Creature
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowCreaturePath+"/"+strconv.Itoa(creatureID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureDisplayMedia returns media for a creature display by ID
func (c *Client) ClassicWoWCreatureDisplayMedia(creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, []byte, error) {
	var (
		dat wowcgd.CreatureDisplayMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMediaCreatureDisplayPath+"/"+strconv.Itoa(creatureDisplayID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// ClassicWoWCreatureFamilyMedia returns media for a creature family by ID
func (c *Client) ClassicWoWCreatureFamilyMedia(creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, []byte, error) {
	var (
		dat wowcgd.CreatureFamilyMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMediaCreatureFamilyPath+"/"+strconv.Itoa(creatureFamilyID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowGuildCrestIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowMediaGuildCrestBorderPath+"/"+strconv.Itoa(borderID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowMediaGuildCrestEmblemPath+"/"+strconv.Itoa(emblemID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowItemClassIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowItemClassPath+"/"+strconv.Itoa(itemClassID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowItemClassPath+"/"+strconv.Itoa(itemClassID)+"/"+wowItemSubclassPath+"/"+strconv.Itoa(itemSubclassID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowDataItemPath+"/"+strconv.Itoa(itemID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowMediaItemPath+"/"+strconv.Itoa(itemID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowPlayableClassIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowPlayableClassPath+"/"+strconv.Itoa(classID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowMediaPlayableClassPath+"/"+strconv.Itoa(playableClassID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowRaceIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowRacePath+"/"+strconv.Itoa(raceID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowPowerTypeIndexPath+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
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

	b, err = c.getURLBody(c.apiURL+wowPowerTypePath+"/"+strconv.Itoa(powerTypeID)+"?"+localeQuery+c.locale.String(), c.staticClassicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
