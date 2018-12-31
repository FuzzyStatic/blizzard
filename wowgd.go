package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/wowgd"
)

const (
	dataWowPath                        = dataPath + wowPath
	indexPath                          = "/index"
	wowConnectedRealmPath              = dataWowPath + "/connected-realm"
	wowConnectedRealmIndexPath         = wowConnectedRealmPath + indexPath
	wowKeystoneAffixPath               = dataWowPath + "/keystone-affix"
	wowKeystoneAffixIndexPath          = wowKeystoneAffixPath + indexPath
	wowLeaderboardHallOfFamePath       = dataWowPath + "/leaderboard/hall-of-fame"
	wowMythicKeystonePath              = dataWowPath + "/mythic-keystone"
	wowMythicKeystoneDungeonPath       = wowMythicKeystonePath + "/dungeon"
	wowMythicKeystoneDungeonIndexPath  = wowMythicKeystoneDungeonPath + indexPath
	wowMythicKeystoneIndexPath         = wowMythicKeystonePath + indexPath
	periodPath                         = "/period"
	wowMythicKeystonePeriodPath        = wowMythicKeystonePath + periodPath
	wowMythicKeystonePeriodIndexPath   = wowMythicKeystonePeriodPath + indexPath
	wowMythicKeystoneSeasonPath        = wowMythicKeystonePath + "/season"
	wowMythicKeystoneSeasonIndexPath   = wowMythicKeystoneSeasonPath + indexPath
	mythicLeaderboardPath              = "/mythic-leaderboard"
	wowPlayableClassesPath             = dataWowPath + "/playable-classes"
	wowPlayableClassesIndexPath        = wowPlayableClassesPath + indexPath
	pvpTalentSlotsPath                 = "/pvp-talent-slots"
	wowPlayableSpecializationPath      = dataWowPath + "/playable-specialization"
	wowPlayableSpecializationIndexPath = wowPlayableSpecializationPath + indexPath
	wowPowerTypePath                   = dataWowPath + "/power-type"
	wowPowerTypeIndexPath              = wowPowerTypePath + indexPath
	wowRacePath                        = dataWowPath + "/race"
	wowRaceIndexPath                   = wowRacePath + indexPath
	wowRealmPath                       = dataWowPath + "/realm"
	wowRealmIndexPath                  = wowRealmPath + indexPath
	wowRegionPath                      = dataWowPath + "/region"
	wowRegionIndexPath                 = wowRegionPath + indexPath
	wowTokenIndexPath                  = dataWowPath + "/token/index"
)

// WoWConnectedRealmIndex returns an index of connected realms
func (c *Client) WoWConnectedRealmIndex() (*wowgd.ConnectedRealmIndex, error) {
	var (
		dat wowgd.ConnectedRealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWConnectedRealm returns a single connected realm by ID
func (c *Client) WoWConnectedRealm(connectedRealmID int) (*wowgd.ConnectedRealm, error) {
	var (
		dat wowgd.ConnectedRealm
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmPath+"/"+strconv.Itoa(connectedRealmID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneAffixIndex returns an index of Keystone affixes
func (c *Client) WoWMythicKeystoneAffixIndex() (*wowgd.MythicKeystoneAffixIndex, error) {
	var (
		dat wowgd.MythicKeystoneAffixIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowKeystoneAffixIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneAffix returns a single connected realm by ID
func (c *Client) WoWMythicKeystoneAffix(keystoneAffixID int) (*wowgd.MythicKeystoneAffix, error) {
	var (
		dat wowgd.MythicKeystoneAffix
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowKeystoneAffixPath+"/"+strconv.Itoa(keystoneAffixID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicRaidLeaderboard returns the leaderboard for a given raid and faction
func (c *Client) WoWMythicRaidLeaderboard(raid, faction string) (*wowgd.MythicRaidLeaderboard, error) {
	var (
		dat wowgd.MythicRaidLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowLeaderboardHallOfFamePath+"/"+raid+"/"+faction+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneDungeonIndex returns an index of Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneDungeonIndex() (*wowgd.MythicKeystoneDungeonIndex, error) {
	var (
		dat wowgd.MythicKeystoneDungeonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneDungeonIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneDungeon returns a Mythic Keystone dungeon by ID
func (c *Client) WoWMythicKeystoneDungeon(dungeonID int) (*wowgd.MythicKeystoneDungeon, error) {
	var (
		dat wowgd.MythicKeystoneDungeon
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneDungeonPath+"/"+strconv.Itoa(dungeonID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneIndex returns n index of links to other documents related to Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneIndex() (*wowgd.MythicKeystoneIndex, error) {
	var (
		dat wowgd.MythicKeystoneIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystonePeriodIndex returns an index of Mythic Keystone periods
func (c *Client) WoWMythicKeystonePeriodIndex() (*wowgd.MythicKeystonePeriodIndex, error) {
	var (
		dat wowgd.MythicKeystonePeriodIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystonePeriodIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystonePeriod returns a single connected realm by ID
func (c *Client) WoWMythicKeystonePeriod(keystonePeriodID int) (*wowgd.MythicKeystonePeriod, error) {
	var (
		dat wowgd.MythicKeystonePeriod
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystonePeriodPath+"/"+strconv.Itoa(keystonePeriodID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneSeasonIndex returns an index of Keystone affixes
func (c *Client) WoWMythicKeystoneSeasonIndex() (*wowgd.MythicKeystoneSeasonIndex, error) {
	var (
		dat wowgd.MythicKeystoneSeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneSeasonIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMythicKeystoneSeason returns a single connected realm by ID
func (c *Client) WoWMythicKeystoneSeason(seasonID int) (*wowgd.MythicKeystoneSeason, error) {
	var (
		dat wowgd.MythicKeystoneSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneSeasonPath+"/"+strconv.Itoa(seasonID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
