/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:06
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:38:06
 */

package worldofwarcraft

import (
	"fmt"
	"go-blizzard/blizzard"
	"testing"

	"github.com/spf13/viper"
)

var w *WorldOfWarcraft

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

	w = New(
		blizzard.Auth{
			AccessToken: accessToken,
			APIKey:      apiKey,
		},
		blizzard.US,
	)
}

func TestGetConnectedRealmIndexJSON(t *testing.T) {
	var (
		connectedRealmIndexJSON *[]byte
		connectedRealmIndex     *ConnectedRealmIndex
		err                     error
	)

	connectedRealmIndexJSON, err = w.GetConnectedRealmIndexJSON()
	if err != nil {
		t.Fail()
	}

	connectedRealmIndex, err = w.GetConnectedRealmIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*connectedRealmIndexJSON))
	fmt.Println(*connectedRealmIndex)
}
