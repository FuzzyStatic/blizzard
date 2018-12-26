package blizzard

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	clientID := viper.GetString("authentication.client_id")
	clientSecret := viper.GetString("authentication.client_secret")

	c = New(clientID, clientSecret, US)
}

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

	c.UpdateAccessTokenIfExp()
	fmt.Printf("%+v\n", c.oauth)

	c.oauth.AccessTokenRequest.ExpiresIn = 0
	c.oauth.ExpiresAt = time.Now().UTC()

	fmt.Printf("%+v\n", c.oauth)
	c.UpdateAccessTokenIfExp()

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

func TestUserInfoParam(t *testing.T) {
	body, err := c.UserInfoParam()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%s\n", body)
}

func TestTokenValidation(t *testing.T) {
	dat, err := c.TokenValidation()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
