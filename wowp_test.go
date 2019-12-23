package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWCharacterAchievementsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAchievementsSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAppearanceSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAppearanceSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCollectionsIndex(t *testing.T) {
	dat, _, err := c.WoWCharacterCollectionsIndex("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMountsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterMountsCollectionSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPetsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterPetsCollectionSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfile("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfileSeason("illidan", "wildz", 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
