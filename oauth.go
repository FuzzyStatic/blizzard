package blizzard

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/FuzzyStatic/blizzard/oauth"
)

const (
	oauthPath      = "/oauth"
	tokenPath      = "/token"
	userInfoPath   = "/userinfo"
	checkTokenPath = "/check_token"
	tokenQuery     = "token="
	grantType      = `grant_type=client_credentials`
)

// OAuth credentials and access token to access Blizzard API
type OAuth struct {
	ClientID           string
	ClientSecret       string
	AccessTokenRequest oauth.AccessTokenRequest
	ExpiresAt          time.Time
}

// AccessTokenReq retrieves new Access Token
func (c *Client) AccessTokenReq() error {
	var (
		req     *http.Request
		res     *http.Response
		reqBody *strings.Reader
		b       []byte
		err     error
	)

	reqBody = strings.NewReader(grantType)

	req, err = http.NewRequest("POST", c.oauthURL+oauthPath+tokenPath, reqBody)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.oauth.ClientID, c.oauth.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = c.client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &c.oauth.AccessTokenRequest)
	if err != nil {
		return err
	}

	c.oauth.ExpiresAt = time.Now().UTC().Add(time.Second * time.Duration(c.oauth.AccessTokenRequest.ExpiresIn))

	return nil
}

// updateAccessTokenIfExp updates Access Token if expired
func (c *Client) updateAccessTokenIfExp() error {
	var err error

	if c.oauth.ExpiresAt.Sub(time.Now().UTC()) < 60 {
		err = c.AccessTokenReq()
		if err != nil {
			return err
		}
	}

	return nil
}

// UserInfoHeader teturns basic information about the user associated with the current bearer token
func (c *Client) UserInfoHeader() ([]byte, error) {
	var (
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	err = c.updateAccessTokenIfExp()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", c.oauthURL+oauthPath+userInfoPath, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.oauth.AccessTokenRequest.AccessToken)
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

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
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
		return nil, nil, err
	}

	req, err = http.NewRequest("GET", c.oauthURL+oauthPath+checkTokenPath+"?"+tokenQuery+c.oauth.AccessTokenRequest.AccessToken, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	res, err = c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}
