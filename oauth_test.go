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

	fmt.Printf("%+v\n", c.oauth)
}

func TestUpdateAccessTokenIfExp(t *testing.T) {
	err := c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	c.updateAccessTokenIfExp()
	fmt.Printf("%+v\n", c.oauth)

	c.oauth.AccessTokenRequest.ExpiresIn = 0
	c.oauth.ExpiresAt = time.Now().UTC()

	fmt.Printf("%+v\n", c.oauth)
	c.updateAccessTokenIfExp()

	fmt.Printf("%+v\n", c.oauth)
}

func TestUserInfoHeader(t *testing.T) {
	body, err := c.UserInfoHeader()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%s\n", body)
}

func TestTokenValidation(t *testing.T) {
	dat, _, err := c.TokenValidation()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}
