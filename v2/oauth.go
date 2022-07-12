package blizzard

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/FuzzyStatic/blizzard/v2/oauth"
	"golang.org/x/oauth2"
)

// OAuth credentials and access token to access Blizzard API
type OAuth struct {
	ClientID     string
	ClientSecret string
	Token        *oauth2.Token
}

// AuthorizeConfig returns OAuth2 config
func (c *Client) AuthorizeConfig(redirectURI string, profiles ...oauth.Profile) oauth2.Config {
	var scopes []string

	for _, profile := range profiles {
		scopes = append(scopes, string(profile))
	}

	c.authorizedCfg = oauth2.Config{
		ClientID:     c.oauth.ClientID,
		ClientSecret: c.oauth.ClientSecret,
		Scopes:       scopes,
		RedirectURL:  redirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.oauthHost + "/oauth/authorize",
			TokenURL: c.oauthHost + "/oauth/token",
		},
	}

	return c.authorizedCfg
}

// AccessTokenRequest retrieves new OAuth2 Token
func (c *Client) AccessTokenRequest(ctx context.Context) error {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	req, err = http.NewRequestWithContext(ctx, "POST", c.oauthHost+"/oauth/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Set("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(c.oauth.ClientID, c.oauth.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &c.oauth.Token)
	if err != nil {
		return err
	}

	return nil
}

// UserInfoHeader returns basic information about the user associated with the current bearer token
func (c *Client) UserInfoHeader(token *oauth2.Token) (*oauth.UserInfo, []byte, error) {
	var (
		dat    oauth.UserInfo
		req    *http.Request
		client *http.Client
		res    *http.Response
		b      []byte
		err    error
	)

	req, err = http.NewRequest("GET", c.oauthHost+"/oauth/userinfo", nil)
	if err != nil {
		return &dat, b, err
	}

	client = c.authorizedCfg.Client(context.Background(), token)

	res, err = client.Do(req)
	if err != nil {
		return &dat, b, err
	}
	defer res.Body.Close()

	b, err = io.ReadAll(res.Body)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// TokenValidation verify that a given bearer token is valid and retrieve metadata about the token including the client_id used to create the token, expiration timestamp, and scopes granted to the token
func (c *Client) TokenValidation(ctx context.Context, token *oauth2.Token) (*oauth.TokenValidation, []byte, error) {
	var (
		dat oauth.TokenValidation
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	req, err = http.NewRequestWithContext(ctx, "GET", c.oauthHost+fmt.Sprintf("/oauth/check_token?token=%s", token.AccessToken), nil)
	if err != nil {
		return &dat, b, err
	}

	req.Header.Set("Accept", "application/json")

	res, err = c.httpClient.Do(req)
	if err != nil {
		return &dat, b, err
	}
	defer res.Body.Close()

	b, err = io.ReadAll(res.Body)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
