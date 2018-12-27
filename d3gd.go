package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/d3gd"
)

const (
	dataD3Path                  = dataPath + "/d3"
	seasonPath                  = dataD3Path + "/season"
	leaderboardPath             = "/leaderboard"
	achievementPointsPath       = "/achievement-points"
	riftHardcoreBarbarianPath   = "/rift-hardcore-barbarian"
	riftBarbarianPath           = "/rift-barbarian"
	riftHardcoreCrusaderPath    = "/rift-hardcore-crusader"
	riftCrusaderPath            = "/rift-crusader"
	riftHardcoreDHPath          = "/rift-hardcore-dh"
	riftDHPath                  = "/rift-dh"
	riftHardcoreMonkPath        = "/rift-hardcore-monk"
	riftMonkPath                = "/rift-monk"
	riftHardcoreNecromancerPath = "/rift-hardcore-necromancer"
	riftNecromancerPath         = "/rift-necromancer"
	riftHardcoreWizardPath      = "/rift-hardcore-wizard"
	riftWizardPath              = "/rift-wizard"
	riftHardcoreWDPath          = "/rift-hardcore-wd"
	riftWDPath                  = "/rift-wd"
	riftHardcoreTeam2Path       = "/rift-hardcore-team-2"
	riftTeam2Path               = "/rift-team-2"
	riftHardcoreTeam3Path       = "/rift-hardcore-team-3"
	riftTeam3Path               = "/rift-team-3"
	riftHardcoreTeam4Path       = "/rift-hardcore-team-4"
	riftTeam4Path               = "/rift-team-4"
	eraPath                     = dataD3Path + "/era"
)

// D3SeasonIndex returns an index of seasons
func (c *Config) D3SeasonIndex() (*d3gd.SeasonIndex, error) {
	var (
		dat d3gd.SeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Season returns season data
func (c *Config) D3Season(seasonID int) (*d3gd.Season, error) {
	var (
		dat d3gd.Season
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboard returns leaderboard data for season and leaderboard
func (c *Config) D3SeasonLeaderboard(seasonID int, leaderboard string) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + "/" + leaderboard + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardAchievementPoints returns achievement points leaderboard data for season
func (c *Config) D3SeasonLeaderboardAchievementPoints(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + achievementPointsPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreBarbarian(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreBarbarianPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Config) D3SeasonLeaderboardBarbarian(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftBarbarianPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreCrusader(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreCrusaderPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardCrusader returns crusader leaderboard data for season
func (c *Config) D3SeasonLeaderboardCrusader(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftCrusaderPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreDemonHunter(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreDHPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Config) D3SeasonLeaderboardDemonHunter(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftDHPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreMonk(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreMonkPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardMonk returns monk leaderboard data for season
func (c *Config) D3SeasonLeaderboardMonk(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftMonkPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreNecromancer(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreNecromancerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Config) D3SeasonLeaderboardNecromancer(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftNecromancerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreWizard(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreWizardPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardWizard returns wizard leaderboard data for season
func (c *Config) D3SeasonLeaderboardWizard(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftWizardPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreWitchDoctor(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreWDPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Config) D3SeasonLeaderboardWitchDoctor(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftWDPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreTeam2(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreTeam2Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardTeam2(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftTeam2Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreTeam3(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreTeam3Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardTeam3(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftTeam3Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardHardcoreTeam4(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftHardcoreTeam4Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3SeasonLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Config) D3SeasonLeaderboardTeam4(seasonID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + seasonPath + "/" + strconv.Itoa(seasonID) + leaderboardPath + riftTeam4Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraIndex returns an index of seasons
func (c *Config) D3EraIndex() (*d3gd.EraIndex, error) {
	var (
		dat d3gd.EraIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Era returns season data
func (c *Config) D3Era(eraID int) (*d3gd.Era, error) {
	var (
		dat d3gd.Era
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreBarbarian(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreBarbarianPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Config) D3EraLeaderboardBarbarian(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftBarbarianPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreCrusader(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreCrusaderPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardCrusader returns crusader leaderboard data for season
func (c *Config) D3EraLeaderboardCrusader(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftCrusaderPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreDemonHunter(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreDHPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Config) D3EraLeaderboardDemonHunter(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftDHPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreMonk(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreMonkPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardMonk returns monk leaderboard data for season
func (c *Config) D3EraLeaderboardMonk(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftMonkPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreNecromancer(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreNecromancerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Config) D3EraLeaderboardNecromancer(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftNecromancerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreWizard(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreWizardPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardWizard returns wizard leaderboard data for season
func (c *Config) D3EraLeaderboardWizard(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftWizardPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreWitchDoctor(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreWDPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Config) D3EraLeaderboardWitchDoctor(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftWDPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreTeam2(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreTeam2Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Config) D3EraLeaderboardTeam2(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftTeam2Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreTeam3(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreTeam3Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Config) D3EraLeaderboardTeam3(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftTeam3Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Config) D3EraLeaderboardHardcoreTeam4(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftHardcoreTeam4Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3EraLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Config) D3EraLeaderboardTeam4(eraID int) (*d3gd.Leaderboard, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + eraPath + "/" + strconv.Itoa(eraID) + leaderboardPath + riftTeam4Path + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
