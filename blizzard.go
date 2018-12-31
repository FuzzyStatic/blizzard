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
	client           *http.Client
	oauth            OAuth
	oauthURL         string
	apiURL           string
	dynamicNamespace string
	profileNamespace string
	staticNamespace  string
	locale           Locale
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

func (region Region) String() string {
	var rr = []string{
		"",
		"us",
		"eu",
		"kr",
		"tw",
		"cn",
	}

	return rr[region]
}

// Locale generic locale string
// erfsf, enUS, esMX, ptBR, enGB, esES, frFR, ruRU, deDE, ptPT, itIT, koKR, zhTW, zhCN
type Locale string

func (locale Locale) String() string {
	return string(locale)
}

// Locale constants
const (
	enUS = Locale("en_US")
	esMX = Locale("es_MX")
	ptBR = Locale("pt_BR")
	enGB = Locale("en_GB")
	esES = Locale("es_ES")
	frFR = Locale("fr_FR")
	ruRU = Locale("ru_RU")
	deDE = Locale("de_DE")
	ptPT = Locale("pt_PT")
	itIT = Locale("it_IT")
	koKR = Locale("ko_KR")
	zhTW = Locale("zh_TW")
	zhCN = Locale("zh_CN")
)

// Path constants
const (
	localeQuery = "locale="
	dataPath    = "/data"
	profilePath = "/profile"
)

// New create new Blizzard structure. This structure will be used to acquire your access token and make API calls.
func New(clientID, clientSecret string, region Region, locale Locale) *Config {
	var c = Config{
		client: &http.Client{
			Timeout: time.Second * time.Duration(60),
		},
		oauth: OAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			ExpiresAt:    time.Now(),
		},
		locale: locale,
	}

	switch region {
	case CN:
		c.oauthURL = "https://www.battlenet.com.cn"
		c.apiURL = "https://gateway.battlenet.com.cn"
		c.dynamicNamespace = "dynamic-zh"
		c.profileNamespace = "profile-zh"
		c.staticNamespace = "static-zh"
	default:
		c.oauthURL = fmt.Sprintf("https://%s.battle.net/", region)
		c.apiURL = fmt.Sprintf("https://%s.api.blizzard.com", region)
		c.dynamicNamespace = fmt.Sprintf("dynamic-%s", region)
		c.profileNamespace = fmt.Sprintf("profile-%s", region)
		c.staticNamespace = fmt.Sprintf("static-%s", region)
	}

	return &c
}

// getURLBody processes simple GET request based on URL
func (c *Config) getURLBody(url, namespace string) ([]byte, error) {
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

	req.Header.Set("Authorization", "Bearer "+c.oauth.AccessTokenRequest.AccessToken)
	req.Header.Set("Accept", "application/json")

	if namespace != "" {
		req.Header.Set("Battlenet-Namespace", namespace)
	}

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
