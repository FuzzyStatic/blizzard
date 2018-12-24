/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:33
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 19:46:27
 */

package diablo3

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard"
	"github.com/spf13/viper"
)

var d *Diablo3

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	accessToken := viper.GetString("API.access_token")
	apiKey := viper.GetString("API.api_key")

	d = New(
		blizzard.New(
			blizzard.Auth{
				AccessToken: accessToken,
				APIKey:      apiKey,
			},
			blizzard.US),
	)
}

func TestGetSeasonIndex(t *testing.T) {
	var (
		seasonIndexJSON []byte
		seasonIndex     *SeasonIndex
		err             error
	)

	seasonIndexJSON, err = d.GetSeasonIndexJSON()
	if err != nil {
		t.Fail()
	}

	seasonIndex, err = d.GetSeasonIndex()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(seasonIndexJSON))
	fmt.Printf("%+v", seasonIndex)
}

func TestGetSeason(t *testing.T) {
	var (
		seasonIndex *SeasonIndex
		seasonJSON  []byte
		season      *Season
		err         error
	)

	seasonIndex, err = d.GetSeasonIndex()
	if err != nil {
		t.Fail()
	}

	seasonJSON, err = d.GetSeasonJSON(seasonIndex.CurrentSeason)
	if err != nil {
		t.Fail()
	}

	season, err = d.GetSeason(seasonIndex.CurrentSeason)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(seasonJSON))
	fmt.Printf("%+v", season)
}

func TestGetSeasonAchievementPoints(t *testing.T) {
	var (
		seasonIndex                 *SeasonIndex
		seasonAchievementPointsJSON []byte
		seasonAchievementPoints     *SeasonAchievementPoints
		err                         error
	)

	seasonIndex, err = d.GetSeasonIndex()
	if err != nil {
		t.Fail()
	}

	seasonAchievementPointsJSON, err = d.GetSeasonAchievementPointsJSON(seasonIndex.CurrentSeason)
	if err != nil {
		t.Fail()
	}

	seasonAchievementPoints, err = d.GetSeasonAchievementPoints(seasonIndex.CurrentSeason)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(seasonAchievementPointsJSON))
	fmt.Printf("%+v", seasonAchievementPoints)
}

func TestGetSeasonRift(t *testing.T) {
	var (
		seasonRiftJSON []byte
		seasonRift     *SeasonRift
		err            error
	)

	seasonRiftJSON, err = d.GetCurrentSeasonNecromancerRiftJSON(true)
	if err != nil {
		t.Fail()
	}

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(seasonRiftJSON))
	fmt.Println(seasonRift)
}

func TestGetSeasonRiftRowFromRank(t *testing.T) {
	var (
		seasonRift            *SeasonRift
		seasonRiftRowFromRank *Row
		err                   error
	)

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	seasonRiftRowFromRank, err = seasonRift.GetSeasonRiftRowFromRank(1)
	if err != nil {
		t.Fail()
	}

	fmt.Printf("%+v", seasonRiftRowFromRank)
}

func TestGetSeasonRowPlayer(t *testing.T) {
	var (
		seasonRift            *SeasonRift
		seasonRiftRowFromRank *Row
		seasonRowPlayer       *Player
		err                   error
	)

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	seasonRiftRowFromRank, err = seasonRift.GetSeasonRiftRowFromRank(1)
	if err != nil {
		t.Fail()
	}

	seasonRowPlayer, err = seasonRiftRowFromRank.GetSeasonRowPlayer()
	if err != nil {
		t.Fail()
	}

	fmt.Printf("%+v", seasonRowPlayer)
}

func TestGetHeroBattleTagFromSeasonRowPlayer(t *testing.T) {
	var (
		seasonRift                       *SeasonRift
		seasonRiftRowFromRank            *Row
		seasonRowPlayer                  *Player
		heroBattleTagFromSeasonRowPlayer string
		err                              error
	)

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	seasonRiftRowFromRank, err = seasonRift.GetSeasonRiftRowFromRank(1)
	if err != nil {
		t.Fail()
	}

	seasonRowPlayer, err = seasonRiftRowFromRank.GetSeasonRowPlayer()
	if err != nil {
		t.Fail()
	}

	heroBattleTagFromSeasonRowPlayer, err = seasonRowPlayer.GetHeroBattleTagFromSeasonRowPlayer()
	if err != nil {
		t.Fail()
	}

	fmt.Println(heroBattleTagFromSeasonRowPlayer)
}

func TestGetHeroIDFromSeasonRowPlayer(t *testing.T) {
	var (
		seasonRift                *SeasonRift
		seasonRiftRowFromRank     *Row
		seasonRowPlayer           *Player
		heroIDFromSeasonRowPlayer int
		err                       error
	)

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	seasonRiftRowFromRank, err = seasonRift.GetSeasonRiftRowFromRank(1)
	if err != nil {
		t.Fail()
	}

	seasonRowPlayer, err = seasonRiftRowFromRank.GetSeasonRowPlayer()
	if err != nil {
		t.Fail()
	}

	heroIDFromSeasonRowPlayer, err = seasonRowPlayer.GetHeroIDFromSeasonRowPlayer()
	if err != nil {
		t.Fail()
	}

	fmt.Println(heroIDFromSeasonRowPlayer)
}

