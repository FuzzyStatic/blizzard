/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:06
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 13:06:39
 */

package worldofwarcraft

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard"
	"github.com/spf13/viper"
)

var w *WorldOfWarcraft

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	accessToken := viper.GetString("API.access_token")
	apiKey := viper.GetString("API.api_key")

	w = New(
		blizzard.Auth{
			AccessToken: accessToken,
			APIKey:      apiKey,
		},
		blizzard.US,
	)
}

func TestGetConnectedRealmIndex(t *testing.T) {
	var (
		connectedRealmIndexJSON *[]byte
		connectedRealmIndex     *ConnectedRealmIndex
		err                     error
	)

	connectedRealmIndexJSON, err = w.GetConnectedRealmIndexJSON()
	if err != nil {
		t.Fail()
	}

	connectedRealmIndex, err = w.GetConnectedRealmIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*connectedRealmIndexJSON))
	fmt.Println(*connectedRealmIndex)
}

func TestGetConnectedRealm(t *testing.T) {
	var (
		connectedRealmJSON *[]byte
		connectedRealm     *ConnectedRealm
		err                error
	)

	connectedRealmJSON, err = w.GetConnectedRealmJSON(3729)
	if err != nil {
		t.Fail()
	}

	connectedRealm, err = w.GetConnectedRealm(3729)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*connectedRealmJSON))
	fmt.Println(*connectedRealm)
}

func TestGetMythicLeaderboardIndex(t *testing.T) {
	var (
		mythicLeaderboardIndexJSON *[]byte
		mythicLeaderboardIndex     *MythicLeaderboardIndex
		err                        error
	)

	mythicLeaderboardIndexJSON, err = w.GetMythicLeaderboardIndexJSON(3729)
	if err != nil {
		t.Fail()
	}

	mythicLeaderboardIndex, err = w.GetMythicLeaderboardIndex(3729)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*mythicLeaderboardIndexJSON))
	fmt.Println(*mythicLeaderboardIndex)
}

func TestGetMythicLeaderboard(t *testing.T) {
	var (
		mythicLeaderboardJSON *[]byte
		mythicLeaderboard     *MythicLeaderboard
		err                   error
	)

	mythicLeaderboardJSON, err = w.GetMythicLeaderboardJSON(3729, 198, 627)
	if err != nil {
		t.Fail()
	}

	mythicLeaderboard, err = w.GetMythicLeaderboard(3729, 198, 627)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*mythicLeaderboardJSON))
	fmt.Println(*mythicLeaderboard)
}

func TestGetRealmIndex(t *testing.T) {
	var (
		realmIndexJSON *[]byte
		realmIndex     *RealmIndex
		err            error
	)

	realmIndexJSON, err = w.GetRealmIndexJSON()
	if err != nil {
		t.Fail()
	}

	realmIndex, err = w.GetRealmIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*realmIndexJSON))
	fmt.Println(*realmIndex)
}

func TestGetRealm(t *testing.T) {
	var (
		realmJSON *[]byte
		realm     *Realm
		err       error
	)

	realmJSON, err = w.GetRealmJSON(59)
	if err != nil {
		t.Fail()
	}

	realm, err = w.GetRealm(59)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*realmJSON))
	fmt.Println(*realm)

	realmJSON, err = w.GetRealmJSON("malganis")
	if err != nil {
		t.Fail()
	}

	realm, err = w.GetRealm("malganis")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*realmJSON))
	fmt.Println(*realm)
}

func TestGetRegionIndex(t *testing.T) {
	var (
		regionIndexJSON *[]byte
		regionIndex     *RegionIndex
		err             error
	)

	regionIndexJSON, err = w.GetRegionIndexJSON()
	if err != nil {
		t.Fail()
	}

	regionIndex, err = w.GetRegionIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*regionIndexJSON))
	fmt.Println(*regionIndex)
}

func TestGetRegion(t *testing.T) {
	var (
		regionJSON *[]byte
		region     *Region
		err        error
	)

	regionJSON, err = w.GetRegionJSON(1)
	if err != nil {
		t.Fail()
	}

	region, err = w.GetRegion(1)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*regionJSON))
	fmt.Println(*region)
}

func TestGetTokenIndex(t *testing.T) {
	var (
		tokenIndexJSON *[]byte
		tokenIndex     *TokenIndex
		err            error
	)

	tokenIndexJSON, err = w.GetTokenIndexJSON()
	if err != nil {
		t.Fail()
	}

	tokenIndex, err = w.GetTokenIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*tokenIndexJSON))
	fmt.Println(*tokenIndex)
}

func TestGetAchievement(t *testing.T) {
	var (
		achievementJSON *[]byte
		achievement     *Achievement
		err             error
	)

	achievementJSON, err = w.GetAchievementJSON(2144)
	if err != nil {
		t.Fail()
	}

	achievement, err = w.GetAchievement(2144)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*achievementJSON))
	fmt.Println(*achievement)
}

func TestGetAuctionData(t *testing.T) {
	var (
		auctionDataJSON *[]byte
		auctionData     *AuctionData
		err             error
	)

	auctionDataJSON, err = w.GetAuctionDataJSON("malganis")
	if err != nil {
		t.Fail()
	}

	auctionData, err = w.GetAuctionData("malganis")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*auctionDataJSON))
	fmt.Println(*auctionData)
}

