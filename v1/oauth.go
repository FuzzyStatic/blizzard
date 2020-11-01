package blizzard

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/FuzzyStatic/blizzard/oauth"
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
			AuthURL:  c.oauthURL + "/oauth/authorize",
			TokenURL: c.oauthURL + "/oauth/token",
		},
	}

	return c.authorizedCfg
}

// AccessTokenRequest retrieves new OAuth2 Token
func (c *Client) AccessTokenRequest() error {
	var (
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	req, err = http.NewRequest("POST", c.oauthURL+"/oauth/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &c.oauth.Token)
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

	req, err = http.NewRequest("GET", c.oauthURL+"/oauth/userinfo", nil)
	if err != nil {
		return &dat, b, err
	}

	client = c.authorizedCfg.Client(context.Background(), token)

	res, err = client.Do(req)
	if err != nil {
		return &dat, b, err
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
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
func (c *Client) TokenValidation(token *oauth2.Token) (*oauth.TokenValidation, []byte, error) {
	var (
		dat oauth.TokenValidation
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	req, err = http.NewRequest("GET", c.oauthURL+fmt.Sprintf("/oauth/check_token?token=%s", token.AccessToken), nil)
	if err != nil {
		return &dat, b, err
	}

	req.Header.Set("Accept", "application/json")

	res, err = c.client.Do(req)
	if err != nil {
		return &dat, b, err
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
