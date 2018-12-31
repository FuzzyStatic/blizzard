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
	dat, err := c.WoWMythicKeystoneDungeon(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

/*

func TestWoWMythicKeystoneAffixIndex(t *testing.T) {
	dat, err := c.WoWMythicKeystoneAffixIndex()
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

*/
