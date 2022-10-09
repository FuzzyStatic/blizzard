package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wow"
)

func TestClassicWoWConnectedRealmsIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWConnectedRealmsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWConnectedRealm(t *testing.T) {
	dat, _, err := usClient.ClassicWoWConnectedRealm(context.Background(), 4388)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureFamiliesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamily(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureFamily(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureType(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureType(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreature(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreature(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureDisplayMedia(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWCreatureFamilyMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWGuildCrestComponentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWGuildCrestBorderMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWGuildCrestEmblemMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemClassesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWItemClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemClass(t *testing.T) {
	dat, _, err := usClient.ClassicWoWItemClass(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemSubclass(t *testing.T) {
	dat, _, err := usClient.ClassicWoWItemSubclass(context.Background(), 0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItem(t *testing.T) {
	dat, _, err := usClient.ClassicWoWItem(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWItemMedia(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPlayableClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClass(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPlayableClass(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassMedia(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPlayableClassMedia(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPlayableRacesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRace(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPlayableRace(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPowerTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerType(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPowerType(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPSeasonsIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPSeasonsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPSeason(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPSeason(context.Background(), 4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPRegionIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPRegionIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPRegionalSeasonIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPRegionalSeasonIndex(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPRegionalSeason(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPRegionalSeason(context.Background(), 1, 4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPLeaderboardsIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPLeaderboardsIndex(context.Background(), 1, 4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPLeaderboards(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPLeaderboards(context.Background(), 1, 4, wow.Bracket2v2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = usClient.ClassicWoWPvPLeaderboards(context.Background(), 1, 4, wow.Bracket3v3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPvPRewardsIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWPvPRewardsIndex(context.Background(), 1, 4)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRealmIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWRealmIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRealm(t *testing.T) {
	dat, _, err := usClient.ClassicWoWRealm(context.Background(), "westfall")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRegionIndex(t *testing.T) {
	dat, _, err := usClient.ClassicWoWRegionIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRegion(t *testing.T) {
	dat, _, err := usClient.ClassicWoWRegion(context.Background(), 41)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
