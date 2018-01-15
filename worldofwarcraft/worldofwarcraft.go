/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:37:59
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:30:02
 */

// Package worldofwarcraft is a client library to use Blizzard World of Warcraft API calls.
package worldofwarcraft

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FuzzyStatic/blizzard"
)

// WorldOfWarcraft regional API URLs, locale, access token, api key
type WorldOfWarcraft struct {
	Auth         blizzard.Auth
	Locale       string
	Namespace    string
	DataURL      string
	CommunityURL string
}

// New create new WorldOfWarcraft structure
func New(auth blizzard.Auth, region blizzard.Region) *WorldOfWarcraft {
	var w = WorldOfWarcraft{Auth: auth}

	switch region {
	case blizzard.EU:
		w.Locale = "en_GB"
		w.Namespace = "dynamic-eu"
		w.DataURL = "https://eu.api.battle.net/data/wow"
		w.CommunityURL = "https://eu.api.battle.net/wow"
	case blizzard.KR:
		w.Locale = "ko_KR"
		w.Namespace = "dynamic-kr"
		w.DataURL = "https://kr.api.battle.net/data/wow"
		w.CommunityURL = "https://kr.api.battle.net/wow"
	case blizzard.TW:
		w.Locale = "zh_TW"
		w.Namespace = "dynamic-tw"
		w.DataURL = "https://tw.api.battle.net/data/wow"
		w.CommunityURL = "https://tw.api.battle.net/wow"
	case blizzard.US:
		w.Locale = "en_US"
		w.Namespace = "dynamic-us"
		w.DataURL = "https://us.api.battle.net/data/wow"
		w.CommunityURL = "https://us.api.battle.net/wow"
	default: // USA! USA!
		w.Locale = "en_US"
		w.Namespace = "dynamic-us"
		w.DataURL = "https://us.api.battle.net/data/wow"
		w.CommunityURL = "https://us.api.battle.net/wow"
	}

	return &w
}

// GetConnectedRealmIndexJSON gets connected realm index JSON information
func (w *WorldOfWarcraft) GetConnectedRealmIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + connectedRealmPath + "/?" + namespaceQuery + w.Namespace + "&" + localeQuery +
		w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetConnectedRealmIndex puts connected realm index information into ConnectedRealmIndex structure
func (w *WorldOfWarcraft) GetConnectedRealmIndex() (*ConnectedRealmIndex, error) {
	var (
		connectedRealmIndex ConnectedRealmIndex
		json                *[]byte
		err                 error
	)

	json, err = w.GetConnectedRealmIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &connectedRealmIndex)
	if err != nil {
		return nil, err
	}

	return &connectedRealmIndex, nil
}

