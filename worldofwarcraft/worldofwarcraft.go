/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:37:59
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-08 21:54:04
 */

package worldofwarcraft

import (
	"errors"
	"fmt"
	"go-blizzard"
	"strconv"
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
