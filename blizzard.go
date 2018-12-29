// Package blizzard is a client library designed to make calling and processing Blizzard Game APIs simple
package blizzard

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// For testing
var c *Config

// Config regional API URLs, locale, access token, api key
type Config struct {
	client    *http.Client
	oauth     OAuth
	region    Region
	oauthURL  string
	apiURL    string
	namespace string
	locale    string
}

// Region type
type Region int

// Region constants (1=US, 2=EU, 3=KO and TW, 5=CN) DO NOT REARRANGE
const (
	_ Region = iota
	US
	EU
	KR
	TW
	CN
)

// Path constants
const (
	localeQuery = "locale="
	dataPath    = "/data"
	profilePath = "/profile"
)

// New create new Blizzard structure. This structure will be used to acquire your access token and make API calls.
func New(clientID, clientSecret string, region Region) *Config {
	var c = Config{
		client: &http.Client{
			Timeout: time.Second * time.Duration(60),
		},
		oauth: OAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			ExpiresAt:    time.Now(),
		},
		region: region,
	}

	switch c.region {
	case CN:
		c.oauthURL = "https://www.battlenet.com.cn"
		c.apiURL = "https://gateway.battlenet.com.cn"
		c.namespace = "dynamic-zh"
		c.locale = "zh_CN"
	case EU:
		c.oauthURL = "https://eu.battle.net"
		c.apiURL = "https://eu.api.blizzard.com"
		c.namespace = "dynamic-eu"
		c.locale = "en_GB"
	case KR:
		c.oauthURL = "https://apac.battle.net"
		c.apiURL = "https://apac.api.blizzard.com"
		c.namespace = "dynamic-kr"
		c.locale = "ko_KR"
	case TW:
		c.oauthURL = "https://apac.battle.net"
		c.apiURL = "https://apac.api.blizzard.com"
		c.namespace = "dynamic-tw"
		c.locale = "zh_TW"
	case US:
		c.oauthURL = "https://us.battle.net"
		c.apiURL = "https://us.api.blizzard.com"
		c.namespace = "dynamic-us"
		c.locale = "en_US"
	default: // USA! USA!
		c.oauthURL = "https://us.battle.net"
		c.apiURL = "https://us.api.blizzard.com"
		c.namespace = "dynamic-us"
		c.locale = "en_US"
	}

	c.AccessTokenReq()

	return &c
}

// getURLBody processes simple GET request based on URL
func (c *Config) getURLBody(url string) ([]byte, error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	err = c.updateAccessTokenIfExp()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(url)

	req.Header.Set("Authorization", "Bearer "+c.oauth.AccessTokenRequest.AccessToken)
	req.Header.Set("Battlenet-Namespace", c.namespace)
	req.Header.Set("Accept", "application/json")

	res, err = c.client.Do(req)
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
