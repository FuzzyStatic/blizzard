// Package blizzard is a client library designed to make calling and processing Blizzard Game APIs simple
package blizzard

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/FuzzyStatic/blizzard/v3/wowp"
	"github.com/FuzzyStatic/blizzard/v3/wowsearch"
	"github.com/avast/retry-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/time/rate"
)

// RateLimitConfig contains values for rate limiter configuration.
type RateLimitConfig struct {
	// Enabled whether the rate limiter is enabled.
	Enabled bool
	// Rate  how many requests per second can be made, before throttling occurs.
	Rate rate.Limit
	// Burst the maximum number of burst requests that are allowed before they too become subject to rate limiting.
	Burst int
}

// BlizzardRateLimit rate limiter configuration fitting Blizzard's rate limits (36k/hour 100 burst)
var BlizzardRateLimit = RateLimitConfig{
	Enabled: true,
	Rate:    10,
	Burst:   100,
}

// RetriesConfig contains values for retry attempts when Blizzard's API returns 429 error.
type RetriesConfig struct {
	// Enabled whether the retry feature is enabled.
	Enabled bool
	// Attempts how many retries are attempted before the request fails with an error.
	Attempts uint
	// Delay how much time must pass before another try is attempted
	Delay time.Duration
}

// Config contains values for Blizzard client creation
type Config struct {
	// ClientID is the client ID value from a Blizzard developer
	// account.
	ClientID string `validate:"required"`

	// ClientSecret is the client secret value from a Blizzard
	// developer account.
	ClientSecret string `validate:"required"`

	// HTTP client used for Blizzard API calls.
	HTTPClient *http.Client `validate:"required"`

	// Region - API data is limited to specific regions. For example,
	// US APIs accessed through us.battle.net only contain data from
	// US battlegroups and realms. Locale support is limited to
	// locations supported on Blizzard community game sites.
	// Game information is different from region to region. For
	// example, a user that has both US and EU WoW accounts has
	// different characters, achievements, and other information
	// in each region. Likewise, a D3 user who has a single license
	// for D3 on their account has a different character list for
	// each region in which the game operates.
	Region Region `validate:"required"`

	// Locale - All available API resources provide localized strings
	// using the locale query string parameter. Supported locales vary
	// from region to region and align with those supported on Blizzard
	// community sites.
	Locale Locale `validate:"required"`

	// RateLimit configures the rate limiter. It allows limiting outgoing
	// requests to Blizzard's API before they get limited by Blizzard.
	// The rate limiter is disabled by default.
	RateLimit RateLimitConfig

	// Retries configures request retries. It automatically retries a request
	// when Blizzard's API responds with rate limiting (Error 429 Too Many Requests).
	// Request retries are disabled by default.
	Retries RetriesConfig
}

// Client regional API URLs, locale, client ID, client secret
type Client struct {
	cfg                                             Config
	httpClient                                      *http.Client
	clntCredCfg                                     clientcredentials.Config
	authorizedCfg                                   oauth2.Config
	oauth                                           OAuth
	oauthHost                                       string
	apiHost                                         string
	dynamicNamespace, staticNamespace               string
	profileNamespace                                string
	dynamicClassicNamespace, staticClassicNamespace string
	region                                          Region
	locale                                          Locale
	ratelimiter                                     *rate.Limiter
	retryopts                                       []retry.Option
}

//go:generate stringer -type=Region -linecomment

// Region type
type Region int

// Region constants (1=US, 2=EU, 3=KO and TW, 5=CN) DO NOT REARRANGE
const (
	_  Region = iota
	US        // us
	EU        // eu
	KR        // kr
	TW        // tw
	CN        // cn
)

// Locale generic locale string
type Locale string

// Locale constants
const (
	// North America locales
	EnUS = Locale("en_US")
	EsMX = Locale("es_MX")
	PtBR = Locale("pt_BR")

	// Europe locales
	EnGB = Locale("en_GB")
	EsES = Locale("es_ES")
	FrFR = Locale("fr_FR")
	RuRU = Locale("ru_RU")
	PtPT = Locale("pt_PT")
	DeDE = Locale("de_DE")
	ItIT = Locale("it_IT")

	// Korea locales
	KoKR = Locale("ko_KR")

	// Taiwan locales
	ZhTW = Locale("zh_TW")

	// China locales
	ZhCN = Locale("zh_CN")
)

func (locale Locale) String() string {
	return string(locale)
}

// NewClient create new Blizzard structure. This structure will be used to acquire your access token and make API calls.
func NewClient(cfg Config) (*Client, error) {
	var (
		c   Client
		v   *validator.Validate
		err error
	)

	v = validator.New()

	err = v.Struct(cfg)
	if err != nil {
		return nil, err
	}

	c.cfg = cfg
	c.clntCredCfg = clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	}

	c.oauth = OAuth{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
	}

	err = c.SetRegionParameters(cfg.Region, cfg.Locale)
	if err != nil {
		return nil, err
	}

	if ratelimit := cfg.RateLimit; ratelimit.Enabled {
		c.ratelimiter = rate.NewLimiter(ratelimit.Rate, ratelimit.Burst)
	}

	if retries := cfg.Retries; retries.Enabled {
		c.retryopts = []retry.Option{
			retry.Attempts(retries.Attempts),
			retry.Delay(retries.Delay),
			retry.DelayType(retry.BackOffDelay),
			retry.MaxJitter(0),
			retry.RetryIf(func(err error) bool {
				return err.Error() == "429 Too Many Requests"
			}),
		}
	}

	return &c, nil
}

