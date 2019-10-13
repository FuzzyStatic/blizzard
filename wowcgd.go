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

// WoWClassicCreatureFamiliesIndex returns an index of creature families
func (c *Client) WoWClassicCreatureFamiliesIndex() (*wowcgd.CreatureFamiliesIndex, []byte, error) {
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

// WoWClassicCreatureFamily returns a creature family by ID
func (c *Client) WoWClassicCreatureFamily(creatureFamilyID int) (*wowcgd.CreatureFamily, []byte, error) {
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

// WoWClassicCreatureTypesIndex returns an index of creature types
func (c *Client) WoWClassicCreatureTypesIndex() (*wowcgd.CreatureTypesIndex, []byte, error) {
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

// WoWClassicCreatureType returns a creature type by ID
func (c *Client) WoWClassicCreatureType(creatureTypeID int) (*wowcgd.CreatureType, []byte, error) {
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

// WoWClassicCreature returns a creature type by ID
func (c *Client) WoWClassicCreature(creatureID int) (*wowcgd.Creature, []byte, error) {
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

// WoWClassicCreatureDisplayMedia returns media for a creature display by ID
func (c *Client) WoWClassicCreatureDisplayMedia(creatureDisplayID int) (*wowcgd.CreatureDisplayMedia, []byte, error) {
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

// WoWClassicCreatureFamilyMedia returns media for a creature family by ID
func (c *Client) WoWClassicCreatureFamilyMedia(creatureFamilyID int) (*wowcgd.CreatureFamilyMedia, []byte, error) {
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
