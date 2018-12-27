package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/sc2c"
)

const (
	sc2Path                = "/sc2"
	staticPath             = "/static"
	sc2StaticProfilePath   = sc2Path + staticPath + "/profile"
	metadataPath           = "/metadata"
	sc2MetadataProfilePath = sc2Path + metadataPath + "/profile"
	sc2ProfilePath         = sc2Path + "/profile"
	ladderPath             = "/ladder"
	ladderSummaryPath      = ladderPath + "/summary"
)

// SC2GetStaticProfile returns all static SC2 profile data (achievements, categories, criteria, and rewards)
func (c *Config) SC2GetStaticProfile(region Region) (*sc2c.StaticProfile, error) {
	var (
		dat sc2c.StaticProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2StaticProfilePath + "/" + strconv.Itoa(int(region)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetMetadataProfile returns all SC2 profile metadata
func (c *Config) SC2GetMetadataProfile(region Region, realmID, profileID int) (*sc2c.MetadataProfile, error) {
	var (
		dat sc2c.MetadataProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2MetadataProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetProfile returns all SC2 profile data
func (c *Config) SC2GetProfile(region Region, realmID, profileID int) (*sc2c.Profile, error) {
	var (
		dat sc2c.Profile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2ProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetProfileLadderSummary returns SC2 profile ladder summary
func (c *Config) SC2GetProfileLadderSummary(region Region, realmID, profileID int) (*sc2c.LadderSummary, error) {
	var (
		dat sc2c.LadderSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2ProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + ladderSummaryPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetProfileLadder returns SC2 profile ladder data
func (c *Config) SC2GetProfileLadder(region Region, realmID, profileID, ladderID int) (*sc2c.Ladder, error) {
	var (
		dat sc2c.Ladder
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2ProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + ladderPath + "/" + strconv.Itoa(ladderID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
