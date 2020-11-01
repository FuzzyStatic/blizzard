// Package blizzard is a client library designed to make calling and processing Blizzard Game APIs simple
package blizzard

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// For testing
var c *Client

// Client regional API URLs, locale, client ID, client secret
type Client struct {
	client                 *http.Client
	cfg                    clientcredentials.Config
	authorizedCfg          oauth2.Config
	oauth                  OAuth
	oauthURL               string
	apiURL                 string
	dynamicNamespace       string
	profileNamespace       string
	staticNamespace        string
	staticClassicNamespace string
	region                 Region
	locale                 Locale
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
// enUS, esMX, ptBR, enGB, esES, frFR, ruRU, deDE, ptPT, itIT, koKR, zhTW, zhCN
type Locale string

func (locale Locale) String() string {
	return string(locale)
}

// Locale constants
const (
	DeDE = Locale("de_DE")
	EnUS = Locale("en_US")
	EsES = Locale("es_ES")
	EsMX = Locale("es_MX")
	FrFR = Locale("fr_FR")
	ItIT = Locale("it_IT")
	JaJP = Locale("ja_JP")
	KoKR = Locale("ko_KR")
	PlPL = Locale("pl_PL")
	PtBR = Locale("pt_BR")
	RuRU = Locale("ru_RU")
	ThTH = Locale("th_TH")
	ZhCN = Locale("zh_CN")
	ZhTW = Locale("zh_TW")
)

// NewClient create new Blizzard structure. This structure will be used to acquire your access token and make API calls.
func NewClient(clientID, clientSecret string, region Region, locale Locale) *Client {
	var c = Client{
		oauth: OAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		locale: locale,
	}

	c.cfg = clientcredentials.Config{
		ClientID:     c.oauth.ClientID,
		ClientSecret: c.oauth.ClientSecret,
	}

	c.SetRegion(region)

	return &c
}

// SetLocale changes the Locale of the client
func (c *Client) SetLocale(locale Locale) {
	c.locale = locale
}

// SetRegion changes the Region of the client
func (c *Client) SetRegion(region Region) {
	c.region = region

	switch region {
	case CN:
		c.oauthURL = "https://www.battlenet.com.cn"
		c.apiURL = "https://gateway.battlenet.com.cn"
		c.dynamicNamespace = "dynamic-zh"
		c.profileNamespace = "profile-zh"
		c.staticNamespace = "static-zh"
		c.staticClassicNamespace = "static-classic-zh"
	default:
		c.oauthURL = fmt.Sprintf("https://%s.battle.net", region)
		c.apiURL = fmt.Sprintf("https://%s.api.blizzard.com", region)
		c.dynamicNamespace = fmt.Sprintf("dynamic-%s", region)
		c.profileNamespace = fmt.Sprintf("profile-%s", region)
		c.staticNamespace = fmt.Sprintf("static-%s", region)
		c.staticClassicNamespace = fmt.Sprintf("static-classic-%s", region)
	}

	c.cfg.TokenURL = c.oauthURL + "/oauth/token"
	c.client = c.cfg.Client(context.Background())
}

// getURLBody processes simple GET request based on URL
func (c *Client) getURLBody(ctx context.Context, url, namespace string) ([]byte, error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if namespace != "" {
		req.Header.Set("Battlenet-Namespace", namespace)
	}

	res, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	return body, nil
}