// GetConnectedRealmJSON gets specified connected realm JSON information
func (w *WorldOfWarcraft) GetConnectedRealmJSON(connectRealmID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + connectedRealmPath + "/" + strconv.Itoa(connectRealmID) + "?" + namespaceQuery +
		w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetConnectedRealm puts specified connected realm information into ConnectedRealm structure
func (w *WorldOfWarcraft) GetConnectedRealm(connectRealmID int) (*ConnectedRealm, error) {
	var (
		connectedRealm ConnectedRealm
		json           *[]byte
		err            error
	)

	json, err = w.GetConnectedRealmJSON(connectRealmID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &connectedRealm)
	if err != nil {
		return nil, err
	}

	return &connectedRealm, nil
}

// GetMythicLeaderboardIndexJSON gets specified connected realm mythic leaderboard index JSON information
func (w *WorldOfWarcraft) GetMythicLeaderboardIndexJSON(connectRealmID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + connectedRealmPath + "/" + strconv.Itoa(connectRealmID) + mythicLeaderPath +
		"/?" + namespaceQuery + w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery +
		w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetMythicLeaderboardIndex puts specified connected realm mythic leaderboard index information into MythicLeaderboardIndex structure
func (w *WorldOfWarcraft) GetMythicLeaderboardIndex(connectRealmID int) (*MythicLeaderboardIndex, error) {
	var (
		mythicLeaderboardIndex MythicLeaderboardIndex
		json                   *[]byte
		err                    error
	)

	json, err = w.GetMythicLeaderboardIndexJSON(connectRealmID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &mythicLeaderboardIndex)
	if err != nil {
		return nil, err
	}

	return &mythicLeaderboardIndex, nil
}

// GetMythicLeaderboardJSON gets specified connected realm and dungeon mythic leaderboard JSON information for period
func (w *WorldOfWarcraft) GetMythicLeaderboardJSON(connectRealmID, dungeonID, period int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + connectedRealmPath + "/" + strconv.Itoa(connectRealmID) + mythicLeaderPath +
		"/" + strconv.Itoa(dungeonID) + periodPath + "/" + strconv.Itoa(period) + "?" + namespaceQuery +
		w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetMythicLeaderboard puts specified connected realm and dungeon mythic leaderboard information into MythicLeaderboardIndex structure for period
func (w *WorldOfWarcraft) GetMythicLeaderboard(connectRealmID, dungeonID, period int) (*MythicLeaderboard, error) {
	var (
		mythicLeaderboard MythicLeaderboard
		json              *[]byte
		err               error
	)

	json, err = w.GetMythicLeaderboardJSON(connectRealmID, dungeonID, period)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &mythicLeaderboard)
	if err != nil {
		return nil, err
	}

	return &mythicLeaderboard, nil
}

// GetRealmIndexJSON gets realm index JSON information
func (w *WorldOfWarcraft) GetRealmIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + realmPath + "/?" + namespaceQuery + w.Namespace + "&" + localeQuery +
		w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRealmIndex puts realm index information into RealmIndex structure
func (w *WorldOfWarcraft) GetRealmIndex() (*RealmIndex, error) {
	var (
		realmIndex RealmIndex
		json       *[]byte
		err        error
	)

	json, err = w.GetRealmIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &realmIndex)
	if err != nil {
		return nil, err
	}

	return &realmIndex, nil
}

// GetRealmJSON gets specified realm JSON information
// Values accepted: realm ID or realm slug
func (w *WorldOfWarcraft) GetRealmJSON(realmID interface{}) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	if v, ok := realmID.(int); ok {
		url = w.DataURL + realmPath + "/" + strconv.Itoa(v) + "?" + namespaceQuery +
			w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken
	}

	if v, ok := realmID.(string); ok {
		url = w.DataURL + realmPath + "/" + v + "?" + namespaceQuery +
			w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken
	}

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRealm puts specified realm information into Realm structure
// Values accepted: realm ID or realm slug
func (w *WorldOfWarcraft) GetRealm(realmID interface{}) (*Realm, error) {
	var (
		realm Realm
		json  *[]byte
		err   error
	)

	json, err = w.GetRealmJSON(realmID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &realm)
	if err != nil {
		return nil, err
	}

	return &realm, nil
}

// GetRegionIndexJSON gets region index JSON information
func (w *WorldOfWarcraft) GetRegionIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + regionPath + "/?" + namespaceQuery + w.Namespace + "&" + localeQuery +
		w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRegionIndex puts region index information into RealmIndex structure
func (w *WorldOfWarcraft) GetRegionIndex() (*RegionIndex, error) {
	var (
		regionIndex RegionIndex
		json        *[]byte
		err         error
	)

	json, err = w.GetRegionIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &regionIndex)
	if err != nil {
		return nil, err
	}

	return &regionIndex, nil
}

// GetRegionJSON gets specified region JSON information
func (w *WorldOfWarcraft) GetRegionJSON(regionID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + regionPath + "/" + strconv.Itoa(regionID) + "?" + namespaceQuery +
		w.Namespace + "&" + localeQuery + w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRegion puts specified region information into Region structure
func (w *WorldOfWarcraft) GetRegion(regionID int) (*Region, error) {
	var (
		region Region
		json   *[]byte
		err    error
	)

	json, err = w.GetRegionJSON(regionID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &region)
	if err != nil {
		return nil, err
	}

	return &region, nil
}

// GetTokenIndexJSON gets token index JSON information
func (w *WorldOfWarcraft) GetTokenIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.DataURL + tokenPath + "/?" + namespaceQuery + w.Namespace + "&" + localeQuery +
		w.Locale + "&" + accessTokenQuery + w.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetTokenIndex puts token index information into RealmIndex structure
func (w *WorldOfWarcraft) GetTokenIndex() (*TokenIndex, error) {
	var (
		tokenIndex TokenIndex
		json       *[]byte
		err        error
	)

	json, err = w.GetTokenIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &tokenIndex)
	if err != nil {
		return nil, err
	}

	return &tokenIndex, nil
}