func TestGetAuctions(t *testing.T) {
	var (
		auctionsArr []*Auctions
		err         error
	)

	auctionsArr, err = w.GetAuctions("malganis")
	if err != nil {
		t.Fail()
	}

	for _, auctions := range auctionsArr {
		fmt.Println(*auctions)
	}
}

func TestGetBossIndex(t *testing.T) {
	var (
		bossIndexJSON *[]byte
		bossIndex     *BossIndex
		err           error
	)

	bossIndexJSON, err = w.GetBossIndexJSON()
	if err != nil {
		t.Fail()
	}

	bossIndex, err = w.GetBossIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*bossIndexJSON))
	fmt.Println(*bossIndex)
}

func TestGetBoss(t *testing.T) {
	var (
		bossJSON *[]byte
		boss     *Boss
		err      error
	)

	bossJSON, err = w.GetBossJSON(24723)
	if err != nil {
		t.Fail()
	}

	boss, err = w.GetBoss(24723)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*bossJSON))
	fmt.Println(*boss)
}

func TestGetChallengeRegionLeaderboard(t *testing.T) {
	var (
		challengeRegionLeaderboardJSON *[]byte
		challengeRegionLeaderboard     *ChallengeRegionLeaderboard
		err                            error
	)

	challengeRegionLeaderboardJSON, err = w.GetChallengeRegionLeaderboardJSON()
	if err != nil {
		t.Fail()
	}

	challengeRegionLeaderboard, err = w.GetChallengeRegionLeaderboard()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*challengeRegionLeaderboardJSON))
	fmt.Println(*challengeRegionLeaderboard)
}

