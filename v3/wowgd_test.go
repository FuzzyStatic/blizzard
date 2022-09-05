package blizzard

import (
	"context"
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wowgd"
	"github.com/FuzzyStatic/blizzard/wowsearch"
)

func TestWoWAchievementCategoriesIndex(t *testing.T) {
	dat, _, err := usClient.WoWAchievementCategoriesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementCategory(t *testing.T) {
	dat, _, err := usClient.WoWAchievementCategory(context.Background(), 81)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementIndex(t *testing.T) {
	dat, _, err := usClient.WoWAchievementIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievement(t *testing.T) {
	dat, _, err := usClient.WoWAchievement(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementMedia(t *testing.T) {
	dat, _, err := usClient.WoWAchievementMedia(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAuctions(t *testing.T) {
	dat, _, err := usClient.WoWAuctions(context.Background(), 1138)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceIndex(t *testing.T) {
	dat, _, err := usClient.WoWAzeriteEssenceIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssence(t *testing.T) {
	dat, _, err := usClient.WoWAzeriteEssence(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceMedia(t *testing.T) {
	dat, _, err := usClient.WoWAzeriteEssenceMedia(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealmsIndex(t *testing.T) {
	dat, _, err := usClient.WoWConnectedRealmsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealm(t *testing.T) {
	dat, _, err := usClient.WoWConnectedRealm(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantsIndex(t *testing.T) {
	dat, _, err := usClient.WoWCovenantsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenant(t *testing.T) {
	dat, _, err := usClient.WoWCovenant(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantMedia(t *testing.T) {
	dat, _, err := usClient.WoWCovenantMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantSoulbindsIndex(t *testing.T) {
	dat, _, err := usClient.WoWCovenantSoulbindsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantSoulbind(t *testing.T) {
	dat, _, err := usClient.WoWCovenantSoulbind(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantConduitsIndex(t *testing.T) {
	dat, _, err := usClient.WoWCovenantConduitsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCovenantConduit(t *testing.T) {
	dat, _, err := usClient.WoWCovenantConduit(context.Background(), 13)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := usClient.WoWCreatureFamiliesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamily(t *testing.T) {
	dat, _, err := usClient.WoWCreatureFamily(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := usClient.WoWCreatureTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureType(t *testing.T) {
	dat, _, err := usClient.WoWCreatureType(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreature(t *testing.T) {
	dat, _, err := usClient.WoWCreature(context.Background(), 30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := usClient.WoWCreatureDisplayMedia(context.Background(), 30221)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := usClient.WoWCreatureFamilyMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := usClient.WoWGuildCrestComponentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := usClient.WoWGuildCrestBorderMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := usClient.WoWGuildCrestEmblemMedia(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClassesIndex(t *testing.T) {
	dat, _, err := usClient.WoWItemClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClass(t *testing.T) {
	dat, _, err := usClient.WoWItemClass(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSetsIndex(t *testing.T) {
	dat, _, err := usClient.WoWItemSetsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSet(t *testing.T) {
	dat, _, err := usClient.WoWItemClass(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSubclass(t *testing.T) {
	dat, _, err := usClient.WoWItemSubclass(context.Background(), 0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItem(t *testing.T) {
	dat, _, err := usClient.WoWItem(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSearch(t *testing.T) {
	dat, _, err := usClient.WoWItemSearch(context.Background(),
		wowsearch.Page(1),
		wowsearch.PageSize(5),
		wowsearch.OrderBy("level:desc"),
		wowsearch.Field().
			AND("item_class.id", "0").
			AND("item_subclass.id", "5||7||8||2||3||0||1||9").
			NOT("sell_price", "0"))

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		t.Logf("%+v\n", dat)
	}
}

func TestWoWItemMedia(t *testing.T) {
	dat, _, err := usClient.WoWItemMedia(context.Background(), 19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalExpansionsIndex(t *testing.T) {
	dat, _, err := usClient.WoWJournalExpansionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalExpansion(t *testing.T) {
	dat, _, err := usClient.WoWJournalExpansion(context.Background(), 396)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalEncountersIndex(t *testing.T) {
	dat, _, err := usClient.WoWJournalEncountersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalEncounter(t *testing.T) {
	dat, _, err := usClient.WoWJournalEncounter(context.Background(), 396)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstancesIndex(t *testing.T) {
	dat, _, err := usClient.WoWJournalInstancesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstance(t *testing.T) {
	dat, _, err := usClient.WoWJournalInstance(context.Background(), 63)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWJournalInstanceMedia(t *testing.T) {
	dat, _, err := usClient.WoWJournalInstanceMedia(context.Background(), 968)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMediaSearch(t *testing.T) {
	dat, _, err := usClient.WoWMediaSearch(context.Background(), "item", "id", 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingIndex(t *testing.T) {
	dat, _, err := usClient.WoWModifiedCraftingIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingCategoryIndex(t *testing.T) {
	dat, _, err := usClient.WoWModifiedCraftingCategoryIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingCategory(t *testing.T) {
	dat, _, err := usClient.WoWModifiedCraftingCategory(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingReagentSlotTypeIndex(t *testing.T) {
	dat, _, err := usClient.WoWModifiedCraftingReagentSlotTypeIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWModifiedCraftingReagentSlotType(t *testing.T) {
	dat, _, err := usClient.WoWModifiedCraftingReagentSlotType(context.Background(), 16)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMountIndex(t *testing.T) {
	dat, _, err := usClient.WoWMountIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMount(t *testing.T) {
	dat, _, err := usClient.WoWMount(context.Background(), 6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffixIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneAffixIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffix(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneAffix(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffixMedia(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneAffixMedia(context.Background(), 3)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeonIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneDungeonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeon(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneDungeon(context.Background(), 375)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriodIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystonePeriodIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriod(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystonePeriod(context.Background(), 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeasonIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneSeasonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeason(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneSeason(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboardIndex(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneLeaderboardIndex(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboard(t *testing.T) {
	dat, _, err := usClient.WoWMythicKeystoneLeaderboard(context.Background(), 11, 197, 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicRaidLeaderboard(t *testing.T) {
	dat, _, err := usClient.WoWMythicRaidLeaderboard(context.Background(), "uldir", "alliance")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetIndex(t *testing.T) {
	dat, _, err := usClient.WoWPetIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPet(t *testing.T) {
	dat, _, err := usClient.WoWPet(context.Background(), 39)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetMedia(t *testing.T) {
	dat, _, err := usClient.WoWPetMedia(context.Background(), 39)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilityIndex(t *testing.T) {
	dat, _, err := usClient.WoWPetAbilityIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilty(t *testing.T) {
	dat, _, err := usClient.WoWPetAbility(context.Background(), 110)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetAbilityMedia(t *testing.T) {
	dat, _, err := usClient.WoWPetAbilityMedia(context.Background(), 110)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := usClient.WoWPlayableClassesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClass(t *testing.T) {
	dat, _, err := usClient.WoWPlayableClass(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassMedia(t *testing.T) {
	dat, _, err := usClient.WoWPlayableClassMedia(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassPvPTalentSlots(t *testing.T) {
	dat, _, err := usClient.WoWPlayableClassPvPTalentSlots(context.Background(), 7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := usClient.WoWPlayableRacesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRace(t *testing.T) {
	dat, _, err := usClient.WoWPlayableRace(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecializationIndex(t *testing.T) {
	dat, _, err := usClient.WoWPlayableSpecializationIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecialization(t *testing.T) {
	dat, _, err := usClient.WoWPlayableSpecialization(context.Background(), 262)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecializationMedia(t *testing.T) {
	dat, _, err := usClient.WoWPlayableSpecializationMedia(context.Background(), 262)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := usClient.WoWPowerTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerType(t *testing.T) {
	dat, _, err := usClient.WoWPowerType(context.Background(), 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfessionsIndex(t *testing.T) {
	dat, _, err := usClient.WoWProfessionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfession(t *testing.T) {
	dat, _, err := usClient.WoWProfession(context.Background(), 164)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfessionMedia(t *testing.T) {
	dat, _, err := usClient.WoWProfessionMedia(context.Background(), 164)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWProfessionSkillTier(t *testing.T) {
	dat, _, err := usClient.WoWProfessionSkillTier(context.Background(), 164, 2477)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRecipe(t *testing.T) {
	dat, _, err := usClient.WoWRecipe(context.Background(), 1631)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRecipeMedia(t *testing.T) {
	dat, _, err := usClient.WoWRecipeMedia(context.Background(), 1631)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeasonIndex(t *testing.T) {
	dat, _, err := usClient.WoWPvPSeasonIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeason(t *testing.T) {
	dat, _, err := usClient.WoWPvPSeason(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboardsIndex(t *testing.T) {
	dat, _, err := usClient.WoWPvPLeaderboardsIndex(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboard(t *testing.T) {
	dat, _, err := usClient.WoWPvPLeaderboard(context.Background(), 27, wowgd.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPRewardsIndex(t *testing.T) {
	dat, _, err := usClient.WoWPvPRewardsIndex(context.Background(), 27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTierMedia(t *testing.T) {
	dat, _, err := usClient.WoWPvPTierMedia(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTiersIndex(t *testing.T) {
	dat, _, err := usClient.WoWPvPTiersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTier(t *testing.T) {
	dat, _, err := usClient.WoWPvPTier(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestsIndex(t *testing.T) {
	dat, _, err := usClient.WoWQuestsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuest(t *testing.T) {
	dat, _, err := usClient.WoWQuest(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestCategoriesIndex(t *testing.T) {
	dat, _, err := usClient.WoWQuestCategoriesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestCategory(t *testing.T) {
	dat, _, err := usClient.WoWQuestCategory(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestAreasIndex(t *testing.T) {
	dat, _, err := usClient.WoWQuestAreasIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestArea(t *testing.T) {
	dat, _, err := usClient.WoWQuestArea(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestTypesIndex(t *testing.T) {
	dat, _, err := usClient.WoWQuestTypesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWQuestType(t *testing.T) {
	dat, _, err := usClient.WoWQuestType(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealmIndex(t *testing.T) {
	dat, _, err := usClient.WoWRealmIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealm(t *testing.T) {
	dat, _, err := usClient.WoWRealm(context.Background(), "tichondrius")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegionIndex(t *testing.T) {
	dat, _, err := usClient.WoWRegionIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegion(t *testing.T) {
	dat, _, err := usClient.WoWRegion(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFactionsIndex(t *testing.T) {
	dat, _, err := usClient.WoWReputationFactionsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFaction(t *testing.T) {
	dat, _, err := usClient.WoWReputationFaction(context.Background(), 21)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiersIndex(t *testing.T) {
	dat, _, err := usClient.WoWReputationTiersIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiers(t *testing.T) {
	dat, _, err := usClient.WoWReputationTiers(context.Background(), 2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWSpell(t *testing.T) {
	dat, _, err := usClient.WoWSpell(context.Background(), 196607)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWSpellMedia(t *testing.T) {
	dat, _, err := usClient.WoWSpellMedia(context.Background(), 196607)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTalentsIndex(t *testing.T) {
	dat, _, err := usClient.WoWTalentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTalent(t *testing.T) {
	dat, _, err := usClient.WoWTalent(context.Background(), 23106)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTalentsIndex(t *testing.T) {
	dat, _, err := usClient.WoWPvPTalentsIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTalent(t *testing.T) {
	dat, _, err := usClient.WoWPvPTalent(context.Background(), 11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentTreeIndex(t *testing.T) {
	dat, _, err := usClient.WoWTechTalentTreeIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentTree(t *testing.T) {
	dat, _, err := usClient.WoWTechTalentTree(context.Background(), 276)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentIndex(t *testing.T) {
	dat, _, err := usClient.WoWTechTalentIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalent(t *testing.T) {
	dat, _, err := usClient.WoWTechTalent(context.Background(), 863)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTechTalentMedia(t *testing.T) {
	dat, _, err := usClient.WoWTechTalentMedia(context.Background(), 863)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitlesIndex(t *testing.T) {
	dat, _, err := usClient.WoWTitlesIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitle(t *testing.T) {
	dat, _, err := usClient.WoWTitle(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWToken(t *testing.T) {
	dat, _, err := usClient.WoWToken(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
