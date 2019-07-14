package blizzard

import (
	"fmt"
	"testing"
	"time"
)

func TestAccessTokenReq(t *testing.T) {
	err := c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}
}

func TestUpdateAccessTokenIfExp(t *testing.T) {
	err := c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = c.updateAccessTokenIfExp()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}

	c.oauth.AccessTokenRequest.ExpiresIn = 0
	c.oauth.ExpiresAt = time.Now().UTC()

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}

	err = c.updateAccessTokenIfExp()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}
}

func TestUserInfoHeader(t *testing.T) {
	body, err := c.UserInfoHeader()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%s\n", body)
	}
}

func TestTokenValidation(t *testing.T) {
	dat, _, err := c.TokenValidation()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