func TestGetCharacter(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacter("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithAchievements(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithAchievementsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithAchievements("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithAppearance(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithAppearanceJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithAppearance("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithFeed(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithFeedJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithFeed("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithGuild(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithGuildJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithGuild("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithHunterPets(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithHunterPetsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithHunterPets("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithItems(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithItemsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithItems("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithMounts(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithMountsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithMounts("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithPets(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithPetsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithPets("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithPetSlots(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithPetSlotsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithPetSlots("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithProfessions(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithProfessionsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithProfessions("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithProgression(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithProgressionJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithProgression("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithPVP(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithPVPJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithPVP("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithQuests(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithQuestsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithQuests("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithReputation(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithReputationJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithReputation("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithStatistics(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithStatisticsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithStatistics("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithStats(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithStatsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithStats("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithTalents(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithTalentsJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithTalents("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithTitles(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithTitlesJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithTitles("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithAudit(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithAuditJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithAudit("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetCharacterWithAll(t *testing.T) {
	var (
		characterJSON *[]byte
		character     *Character
		err           error
	)

	characterJSON, err = w.GetCharacterWithAllJSON("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	character, err = w.GetCharacterWithAll("illidan", "flowbs")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*characterJSON))
	fmt.Println(*character)
}

func TestGetGuild(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuild("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetGuildWithMembers(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildWithMembersJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuildWithMembers("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetGuildWithAchievements(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildWithAchievementsJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuildWithAchievements("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetGuildWithNews(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildWithNewsJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuildWithNews("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetGuildWithChallenge(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildWithChallengeJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuildWithChallenge("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetGuildWithAll(t *testing.T) {
	var (
		guildJSON *[]byte
		guild     *Guild
		err       error
	)

	guildJSON, err = w.GetGuildWithAllJSON("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	guild, err = w.GetGuildWithAll("illidan", "limit")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*guildJSON))
	fmt.Println(*guild)
}

func TestGetItem(t *testing.T) {
	var (
		itemJSON *[]byte
		item     *Item
		err      error
	)

	itemJSON, err = w.GetItemJSON(18803)
	if err != nil {
		t.Fail()
	}

	item, err = w.GetItem(18803)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*itemJSON))
	fmt.Println(*item)
}

func TestGetSet(t *testing.T) {
	var (
		setJSON *[]byte
		set     *Set
		err     error
	)

	setJSON, err = w.GetSetJSON(1060)
	if err != nil {
		t.Fail()
	}

	set, err = w.GetSet(1060)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*setJSON))
	fmt.Println(*set)
}

func TestGetMountIndex(t *testing.T) {
	var (
		mountIndexJSON *[]byte
		mountIndex     *MountIndex
		err            error
	)

	mountIndexJSON, err = w.GetMountIndexJSON()
	if err != nil {
		t.Fail()
	}

	mountIndex, err = w.GetMountIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*mountIndexJSON))
	fmt.Println(*mountIndex)
}

func TestGetPetIndex(t *testing.T) {
	var (
		petIndexJSON *[]byte
		petIndex     *PetIndex
		err          error
	)

	petIndexJSON, err = w.GetPetIndexJSON()
	if err != nil {
		t.Fail()
	}

	petIndex, err = w.GetPetIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*petIndexJSON))
	fmt.Println(*petIndex)
}

func TestGetPetAbility(t *testing.T) {
	var (
		petAbilityJSON *[]byte
		petAbility     *PetAbility
		err            error
	)

	petAbilityJSON, err = w.GetPetAbilityJSON(640)
	if err != nil {
		t.Fail()
	}

	petAbility, err = w.GetPetAbility(640)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*petAbilityJSON))
	fmt.Println(*petAbility)
}

func TestGetPetSpecies(t *testing.T) {
	var (
		petSpeciesJSON *[]byte
		petSpecies     *PetSpecies
		err            error
	)

	petSpeciesJSON, err = w.GetPetSpeciesJSON(258)
	if err != nil {
		t.Fail()
	}

	petSpecies, err = w.GetPetSpecies(258)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*petSpeciesJSON))
	fmt.Println(*petSpecies)
}

func TestGetPetStats(t *testing.T) {
	var (
		petStatsJSON *[]byte
		petStats     *PetStats
		err          error
	)

	petStatsJSON, err = w.GetPetStatsJSON(258, 25, 5, 4)
	if err != nil {
		t.Fail()
	}

	petStats, err = w.GetPetStats(258, 25, 5, 4)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*petStatsJSON))
	fmt.Println(*petStats)
}

func TestGet2v2Leaderboard(t *testing.T) {
	var (
		leaderboardJSON *[]byte
		leaderboard     *Leaderboard
		err             error
	)

	leaderboardJSON, err = w.Get2v2LeaderboardJSON()
	if err != nil {
		t.Fail()
	}

	leaderboard, err = w.Get2v2Leaderboard()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*leaderboardJSON))
	fmt.Println(*leaderboard)
}

func TestGet3v3Leaderboard(t *testing.T) {
	var (
		leaderboardJSON *[]byte
		leaderboard     *Leaderboard
		err             error
	)

	leaderboardJSON, err = w.Get3v3LeaderboardJSON()
	if err != nil {
		t.Fail()
	}

	leaderboard, err = w.Get3v3Leaderboard()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*leaderboardJSON))
	fmt.Println(*leaderboard)
}

func TestGet5v5Leaderboard(t *testing.T) {
	var (
		leaderboardJSON *[]byte
		leaderboard     *Leaderboard
		err             error
	)

	leaderboardJSON, err = w.Get5v5LeaderboardJSON()
	if err != nil {
		t.Fail()
	}

	leaderboard, err = w.Get5v5Leaderboard()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*leaderboardJSON))
	fmt.Println(*leaderboard)
}

func TestGetRBGLeaderboard(t *testing.T) {
	var (
		leaderboardJSON *[]byte
		leaderboard     *Leaderboard
		err             error
	)

	leaderboardJSON, err = w.GetRBGLeaderboardJSON()
	if err != nil {
		t.Fail()
	}

	leaderboard, err = w.GetRBGLeaderboard()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*leaderboardJSON))
	fmt.Println(*leaderboard)
}

func TestGetQuest(t *testing.T) {
	var (
		questJSON *[]byte
		quest     *Quest
		err       error
	)

	questJSON, err = w.GetQuestJSON(13146)
	if err != nil {
		t.Fail()
	}

	quest, err = w.GetQuest(13146)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*questJSON))
	fmt.Println(*quest)
}

func TestGetRealmStatus(t *testing.T) {
	var (
		realmStatusJSON *[]byte
		realmStatus     *RealmStatus
		err             error
	)

	realmStatusJSON, err = w.GetRealmStatusJSON()
	if err != nil {
		t.Fail()
	}

	realmStatus, err = w.GetRealmStatus()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*realmStatusJSON))
	fmt.Println(*realmStatus)
}

func TestGetRecipe(t *testing.T) {
	var (
		recipeJSON *[]byte
		recipe     *Recipe
		err        error
	)

	recipeJSON, err = w.GetRecipeJSON(33994)
	if err != nil {
		t.Fail()
	}

	recipe, err = w.GetRecipe(33994)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*recipeJSON))
	fmt.Println(*recipe)
}

func TestGetSpell(t *testing.T) {
	var (
		spellJSON *[]byte
		spell     *Spell
		err       error
	)

	spellJSON, err = w.GetSpellJSON(4900)
	if err != nil {
		t.Fail()
	}

	spell, err = w.GetSpell(4900)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*spellJSON))
	fmt.Println(*spell)
}

func TestGetZoneIndex(t *testing.T) {
	var (
		zoneIndexJSON *[]byte
		zoneIndex     *ZoneIndex
		err           error
	)

	zoneIndexJSON, err = w.GetZoneIndexJSON()
	if err != nil {
		t.Fail()
	}

	zoneIndex, err = w.GetZoneIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*zoneIndexJSON))
	fmt.Println(*zoneIndex)
}

func TestGetZone(t *testing.T) {
	var (
		zoneJSON *[]byte
		zone     *Zone
		err      error
	)

	zoneJSON, err = w.GetZoneJSON(4131)
	if err != nil {
		t.Fail()
	}

	zone, err = w.GetZone(4131)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*zoneJSON))
	fmt.Println(*zone)
}
