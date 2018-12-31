package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWConnectedRealmIndex(t *testing.T) {
	dat, err := c.WoWConnectedRealmIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWConnectedRealm(t *testing.T) {
	dat, err := c.WoWConnectedRealm(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneAffixIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneAffixIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneAffix(t *testing.T) {
	dat, err := c.WoWMythicKeystoneAffix(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicRaidLeaderboard(t *testing.T) {
	dat, err := c.WoWMythicRaidLeaderboard("uldir", "alliance")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneDungeonIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneDungeonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneDungeon(t *testing.T) {
	dat, err := c.WoWMythicKeystoneDungeon(353)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystonePeriodIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystonePeriodIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystonePeriod(t *testing.T) {
	dat, err := c.WoWMythicKeystonePeriod(641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneSeasonIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneSeasonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneSeason(t *testing.T) {
	dat, err := c.WoWMythicKeystoneSeason(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneLeaderboardIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneLeaderboardIndex(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWMythicKeystoneLeaderboard(t *testing.T) {
	dat, err := c.WoWMythicKeystoneLeaderboard(11, 197, 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPlayableClassesIndex(t *testing.T) {
	dat, err := c.WoWPlayableClassesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPlayableClass(t *testing.T) {
	dat, err := c.WoWPlayableClass(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPlayableClassPVPTalentSlots(t *testing.T) {
	dat, err := c.WoWPlayableClassPVPTalentSlots(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPowerTypesIndex(t *testing.T) {
	dat, err := c.WoWPowerTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPowerType(t *testing.T) {
	dat, err := c.WoWPowerType(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPlayableRaceIndex(t *testing.T) {
	dat, err := c.WoWPlayableRacesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWPlayableRace(t *testing.T) {
	dat, err := c.WoWPlayableRace(2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRealmIndex(t *testing.T) {
	dat, err := c.WoWRealmIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRealm(t *testing.T) {
	dat, err := c.WoWRealm("tichondrius")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRegionIndex(t *testing.T) {
	dat, err := c.WoWRegionIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWRegion(t *testing.T) {
	dat, err := c.WoWRegion(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWToken(t *testing.T) {
	dat, err := c.WoWToken()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
