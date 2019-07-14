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
	mythicLeaderboardIndexPath         = mythicLeaderboardPath + indexPath
	wowPlayableClassPath               = dataWowPath + "/playable-class"
	wowPlayableClassIndexPath          = wowPlayableClassPath + indexPath
	pvpTalentSlotsPath                 = "/pvp-talent-slots"
	wowPlayableSpecializationPath      = dataWowPath + "/playable-specialization"
	wowPlayableSpecializationIndexPath = wowPlayableSpecializationPath + indexPath
	wowPowerTypePath                   = dataWowPath + "/power-type"
	wowPowerTypeIndexPath              = wowPowerTypePath + indexPath
	wowRacePath                        = dataWowPath + "/playable-race"
	wowRaceIndexPath                   = wowRacePath + indexPath
	wowRealmPath                       = dataWowPath + "/realm"
	wowRealmIndexPath                  = wowRealmPath + indexPath
	wowRegionPath                      = dataWowPath + "/region"
	wowRegionIndexPath                 = wowRegionPath + indexPath
	wowTokenIndexPath                  = dataWowPath + "/token/index"
)

// WoWConnectedRealmIndex returns an index of connected realms
func (c *Client) WoWConnectedRealmIndex() (*wowgd.ConnectedRealmIndex, []byte, error) {
	var (
		dat wowgd.ConnectedRealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWConnectedRealm returns a single connected realm by ID
func (c *Client) WoWConnectedRealm(connectedRealmID int) (*wowgd.ConnectedRealm, []byte, error) {
	var (
		dat wowgd.ConnectedRealm
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmPath+"/"+strconv.Itoa(connectedRealmID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneAffixIndex returns an index of Keystone affixes
func (c *Client) WoWMythicKeystoneAffixIndex() (*wowgd.MythicKeystoneAffixIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneAffixIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowKeystoneAffixIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneAffix returns a single connected realm by ID
func (c *Client) WoWMythicKeystoneAffix(keystoneAffixID int) (*wowgd.MythicKeystoneAffix, []byte, error) {
	var (
		dat wowgd.MythicKeystoneAffix
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowKeystoneAffixPath+"/"+strconv.Itoa(keystoneAffixID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicRaidLeaderboard returns the leaderboard for a given raid and faction
func (c *Client) WoWMythicRaidLeaderboard(raid, faction string) (*wowgd.MythicRaidLeaderboard, []byte, error) {
	var (
		dat wowgd.MythicRaidLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowLeaderboardHallOfFamePath+"/"+raid+"/"+faction+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneDungeonIndex returns an index of Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneDungeonIndex() (*wowgd.MythicKeystoneDungeonIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneDungeonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneDungeonIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneDungeon returns a Mythic Keystone dungeon by ID
func (c *Client) WoWMythicKeystoneDungeon(dungeonID int) (*wowgd.MythicKeystoneDungeon, []byte, error) {
	var (
		dat wowgd.MythicKeystoneDungeon
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneDungeonPath+"/"+strconv.Itoa(dungeonID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneIndex returns n index of links to other documents related to Mythic Keystone dungeons
func (c *Client) WoWMythicKeystoneIndex() (*wowgd.MythicKeystoneIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystonePeriodIndex returns an index of Mythic Keystone periods
func (c *Client) WoWMythicKeystonePeriodIndex() (*wowgd.MythicKeystonePeriodIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystonePeriodIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystonePeriodIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystonePeriod returns a Mythic Keystone period by ID
func (c *Client) WoWMythicKeystonePeriod(periodID int) (*wowgd.MythicKeystonePeriod, []byte, error) {
	var (
		dat wowgd.MythicKeystonePeriod
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystonePeriodPath+"/"+strconv.Itoa(periodID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneSeasonIndex returns an index of Mythic Keystone seasons
func (c *Client) WoWMythicKeystoneSeasonIndex() (*wowgd.MythicKeystoneSeasonIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneSeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneSeasonIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneSeason returns a Mythic Keystone season by ID
func (c *Client) WoWMythicKeystoneSeason(seasonID int) (*wowgd.MythicKeystoneSeason, []byte, error) {
	var (
		dat wowgd.MythicKeystoneSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMythicKeystoneSeasonPath+"/"+strconv.Itoa(seasonID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneLeaderboardIndex returns an index of Mythic Keystone Leaderboard dungeon instances for a connected realm
func (c *Client) WoWMythicKeystoneLeaderboardIndex(connectedRealmID int) (*wowgd.MythicKeystoneLeaderboardIndex, []byte, error) {
	var (
		dat wowgd.MythicKeystoneLeaderboardIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmPath+"/"+strconv.Itoa(connectedRealmID)+"/"+mythicLeaderboardIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneLeaderboard returns a weekly Mythic Keystone Leaderboard by period
func (c *Client) WoWMythicKeystoneLeaderboard(connectedRealmID, dungeonID, period int) (*wowgd.MythicKeystoneLeaderboard, []byte, error) {
	var (
		dat wowgd.MythicKeystoneLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowConnectedRealmPath+"/"+strconv.Itoa(connectedRealmID)+"/"+mythicLeaderboardPath+"/"+strconv.Itoa(dungeonID)+periodPath+"/"+strconv.Itoa(period)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableClassesIndex returns an index of playable classes
func (c *Client) WoWPlayableClassesIndex() (*wowgd.PlayableClassesIndex, []byte, error) {
	var (
		dat wowgd.PlayableClassesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPlayableClassIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableClass returns a playable class by ID
func (c *Client) WoWPlayableClass(classID int) (*wowgd.PlayableClass, []byte, error) {
	var (
		dat wowgd.PlayableClass
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPlayableClassPath+"/"+strconv.Itoa(classID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableClassPVPTalentSlots returns the PvP talent slots for a playable class by ID
func (c *Client) WoWPlayableClassPVPTalentSlots(classID int) (*wowgd.PlayableClassPVPTalentSlots, []byte, error) {
	var (
		dat wowgd.PlayableClassPVPTalentSlots
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPlayableClassPath+"/"+strconv.Itoa(classID)+pvpTalentSlotsPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecializationIndex returns an index of playable specializations
func (c *Client) WoWPlayableSpecializationIndex() (*wowgd.PlayableSpecializationIndex, []byte, error) {
	var (
		dat wowgd.PlayableSpecializationIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPlayableSpecializationIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecialization returns a playable specialization by ID
func (c *Client) WoWPlayableSpecialization(specID int) (*wowgd.PlayableSpecialization, []byte, error) {
	var (
		dat wowgd.PlayableSpecialization
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPlayableSpecializationPath+"/"+strconv.Itoa(specID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPowerTypesIndex returns an index of power types
func (c *Client) WoWPowerTypesIndex() (*wowgd.PowerTypesIndex, []byte, error) {
	var (
		dat wowgd.PowerTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPowerTypeIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPowerType returns a power type by ID
func (c *Client) WoWPowerType(powerTypeID int) (*wowgd.PowerType, []byte, error) {
	var (
		dat wowgd.PowerType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPowerTypePath+"/"+strconv.Itoa(powerTypeID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableRacesIndex returns an index of races
func (c *Client) WoWPlayableRacesIndex() (*wowgd.PlayableRacesIndex, []byte, error) {
	var (
		dat wowgd.PlayableRacesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRaceIndexPath+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPlayableRace returns a race by ID
func (c *Client) WoWPlayableRace(raceID int) (*wowgd.PlayableRace, []byte, error) {
	var (
		dat wowgd.PlayableRace
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRacePath+"/"+strconv.Itoa(raceID)+"?"+localeQuery+c.locale.String(), c.staticNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRealmIndex returns an index of realms
func (c *Client) WoWRealmIndex() (*wowgd.RealmIndex, []byte, error) {
	var (
		dat wowgd.RealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRealmIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRealm returns a single realm by slug or ID
func (c *Client) WoWRealm(realmSlug string) (*wowgd.Realm, []byte, error) {
	var (
		dat wowgd.Realm
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRealmPath+"/"+realmSlug+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRegionIndex returns an index of regions
func (c *Client) WoWRegionIndex() (*wowgd.RegionIndex, []byte, error) {
	var (
		dat wowgd.RegionIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRegionIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRegion returns a single region by ID
func (c *Client) WoWRegion(regionID int) (*wowgd.Region, []byte, error) {
	var (
		dat wowgd.Region
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRegionPath+"/"+strconv.Itoa(regionID)+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWToken returns the WoW Token index
func (c *Client) WoWToken() (*wowgd.Token, []byte, error) {
	var (
		dat wowgd.Token
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowTokenIndexPath+"?"+localeQuery+c.locale.String(), c.dynamicNamespace)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}
