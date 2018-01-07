/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:37:59
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:37:59
 */

package worldofwarcraft

import (
	"errors"
	"go-blizzard/blizzard"
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
