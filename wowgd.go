package blizzard

import (
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/wowgd"
)

// WoWAchievementCategoriesIndex returns an index of achievement categories.
func (c *Client) WoWAchievementCategoriesIndex() (*wowgd.AchievementCategoriesIndex, []byte, error) {
	var (
		dat wowgd.AchievementCategoriesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/achievement-category/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementCategory returns an achievement category by ID.
func (c *Client) WoWAchievementCategory(achievementCategoryID int) (*wowgd.AchievementCategory, []byte, error) {
	var (
		dat wowgd.AchievementCategory
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/achievement-category/%d?locale=%s", achievementCategoryID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementIndex returns an index of achievements.
func (c *Client) WoWAchievementIndex() (*wowgd.AchievementIndex, []byte, error) {
	var (
		dat wowgd.AchievementIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/achievement/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievement returns an achievement category by ID.
func (c *Client) WoWAchievement(achievementID int) (*wowgd.Achievement, []byte, error) {
	var (
		dat wowgd.Achievement
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/achievement/%d?locale=%s", achievementID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAchievementMedia returns media for an achievement by ID.
func (c *Client) WoWAchievementMedia(achievementID int) (*wowgd.AchievementMedia, []byte, error) {
	var (
		dat wowgd.AchievementMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/achievement/%d?locale=%s", achievementID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssenceIndex returns an index of azerite essences.
func (c *Client) WoWAzeriteEssenceIndex() (*wowgd.AzeriteEssenceIndex, []byte, error) {
	var (
		dat wowgd.AzeriteEssenceIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/azerite-essence/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssencereturns an azerite essence by ID.
func (c *Client) WoWAzeriteEssence(azeriteEssenceID int) (*wowgd.AzeriteEssence, []byte, error) {
	var (
		dat wowgd.AzeriteEssence
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/azerite-essence/%d?locale=%s", azeriteEssenceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWAzeriteEssenceMedia returns media for an azerite essence by ID.
func (c *Client) WoWAzeriteEssenceMedia(azeriteEssenceID int) (*wowgd.AzeriteEssenceMedia, []byte, error) {
	var (
		dat wowgd.AzeriteEssenceMedia
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/media/azerite-essence/%d?locale=%s", azeriteEssenceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWConnectedRealmIndex returns an index of connected realms
func (c *Client) WoWConnectedRealmIndex() (*wowgd.ConnectedRealmIndex, []byte, error) {
	var (
		dat wowgd.ConnectedRealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/connected-realm/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d?locale=%s", connectedRealmID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/keystone-affix/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/keystone-affix/%d?locale=%s", keystoneAffixID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/leaderboard/hall-of-fame/%s/%s?locale=%s", raid, faction, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/dungeon/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/dungeon/%d?locale=%s", dungeonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/period/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/period/%d?locale=%s", periodID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/season/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/mythic-keystone/season/%d?locale=%s", seasonID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/index?locale=%s", connectedRealmID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/connected-realm/%d/mythic-leaderboard/%d/period/%d?locale=%s", connectedRealmID, dungeonID, period, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-class/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d?locale=%s", classID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableClassPvPTalentSlots returns the PvP talent slots for a playable class by ID
func (c *Client) WoWPlayableClassPvPTalentSlots(classID int) (*wowgd.PlayableClassPvPTalentSlots, []byte, error) {
	var (
		dat wowgd.PlayableClassPvPTalentSlots
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-class/%d/pvp-talent-slots?locale=%s", classID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecializationIndex returns an index of playable specializations.
func (c *Client) WoWPlayableSpecializationIndex() (*wowgd.PlayableSpecializationIndex, []byte, error) {
	var (
		dat wowgd.PlayableSpecializationIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-specialization/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableSpecialization returns a playable specialization by ID.
func (c *Client) WoWPlayableSpecialization(specID int) (*wowgd.PlayableSpecialization, []byte, error) {
	var (
		dat wowgd.PlayableSpecialization
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-specialization/%d?locale=%s", specID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPowerTypesIndex returns an index of power types.
func (c *Client) WoWPowerTypesIndex() (*wowgd.PowerTypesIndex, []byte, error) {
	var (
		dat wowgd.PowerTypesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/power-type/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPowerType returns a power type by ID.
func (c *Client) WoWPowerType(powerTypeID int) (*wowgd.PowerType, []byte, error) {
	var (
		dat wowgd.PowerType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/power-type/%d?locale=%s", powerTypeID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableRacesIndex returns an index of races.
func (c *Client) WoWPlayableRacesIndex() (*wowgd.PlayableRacesIndex, []byte, error) {
	var (
		dat wowgd.PlayableRacesIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-race/index?locale=%s", c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWPlayableRace returns a race by ID.
func (c *Client) WoWPlayableRace(raceID int) (*wowgd.PlayableRace, []byte, error) {
	var (
		dat wowgd.PlayableRace
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/playable-race/%d?locale=%s", raceID, c.locale), c.staticNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWRealmIndex returns an index of realms.
func (c *Client) WoWRealmIndex() (*wowgd.RealmIndex, []byte, error) {
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

// WoWRealm returns a single realm by slug or ID.
func (c *Client) WoWRealm(realmSlug string) (*wowgd.Realm, []byte, error) {
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

// WoWRegionIndex returns an index of regions.
func (c *Client) WoWRegionIndex() (*wowgd.RegionIndex, []byte, error) {
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

// WoWRegion returns a single region by ID.
func (c *Client) WoWRegion(regionID int) (*wowgd.Region, []byte, error) {
	var (
		dat wowgd.Region
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/region/%d?locale=%s", regionID, c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
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

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/wow/token/index?locale=%s", c.locale), c.dynamicNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
