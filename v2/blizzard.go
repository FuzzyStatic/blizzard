// Package blizzard is a client library designed to make calling and processing Blizzard Game APIs simple
package blizzard

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/FuzzyStatic/blizzard/v2/wowsearch"
	"github.com/avast/retry-go"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/time/rate"
)

// For testing
var c *Client

type ClientOpts interface {
	Apply(c *Client)
}

// Client regional API URLs, locale, client ID, client secret
type Client struct {
	httpClient                                      *http.Client
	cfg                                             clientcredentials.Config
	authorizedCfg                                   oauth2.Config
	oauth                                           OAuth
	oauthHost                                       string
	apiHost                                         string
	dynamicNamespace, staticNamespace               string
	profileNamespace                                string
	dynamicClassicNamespace, staticClassicNamespace string
	region                                          Region
	locale                                          Locale
	retryopts                                       []retry.Option
	ratelimiter                                     *rate.Limiter
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
func NewClient(clientID, clientSecret string, region Region, locale Locale, opts ...ClientOpts) *Client {
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

	for _, opt := range opts {
		opt.Apply(&c)
	}

	if c.ratelimiter == nil {
		c.ratelimiter = rate.NewLimiter(rate.Every(1*time.Second/100), 10)
	}

	if len(c.retryopts) == 0 {
		c.retryopts = []retry.Option{
			retry.Attempts(3),
			retry.Delay(100 * time.Millisecond),
			retry.DelayType(retry.BackOffDelay),
			retry.MaxJitter(0),
			retry.RetryIf(func(err error) bool {
				switch {
				case err.Error() == "429 Too Many Requests":
					return true // recoverable error, retry
				case err.Error() == "403 Forbidden":
					return false
				case err.Error() == "404 Not Found":
					return false
				default:
					return false // unhandled error
				}
			}),
		}
	}

	return &c
}

// GetLocale returns the Locale of the client
func (c *Client) GetLocale() Locale {
	return c.locale
}

// SetLocale changes the Locale of the client
func (c *Client) SetLocale(locale Locale) {
	c.locale = locale
}

// GetRegion returns the Region of the client
func (c *Client) GetRegion() Region {
	return c.region
}

// SetRegion changes the Region of the client
func (c *Client) SetRegion(region Region) {
	c.region = region

	switch region {
	case CN:
		c.oauthHost = "https://www.battlenet.com.cn"
		c.apiHost = "https://gateway.battlenet.com.cn"
		c.dynamicNamespace = "dynamic-zh"
		c.dynamicClassicNamespace = "dynamic-classic-zh"
		c.profileNamespace = "profile-zh"
		c.staticNamespace = "static-zh"
		c.staticClassicNamespace = "static-classic-zh"
	default:
		c.oauthHost = fmt.Sprintf("https://%s.battle.net", region)
		c.apiHost = fmt.Sprintf("https://%s.api.blizzard.com", region)
		c.dynamicNamespace = fmt.Sprintf("dynamic-%s", region)
		c.dynamicClassicNamespace = fmt.Sprintf("dynamic-classic-%s", region)
		c.profileNamespace = fmt.Sprintf("profile-%s", region)
		c.staticNamespace = fmt.Sprintf("static-%s", region)
		c.staticClassicNamespace = fmt.Sprintf("static-classic-%s", region)
	}

	c.cfg.TokenURL = c.oauthHost + "/oauth/token"

	defaultTransport := &http.Transport{
		Dial:                (&net.Dialer{KeepAlive: 10 * time.Second}).Dial,
		MaxIdleConns:        6,
		MaxIdleConnsPerHost: 2,
	}

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: defaultTransport,
	}
	ctx := context.WithValue(context.TODO(), oauth2.HTTPClient, httpClient)
	c.httpClient = c.cfg.Client(ctx)
}

// GetOAuthHost returns the OAuth host of the client
func (c *Client) GetOAuthHost() string {
	return c.oauthHost
}

// GetAPIHost returns the API host of the client
func (c *Client) GetAPIHost() string {
	return c.apiHost
}

