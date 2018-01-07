/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:37:59
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 18:45:06
 */

package worldofwarcraft

import (
	"errors"
	"go-blizzard/blizzard"
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

	if json, err = w.GetConnectedRealmIndexJSON(); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &connectedRealmIndex); err != nil {
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

	if json, err = w.GetConnectedRealmJSON(connectRealmID); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &connectedRealm); err != nil {
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

	if json, err = w.GetMythicLeaderboardIndexJSON(connectRealmID); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &mythicLeaderboardIndex); err != nil {
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

	if json, err = w.GetMythicLeaderboardJSON(connectRealmID, dungeonID, period); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &mythicLeaderboard); err != nil {
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

	if json, err = w.GetRealmIndexJSON(); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &realmIndex); err != nil {
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

	if json, err = w.GetRealmJSON(realmID); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &realm); err != nil {
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

	if json, err = w.GetRegionIndexJSON(); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &regionIndex); err != nil {
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

	if json, err = w.GetRegionJSON(regionID); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &region); err != nil {
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

	if json, err = w.GetTokenIndexJSON(); err != nil {
		return nil, err
	}

	if err = blizzard.GetStruct(json, &tokenIndex); err != nil {
		return nil, err
	}

	return &tokenIndex, nil
}
