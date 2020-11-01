package blizzard

import (
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v1/d3gd"
)

// D3SeasonIndex returns an index of seasons
func (c *Client) D3SeasonIndex() (*d3gd.SeasonIndex, []byte, error) {
	var (
		dat d3gd.SeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/?locale=%s", c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Season returns season data
func (c *Client) D3Season(seasonID int) (*d3gd.Season, []byte, error) {
	var (
		dat d3gd.Season
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboard returns leaderboard data for season and leaderboard
func (c *Client) D3SeasonLeaderboard(seasonID int, leaderboard string) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/%s?locale=%s", seasonID, leaderboard, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardAchievementPoints returns achievement points leaderboard data for season
func (c *Client) D3SeasonLeaderboardAchievementPoints(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/achievement-points?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreBarbarian(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-barbarian?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardBarbarian(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-barbarian?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreCrusader(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-crusader?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardCrusader returns crusader leaderboard data for season
func (c *Client) D3SeasonLeaderboardCrusader(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-crusader?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreDemonHunter(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-dh?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardDemonHunter(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-dh?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreMonk(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-monk?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardMonk returns monk leaderboard data for season
func (c *Client) D3SeasonLeaderboardMonk(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-monk?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreNecromancer(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-necromancer?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Client) D3SeasonLeaderboardNecromancer(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-necromancer?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreWizard(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wizard?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardWizard returns wizard leaderboard data for season
func (c *Client) D3SeasonLeaderboardWizard(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wizard?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreWitchDoctor(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wd?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Client) D3SeasonLeaderboardWitchDoctor(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wd?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam2(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-2?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam2(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-2?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam3(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-3?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam3(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-3?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam4(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-4?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3SeasonLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam4(seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-4?locale=%s", seasonID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraIndex returns an index of seasons
func (c *Client) D3EraIndex() (*d3gd.EraIndex, []byte, error) {
	var (
		dat d3gd.EraIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/?locale=%s", c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Era returns season data
func (c *Client) D3Era(eraID int) (*d3gd.Era, []byte, error) {
	var (
		dat d3gd.Era
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreBarbarian(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-barbarian?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardBarbarian(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-barbarian?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreCrusader(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-crusader?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardCrusader returns crusader leaderboard data for season
func (c *Client) D3EraLeaderboardCrusader(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-crusader?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreDemonHunter(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-dh?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardDemonHunter(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-dh?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreMonk(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-monk?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardMonk returns monk leaderboard data for season
func (c *Client) D3EraLeaderboardMonk(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-monk?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreNecromancer(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-necromancer?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Client) D3EraLeaderboardNecromancer(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-necromancer?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreWizard(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wizard?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardWizard returns wizard leaderboard data for season
func (c *Client) D3EraLeaderboardWizard(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wizard?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreWitchDoctor(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wd?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Client) D3EraLeaderboardWitchDoctor(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wd?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam2(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-2?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam2(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-2?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam3(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-3?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam3(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-3?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam4(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-4?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3EraLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam4(eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-4?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
