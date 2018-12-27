package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/sc2c"
)

const (
	sc2Path                   = "/sc2"
	staticPath                = "/static"
	sc2StaticProfilePath      = sc2Path + staticPath + profilePath
	metadataPath              = "/metadata"
	sc2MetadataProfilePath    = sc2Path + metadataPath + profilePath
	sc2ProfilePath            = sc2Path + profilePath
	ladderPath                = "/ladder"
	ladderSummaryPath         = ladderPath + "/summary"
	sc2GrandmasterPath        = sc2Path + ladderPath + "/grandmaster"
	sc2SeasonPath             = sc2Path + ladderPath + "/season"
	sc2PlayerPath             = sc2Path + "/player"
	legacyPath                = "/legacy"
	sc2LegacyProfilePath      = sc2Path + legacyPath + profilePath
	laddersPath               = "/ladders"
	matchesPath               = "/matches"
	sc2LegacyLadderPath       = sc2Path + legacyPath + ladderPath
	sc2LegacyDataPath         = sc2Path + legacyPath + dataPath
	sc2LegacyAchievementsPath = sc2LegacyDataPath + "/achievements"
	sc2LegacyRewardsPath      = sc2LegacyDataPath + "/rewards"
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

// SC2GetLadderGrandmaster returns SC2 ladder grandmaster for current season
func (c *Config) SC2GetLadderGrandmaster(region Region) (*sc2c.LadderGrandmaster, error) {
	var (
		dat sc2c.LadderGrandmaster
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2GrandmasterPath + "/" + strconv.Itoa(int(region)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLadderSeason returns SC2 ladder current season
func (c *Config) SC2GetLadderSeason(region Region) (*sc2c.LadderSeason, error) {
	var (
		dat sc2c.LadderSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2SeasonPath + "/" + strconv.Itoa(int(region)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetPlayer returns data about player using account ID
func (c *Config) SC2GetPlayer(accountID int) (*sc2c.Player, error) {
	var (
		dat sc2c.Player
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2PlayerPath + "/" + strconv.Itoa(accountID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyProfile returns all SC2 legacy profile data
func (c *Config) SC2GetLegacyProfile(region Region, realmID, profileID int) (*sc2c.LegacyProfile, error) {
	var (
		dat sc2c.LegacyProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyProfileLadders returns all SC2 legacy profile ladder data
func (c *Config) SC2GetLegacyProfileLadders(region Region, realmID, profileID int) (*sc2c.LegacyProfileLadders, error) {
	var (
		dat sc2c.LegacyProfileLadders
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + laddersPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyProfileMatches returns all SC2 legacy profile matches data
func (c *Config) SC2GetLegacyProfileMatches(region Region, realmID, profileID int) (*sc2c.LegacyProfileMatches, error) {
	var (
		dat sc2c.LegacyProfileMatches
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyProfilePath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(realmID) + "/" + strconv.Itoa(profileID) + matchesPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyLadder returns SC2 legacy ladder data
func (c *Config) SC2GetLegacyLadder(region Region, ladderID int) (*sc2c.LegacyLadder, error) {
	var (
		dat sc2c.LegacyLadder
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyLadderPath + "/" + strconv.Itoa(int(region)) + "/" + strconv.Itoa(ladderID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyAchievements returns SC2 legacy achievements for region
func (c *Config) SC2GetLegacyAchievements(region Region) (*sc2c.LegacyAchievements, error) {
	var (
		dat sc2c.LegacyAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyAchievementsPath + "/" + strconv.Itoa(int(region)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// SC2GetLegacyRewards returns SC2 legacy rewards for region
func (c *Config) SC2GetLegacyRewards(region Region) (*sc2c.LegacyRewards, error) {
	var (
		dat sc2c.LegacyRewards
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LegacyRewardsPath + "/" + strconv.Itoa(int(region)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
