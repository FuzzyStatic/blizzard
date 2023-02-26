package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v3/d3gd"
)

// D3SeasonIndex returns an index of seasons
func (c *Client) D3SeasonIndex(ctx context.Context) (*d3gd.SeasonIndex, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		"/data/d3/season/",
		&d3gd.SeasonIndex{},
	)
	return dat.(*d3gd.SeasonIndex), header, err
}

// D3Season returns season data
func (c *Client) D3Season(ctx context.Context, seasonID int) (*d3gd.Season, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d", seasonID),
		&d3gd.Season{},
	)
	return dat.(*d3gd.Season), header, err
}

// D3SeasonLeaderboard returns leaderboard data for season and leaderboard
func (c *Client) D3SeasonLeaderboard(ctx context.Context, seasonID int, leaderboard string) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/%s", seasonID, leaderboard),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardAchievementPoints returns achievement points leaderboard data for season
func (c *Client) D3SeasonLeaderboardAchievementPoints(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/achievement-points", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreBarbarian(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-barbarian", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardBarbarian(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-barbarian", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreCrusader(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-crusader", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardCrusader returns crusader leaderboard data for season
func (c *Client) D3SeasonLeaderboardCrusader(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-crusader", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreDemonHunter(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-dh", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Client) D3SeasonLeaderboardDemonHunter(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-dh", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreMonk(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-monk", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardMonk returns monk leaderboard data for season
func (c *Client) D3SeasonLeaderboardMonk(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-monk", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreNecromancer(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-necromancer", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Client) D3SeasonLeaderboardNecromancer(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-necromancer", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreWizard(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wizard", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardWizard returns wizard leaderboard data for season
func (c *Client) D3SeasonLeaderboardWizard(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wizard", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreWitchDoctor(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-wd", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Client) D3SeasonLeaderboardWitchDoctor(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-wd", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam2(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-2", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam2(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-2", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam3(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-3", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam3(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-3", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardHardcoreTeam4(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-hardcore-team-4", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3SeasonLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Client) D3SeasonLeaderboardTeam4(ctx context.Context, seasonID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/season/%d/leaderboard/rift-team-4", seasonID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraIndex returns an index of seasons
func (c *Client) D3EraIndex(ctx context.Context) (*d3gd.EraIndex, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		"/data/d3/era/",
		&d3gd.EraIndex{},
	)
	return dat.(*d3gd.EraIndex), header, err
}

// D3Era returns season data
func (c *Client) D3Era(ctx context.Context, eraID int) (*d3gd.Era, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d", eraID),
		&d3gd.Era{},
	)
	return dat.(*d3gd.Era), header, err
}

// D3EraLeaderboardHardcoreBarbarian returns hardcore barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreBarbarian(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-barbarian", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardBarbarian returns barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardBarbarian(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-barbarian", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreCrusader returns hardcore crusader leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreCrusader(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-crusader", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardCrusader returns crusader leaderboard data for season
func (c *Client) D3EraLeaderboardCrusader(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-crusader", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreDemonHunter returns hardcore demon hunter leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreDemonHunter(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-dh", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardDemonHunter returns barbarian leaderboard data for season
func (c *Client) D3EraLeaderboardDemonHunter(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-dh", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreMonk returns hardcore monk leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreMonk(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-monk", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardMonk returns monk leaderboard data for season
func (c *Client) D3EraLeaderboardMonk(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-monk", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreNecromancer returns hardcore necromancer leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreNecromancer(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-necromancer", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardNecromancer returns necromancer leaderboard data for season
func (c *Client) D3EraLeaderboardNecromancer(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-necromancer", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreWizard returns hardcore wizard leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreWizard(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wizard", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardWizard returns wizard leaderboard data for season
func (c *Client) D3EraLeaderboardWizard(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wizard", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreWitchDoctor returns hardcore witch doctor leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreWitchDoctor(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-wd", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardWitchDoctor returns witch doctor leaderboard data for season
func (c *Client) D3EraLeaderboardWitchDoctor(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-wd", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreTeam2 returns hardcore 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam2(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-2", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardTeam2 returns 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam2(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-2", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreTeam3 returns hardcore 2 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam3(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-3", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardTeam3 returns 3 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam3(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-3", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardHardcoreTeam4 returns hardcore 4 team leaderboard data for season
func (c *Client) D3EraLeaderboardHardcoreTeam4(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-hardcore-team-4", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}

// D3EraLeaderboardTeam4 returns 4 team leaderboard data for season
func (c *Client) D3EraLeaderboardTeam4(ctx context.Context, eraID int) (*d3gd.Leaderboard, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/data/d3/era/%d/leaderboard/rift-team-4", eraID),
		&d3gd.Leaderboard{},
	)
	return dat.(*d3gd.Leaderboard), header, err
}
