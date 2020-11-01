package blizzard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/d3gd"
)

// D3SeasonIndex returns an index of seasons
func (c *Client) D3SeasonIndex(ctx context.Context) (*d3gd.SeasonIndex, []byte, error) {
	var (
		dat d3gd.SeasonIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/?locale=%s", c.locale), "")
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
func (c *Client) D3Season(ctx context.Context, seasonID int) (*d3gd.Season, []byte, error) {
	var (
		dat d3gd.Season
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboard(ctx context.Context, seasonID int, leaderboard string) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/%s?locale=%s", seasonID, leaderboard, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardAchievementPoints(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/achievement-points?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreBarbarian(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-barbarian?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardBarbarian(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-barbarian?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreCrusader(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-crusader?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardCrusader(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-crusader?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreDemonHunter(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-dh?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardDemonHunter(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-dh?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreMonk(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-monk?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardMonk(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-monk?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreNecromancer(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-necromancer?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardNecromancer(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-necromancer?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreWizard(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wizard?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardWizard(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wizard?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreWitchDoctor(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wd?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardWitchDoctor(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wd?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreTeam2(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-2?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardTeam2(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-2?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreTeam3(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-3?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardTeam3(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-3?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardHardcoreTeam4(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-4?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3SeasonLeaderboardTeam4(ctx context.Context, seasonID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-4?locale=%s", seasonID, c.locale), "")
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
func (c *Client) D3EraIndex(ctx context.Context) (*d3gd.EraIndex, []byte, error) {
	var (
		dat d3gd.EraIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/?locale=%s", c.locale), "")
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
func (c *Client) D3Era(ctx context.Context, eraID int) (*d3gd.Era, []byte, error) {
	var (
		dat d3gd.Era
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreBarbarian(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-barbarian?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardBarbarian(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-barbarian?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreCrusader(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-crusader?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardCrusader(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-crusader?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreDemonHunter(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-dh?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardDemonHunter(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-dh?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreMonk(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-monk?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardMonk(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-monk?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreNecromancer(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-necromancer?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardNecromancer(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-necromancer?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreWizard(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wizard?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardWizard(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wizard?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreWitchDoctor(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wd?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardWitchDoctor(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wd?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreTeam2(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-2?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardTeam2(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-2?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreTeam3(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-3?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardTeam3(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-3?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardHardcoreTeam4(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-4?locale=%s", eraID, c.locale), "")
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
func (c *Client) D3EraLeaderboardTeam4(ctx context.Context, eraID int) (*d3gd.Leaderboard, []byte, error) {
	var (
		dat d3gd.Leaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-4?locale=%s", eraID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
