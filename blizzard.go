/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:40:31
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-16 19:32:55
 */

// Package blizzard is the top level library needed to use the
// API calls for Config games within the subpackages.
package blizzard

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// JSON interface for struture creation
type JSON interface {
	JSON2Struct([]byte) error
}

// Auth contains access token and api key
type Auth struct {
	AccessToken string
	APIKey      string
}

// Config regional API URLs, locale, access token, api key
type Config struct {
	Client       *http.Client
	Auth         Auth
	Region       Region
	CommunityURL string
}

// New create new WorldOfWarcraft structure
func New(auth Auth, region Region) *Config {
	var c = Config{
		Client: &http.Client{
			Timeout: time.Second * time.Duration(60),
		},
		Auth:   auth,
		Region: region,
	}

	switch c.Region {
	case EU:
		c.CommunityURL = "https://eu.api.battle.net"
	case KR:
		c.CommunityURL = "https://kr.api.battle.net"
	case SEA:
		c.CommunityURL = "https://sea.api.battle.net"
	case TW:
		c.CommunityURL = "https://tb.api.battle.net"
	case US:
		c.CommunityURL = "https://us.api.battle.net"
	default: // USA! USA!
		c.CommunityURL = "https://us.api.battle.net"
	}

	return &c
}

// GetStruct returns structure of JSON interface
func GetStruct(b []byte, v JSON) error {
	return v.JSON2Struct(b)
}

// GetURLBody fills body of url
func (c *Config) GetURLBody(url string) ([]byte, error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err = c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New(res.Status)
	}

	return body, nil
}

// GetAccountUserJSON gets account user JSON information
func (c *Config) GetAccountUserJSON() ([]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = c.CommunityURL + accountPath + userPath + "?" + accessTokenQuery + c.Auth.AccessToken

	json, err = c.GetURLBody(url)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return json, nil
}

// GetAccountUser puts account user info into AccountUser structure
func (c *Config) GetAccountUser() (*AccountUser, error) {
	var (
		accountUser AccountUser
		json        []byte
		err         error
	)

	json, err = c.GetAccountUserJSON()
	if err != nil {
		return nil, err
	}

	err = GetStruct(json, &accountUser)
	if err != nil {
		return nil, err
	}

	return &accountUser, nil
}
