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

func TestD3GetSeasonIndex(t *testing.T) {
	dat, err := c.D3GetSeasonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeason(t *testing.T) {
	dat, err := c.D3GetSeason(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboard(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboard(15, "74987248527082")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardAchievementPoints(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardAchievementPoints(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardBarbarian(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreCrusader(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardCrusader(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardDemonHunter(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreMonk(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardMonk(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardNecromancer(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreWizard(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardWizard(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetSeasonLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardWitchDoctor(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardHardcoreTeam2(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardTeam2(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetSeasonLeaderboardHardcoreTeam3(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardTeam3(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetSeasonLeaderboardHardcoreTeam4(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardHardcoreTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetSeasonLeaderboardTeam4(t *testing.T) {
	dat, err := c.D3GetSeasonLeaderboardTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraIndex(t *testing.T) {
	dat, err := c.D3GetEraIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEra(t *testing.T) {
	dat, err := c.D3GetEra(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboard(t *testing.T) {
	dat, err := c.D3GetEraLeaderboard(15, "74987248527082")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardAchievementPoints(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardAchievementPoints(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardBarbarian(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreCrusader(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardCrusader(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardDemonHunter(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreMonk(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardMonk(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardNecromancer(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreWizard(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardWizard(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetEraLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardWitchDoctor(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardHardcoreTeam2(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardTeam2(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetEraLeaderboardHardcoreTeam3(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardTeam3(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3GetEraLeaderboardHardcoreTeam4(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardHardcoreTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEraLeaderboardTeam4(t *testing.T) {
	dat, err := c.D3GetEraLeaderboardTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
