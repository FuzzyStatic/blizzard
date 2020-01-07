package blizzard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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

	cfg := oauth2.Config{
		ClientID:     c.oauth.ClientID,
		ClientSecret: c.oauth.ClientSecret,
		Scopes:       scopes,
		RedirectURL:  redirectURI,
		// This points to our Authorization Server
		// if our Client ID and Client Secret are valid
		// it will attempt to authorize our user
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.oauthURL + "/oauth/authorize",
			TokenURL: c.oauthURL + "/oauth/token",
		},
	}

	return cfg
}

// func (c *Client) SetToken(token *oauth2.Token) {
// 	c.oauth.Token = token
// }

// Token retrieves new OAuth2 Token
func (c *Client) Token() error {
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

	req.SetBasicAuth(c.oauth.ClientID, c.oauth.ClientSecret)
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

// updateAccessTokenIfExp updates Access Token if expired
func (c *Client) updateAccessTokenIfExp() error {
	var err error

	if c.oauth.Token.Expiry.Sub(time.Now().UTC()) < 60 {
		err = c.Token()
		if err != nil {
			return err
		}
	}

	return nil
}

// UserInfoHeader teturns basic information about the user associated with the current bearer token
func (c *Client) UserInfoHeader(token *oauth2.Token) (*oauth.UserInfo, []byte, error) {
	var (
		dat oauth.UserInfo
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	req, err = http.NewRequest("GET", c.oauthURL+"/oauth/userinfo", nil)
	if err != nil {
		return &dat, b, err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
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

	fmt.Println(string(b))

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// TokenValidation verify that a given bearer token is valid and retrieve metadata about the token including the client_id used to create the token, expiration timestamp, and scopes granted to the token
func (c *Client) TokenValidation() (*oauth.TokenValidation, []byte, error) {
	var (
		dat oauth.TokenValidation
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	err = c.updateAccessTokenIfExp()
	if err != nil {
		return &dat, b, err
	}

	req, err = http.NewRequest("GET", c.oauthURL+fmt.Sprintf("/oauth/check_token?token=%s", c.oauth.Token.AccessToken), nil)
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
