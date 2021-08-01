package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestAccessTokenRequest(t *testing.T) {
	err := usClient.AccessTokenRequest(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", usClient.oauth)
	}
}

// func TestUserInfoHeader(t *testing.T) {
// 	err := usClient.Token()
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// 	dat, _, err := usClient.UserInfoHeader(usClient.oauth.Token)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// 	if printOutput != "" {
// 		fmt.Printf("%+v\n", dat)
// 	}
// }

func TestTokenValidation(t *testing.T) {
	dat, _, err := usClient.TokenValidation(context.Background(), usClient.oauth.Token)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
