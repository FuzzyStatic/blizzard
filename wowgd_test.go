package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/wowgd"
)

func TestWoWAchievementCategoriesIndex(t *testing.T) {
	dat, _, err := c.WoWAchievementCategoriesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementCategory(t *testing.T) {
	dat, _, err := c.WoWAchievementCategory(81)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementIndex(t *testing.T) {
	dat, _, err := c.WoWAchievementIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievement(t *testing.T) {
	dat, _, err := c.WoWAchievement(6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAchievementMedia(t *testing.T) {
	dat, _, err := c.WoWAchievementMedia(6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceIndex(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssenceIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssence(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssence(6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWAzeriteEssenceMedia(t *testing.T) {
	dat, _, err := c.WoWAzeriteEssenceMedia(6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealmIndex(t *testing.T) {
	dat, _, err := c.WoWConnectedRealmIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWConnectedRealm(t *testing.T) {
	dat, _, err := c.WoWConnectedRealm(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := c.WoWCreatureFamiliesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamily(t *testing.T) {
	dat, _, err := c.WoWCreatureFamily(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := c.WoWCreatureTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureType(t *testing.T) {
	dat, _, err := c.WoWCreatureType(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreature(t *testing.T) {
	dat, _, err := c.WoWCreature(30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := c.WoWCreatureDisplayMedia(30221)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := c.WoWCreatureFamilyMedia(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := c.WoWGuildCrestComponentsIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := c.WoWGuildCrestBorderMedia(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := c.WoWGuildCrestEmblemMedia(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClassesIndex(t *testing.T) {
	dat, _, err := c.WoWItemClassesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemClass(t *testing.T) {
	dat, _, err := c.WoWItemClass(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemSubclass(t *testing.T) {
	dat, _, err := c.WoWItemSubclass(0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItem(t *testing.T) {
	dat, _, err := c.WoWItem(19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWItemMedia(t *testing.T) {
	dat, _, err := c.WoWItemMedia(19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMountIndex(t *testing.T) {
	dat, _, err := c.WoWMountIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMount(t *testing.T) {
	dat, _, err := c.WoWMount(6)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffixIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneAffixIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneAffix(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneAffix(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicRaidLeaderboard(t *testing.T) {
	dat, _, err := c.WoWMythicRaidLeaderboard("uldir", "alliance")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeonIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneDungeonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneDungeon(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneDungeon(353)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriodIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystonePeriodIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystonePeriod(t *testing.T) {
	dat, _, err := c.WoWMythicKeystonePeriod(641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeasonIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneSeasonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneSeason(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneSeason(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboardIndex(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneLeaderboardIndex(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWMythicKeystoneLeaderboard(t *testing.T) {
	dat, _, err := c.WoWMythicKeystoneLeaderboard(11, 197, 641)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPetIndex(t *testing.T) {
	dat, _, err := c.WoWPetIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPet(t *testing.T) {
	dat, _, err := c.WoWPet(39)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableClassesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClass(t *testing.T) {
	dat, _, err := c.WoWPlayableClass(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableClassPvPTalentSlots(t *testing.T) {
	dat, _, err := c.WoWPlayableClassPvPTalentSlots(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableRacesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableRace(t *testing.T) {
	dat, _, err := c.WoWPlayableRace(2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecializationIndex(t *testing.T) {
	dat, _, err := c.WoWPlayableSpecializationIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPlayableSpecialization(t *testing.T) {
	dat, _, err := c.WoWPlayableSpecialization(262)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := c.WoWPowerTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPowerType(t *testing.T) {
	dat, _, err := c.WoWPowerType(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeasonIndex(t *testing.T) {
	dat, _, err := c.WoWPvPSeasonIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPSeason(t *testing.T) {
	dat, _, err := c.WoWPvPSeason(27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboardsIndex(t *testing.T) {
	dat, _, err := c.WoWPvPLeaderboardsIndex(27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPLeaderboard(t *testing.T) {
	dat, _, err := c.WoWPvPLeaderboard(27, wowgd.BracketRBG)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPRewardsIndex(t *testing.T) {
	dat, _, err := c.WoWPvPRewardsIndex(27)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTierMedia(t *testing.T) {
	dat, _, err := c.WoWPvPTierMedia(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTiersIndex(t *testing.T) {
	dat, _, err := c.WoWPvPTiersIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWPvPTier(t *testing.T) {
	dat, _, err := c.WoWPvPTier(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFactionsIndex(t *testing.T) {
	dat, _, err := c.WoWReputationFactionsIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationFaction(t *testing.T) {
	dat, _, err := c.WoWReputationFaction(21)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiersIndex(t *testing.T) {
	dat, _, err := c.WoWReputationTiersIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWReputationTiers(t *testing.T) {
	dat, _, err := c.WoWReputationTiers(2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitlesIndex(t *testing.T) {
	dat, _, err := c.WoWTitlesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWTitle(t *testing.T) {
	dat, _, err := c.WoWTitle(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealmIndex(t *testing.T) {
	dat, _, err := c.WoWRealmIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRealm(t *testing.T) {
	dat, _, err := c.WoWRealm("tichondrius")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegionIndex(t *testing.T) {
	dat, _, err := c.WoWRegionIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWRegion(t *testing.T) {
	dat, _, err := c.WoWRegion(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWToken(t *testing.T) {
	dat, _, err := c.WoWToken()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
