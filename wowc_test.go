package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wowc"
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
			wowc.FieldCharacterAchievements,
			wowc.FieldCharacterAppearance,
			wowc.FieldCharacterAudit,
			wowc.FieldCharacterFeed,
			wowc.FieldCharacterGuild,
			wowc.FieldCharacterItems,
			wowc.FieldCharacterMounts,
			wowc.FieldCharacterPVP,
			wowc.FieldCharacterPetSlots,
			wowc.FieldCharacterPets,
			wowc.FieldCharacterProfessions,
			wowc.FieldCharacterProgression,
			wowc.FieldCharacterQuests,
			wowc.FieldCharacterReputation,
			wowc.FieldCharacterStatistics,
			wowc.FieldCharacterStats,
			wowc.FieldCharacterTalents,
			wowc.FieldCharacterTitle,
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
			wowc.FieldGuildAchievements,
			wowc.FieldGuildChallenge,
			wowc.FieldGuildMembers,
			wowc.FieldGuildNews,
		},
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWItem(t *testing.T) {
	dat, err := c.WoWItem(18803)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWItemSet(t *testing.T) {
	dat, err := c.WoWItemSet(1060)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMountMasterList(t *testing.T) {
	dat, err := c.WoWMountMasterList()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPetMasterList(t *testing.T) {
	dat, err := c.WoWPetMasterList()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPetAbility(t *testing.T) {
	dat, err := c.WoWPetAbility(640)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPetSpecies(t *testing.T) {
	dat, err := c.WoWPetSpecies(258)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPetStats(t *testing.T) {
	dat, err := c.WoWPetStats(258, 25, 5, 4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPVPLeaderboard(t *testing.T) {
	dat, err := c.WoWPVPLeaderboard(wowc.Bracket3v3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWQuest(t *testing.T) {
	dat, err := c.WoWQuest(13146)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRealmStatus(t *testing.T) {
	dat, err := c.WoWRealmStatus()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRecipe(t *testing.T) {
	dat, err := c.WoWRecipe(33994)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWSpell(t *testing.T) {
	dat, err := c.WoWSpell(17086)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWZoneMasterList(t *testing.T) {
	dat, err := c.WoWZoneMasterList()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWZone(t *testing.T) {
	dat, err := c.WoWZone(4131)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRegionBattlegroups(t *testing.T) {
	dat, err := c.WoWRegionBattlegroups()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWCharacterRaces(t *testing.T) {
	dat, err := c.WoWCharacterRaces()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWCharacterClasses(t *testing.T) {
	dat, err := c.WoWCharacterClasses()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWCharacterAchievements(t *testing.T) {
	dat, err := c.WoWCharacterAchievements()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWGuildRewards(t *testing.T) {
	dat, err := c.WoWGuildRewards()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWGuildPerks(t *testing.T) {
	dat, err := c.WoWGuildPerks()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWGuildAchievements(t *testing.T) {
	dat, err := c.WoWGuildAchievements()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
