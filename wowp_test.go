package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wow"
)

func TestWoWCharacterAchievementsSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterAchievementsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAchievementsStatistics(t *testing.T) {
	dat, _, err := usClient.WoWCharacterAchievementsStatistics(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterAppearanceSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterAppearanceSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCollectionsIndex(t *testing.T) {
	dat, _, err := usClient.WoWCharacterCollectionsIndex(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMountsCollectionSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterMountsCollectionSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPetsCollectionSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterPetsCollectionSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEncountersSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterEncountersSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterDungeons(t *testing.T) {
	dat, _, err := usClient.WoWCharacterDungeons(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterRaids(t *testing.T) {
	dat, _, err := usClient.WoWCharacterRaids(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterEquipmentSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterEquipmentSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterHunterPetsSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterHunterPetsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMediaSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterMediaSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestMythicKeystoneProfileIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneProfileIndex(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeasonDetails(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneSeasonDetails(context.Background(), "illidan", "covlol", 3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterProfileSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterProfileStatus(t *testing.T) {
	dat, _, err := usClient.WoWCharacterProfileStatus(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, _, err := usClient.WoWCharacterMythicKeystoneProfile(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, _, err := usClient.WoWCharacterMythicKeystoneProfileSeason(context.Background(), "illidan", "covlol", 3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPvPBracketStatistics(t *testing.T) {
	dat, _, err := usClient.WoWCharacterPvPBracketStatistics(context.Background(), "tichondrius", "tmpikaboo", wow.Bracket2v2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = usClient.WoWCharacterPvPBracketStatistics(context.Background(), "tichondrius", "tmpikaboo", wow.Bracket3v3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}

	dat, _, err = usClient.WoWCharacterPvPBracketStatistics(context.Background(), "icecrown", "soballer", wow.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterPvPSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterPvPSummary(context.Background(), "tichondrius", "tmpikaboo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterQuests(t *testing.T) {
	dat, _, err := usClient.WoWCharacterQuests(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterCompletedQuests(t *testing.T) {
	dat, _, err := usClient.WoWCharacterCompletedQuests(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterReputationsSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterReputationsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterSoulbinds(t *testing.T) {
	dat, _, err := usClient.WoWCharacterSoulbinds(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterSpecializationsSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterSpecializationsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterStatisticsSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterStatisticsSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterTitlesSummary(t *testing.T) {
	dat, _, err := usClient.WoWCharacterTitlesSummary(context.Background(), "illidan", "covlol")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuild(t *testing.T) {
	dat, _, err := usClient.WoWGuild(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildActivity(t *testing.T) {
	dat, _, err := usClient.WoWGuildActivity(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildAchievements(t *testing.T) {
	dat, _, err := usClient.WoWGuildAchievements(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildRoster(t *testing.T) {
	dat, _, err := usClient.WoWGuildRoster(context.Background(), "illidan", "Complexity Limit")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