// GetAchievementJSON gets specified achievement JSON information
func (w *WorldOfWarcraft) GetAchievementJSON(achievementID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + achievementPath + "/" + strconv.Itoa(achievementID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	fmt.Println(url)

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetAchievement puts specified achievement information into Achievement structure
func (w *WorldOfWarcraft) GetAchievement(achievementID int) (*Achievement, error) {
	var (
		achievement Achievement
		json        *[]byte
		err         error
	)

	json, err = w.GetAchievementJSON(achievementID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &achievement)
	if err != nil {
		return nil, err
	}

	return &achievement, nil
}

// GetAuctionDataJSON gets auction data JSON information for realm
func (w *WorldOfWarcraft) GetAuctionDataJSON(realmSlug string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + auctionDataPath + "/" + realmSlug + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetAuctionData puts auction data information into AuctionData structure for realm
func (w *WorldOfWarcraft) GetAuctionData(realmSlug string) (*AuctionData, error) {
	var (
		auctionData AuctionData
		json        *[]byte
		err         error
	)

	json, err = w.GetAuctionDataJSON(realmSlug)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &auctionData)
	if err != nil {
		return nil, err
	}

	return &auctionData, nil
}

// GetAuctions puts all auctions into array of Auctions structures for realm
func (w *WorldOfWarcraft) GetAuctions(realmSlug string) ([]*Auctions, error) {
	var (
		auctionsArr []*Auctions
		auctionData *AuctionData
		err         error
	)

	auctionData, err = w.GetAuctionData(realmSlug)
	if err != nil {
		return nil, err
	}

	for _, file := range auctionData.Files {
		var (
			json     []byte
			auctions Auctions
		)

		err = blizzard.GetURLBody(file.URL, &json)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		err = blizzard.GetStruct(&json, &auctions)
		if err != nil {
			return nil, err
		}

		auctionsArr = append(auctionsArr, &auctions)
	}

	return auctionsArr, err
}

// GetBossIndexJSON gets boss index JSON information
func (w *WorldOfWarcraft) GetBossIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + bossPath + "/?" + localeQuery + w.Locale + "&" + apiKeyQuery +
		w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetBossIndex puts boss index information into RealmIndex structure
func (w *WorldOfWarcraft) GetBossIndex() (*BossIndex, error) {
	var (
		bossIndex BossIndex
		json      *[]byte
		err       error
	)

	json, err = w.GetBossIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &bossIndex)
	if err != nil {
		return nil, err
	}

	return &bossIndex, nil
}

// GetBossJSON gets specified boss JSON information
func (w *WorldOfWarcraft) GetBossJSON(bossID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + bossPath + "/" + strconv.Itoa(bossID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetBoss puts specified boss information into Boss structure
func (w *WorldOfWarcraft) GetBoss(bossID int) (*Boss, error) {
	var (
		boss Boss
		json *[]byte
		err  error
	)

	json, err = w.GetBossJSON(bossID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &boss)
	if err != nil {
		return nil, err
	}

	return &boss, nil
}

// GetChallengeRealmLeaderboardJSON gets specified challenge realm leaderboard JSON information
func (w *WorldOfWarcraft) GetChallengeRealmLeaderboardJSON(realmSlug string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + challengePath + "/" + realmSlug + "?" + localeQuery + w.Locale +
		"&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetChallengeRealmLeaderboard puts specified challenge realm leaderboard into ChallengeRealmLeaderboard structure
func (w *WorldOfWarcraft) GetChallengeRealmLeaderboard(realmSlug string) (*ChallengeRealmLeaderboard, error) {
	var (
		challengeRealmLeaderboard ChallengeRealmLeaderboard
		json                      *[]byte
		err                       error
	)

	json, err = w.GetChallengeRealmLeaderboardJSON(realmSlug)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &challengeRealmLeaderboard)
	if err != nil {
		return nil, err
	}

	return &challengeRealmLeaderboard, nil
}

// GetChallengeRegionLeaderboardJSON gets specified challenge region leaderboard JSON information
func (w *WorldOfWarcraft) GetChallengeRegionLeaderboardJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + challengePath + "/" + regionPath + "?" + localeQuery + w.Locale +
		"&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetChallengeRegionLeaderboard puts specified challenge region leaderboard into ChallengeRegionLeaderboard structure
func (w *WorldOfWarcraft) GetChallengeRegionLeaderboard() (*ChallengeRegionLeaderboard, error) {
	var (
		challengeRegionLeaderboard ChallengeRegionLeaderboard
		json                       *[]byte
		err                        error
	)

	json, err = w.GetChallengeRegionLeaderboardJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &challengeRegionLeaderboard)
	if err != nil {
		return nil, err
	}

	return &challengeRegionLeaderboard, nil
}

// GetCharacterJSON gets specified character JSON information
func (w *WorldOfWarcraft) GetCharacterJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacter puts specified character info into Character structure
func (w *WorldOfWarcraft) GetCharacter(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithAchievementsJSON gets specified character with achievements JSON information
func (w *WorldOfWarcraft) GetCharacterWithAchievementsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		achievementsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithAchievements puts character info with achievements into Character structure
func (w *WorldOfWarcraft) GetCharacterWithAchievements(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithAchievementsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithAppearanceJSON gets specified character with appearance JSON information
func (w *WorldOfWarcraft) GetCharacterWithAppearanceJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		appearanceField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithAppearance puts character info with appearance into Character structure
func (w *WorldOfWarcraft) GetCharacterWithAppearance(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithAppearanceJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithFeedJSON gets specified character with feed JSON information
func (w *WorldOfWarcraft) GetCharacterWithFeedJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		feedField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithFeed puts character info with feed into Character structure
func (w *WorldOfWarcraft) GetCharacterWithFeed(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithFeedJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithGuildJSON gets specified character with guild JSON information
func (w *WorldOfWarcraft) GetCharacterWithGuildJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		guildField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithGuild puts character info with guild into Character structure
func (w *WorldOfWarcraft) GetCharacterWithGuild(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithGuildJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithHunterPetsJSON gets specified character with hunterPets JSON information
func (w *WorldOfWarcraft) GetCharacterWithHunterPetsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		hunterPetsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithHunterPets puts character info with hunterPets into Character structure
func (w *WorldOfWarcraft) GetCharacterWithHunterPets(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithHunterPetsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithItemsJSON gets specified character with items JSON information
func (w *WorldOfWarcraft) GetCharacterWithItemsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		itemsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithItems puts character info with items into Character structure
func (w *WorldOfWarcraft) GetCharacterWithItems(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithItemsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithMountsJSON gets specified character with mounts JSON information
func (w *WorldOfWarcraft) GetCharacterWithMountsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		mountsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithMounts puts character info with mounts into Character structure
func (w *WorldOfWarcraft) GetCharacterWithMounts(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithMountsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithPetsJSON gets specified character with pets JSON information
func (w *WorldOfWarcraft) GetCharacterWithPetsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		petsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithPets puts character info with pets into Character structure
func (w *WorldOfWarcraft) GetCharacterWithPets(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithPetsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithPetSlotsJSON gets specified character with petSlots JSON information
func (w *WorldOfWarcraft) GetCharacterWithPetSlotsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		petSlotsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithPetSlots puts character info with petSlots into Character structure
func (w *WorldOfWarcraft) GetCharacterWithPetSlots(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithPetSlotsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithProfessionsJSON gets specified character with professions JSON information
func (w *WorldOfWarcraft) GetCharacterWithProfessionsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		professionsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithProfessions puts character info with professions into Character structure
func (w *WorldOfWarcraft) GetCharacterWithProfessions(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithProfessionsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithProgressionJSON gets specified character with progression JSON information
func (w *WorldOfWarcraft) GetCharacterWithProgressionJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		progressionField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithProgression puts character info with progression into Character structure
func (w *WorldOfWarcraft) GetCharacterWithProgression(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithProgressionJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithPVPJSON gets specified character with pvp JSON information
func (w *WorldOfWarcraft) GetCharacterWithPVPJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		pvpField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithPVP puts character info with pvp into Character structure
func (w *WorldOfWarcraft) GetCharacterWithPVP(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithPVPJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithQuestsJSON gets specified character with quests JSON information
func (w *WorldOfWarcraft) GetCharacterWithQuestsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		questsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithQuests puts character info with quests into Character structure
func (w *WorldOfWarcraft) GetCharacterWithQuests(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithQuestsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithReputationJSON gets specified character with reputation JSON information
func (w *WorldOfWarcraft) GetCharacterWithReputationJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		reputationField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithReputation puts character info with reputation into Character structure
func (w *WorldOfWarcraft) GetCharacterWithReputation(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithReputationJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithStatisticsJSON gets specified character with statistics JSON information
func (w *WorldOfWarcraft) GetCharacterWithStatisticsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		statisticsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithStatistics puts character info with statistics into Character structure
func (w *WorldOfWarcraft) GetCharacterWithStatistics(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithStatisticsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithStatsJSON gets specified character with stats JSON information
func (w *WorldOfWarcraft) GetCharacterWithStatsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		statsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithStats puts character info with stats into Character structure
func (w *WorldOfWarcraft) GetCharacterWithStats(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithStatsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithTalentsJSON gets specified character with talents JSON information
func (w *WorldOfWarcraft) GetCharacterWithTalentsJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		talentsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithTalents puts character info with talents into Character structure
func (w *WorldOfWarcraft) GetCharacterWithTalents(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithTalentsJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithTitlesJSON gets specified character with titles JSON information
func (w *WorldOfWarcraft) GetCharacterWithTitlesJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		titlesField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithTitles puts character info with titles into Character structure
func (w *WorldOfWarcraft) GetCharacterWithTitles(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithTitlesJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithAuditJSON gets specified character with audit JSON information
func (w *WorldOfWarcraft) GetCharacterWithAuditJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		auditField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithAudit puts character info with audit into Character structure
func (w *WorldOfWarcraft) GetCharacterWithAudit(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithAuditJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetCharacterWithAllJSON gets specified character with all fields JSON information
func (w *WorldOfWarcraft) GetCharacterWithAllJSON(realm, characterName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + characterPath + "/" + realm + "/" + characterName + "?" + fieldsQuery +
		achievementsField + "," + appearanceField + "," + feedField + "," + guildField + "," +
		hunterPetsField + "," + itemsField + "," + mountsField + "," + petsField + "," +
		petSlotsField + "," + professionsField + "," + progressionField + "," + pvpField + "," +
		questsField + "," + reputationField + "," + statisticsField + "," + statsField + "," +
		talentsField + "," + titlesField + "," + auditField + "&" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetCharacterWithAll puts character info with all fields into Character structure
func (w *WorldOfWarcraft) GetCharacterWithAll(realm, characterName string) (*Character, error) {
	var (
		character Character
		json      *[]byte
		err       error
	)

	json, err = w.GetCharacterWithAllJSON(realm, characterName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

// GetGuildJSON gets specified guild JSON information
func (w *WorldOfWarcraft) GetGuildJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuild puts guild info into Guild structure
func (w *WorldOfWarcraft) GetGuild(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuildWithMembersJSON gets specified guild with members JSON information
func (w *WorldOfWarcraft) GetGuildWithMembersJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + fieldsQuery +
		membersField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuildWithMembers puts guild info with members into Guild structure
func (w *WorldOfWarcraft) GetGuildWithMembers(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildWithMembersJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuildWithAchievementsJSON gets specified guild with achievements JSON information
func (w *WorldOfWarcraft) GetGuildWithAchievementsJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + fieldsQuery +
		achievementsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuildWithAchievements puts guild info with achievements into Guild structure
func (w *WorldOfWarcraft) GetGuildWithAchievements(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildWithAchievementsJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuildWithNewsJSON gets specified guild with news JSON information
func (w *WorldOfWarcraft) GetGuildWithNewsJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + fieldsQuery +
		newsField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuildWithNews puts guild info with news into Guild structure
func (w *WorldOfWarcraft) GetGuildWithNews(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildWithNewsJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuildWithChallengeJSON gets specified guild with challenge JSON information
func (w *WorldOfWarcraft) GetGuildWithChallengeJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + fieldsQuery +
		challengeField + "&" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuildWithChallenge puts guild info with challenge into Guild structure
func (w *WorldOfWarcraft) GetGuildWithChallenge(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildWithChallengeJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuildWithAllJSON gets specified guild with all fields JSON information
func (w *WorldOfWarcraft) GetGuildWithAllJSON(realm, guildName string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + guildPath + "/" + realm + "/" + guildName + "?" + fieldsQuery +
		membersField + "," + achievementsField + "," + newsField + "," + challengeField + "&" +
		localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetGuildWithAll puts guild info with all into Guild structure
func (w *WorldOfWarcraft) GetGuildWithAll(realm, guildName string) (*Guild, error) {
	var (
		guild Guild
		json  *[]byte
		err   error
	)

	json, err = w.GetGuildWithAllJSON(realm, guildName)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetItemJSON gets specified item JSON information
func (w *WorldOfWarcraft) GetItemJSON(itemID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + itemPath + "/" + strconv.Itoa(itemID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetItem puts item info into Item structure
func (w *WorldOfWarcraft) GetItem(itemID int) (*Item, error) {
	var (
		item Item
		json *[]byte
		err  error
	)

	json, err = w.GetItemJSON(itemID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetSetJSON gets specified set JSON information
func (w *WorldOfWarcraft) GetSetJSON(setID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + itemPath + setPath + "/" + strconv.Itoa(setID) + "?" +
		localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetSet puts set info into Set structure
func (w *WorldOfWarcraft) GetSet(setID int) (*Set, error) {
	var (
		set  Set
		json *[]byte
		err  error
	)

	json, err = w.GetSetJSON(setID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &set)
	if err != nil {
		return nil, err
	}

	return &set, nil
}

// GetMountIndexJSON gets mount JSON information
func (w *WorldOfWarcraft) GetMountIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + mountPath + "/?" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetMountIndex puts mount info into MountIndex structure
func (w *WorldOfWarcraft) GetMountIndex() (*MountIndex, error) {
	var (
		mountIndex MountIndex
		json       *[]byte
		err        error
	)

	json, err = w.GetMountIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &mountIndex)
	if err != nil {
		return nil, err
	}

	return &mountIndex, nil
}

// GetPetIndexJSON gets pet JSON information
func (w *WorldOfWarcraft) GetPetIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + petPath + "/?" + localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetPetIndex puts pet info into PetIndex structure
func (w *WorldOfWarcraft) GetPetIndex() (*PetIndex, error) {
	var (
		petIndex PetIndex
		json     *[]byte
		err      error
	)

	json, err = w.GetPetIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &petIndex)
	if err != nil {
		return nil, err
	}

	return &petIndex, nil
}

// GetPetAbilityJSON gets pet ability JSON information
func (w *WorldOfWarcraft) GetPetAbilityJSON(abilityID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + petPath + abilityPath + "/" + strconv.Itoa(abilityID) + "?" +
		localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetPetAbility puts pet ability info into PetAbility structure
func (w *WorldOfWarcraft) GetPetAbility(abilityID int) (*PetAbility, error) {
	var (
		petAbility PetAbility
		json       *[]byte
		err        error
	)

	json, err = w.GetPetAbilityJSON(abilityID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &petAbility)
	if err != nil {
		return nil, err
	}

	return &petAbility, nil
}

// GetPetSpeciesJSON gets pet species JSON information
func (w *WorldOfWarcraft) GetPetSpeciesJSON(speciesID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + petPath + speciesPath + "/" + strconv.Itoa(speciesID) + "?" +
		localeQuery + w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetPetSpecies puts pet species info into PetSpecies structure
func (w *WorldOfWarcraft) GetPetSpecies(speciesID int) (*PetSpecies, error) {
	var (
		petSpecies PetSpecies
		json       *[]byte
		err        error
	)

	json, err = w.GetPetSpeciesJSON(speciesID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &petSpecies)
	if err != nil {
		return nil, err
	}

	return &petSpecies, nil
}

// GetPetStatsJSON gets pet stats JSON information
func (w *WorldOfWarcraft) GetPetStatsJSON(speciesID, level, breedID, qualityID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + petPath + statsPath + "/" + strconv.Itoa(speciesID) + "?" +
		levelQuery + strconv.Itoa(level) + "&" + breedIDQuery + strconv.Itoa(breedID) + "&" +
		qualityIDQuery + strconv.Itoa(qualityID) + "&" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetPetStats puts pet stats info into PetStats structure
func (w *WorldOfWarcraft) GetPetStats(speciesID, level, breedID, qualityID int) (*PetStats, error) {
	var (
		petStats PetStats
		json     *[]byte
		err      error
	)

	json, err = w.GetPetStatsJSON(speciesID, level, breedID, qualityID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &petStats)
	if err != nil {
		return nil, err
	}

	return &petStats, nil
}

// Get2v2LeaderboardJSON gets 2v2 PvP leaderboard JSON information
func (w *WorldOfWarcraft) Get2v2LeaderboardJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + leaderboardPath + "/2v2?" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// Get2v2Leaderboard puts 2v2 PvP leaderboard info into Leaderboard structure
func (w *WorldOfWarcraft) Get2v2Leaderboard() (*Leaderboard, error) {
	var (
		leaderboard Leaderboard
		json        *[]byte
		err         error
	)

	json, err = w.Get2v2LeaderboardJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &leaderboard)
	if err != nil {
		return nil, err
	}

	return &leaderboard, nil
}

// Get3v3LeaderboardJSON gets 3v3 PvP leaderboard JSON information
func (w *WorldOfWarcraft) Get3v3LeaderboardJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + leaderboardPath + "/3v3?" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// Get3v3Leaderboard puts 3v3 PvP leaderboard info into Leaderboard structure
func (w *WorldOfWarcraft) Get3v3Leaderboard() (*Leaderboard, error) {
	var (
		leaderboard Leaderboard
		json        *[]byte
		err         error
	)

	json, err = w.Get3v3LeaderboardJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &leaderboard)
	if err != nil {
		return nil, err
	}

	return &leaderboard, nil
}

// Get5v5LeaderboardJSON gets 5v5 PvP leaderboard JSON information
func (w *WorldOfWarcraft) Get5v5LeaderboardJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + leaderboardPath + "/5v5?" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// Get5v5Leaderboard puts 5v5 PvP leaderboard info into Leaderboard structure
func (w *WorldOfWarcraft) Get5v5Leaderboard() (*Leaderboard, error) {
	var (
		leaderboard Leaderboard
		json        *[]byte
		err         error
	)

	json, err = w.Get5v5LeaderboardJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &leaderboard)
	if err != nil {
		return nil, err
	}

	return &leaderboard, nil
}

// GetRBGLeaderboardJSON gets RBG PvP leaderboard JSON information
func (w *WorldOfWarcraft) GetRBGLeaderboardJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + leaderboardPath + "/rbg?" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRBGLeaderboard puts RBG PvP leaderboard info into Leaderboard structure
func (w *WorldOfWarcraft) GetRBGLeaderboard() (*Leaderboard, error) {
	var (
		leaderboard Leaderboard
		json        *[]byte
		err         error
	)

	json, err = w.GetRBGLeaderboardJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &leaderboard)
	if err != nil {
		return nil, err
	}

	return &leaderboard, nil
}

// GetQuestJSON gets pet quest JSON information
func (w *WorldOfWarcraft) GetQuestJSON(questID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + questPath + "/" + strconv.Itoa(questID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetQuest puts pet quest info into Quest structure
func (w *WorldOfWarcraft) GetQuest(questID int) (*Quest, error) {
	var (
		quest Quest
		json  *[]byte
		err   error
	)

	json, err = w.GetQuestJSON(questID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &quest)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}

// GetRealmStatusJSON gets realm status JSON information
func (w *WorldOfWarcraft) GetRealmStatusJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + realmPath + statusPath + "?" + localeQuery + w.Locale + "&" +
		apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRealmStatus puts realm status info into RealmStatus structure
func (w *WorldOfWarcraft) GetRealmStatus() (*RealmStatus, error) {
	var (
		realmStatus RealmStatus
		json        *[]byte
		err         error
	)

	json, err = w.GetRealmStatusJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &realmStatus)
	if err != nil {
		return nil, err
	}

	return &realmStatus, nil
}

// GetRecipeJSON gets pet recipe JSON information
func (w *WorldOfWarcraft) GetRecipeJSON(recipeID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + recipePath + "/" + strconv.Itoa(recipeID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetRecipe puts pet recipe info into Recipe structure
func (w *WorldOfWarcraft) GetRecipe(recipeID int) (*Recipe, error) {
	var (
		recipe Recipe
		json   *[]byte
		err    error
	)

	json, err = w.GetRecipeJSON(recipeID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

// GetSpellJSON gets pet spell JSON information
func (w *WorldOfWarcraft) GetSpellJSON(spellID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = w.CommunityURL + spellPath + "/" + strconv.Itoa(spellID) + "?" + localeQuery +
		w.Locale + "&" + apiKeyQuery + w.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetSpell puts pet spell info into Spell structure
func (w *WorldOfWarcraft) GetSpell(spellID int) (*Spell, error) {
	var (
		spell Spell
		json  *[]byte
		err   error
	)

	json, err = w.GetSpellJSON(spellID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &spell)
	if err != nil {
		return nil, err
	}

	return &spell, nil
}
