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
