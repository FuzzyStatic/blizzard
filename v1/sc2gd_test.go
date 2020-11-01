package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/v1/sc2gd"
)

func TestSC2LeagueData(t *testing.T) {
	dat, _, err := c.SC2LeagueData(37, sc2gd.LotV1v1, sc2gd.Arranged, sc2gd.Master)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
