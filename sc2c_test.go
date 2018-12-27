package blizzard

import (
	"fmt"
	"testing"

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

func TestSC2GetStaticProfile(t *testing.T) {
	dat, err := c.SC2GetStaticProfile(CN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2GetMetadataProfile(t *testing.T) {
	dat, err := c.SC2GetMetadataProfile(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2GetProfile(t *testing.T) {
	dat, err := c.SC2GetProfile(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2GetProfileLadderSummary(t *testing.T) {
	dat, err := c.SC2GetProfileLadderSummary(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2GetProfileLadder(t *testing.T) {
	dat, err := c.SC2GetProfileLadder(US, 1, 2376042, 194163)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
