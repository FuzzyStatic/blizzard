package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/sc2gd"

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

func TestSC2GetLeagueData(t *testing.T) {
	dat, err := c.SC2GetLeagueData(37, sc2gd.LotV1v1, sc2gd.Arranged, sc2gd.Master)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
