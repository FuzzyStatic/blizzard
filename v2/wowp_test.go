package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/v2/wowp"
)

func TestWoWCharacterAchievementsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAchievementsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAchievementsStatistics(t *testing.T) {
	dat, _, err := c.WoWCharacterAchievementsStatistics(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAppearanceSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterAppearanceSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCollectionsIndex(t *testing.T) {
	dat, _, err := c.WoWCharacterCollectionsIndex(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMountsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterMountsCollectionSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPetsCollectionSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterPetsCollectionSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEncountersSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterEncountersSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterDungeons(t *testing.T) {
	dat, _, err := c.WoWCharacterDungeons(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterRaids(t *testing.T) {
	dat, _, err := c.WoWCharacterRaids(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEquipmentSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterEquipmentSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterHunterPetsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterHunterPetsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMediaSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterMediaSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestMythicKeystoneProfileIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneProfileIndex(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeasonDetails(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneSeasonDetails(context.Background(), "illidan", "covlol", 3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileStatus(t *testing.T) {
	dat, _, err := c.WoWCharacterProfileStatus(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfile(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfileSeason(context.Background(), "illidan", "covlol", 3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPvPBracketStatistics(t *testing.T) {
	dat, _, err := c.WoWCharacterPvPBracketStatistics(context.Background(), "tichondrius", "tmpikaboo", wowp.Bracket2v2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = c.WoWCharacterPvPBracketStatistics(context.Background(), "tichondrius", "tmpikaboo", wowp.Bracket3v3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = c.WoWCharacterPvPBracketStatistics(context.Background(), "icecrown", "soballer", wowp.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPvPSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterPvPSummary(context.Background(), "tichondrius", "tmpikaboo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterQuests(t *testing.T) {
	dat, _, err := c.WoWCharacterQuests(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCompletedQuests(t *testing.T) {
	dat, _, err := c.WoWCharacterCompletedQuests(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterReputationsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterReputationsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterSoulbinds(t *testing.T) {
	dat, _, err := c.WoWCharacterSoulbinds(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterSpecializationsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterSpecializationsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterStatisticsSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterStatisticsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterTitlesSummary(t *testing.T) {
	dat, _, err := c.WoWCharacterTitlesSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuild(t *testing.T) {
	dat, _, err := c.WoWGuild(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildActivity(t *testing.T) {
	dat, _, err := c.WoWGuildActivity(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildAchievements(t *testing.T) {
	dat, _, err := c.WoWGuildAchievements(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildRoster(t *testing.T) {
	dat, _, err := c.WoWGuildRoster(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
