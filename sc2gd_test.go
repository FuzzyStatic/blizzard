package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/v3/sc2gd"
)

func TestSC2LeagueData(t *testing.T) {
	dat, _, err := usClient.SC2LeagueData(context.Background(), 37, sc2gd.LotV1v1, sc2gd.Arranged, sc2gd.Master)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestSC2LadderData(t *testing.T) {
	dat, _, err := usClient.SC2LadderData(context.Background(), 292787)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
