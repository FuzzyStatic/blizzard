/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-16 19:26:42
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:28:40
 */

package blizzard

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

var c *Config

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	accessToken := viper.GetString("API.access_token")
	apiKey := viper.GetString("API.api_key")

	c = New(
		Auth{
			AccessToken: accessToken,
			APIKey:      apiKey,
		},
		US,
	)
}

func TestGetAccountUser(t *testing.T) {
	var (
		accountUserJSON []byte
		accountUser     *AccountUser
		err             error
	)

	accountUserJSON, err = c.GetAccountUserJSON()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	accountUser, err = c.GetAccountUser()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(string(accountUserJSON))
	fmt.Printf("%+v\n", accountUser)
}
