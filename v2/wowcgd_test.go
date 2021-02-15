package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestClassicWoWConnectedRealmsIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWConnectedRealmsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWConnectedRealm(t *testing.T) {
	dat, _, err := c.ClassicWoWConnectedRealm(context.Background(), 4388)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamiliesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamily(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamily(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureType(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureType(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreature(t *testing.T) {
	dat, _, err := c.ClassicWoWCreature(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureDisplayMedia(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamilyMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestComponentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestBorderMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestEmblemMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemClassesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWItemClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemClass(t *testing.T) {
	dat, _, err := c.ClassicWoWItemClass(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemSubclass(t *testing.T) {
	dat, _, err := c.ClassicWoWItemSubclass(context.Background(), 0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItem(t *testing.T) {
	dat, _, err := c.ClassicWoWItem(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWItemMedia(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClass(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClass(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClassMedia(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableRacesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRace(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableRace(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPowerTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerType(t *testing.T) {
	dat, _, err := c.ClassicWoWPowerType(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRealmIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWRealmIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRealm(t *testing.T) {
	dat, _, err := c.ClassicWoWRealm(context.Background(), "westfall")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRegionIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWRegionIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWRegion(t *testing.T) {
	dat, _, err := c.ClassicWoWRegion(context.Background(), 41)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
