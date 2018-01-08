/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:06
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 21:56:27
 */

package worldofwarcraft

import (
	"fmt"
	"go-blizzard/blizzard"
	"testing"

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
