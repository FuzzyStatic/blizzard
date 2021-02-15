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

func TestWoWAuctions(t *testing.T) {
	dat, _, err := c.WoWAuctions(context.Background(), 1138)
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

func TestWoWConnectedRealmsIndex(t *testing.T) {
	dat, _, err := c.WoWConnectedRealmsIndex(context.Background())
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

func TestWoWCovenantsIndex(t *testing.T) {
	dat, _, err := c.WoWCovenantsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenant(t *testing.T) {
	dat, _, err := c.WoWCovenant(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantMedia(t *testing.T) {
	dat, _, err := c.WoWCovenantMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantSoulbindsIndex(t *testing.T) {
	dat, _, err := c.WoWCovenantSoulbindsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantSoulbind(t *testing.T) {
	dat, _, err := c.WoWCovenantSoulbind(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantConduitsIndex(t *testing.T) {
	dat, _, err := c.WoWCovenantConduitsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantConduit(t *testing.T) {
	dat, _, err := c.WoWCovenantConduit(context.Background(), 13)
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

func TestWoWModifiedCraftingIndex(t *testing.T) {
	dat, _, err := c.WoWModifiedCraftingIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingCategoryIndex(t *testing.T) {
	dat, _, err := c.WoWModifiedCraftingCategoryIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingCategory(t *testing.T) {
	dat, _, err := c.WoWModifiedCraftingCategory(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingReagentSlotTypeIndex(t *testing.T) {
	dat, _, err := c.WoWModifiedCraftingReagentSlotTypeIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingReagentSlotType(t *testing.T) {
	dat, _, err := c.WoWModifiedCraftingReagentSlotType(context.Background(), 16)
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

func TestWoWMythicKeystoneAffixMedia(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneAffixMedia(context.Background(), 3)
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
	dat, _, err := c.WoWMythicKeystoneDungeon(context.Background(), 375)
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

func TestWoWProfessionsIndex(t *testing.T) {
	dat, _, err := c.WoWProfessionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfession(t *testing.T) {
	dat, _, err := c.WoWProfession(context.Background(), 164)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfessionMedia(t *testing.T) {
	dat, _, err := c.WoWProfessionMedia(context.Background(), 164)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfessionSkillTier(t *testing.T) {
	dat, _, err := c.WoWProfessionSkillTier(context.Background(), 164, 2477)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRecipe(t *testing.T) {
	dat, _, err := c.WoWRecipe(context.Background(), 1631)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRecipeMedia(t *testing.T) {
	dat, _, err := c.WoWRecipeMedia(context.Background(), 1631)
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

func TestWoWQuestsIndex(t *testing.T) {
	dat, _, err := c.WoWQuestsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuest(t *testing.T) {
	dat, _, err := c.WoWQuest(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestCategoriesIndex(t *testing.T) {
	dat, _, err := c.WoWQuestCategoriesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestCategory(t *testing.T) {
	dat, _, err := c.WoWQuestCategory(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestAreasIndex(t *testing.T) {
	dat, _, err := c.WoWQuestAreasIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestArea(t *testing.T) {
	dat, _, err := c.WoWQuestArea(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestTypesIndex(t *testing.T) {
	dat, _, err := c.WoWQuestTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestType(t *testing.T) {
	dat, _, err := c.WoWQuestType(context.Background(), 1)
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

func TestWoWSpell(t *testing.T) {
	dat, _, err := c.WoWSpell(context.Background(), 196607)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWSpellMedia(t *testing.T) {
	dat, _, err := c.WoWSpellMedia(context.Background(), 196607)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTalentsIndex(t *testing.T) {
	dat, _, err := c.WoWTalentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTalent(t *testing.T) {
	dat, _, err := c.WoWTalent(context.Background(), 23106)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTalentIndex(t *testing.T) {
	dat, _, err := c.WoWPvPTalentIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTalent(t *testing.T) {
	dat, _, err := c.WoWPvPTalent(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentTreeIndex(t *testing.T) {
	dat, _, err := c.WoWTechTalentTreeIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentTree(t *testing.T) {
	dat, _, err := c.WoWTechTalentTree(context.Background(), 272)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentIndex(t *testing.T) {
	dat, _, err := c.WoWTechTalentIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalent(t *testing.T) {
	dat, _, err := c.WoWTechTalent(context.Background(), 863)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentMedia(t *testing.T) {
	dat, _, err := c.WoWTechTalentMedia(context.Background(), 863)
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
