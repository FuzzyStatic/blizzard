package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/v2/wowgd"
)

func TestWoWAchievementCategoriesIndex(t *testing.T) {
	dat, _, err := c.WoWAchievementCategoriesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementCategory(t *testing.T) {
	dat, _, err := c.WoWAchievementCategory(context.Background(), 81)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementIndex(t *testing.T) {
	dat, _, err := c.WoWAchievementIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievement(t *testing.T) {
	dat, _, err := c.WoWAchievement(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementMedia(t *testing.T) {
	dat, _, err := c.WoWAchievementMedia(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceIndex(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssenceIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssence(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssence(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceMedia(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssenceMedia(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealmIndex(t *testing.T) {
	dat, _, err := c.WoWConnectedRealmIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealm(t *testing.T) {
	dat, _, err := c.WoWConnectedRealm(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := c.WoWCreatureFamiliesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamily(t *testing.T) {
	dat, _, err := c.WoWCreatureFamily(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := c.WoWCreatureTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureType(t *testing.T) {
	dat, _, err := c.WoWCreatureType(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreature(t *testing.T) {
	dat, _, err := c.WoWCreature(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := c.WoWCreatureDisplayMedia(context.Background(), 30221)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := c.WoWCreatureFamilyMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := c.WoWGuildCrestComponentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := c.WoWGuildCrestBorderMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := c.WoWGuildCrestEmblemMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClassesIndex(t *testing.T) {
	dat, _, err := c.WoWItemClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClass(t *testing.T) {
	dat, _, err := c.WoWItemClass(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSubclass(t *testing.T) {
	dat, _, err := c.WoWItemSubclass(context.Background(), 0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItem(t *testing.T) {
	dat, _, err := c.WoWItem(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemMedia(t *testing.T) {
	dat, _, err := c.WoWItemMedia(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalExpansionsIndex(t *testing.T) {
	dat, _, err := c.WoWJournalExpansionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalExpansion(t *testing.T) {
	dat, _, err := c.WoWJournalExpansion(context.Background(), 396)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalEncountersIndex(t *testing.T) {
	dat, _, err := c.WoWJournalEncountersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalEncounter(t *testing.T) {
	dat, _, err := c.WoWJournalEncounter(context.Background(), 396)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstancesIndex(t *testing.T) {
	dat, _, err := c.WoWJournalInstancesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstance(t *testing.T) {
	dat, _, err := c.WoWJournalInstance(context.Background(), 63)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstanceMedia(t *testing.T) {
	dat, _, err := c.WoWJournalInstanceMedia(context.Background(), 968)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMountIndex(t *testing.T) {
	dat, _, err := c.WoWMountIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMount(t *testing.T) {
	dat, _, err := c.WoWMount(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffixIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneAffixIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffix(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneAffix(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicRaidLeaderboard(t *testing.T) {
	dat, _, err := c.WoWMythicRaidLeaderboard(context.Background(), "uldir", "alliance")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeonIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneDungeonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeon(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneDungeon(context.Background(), 353)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriodIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystonePeriodIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriod(t *testing.T) {
	dat, _, err := c.WoWMythicKeystonePeriod(context.Background(), 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeasonIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneSeasonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeason(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneSeason(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboardIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneLeaderboardIndex(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboard(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneLeaderboard(context.Background(), 11, 197, 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetIndex(t *testing.T) {
	dat, _, err := c.WoWPetIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPet(t *testing.T) {
	dat, _, err := c.WoWPet(context.Background(), 39)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetMedia(t *testing.T) {
	dat, _, err := c.WoWPetMedia(context.Background(), 39)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilityIndex(t *testing.T) {
	dat, _, err := c.WoWPetAbilityIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilty(t *testing.T) {
	dat, _, err := c.WoWPetAbility(context.Background(), 110)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilityMedia(t *testing.T) {
	dat, _, err := c.WoWPetAbilityMedia(context.Background(), 110)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClass(t *testing.T) {
	dat, _, err := c.WoWPlayableClass(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassPvPTalentSlots(t *testing.T) {
	dat, _, err := c.WoWPlayableClassPvPTalentSlots(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableRacesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRace(t *testing.T) {
	dat, _, err := c.WoWPlayableRace(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecializationIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableSpecializationIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecialization(t *testing.T) {
	dat, _, err := c.WoWPlayableSpecialization(context.Background(), 262)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := c.WoWPowerTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerType(t *testing.T) {
	dat, _, err := c.WoWPowerType(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeasonIndex(t *testing.T) {
	dat, _, err := c.WoWPvPSeasonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeason(t *testing.T) {
	dat, _, err := c.WoWPvPSeason(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboardsIndex(t *testing.T) {
	dat, _, err := c.WoWPvPLeaderboardsIndex(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboard(t *testing.T) {
	dat, _, err := c.WoWPvPLeaderboard(context.Background(), 27, wowgd.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPRewardsIndex(t *testing.T) {
	dat, _, err := c.WoWPvPRewardsIndex(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTierMedia(t *testing.T) {
	dat, _, err := c.WoWPvPTierMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTiersIndex(t *testing.T) {
	dat, _, err := c.WoWPvPTiersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTier(t *testing.T) {
	dat, _, err := c.WoWPvPTier(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFactionsIndex(t *testing.T) {
	dat, _, err := c.WoWReputationFactionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFaction(t *testing.T) {
	dat, _, err := c.WoWReputationFaction(context.Background(), 21)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiersIndex(t *testing.T) {
	dat, _, err := c.WoWReputationTiersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiers(t *testing.T) {
	dat, _, err := c.WoWReputationTiers(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitlesIndex(t *testing.T) {
	dat, _, err := c.WoWTitlesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitle(t *testing.T) {
	dat, _, err := c.WoWTitle(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealmIndex(t *testing.T) {
	dat, _, err := c.WoWRealmIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealm(t *testing.T) {
	dat, _, err := c.WoWRealm(context.Background(), "tichondrius")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegionIndex(t *testing.T) {
	dat, _, err := c.WoWRegionIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegion(t *testing.T) {
	dat, _, err := c.WoWRegion(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWToken(t *testing.T) {
	dat, _, err := c.WoWToken(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
