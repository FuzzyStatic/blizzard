/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-17 19:43:46
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:57:15
 */

package starcraft2

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard"
	"github.com/spf13/viper"
)

var s *Starcraft2

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	accessToken := viper.GetString("API.access_token")
	apiKey := viper.GetString("API.api_key")

	s = New(
		blizzard.Auth{
			AccessToken: accessToken,
			APIKey:      apiKey,
		},
		blizzard.US,
	)
}

func TestGetProfile(t *testing.T) {
	var (
		profileJSON *[]byte
		profile     *Profile
		err         error
	)

	profileJSON, err = s.GetProfileJSON(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	profile, err = s.GetProfile(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*profileJSON))
	fmt.Println(*profile)
}

func TestGetProfileLadders(t *testing.T) {
	var (
		profileLaddersJSON *[]byte
		profileLadders     *ProfileLadders
		err                error
	)

	profileLaddersJSON, err = s.GetProfileLaddersJSON(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	profileLadders, err = s.GetProfileLadders(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*profileLaddersJSON))
	fmt.Println(*profileLadders)
}