// GetLocale returns the Locale of the client
func (c *Client) GetLocale() Locale {
	return c.locale
}

// GetRegion returns the Region of the client
func (c *Client) GetRegion() Region {
	return c.region
}

// SetRegionParameters changes the Region parameters of the client
func (c *Client) SetRegionParameters(region Region, locale Locale) error {
	err := validateRegionLocalePair(region, locale)
	if err != nil {
		return err
	}

	c.region = region
	c.locale = locale

	switch region {
	case CN:
		c.oauthHost = "https://oauth.battlenet.com.cn"
		c.apiHost = "https://gateway.battlenet.com.cn"
		c.dynamicNamespace = "dynamic-zh"
		c.dynamicClassicNamespace = "dynamic-classic-zh"
		c.profileNamespace = "profile-zh"
		c.staticNamespace = "static-zh"
		c.staticClassicNamespace = "static-classic-zh"
	default:
		c.oauthHost = "https://oauth.battle.net"
		c.apiHost = fmt.Sprintf("https://%s.api.blizzard.com", region)
		c.dynamicNamespace = fmt.Sprintf("dynamic-%s", region)
		c.dynamicClassicNamespace = fmt.Sprintf("dynamic-classic-%s", region)
		c.profileNamespace = fmt.Sprintf("profile-%s", region)
		c.staticNamespace = fmt.Sprintf("static-%s", region)
		c.staticClassicNamespace = fmt.Sprintf("static-classic-%s", region)
	}

	c.clntCredCfg.TokenURL = c.oauthHost + "/token"
	c.httpClient = c.clntCredCfg.Client(
		context.WithValue(context.TODO(), oauth2.HTTPClient, c.cfg.HTTPClient),
	)

	return nil
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

// runHttpRequest runs the provided request, performs rate limiting and retries based on client configuration
func (c *Client) runHttpRequest(ctx context.Context, request *http.Request) (*http.Response, error) {
	if c.cfg.Retries.Enabled && ctx.Value("withRetries") != true {
		subContext := context.WithValue(ctx, "withRetries", true)
		options := []retry.Option{
			retry.Context(subContext),
		}
		options = append(options, c.retryopts...)
		var res *http.Response
		err := retry.Do(func() (err error) {
			res, err = c.runHttpRequest(subContext, request)

			if err != nil && res != nil {
				_ = res.Body.Close()
			}

			if res != nil && res.StatusCode >= 400 {
				return errors.New(res.Status)
			}

			return
		}, options...)

		return res, err
	}
	if c.ratelimiter != nil {
		err := c.ratelimiter.Wait(ctx)
		if err != nil {
			return nil, err
		}
	}

	return c.httpClient.Do(request)
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

	res, err := c.runHttpRequest(ctx, req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	header := getHeader(res.Header)
	if err != nil {
		return dat, nil, err
	}

	switch v := dat.(type) {
	case *wowp.CharacterProfileSummary:
		bnSchRev, err := strconv.ParseInt(header.BattlenetSchemaRevision, 10, 64)
		if err != nil {
			return dat, nil, err
		}

		if bnSchRev < 24 {
			fmt.Println("true")
			var temp *wowp.CharacterProfileSummaryPreRev24
			err = json.Unmarshal(body, &temp)
			if err != nil {
				return dat, nil, err
			}

			var activeTitle wowp.ActiveTitlePreRev24
			err = json.Unmarshal(body, &activeTitle)
			if err != nil {
				return dat, nil, err
			}

			wowp.ConvertCharacterProfileSummaryPreRev24(activeTitle.ActiveTitle, temp, v)
			break
		}

		err = json.Unmarshal(body, &dat)
		if err != nil {
			return dat, nil, err
		}
	default:
		err = json.Unmarshal(body, &dat)
		if err != nil {
			return dat, nil, err
		}
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

	res, err := c.runHttpRequest(ctx, req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return dat, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return dat, nil, errors.New(res.Status)
	}

	if err != nil {
		return dat, nil, err
	}

	err = json.Unmarshal(body, &dat)
	if err != nil {
		return dat, nil, err
	}

	return dat, getHeader(res.Header), nil
}

// getStructDataNoNamespace processes simple GET request based on pathAndQuery an returns the structured data.
// Does not use a namespace or Locale
func (c *Client) getStructDataNoNamespaceNoLocale(ctx context.Context, pathAndQuery string, dat interface{}) (interface{}, *Header, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+pathAndQuery, nil)
	if err != nil {
		return dat, nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err := c.runHttpRequest(ctx, req)
	if err != nil {
		return dat, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
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

	return dat, getHeader(res.Header), nil
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

	body, err := io.ReadAll(res.Body)
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

	return dat, getHeader(res.Header), nil
}

func validateRegionLocalePair(region Region, locale Locale) error {
	switch locale {
	case EnUS, EsMX, PtBR:
		if region != US {
			return ErrInvalidLocaleForRegion
		}
	case EnGB, EsES, FrFR, RuRU, PtPT, DeDE, ItIT:
		if region != EU {
			return ErrInvalidLocaleForRegion
		}
	case KoKR:
		if region != KR {
			return ErrInvalidLocaleForRegion
		}
	case ZhTW:
		if region != TW {
			return ErrInvalidLocaleForRegion
		}
	case ZhCN:
		if region != CN {
			return ErrInvalidLocaleForRegion
		}
	default:
		return ErrUnknownLocale
	}

	return nil
}

func formatAccount(account string) string {
	return strings.Replace(account, "#", "-", 1)
}
