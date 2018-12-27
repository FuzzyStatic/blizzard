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

func TestSC2StaticProfile(t *testing.T) {
	dat, err := c.SC2StaticProfile(CN)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2MetadataProfile(t *testing.T) {
	dat, err := c.SC2MetadataProfile(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2Profile(t *testing.T) {
	dat, err := c.SC2Profile(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2ProfileLadderSummary(t *testing.T) {
	dat, err := c.SC2ProfileLadderSummary(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2ProfileLadder(t *testing.T) {
	dat, err := c.SC2ProfileLadder(US, 1, 2376042, 194163)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LadderGrandmaster(t *testing.T) {
	dat, err := c.SC2LadderGrandmaster(EU)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LadderSeason(t *testing.T) {
	dat, err := c.SC2LadderSeason(EU)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2Player(t *testing.T) {
	dat, err := c.SC2Player(1312411)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyProfile(t *testing.T) {
	dat, err := c.SC2LegacyProfile(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyProfileLadders(t *testing.T) {
	dat, err := c.SC2LegacyProfileLadders(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyProfileMatches(t *testing.T) {
	dat, err := c.SC2LegacyProfileMatches(US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyLadder(t *testing.T) {
	dat, err := c.SC2LegacyLadder(US, 194163)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyAchievements(t *testing.T) {
	dat, err := c.SC2LegacyAchievements(US)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestSC2LegacyRewards(t *testing.T) {
	dat, err := c.SC2LegacyRewards(US)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