// GetDynamicNamespace returns the dynamic namespace of the client
func (c *Client) GetDynamicNamespace() string {
	return c.dynamicNamespace
}

// GetDynamicClassicNamespace returns the classic dynamic namespace of the client
func (c *Client) GetDynamicClassicNamespace() string {
	return c.dynamicClassicNamespace
}

// GetProfileNamespace returns the profile namespace of the client
func (c *Client) GetProfileNamespace() string {
	return c.profileNamespace
}

// GetStaticNamespace returns the static namespace of the client
func (c *Client) GetStaticNamespace() string {
	return c.staticNamespace
}

// GetStaticClassicNamespace returns the classic static namespace of the client
func (c *Client) GetStaticClassicNamespace() string {
	return c.staticClassicNamespace
}

// buildSearchParams builds params for searches
func buildSearchParams(opts ...wowsearch.Opt) string {
	if len(opts) == 0 {
		return ""
	}
	var params []string
	for _, opt := range opts {
		opt.Apply(&params)
	}
	return "?" + strings.Join(params, "&")
}

func (c *Client) getRetryOpts(ctx context.Context) []retry.Option {
	opts := []retry.Option{
		retry.Context(ctx),
	}
	return append(opts, c.retryopts...)
}

// getStructData processes simple GET request based on pathAndQuery an returns the structured data.
func (c *Client) getStructData(ctx context.Context, pathAndQuery, namespace string, dat interface{}) (interface{}, *Header, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+pathAndQuery, nil)
	if err != nil {
		return dat, nil, err
	}
	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Set("locale", c.locale.String())
	req.URL.RawQuery = q.Encode()

	if namespace != "" {
		req.Header.Set("Battlenet-Namespace", namespace)
	}

	var res *http.Response
	err = retry.Do(
		func() (err error) {
			if err := c.ratelimiter.Wait(ctx); err != nil {
				return err
			}
			res, err = c.httpClient.Do(req)
			if err != nil && res != nil {
				res.Body.Close()
			}
			return
		},
		c.getRetryOpts(ctx)...,
	)
	if err != nil {
		return dat, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return dat, nil, err
	}

	header, err := getHeader(res.Header)
	if err != nil {
		return dat, nil, err
	}

	return dat, header, nil

}

// getStructDataNoNamespace processes simple GET request based on pathAndQuery an returns the structured data.
// Does not use a namespace.
func (c *Client) getStructDataNoNamespace(ctx context.Context, pathAndQuery string, dat interface{}) (interface{}, *Header, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+pathAndQuery, nil)
	if err != nil {
		return dat, nil, err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Set("locale", c.locale.String())
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return dat, nil, err
	}

	header, err := getHeader(res.Header)
	if err != nil {
		return dat, nil, err
	}

	return dat, header, nil
}

// getStructDataNoNamespace processes simple GET request based on pathAndQuery an returns the structured data.
// Does not use a namespace or Locale
func (c *Client) getStructDataNoNamespaceNoLocale(ctx context.Context, pathAndQuery string, dat interface{}) (interface{}, *Header, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+pathAndQuery, nil)
	if err != nil {
		return dat, nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return dat, nil, err
	}

	header, err := getHeader(res.Header)
	if err != nil {
		return dat, nil, err
	}

	return dat, header, nil
}

// getStructDataOAuth processes simple GET request based on pathAndQuery an returns the structured data.
// Uses OAuth2.
func (c *Client) getStructDataOAuth(ctx context.Context, pathAndQuery, namespace string,
	token *oauth2.Token, dat interface{}) (interface{}, *Header, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+pathAndQuery, nil)
	if err != nil {
		return dat, nil, err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Set("locale", c.locale.String())
	req.URL.RawQuery = q.Encode()

	if namespace != "" {
		req.Header.Set("Battlenet-Namespace", namespace)
	}

	client := c.authorizedCfg.Client(context.Background(), token)

	res, err := client.Do(req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return dat, nil, err
	}

	header, err := getHeader(res.Header)
	if err != nil {
		return dat, nil, err
	}

	return dat, header, nil
}

func formatAccount(account string) string {
	return strings.Replace(account, "#", "-", 1)
}
