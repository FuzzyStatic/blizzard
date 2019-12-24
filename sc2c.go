package blizzard

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/FuzzyStatic/blizzard/sc2c"
)

const (
	sc2Path                   = "/sc2"
	ladderPath                = "/ladder"
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

// SC2StaticProfile returns all static SC2 profile data (achievements, categories, criteria, and rewards)
func (c *Client) SC2StaticProfile(region Region) (*sc2c.StaticProfile, []byte, error) {
	var (
		dat sc2c.StaticProfile
		b   []byte
		err error
	)

	if c.region == CN {
		return &dat, b, errors.New("CN is not a valid region for this call")
	}

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/sc2/static/profile/%d?locale=%s", region, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2MetadataProfile returns all SC2 profile metadata
func (c *Client) SC2MetadataProfile(region Region, realmID, profileID int) (*sc2c.MetadataProfile, []byte, error) {
	var (
		dat sc2c.MetadataProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/sc2/metadata/profile/%d/%d/%d?locale=%s", region, realmID, profileID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2Profile returns all SC2 profile data
func (c *Client) SC2Profile(region Region, realmID, profileID int) (*sc2c.Profile, []byte, error) {
	var (
		dat sc2c.Profile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/sc2/profile/%d/%d/%d?locale=%s", region, realmID, profileID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2ProfileLadderSummary returns SC2 profile ladder summary
func (c *Client) SC2ProfileLadderSummary(region Region, realmID, profileID int) (*sc2c.LadderSummary, []byte, error) {
	var (
		dat sc2c.LadderSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/sc2/profile/%d/%d/%d/ladder/summary?locale=%s", region, realmID, profileID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2ProfileLadder returns SC2 profile ladder data
func (c *Client) SC2ProfileLadder(region Region, realmID, profileID, ladderID int) (*sc2c.Ladder, []byte, error) {
	var (
		dat sc2c.Ladder
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/sc2/profile/%d/%d/%d/ladder/%d?locale=%s", region, realmID, profileID, ladderID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LadderGrandmaster returns SC2 ladder grandmaster for current season
func (c *Client) SC2LadderGrandmaster(region Region) (*sc2c.LadderGrandmaster, []byte, error) {
	var (
		dat sc2c.LadderGrandmaster
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2GrandmasterPath+"/"+strconv.Itoa(int(region))+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LadderSeason returns SC2 ladder current season
func (c *Client) SC2LadderSeason(region Region) (*sc2c.LadderSeason, []byte, error) {
	var (
		dat sc2c.LadderSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2SeasonPath+"/"+strconv.Itoa(int(region))+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2Player returns data about player using account ID
func (c *Client) SC2Player(accountID int) (*sc2c.Player, []byte, error) {
	var (
		dat sc2c.Player
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2PlayerPath+"/"+strconv.Itoa(accountID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyProfile returns all SC2 legacy profile data
func (c *Client) SC2LegacyProfile(region Region, realmID, profileID int) (*sc2c.LegacyProfile, []byte, error) {
	var (
		dat sc2c.LegacyProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyProfilePath+"/"+strconv.Itoa(int(region))+"/"+strconv.Itoa(realmID)+"/"+strconv.Itoa(profileID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyProfileLadders returns all SC2 legacy profile ladder data
func (c *Client) SC2LegacyProfileLadders(region Region, realmID, profileID int) (*sc2c.LegacyProfileLadders, []byte, error) {
	var (
		dat sc2c.LegacyProfileLadders
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyProfilePath+"/"+strconv.Itoa(int(region))+"/"+strconv.Itoa(realmID)+"/"+strconv.Itoa(profileID)+laddersPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyProfileMatches returns all SC2 legacy profile matches data
func (c *Client) SC2LegacyProfileMatches(region Region, realmID, profileID int) (*sc2c.LegacyProfileMatches, []byte, error) {
	var (
		dat sc2c.LegacyProfileMatches
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyProfilePath+"/"+strconv.Itoa(int(region))+"/"+strconv.Itoa(realmID)+"/"+strconv.Itoa(profileID)+matchesPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyLadder returns SC2 legacy ladder data
func (c *Client) SC2LegacyLadder(region Region, ladderID int) (*sc2c.LegacyLadder, []byte, error) {
	var (
		dat sc2c.LegacyLadder
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyLadderPath+"/"+strconv.Itoa(int(region))+"/"+strconv.Itoa(ladderID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyAchievements returns SC2 legacy achievements for region
func (c *Client) SC2LegacyAchievements(region Region) (*sc2c.LegacyAchievements, []byte, error) {
	var (
		dat sc2c.LegacyAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyAchievementsPath+"/"+strconv.Itoa(int(region))+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// SC2LegacyRewards returns SC2 legacy rewards for region
func (c *Client) SC2LegacyRewards(region Region) (*sc2c.LegacyRewards, []byte, error) {
	var (
		dat sc2c.LegacyRewards
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+sc2LegacyRewardsPath+"/"+strconv.Itoa(int(region))+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
