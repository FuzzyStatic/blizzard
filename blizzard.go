/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:40:31
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-16 19:32:55
 */

// Package blizzard is the top level library needed to use the
// API calls for Blizzard games within the subpackages.
package blizzard

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// JSON interface for struture creation
type JSON interface {
	JSON2Struct(*[]byte) error
}

// Auth contains access token and api key
type Auth struct {
	AccessToken string
	APIKey      string
}

// Blizzard regional API URLs, locale, access token, api key
type Blizzard struct {
	Auth
	CommunityURL string
}

// New create new WorldOfWarcraft structure
func New(auth Auth, region Region) *Blizzard {
	var b = Blizzard{Auth: auth}

	switch region {
	case EU:
		b.CommunityURL = "https://eu.api.battle.net"
	case KR:
		b.CommunityURL = "https://kr.api.battle.net"
	case SEA:
		b.CommunityURL = "https://sea.api.battle.net"
	case TW:
		b.CommunityURL = "https://tb.api.battle.net"
	case US:
		b.CommunityURL = "https://us.api.battle.net"
	default: // USA! USA!
		b.CommunityURL = "https://us.api.battle.net"
	}

	return &b
}

// GetStruct returns structure of JSON interface
func GetStruct(b *[]byte, v JSON) error {
	return v.JSON2Struct(b)
}

// GetURLBody fills body of url
func GetURLBody(url string, body *[]byte) error {
	var (
		req *http.Request
		res *http.Response
		err error
	)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	*body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return errors.New(res.Status)
	}

	return nil
}

// GetAccountUserJSON gets account user JSON information
func (b *Blizzard) GetAccountUserJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = b.CommunityURL + accountPath + userPath + "?" + accessTokenQuery + b.AccessToken

	err = GetURLBody(url, &json)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &json, nil
}

// GetAccountUser puts account user info into AccountUser structure
func (b *Blizzard) GetAccountUser() (*AccountUser, error) {
	var (
		accountUser AccountUser
		json        *[]byte
		err         error
	)

	json, err = b.GetAccountUserJSON()
	if err != nil {
		return nil, err
	}

	err = GetStruct(json, &accountUser)
	if err != nil {
		return nil, err
	}

	return &accountUser, nil
}
