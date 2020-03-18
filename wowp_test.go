package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wowp"
)

func TestWoWCharacterAchievementsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAchievementsSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAppearanceSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAppearanceSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCollectionsIndex(t *testing.T) {
	dat, _, err := c.WoWCharacterCollectionsIndex("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMountsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterMountsCollectionSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPetsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterPetsCollectionSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEncountersSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterEncountersSummary("tichondrius", "wizprotector")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterDungeons(t *testing.T) {
	dat, _, err := c.WoWCharacterDungeons("tichondrius", "wizprotector")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterRaids(t *testing.T) {
	dat, _, err := c.WoWCharacterRaids("tichondrius", "wizprotector")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEquipmentSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterEquipmentSummary("illidan", "Wildz")
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
	dat, _, err := c.WoWCharacterMediaSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileStatus(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileStatus("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfile("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfileSeason("illidan", "Wildz", 1)
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

func TestWoWCharacterReputationsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterReputationsSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterSpecializationsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterSpecializationsSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterStatisticsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterStatisticsSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterTitlesSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterTitlesSummary("illidan", "Wildz")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuild(t *testing.T) {
	dat, _, err := c.WoWGuild("illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildAchievements(t *testing.T) {
	dat, _, err := c.WoWGuildAchievements("illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildRoster(t *testing.T) {
	dat, _, err := c.WoWGuildRoster("illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
