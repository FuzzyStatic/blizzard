package blizzard

import (
	"fmt"
	"testing"
)

func TestAccessTokenRequest(t *testing.T) {
	err := c.AccessTokenRequest()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}
}

// func TestUserInfoHeader(t *testing.T) {
// 	err := c.Token()
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// 	dat, _, err := c.UserInfoHeader(c.oauth.Token)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// 	if printOutput != "" {
// 		fmt.Printf("%+v\n", dat)
// 	}
// }

func TestTokenValidation(t *testing.T) {
	dat, _, err := c.TokenValidation(c.oauth.Token)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
