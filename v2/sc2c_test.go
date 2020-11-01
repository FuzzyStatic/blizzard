package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestSC2StaticProfile(t *testing.T) {
	dat, _, err := c.SC2StaticProfile(context.Background(), US)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2MetadataProfile(t *testing.T) {
	dat, _, err := c.SC2MetadataProfile(context.Background(), US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2Profile(t *testing.T) {
	dat, _, err := c.SC2Profile(context.Background(), US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2ProfileLadderSummary(t *testing.T) {
	dat, _, err := c.SC2ProfileLadderSummary(context.Background(), US, 1, 305084)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2ProfileLadder(t *testing.T) {
	dat, _, err := c.SC2ProfileLadder(context.Background(), US, 1, 2376042, 194163)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LadderGrandmaster(t *testing.T) {
	dat, _, err := c.SC2LadderGrandmaster(context.Background(), EU)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LadderSeason(t *testing.T) {
	dat, _, err := c.SC2LadderSeason(context.Background(), EU)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2Player(t *testing.T) {
	dat, _, err := c.SC2Player(context.Background(), 1312411)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyProfile(t *testing.T) {
	dat, _, err := c.SC2LegacyProfile(context.Background(), US, 1, 1655091)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyProfileLadders(t *testing.T) {
	dat, _, err := c.SC2LegacyProfileLadders(context.Background(), US, 1, 1655091)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyProfileMatches(t *testing.T) {
	dat, _, err := c.SC2LegacyProfileMatches(context.Background(), US, 1, 1655091)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyLadder(t *testing.T) {
	dat, _, err := c.SC2LegacyLadder(context.Background(), US, 194163)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyAchievements(t *testing.T) {
	dat, _, err := c.SC2LegacyAchievements(context.Background(), US)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LegacyRewards(t *testing.T) {
	dat, _, err := c.SC2LegacyRewards(context.Background(), US)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
