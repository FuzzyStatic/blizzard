package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestD3SeasonIndex(t *testing.T) {
	dat, _, err := c.D3SeasonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Season(t *testing.T) {
	dat, _, err := c.D3Season(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboard(context.Background(), 15, "74987248527082")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardAchievementPoints(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardAchievementPoints(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreBarbarian(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardBarbarian(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardBarbarian(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreCrusader(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreCrusader(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardCrusader(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardCrusader(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreDemonHunter(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardDemonHunter(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardDemonHunter(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreMonk(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreMonk(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardMonk(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardMonk(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreNecromancer(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardNecromancer(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardNecromancer(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreWizard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreWizard(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardWizard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardWizard(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3SeasonLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreWitchDoctor(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardWitchDoctor(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardWitchDoctor(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardHardcoreTeam2(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam2(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardTeam2(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam2(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3SeasonLeaderboardHardcoreTeam3(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam3(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardTeam3(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam3(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3SeasonLeaderboardHardcoreTeam4(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam4(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3SeasonLeaderboardTeam4(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam4(context.Background(), 15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraIndex(t *testing.T) {
	dat, _, err := c.D3EraIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Era(t *testing.T) {
	dat, _, err := c.D3Era(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreBarbarian(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardBarbarian(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardBarbarian(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreCrusader(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreCrusader(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardCrusader(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardCrusader(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreDemonHunter(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardDemonHunter(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardDemonHunter(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreMonk(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreMonk(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardMonk(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardMonk(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreNecromancer(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardNecromancer(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardNecromancer(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreWizard(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreWizard(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardWizard(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardWizard(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3EraLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreWitchDoctor(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardWitchDoctor(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardWitchDoctor(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardHardcoreTeam2(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam2(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardTeam2(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam2(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3EraLeaderboardHardcoreTeam3(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam3(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardTeam3(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam3(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
func TestD3EraLeaderboardHardcoreTeam4(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam4(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3EraLeaderboardTeam4(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam4(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
