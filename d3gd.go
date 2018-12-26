package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/d3gd"
)

const (
	dataPath                    = "/data"
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

// D3GetSeasonIndex returns an index of seasons
func (c *Config) D3GetSeasonIndex() (*d3gd.SeasonIndex, error) {
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

// D3GetSeason returns season data
func (c *Config) D3GetSeason(seasonID int) (*d3gd.Season, error) {
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

// D3GetSeasonLeaderboard returns leaderboard data for season and leaderboard
func (c *Config) D3GetSeasonLeaderboard(seasonID int, leaderboard string) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardAchievementPoints returns achievement points leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardAchievementPoints(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreBarbarian(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardBarbarian(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreCrusader(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardCrusader returns crusader leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardCrusader(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreDemonHunter(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardDemonHunter(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreMonk(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardMonk returns monk leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardMonk(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreNecromancer(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardNecromancer(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreWizard(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardWizard returns wizard leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardWizard(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreWitchDoctor(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardWitchDoctor(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreTeam2(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardTeam2(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreTeam3(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardTeam3(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardHardcoreTeam4(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetSeasonLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Config) D3GetSeasonLeaderboardTeam4(seasonID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraIndex returns an index of seasons
func (c *Config) D3GetEraIndex() (*d3gd.EraIndex, error) {
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

// D3GetEra returns season data
func (c *Config) D3GetEra(eraID int) (*d3gd.Era, error) {
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

// D3GetEraLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreBarbarian(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Config) D3GetEraLeaderboardBarbarian(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreCrusader(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardCrusader returns crusader leaderboard data for season
func (c *Config) D3GetEraLeaderboardCrusader(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreDemonHunter(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Config) D3GetEraLeaderboardDemonHunter(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreMonk(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardMonk returns monk leaderboard data for season
func (c *Config) D3GetEraLeaderboardMonk(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreNecromancer(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Config) D3GetEraLeaderboardNecromancer(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreWizard(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardWizard returns wizard leaderboard data for season
func (c *Config) D3GetEraLeaderboardWizard(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreWitchDoctor(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Config) D3GetEraLeaderboardWitchDoctor(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreTeam2(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardTeam2(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreTeam3(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardTeam3(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardHardcoreTeam4(eraID int) (*d3gd.Leaderboard, error) {
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

// D3GetEraLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Config) D3GetEraLeaderboardTeam4(eraID int) (*d3gd.Leaderboard, error) {
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