func TestGetHeroBattleTagAndIDFromSeasonRow(t *testing.T) {
	var (
		seasonRift            *SeasonRift
		seasonRiftRowFromRank *Row
		battleTag             string
		heroID                int
		err                   error
	)

	seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
	if err != nil {
		t.Fail()
	}

	seasonRiftRowFromRank, err = seasonRift.GetSeasonRiftRowFromRank(1)
	if err != nil {
		t.Fail()
	}

	battleTag, heroID, err = seasonRiftRowFromRank.GetHeroBattleTagAndIDFromSeasonRow()
	if err != nil {
		t.Fail()
	}

	fmt.Println(battleTag)
	fmt.Println(heroID)
}

func TestGetProfile(t *testing.T) {
	var (
		battleTag   = "FuzzyStatic#1384"
		profileJSON []byte
		profile     *Profile
		err         error
	)

	profileJSON, err = d.GetProfileJSON(battleTag)
	if err != nil {
		t.Fail()
	}

	profile, err = d.GetProfile(battleTag)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(profileJSON))
	fmt.Printf("%+v", profile)
}

func TestGetHeroNamesFromProfile(t *testing.T) {
	var (
		battleTag            = "FuzzyStatic#1384"
		profile              *Profile
		heroNamesFromProfile []string
		err                  error
	)

	profile, err = d.GetProfile(battleTag)
	if err != nil {
		t.Fail()
	}

	heroNamesFromProfile = profile.GetHeroNamesFromProfile()

	fmt.Println(heroNamesFromProfile)
}

func TestGetHeroIDsAndNamesFromProfile(t *testing.T) {
	var (
		battleTag                  = "FuzzyStatic#1384"
		profile                    *Profile
		heroIDsAndNamesFromProfile map[int]string
		err                        error
	)

	profile, err = d.GetProfile(battleTag)
	if err != nil {
		t.Fail()
	}

	heroIDsAndNamesFromProfile = profile.GetHeroIDsAndNamesFromProfile()

	fmt.Println(heroIDsAndNamesFromProfile)
}

func TestGetHero(t *testing.T) {
	var (
		battleTag = "DyS#1311"
		heroID    = 104452451
		heroJSON  []byte
		hero      *Hero
		err       error
	)

	heroJSON, err = d.GetHeroJSON(battleTag, heroID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	hero, err = d.GetHero(battleTag, heroID)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(string(heroJSON))
	fmt.Printf("%+v", hero)
}

func TestGetAllItemTooltipParams(t *testing.T) {
	var (
		battleTag            = "DyS#1311"
		heroID               = 104452451
		hero                 *Hero
		allItemTooltipParams map[string]string
		err                  error
	)

	hero, err = d.GetHero(battleTag, heroID)
	if err != nil {
		t.Fail()
	}

	allItemTooltipParams = hero.GetAllItemTooltipParams()

	fmt.Println(allItemTooltipParams)
}

func TestGetAllItems(t *testing.T) {
	var (
		battleTag = "DyS#1311"
		heroID    = 104452451
		hero      *Hero
		allItems  map[string]*Item
		err       error
	)

	hero, err = d.GetHero(battleTag, heroID)
	if err != nil {
		t.Fail()
	}

	allItems = d.GetAllItems(hero)

	fmt.Println(allItems["LeftFinger"])
}

func TestGetItem(t *testing.T) {
	var (
		tooltipParams = "/item/krysbins-sentence-P6_Unique_Ring_03"
		itemJSON      []byte
		item          *Item
		err           error
	)

	itemJSON, err = d.GetItemJSON(tooltipParams)
	if err != nil {
		t.Fail()
	}

	item, err = d.GetItem(tooltipParams)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(itemJSON))
	fmt.Printf("%+v", item)
}

func TestGetItemAttributesRaw(t *testing.T) {
	var (
		tooltipParams     = "/item/krysbins-sentence-P6_Unique_Ring_03"
		item              *Item
		itemAttributesRaw map[string]float64
		err               error
	)

	item, err = d.GetItem(tooltipParams)
	if err != nil {
		t.Fail()
		return
	}

	itemAttributesRaw = GetItemAttributesRaw(item)

	fmt.Println(itemAttributesRaw)
}

func TestCompareItemsAttributesRaw(t *testing.T) {
	var (
		tooltipParams1            = "/item/krysbins-sentence-P6_Unique_Ring_03"
		item1                     *Item
		tooltipParams2            = "/item/krysbins-sentence-P6_Unique_Ring_03"
		item2                     *Item
		compareItemsAttributesRaw map[string]float64
		err                       error
	)

	item1, err = d.GetItem(tooltipParams1)
	if err != nil {
		t.Fail()
		return
	}

	item2, err = d.GetItem(tooltipParams2)
	if err != nil {
		t.Fail()
		return
	}

	compareItemsAttributesRaw = CompareItemsAttributesRaw(item1, item2)

	fmt.Println(compareItemsAttributesRaw)
}

func TestGetFollower(t *testing.T) {
	var (
		followerJSON []byte
		follower     *Follower
		err          error
	)

	followerJSON, err = d.GetTemplarJSON()
	if err != nil {
		t.Fail()
	}

	follower, err = d.GetTemplar()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(followerJSON))
	fmt.Printf("%+v", follower)
}

func TestGetArtisan(t *testing.T) {
	var (
		artisanJSON []byte
		artisan     *Artisan
		err         error
	)

	artisanJSON, err = d.GetMysticJSON()
	if err != nil {
		t.Fail()
	}

	artisan, err = d.GetMystic()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(artisanJSON))
	fmt.Printf("%+v", artisan)
}
