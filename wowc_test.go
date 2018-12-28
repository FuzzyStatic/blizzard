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

func TestWoWUserCharacters(t *testing.T) {
	dat, err := c.WoWUserCharacters(c.oauth.AccessTokenRequest.AccessToken)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWAchievement(t *testing.T) {
	dat, err := c.WoWAchievement(2144)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWAuctionFiles(t *testing.T) {
	dat, err := c.WoWAuctionFiles("medivh")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWAuctionData(t *testing.T) {
	dat, err := c.WoWAuctionData("medivh")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for _, auctions := range dat {
		fmt.Printf("%+v\n", auctions)
	}
}

func TestWoWBossMasterList(t *testing.T) {
	dat, err := c.WoWBossMasterList()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWBoss(t *testing.T) {
	dat, err := c.WoWBoss(24723)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWChallengeModeRealmLeaderboard(t *testing.T) {
	dat, err := c.WoWChallengeModeRealmLeaderboard("medivh")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWChallengeModeRegionLeaderboard(t *testing.T) {
	dat, err := c.WoWChallengeModeRegionLeaderboard()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWCharacterProfile(t *testing.T) {
	dat, err := c.WoWCharacterProfile("emerald-dream", "Limejelly",
		[]string{
			FieldCharacterAchievements,
			FieldCharacterAppearance,
			FieldCharacterAudit,
			FieldCharacterFeed,
			FieldCharacterGuild,
			FieldCharacterItems,
			FieldCharacterMounts,
			FieldCharacterPVP,
			FieldCharacterPetSlots,
			FieldCharacterPets,
			FieldCharacterProfessions,
			FieldCharacterProgression,
			FieldCharacterQuests,
			FieldCharacterReputation,
			FieldCharacterStatistics,
			FieldCharacterStats,
			FieldCharacterTalents,
			FieldCharacterTitle,
		},
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWGuildProfile(t *testing.T) {
	dat, err := c.WoWGuildProfile("emerald-dream", "nightstalkers",
		[]string{
			FieldGuildAchievements,
			FieldGuildChallenge,
			FieldGuildMembers,
			FieldGuildNews,
		},
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
