package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/sc2c"
)

// SC2StaticProfile returns all static SC2 profile data (ctx context.Context, achievements, categories, criteria, and rewards)
func (c *Client) SC2StaticProfile(ctx context.Context, region Region) (*sc2c.StaticProfile, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/static/profile/%d", region),
		"",
		&sc2c.StaticProfile{},
	)
	return dat.(*sc2c.StaticProfile), b, err
}

// SC2MetadataProfile returns all SC2 profile metadata
func (c *Client) SC2MetadataProfile(ctx context.Context, region Region, realmID, profileID int) (*sc2c.MetadataProfile, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/metadata/profile/%d/%d/%d", region, realmID, profileID),
		"",
		&sc2c.MetadataProfile{},
	)
	return dat.(*sc2c.MetadataProfile), b, err
}

// SC2Profile returns all SC2 profile data
func (c *Client) SC2Profile(ctx context.Context, region Region, realmID, profileID int) (*sc2c.Profile, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/profile/%d/%d/%d", region, realmID, profileID),
		"",
		&sc2c.Profile{},
	)
	return dat.(*sc2c.Profile), b, err
}

// SC2ProfileLadderSummary returns SC2 profile ladder summary
func (c *Client) SC2ProfileLadderSummary(ctx context.Context, region Region, realmID, profileID int) (*sc2c.LadderSummary, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/profile/%d/%d/%d/ladder/summary", region, realmID, profileID),
		"",
		&sc2c.LadderSummary{},
	)
	return dat.(*sc2c.LadderSummary), b, err
}

// SC2ProfileLadder returns SC2 profile ladder data
func (c *Client) SC2ProfileLadder(ctx context.Context, region Region, realmID, profileID, ladderID int) (*sc2c.Ladder, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/profile/%d/%d/%d/ladder/%d", region, realmID, profileID, ladderID),
		"",
		&sc2c.Ladder{},
	)
	return dat.(*sc2c.Ladder), b, err
}

// SC2LadderGrandmaster returns SC2 ladder grandmaster for current season
func (c *Client) SC2LadderGrandmaster(ctx context.Context, region Region) (*sc2c.LadderGrandmaster, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/ladder/grandmaster/%d", region),
		"",
		&sc2c.LadderGrandmaster{},
	)
	return dat.(*sc2c.LadderGrandmaster), b, err
}

// SC2LadderSeason returns SC2 ladder current season
func (c *Client) SC2LadderSeason(ctx context.Context, region Region) (*sc2c.LadderSeason, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/ladder/season/%d", region),
		"",
		&sc2c.LadderSeason{},
	)
	return dat.(*sc2c.LadderSeason), b, err
}

// SC2Player returns data about player using account ID
func (c *Client) SC2Player(ctx context.Context, accountID int) (*sc2c.Player, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/player/%d", accountID),
		"",
		&sc2c.Player{},
	)
	return dat.(*sc2c.Player), b, err
}

// SC2LegacyProfile returns all SC2 legacy profile data
func (c *Client) SC2LegacyProfile(ctx context.Context, region Region, realmID, profileID int) (*sc2c.LegacyProfile, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/profile/%d/%d/%d", region, realmID, profileID),
		"",
		&sc2c.LegacyProfile{},
	)
	return dat.(*sc2c.LegacyProfile), b, err
}

// SC2LegacyProfileLadders returns all SC2 legacy profile ladder data
func (c *Client) SC2LegacyProfileLadders(ctx context.Context, region Region, realmID, profileID int) (*sc2c.LegacyProfileLadders, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/profile/%d/%d/%d/ladders", region, realmID, profileID),
		"",
		&sc2c.LegacyProfileLadders{},
	)
	return dat.(*sc2c.LegacyProfileLadders), b, err
}

// SC2LegacyProfileMatches returns all SC2 legacy profile matches data
func (c *Client) SC2LegacyProfileMatches(ctx context.Context, region Region, realmID, profileID int) (*sc2c.LegacyProfileMatches, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/profile/%d/%d/%d/matches", region, realmID, profileID),
		"",
		&sc2c.LegacyProfileMatches{},
	)
	return dat.(*sc2c.LegacyProfileMatches), b, err
}

// SC2LegacyLadder returns SC2 legacy ladder data
func (c *Client) SC2LegacyLadder(ctx context.Context, region Region, ladderID int) (*sc2c.LegacyLadder, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/ladder/%d/%d", region, ladderID),
		"",
		&sc2c.LegacyLadder{},
	)
	return dat.(*sc2c.LegacyLadder), b, err
}

// SC2LegacyAchievements returns SC2 legacy achievements for region
func (c *Client) SC2LegacyAchievements(ctx context.Context, region Region) (*sc2c.LegacyAchievements, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/data/achievements/%d", region),
		"",
		&sc2c.LegacyAchievements{},
	)
	return dat.(*sc2c.LegacyAchievements), b, err
}

// SC2LegacyRewards returns SC2 legacy rewards for region
func (c *Client) SC2LegacyRewards(ctx context.Context, region Region) (*sc2c.LegacyRewards, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/sc2/legacy/data/rewards/%d", region),
		"",
		&sc2c.LegacyRewards{},
	)
	return dat.(*sc2c.LegacyRewards), b, err
}
