package blizzard

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	clientID := viper.GetString("authentication.client_id")
	clientSecret := viper.GetString("authentication.client_secret")

	c = NewClient(clientID, clientSecret, US, enUS)

	err = c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		return
	}
}
