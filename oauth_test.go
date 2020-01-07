package blizzard

import (
	"fmt"
	"testing"
)

func TestAccessTokenReq(t *testing.T) {
	err := c.Token()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", c.oauth)
	}
}

func TestUpdateAccessTokenIfExp(t *testing.T) {
	err := c.Token()
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

	// c.oauth.Token.ExpiresIn = 0
	// c.oauth.ExpiresAt = time.Now().UTC()

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
	dat, _, err := c.TokenValidation()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
