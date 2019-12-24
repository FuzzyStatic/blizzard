package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wowp"
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

func TestWoWCharacterEquipmentSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterEquipmentSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterHunterPetsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterHunterPetsSummary("illidan", "defraudedguy")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMediaSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterMediaSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileSummary("illidan", "wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileStatus(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileStatus("illidan", "wildz")
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

func TestWoWCharacterPvPBracketStatistics(t *testing.T) {
	dat, _, err := c.WoWCharacterPvPBracketStatistics("tichondrius", "tmpikaboo", wowp.Bracket2v2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = c.WoWCharacterPvPBracketStatistics("tichondrius", "tmpikaboo", wowp.Bracket3v3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = c.WoWCharacterPvPBracketStatistics("icecrown", "soballer", wowp.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPvPSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterPvPSummary("tichondrius", "tmpikaboo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
