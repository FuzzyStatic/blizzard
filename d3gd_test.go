package blizzard

import (
	"fmt"
	"testing"
)

func TestD3SeasonIndex(t *testing.T) {
	dat, _, err := c.D3SeasonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3Season(t *testing.T) {
	dat, _, err := c.D3Season(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboard(15, "74987248527082")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardAchievementPoints(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardAchievementPoints(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardBarbarian(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardBarbarian(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreCrusader(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardCrusader(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardCrusader(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardDemonHunter(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardDemonHunter(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreMonk(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardMonk(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardMonk(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardNecromancer(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardNecromancer(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreWizard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardWizard(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardWizard(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3SeasonLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardWitchDoctor(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardWitchDoctor(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardHardcoreTeam2(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardTeam2(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam2(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3SeasonLeaderboardHardcoreTeam3(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardTeam3(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam3(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3SeasonLeaderboardHardcoreTeam4(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardHardcoreTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3SeasonLeaderboardTeam4(t *testing.T) {
	dat, _, err := c.D3SeasonLeaderboardTeam4(15)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraIndex(t *testing.T) {
	dat, _, err := c.D3EraIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3Era(t *testing.T) {
	dat, _, err := c.D3Era(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreBarbarian(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreBarbarian(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardBarbarian(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardBarbarian(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreCrusader(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreCrusader(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardCrusader(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardCrusader(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreDemonHunter(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreDemonHunter(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardDemonHunter(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardDemonHunter(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreMonk(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreMonk(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardMonk(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardMonk(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreNecromancer(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreNecromancer(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardNecromancer(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardNecromancer(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreWizard(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreWizard(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardWizard(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardWizard(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3EraLeaderboardHardcoreWitchDoctor(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreWitchDoctor(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardWitchDoctor(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardWitchDoctor(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardHardcoreTeam2(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam2(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardTeam2(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam2(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3EraLeaderboardHardcoreTeam3(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam3(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardTeam3(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam3(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
func TestD3EraLeaderboardHardcoreTeam4(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardHardcoreTeam4(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3EraLeaderboardTeam4(t *testing.T) {
	dat, _, err := c.D3EraLeaderboardTeam4(10)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
