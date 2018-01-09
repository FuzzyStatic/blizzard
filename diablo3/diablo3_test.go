/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:33
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-08 21:53:47
 */

package diablo3

import (
	"fmt"
	"go-blizzard"
	"testing"

	"github.com/spf13/viper"
)

var d *Diablo3

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

	d = New(
		blizzard.Auth{
			AccessToken: accessToken,
			APIKey:      apiKey,
		},
		blizzard.US,
	)
}

func TestGetSeasonIndex(t *testing.T) {
	var (
		seasonIndexJSON *[]byte
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

	fmt.Println(string(*seasonIndexJSON))
	fmt.Println(*seasonIndex)
}

func TestGetSeason(t *testing.T) {
	var (
		seasonIndex *SeasonIndex
		seasonJSON  *[]byte
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

	fmt.Println(string(*seasonJSON))
	fmt.Println(*season)
}

func TestGetSeasonAchievementPoints(t *testing.T) {
	var (
		seasonIndex                 *SeasonIndex
		seasonAchievementPointsJSON *[]byte
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

	fmt.Println(string(*seasonAchievementPointsJSON))
	fmt.Println(*seasonAchievementPoints)
}

func TestGetSeasonRift(t *testing.T) {
	var (
		seasonRiftJSON *[]byte
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

	fmt.Println(string(*seasonRiftJSON))
	fmt.Println(*seasonRift)
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

	fmt.Println(*seasonRiftRowFromRank)
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

	fmt.Println(*seasonRowPlayer)
}

func TestGetHeroBattleTagFromSeasonRowPlayer(t *testing.T) {
	var (
		seasonRift                       *SeasonRift
		seasonRiftRowFromRank            *Row
		seasonRowPlayer                  *Player
		heroBattleTagFromSeasonRowPlayer *string
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

	fmt.Println(*heroBattleTagFromSeasonRowPlayer)
}

func TestGetHeroIDFromSeasonRowPlayer(t *testing.T) {
	var (
		seasonRift                *SeasonRift
		seasonRiftRowFromRank     *Row
		seasonRowPlayer           *Player
		heroIDFromSeasonRowPlayer *int
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

	fmt.Println(*heroIDFromSeasonRowPlayer)
}

func TestGetHeroBattleTagAndIDFromSeasonRow(t *testing.T) {
	var (
		seasonRift            *SeasonRift
		seasonRiftRowFromRank *Row
		battleTag             *string
		heroID                *int
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

	fmt.Println(*battleTag)
	fmt.Println(*heroID)
}

func TestGetProfile(t *testing.T) {
	var (
		battleTag   = "FuzzyStatic#1384"
		profileJSON *[]byte
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

	fmt.Println(string(*profileJSON))
	fmt.Println(*profile)
}

func TestGetHeroNamesFromProfile(t *testing.T) {
	var (
		battleTag            = "FuzzyStatic#1384"
		profile              *Profile
		heroNamesFromProfile *[]string
		err                  error
	)

	profile, err = d.GetProfile(battleTag)
	if err != nil {
		t.Fail()
	}

	heroNamesFromProfile = profile.GetHeroNamesFromProfile()

	fmt.Println(*heroNamesFromProfile)
}

func TestGetHeroIDsAndNamesFromProfile(t *testing.T) {
	var (
		battleTag                  = "FuzzyStatic#1384"
		profile                    *Profile
		heroIDsAndNamesFromProfile *map[int]string
		err                        error
	)

	profile, err = d.GetProfile(battleTag)
	if err != nil {
		t.Fail()
	}

	heroIDsAndNamesFromProfile = profile.GetHeroIDsAndNamesFromProfile()

	fmt.Println(*heroIDsAndNamesFromProfile)
}

func TestGetHero(t *testing.T) {
	var (
		battleTag = "DyS#1311"
		heroID    = 93409818
		heroJSON  *[]byte
		hero      *Hero
		err       error
	)

	heroJSON, err = d.GetHeroJSON(battleTag, heroID)
	if err != nil {
		t.Fail()
	}

	hero, err = d.GetHero(battleTag, heroID)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(*heroJSON))
	fmt.Println(*hero)
}

func TestGetAllItemTooltipParams(t *testing.T) {
	var (
		battleTag            = "DyS#1311"
		heroID               = 93409818
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
		heroID    = 93409818
		hero      *Hero
		allItems  map[string]*Item
		err       error
	)

	hero, err = d.GetHero(battleTag, heroID)
	if err != nil {
		t.Fail()
	}

	allItems = d.GetAllItems(hero)

	fmt.Println(*allItems["Bracers"])
}

func TestGetItem(t *testing.T) {
	var (
		tooltipParams = "item/CpkBCIOsrpkPEgcIBBX9OUmaHfRwB60dZiMGUB2cBgDLHWc-7OQdbhfHex0bXc9EMIvaAjiiAkAASAtQElgEYNsCaiwKDAgAEK_U4sSDgICgLxIcCMu4g5cJEgcIBBWkvhzCMIvSAjgAQAFYBJABC4ABRo0BVoz9E6UBZz7s5K0B5hXbDbUBf_lOXbgB6eGb9gfAAROAAgSIAsgBGIWYttsMUAxYAqABhZi22wygAYO-tM8DoAGj4cPOA6ABhaiHwAOgAY6k2PoJoAHg-Pm6Dw"
		itemJSON      *[]byte
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

	fmt.Println(string(*itemJSON))
	fmt.Println(*item)
}

func TestGetItemAttributesRaw(t *testing.T) {
	var (
		tooltipParams     = "item/CpkBCIOsrpkPEgcIBBX9OUmaHfRwB60dZiMGUB2cBgDLHWc-7OQdbhfHex0bXc9EMIvaAjiiAkAASAtQElgEYNsCaiwKDAgAEK_U4sSDgICgLxIcCMu4g5cJEgcIBBWkvhzCMIvSAjgAQAFYBJABC4ABRo0BVoz9E6UBZz7s5K0B5hXbDbUBf_lOXbgB6eGb9gfAAROAAgSIAsgBGIWYttsMUAxYAqABhZi22wygAYO-tM8DoAGj4cPOA6ABhaiHwAOgAY6k2PoJoAHg-Pm6Dw"
		item              *Item
		itemAttributesRaw map[string]float64
		err               error
	)

	item, err = d.GetItem(tooltipParams)
	if err != nil {
		t.Fail()
	}

	itemAttributesRaw = GetItemAttributesRaw(*item)

	fmt.Println(itemAttributesRaw)
}

func TestCompareItemsAttributesRaw(t *testing.T) {
	var (
		tooltipParams1            = "item/CpkBCIOsrpkPEgcIBBX9OUmaHfRwB60dZiMGUB2cBgDLHWc-7OQdbhfHex0bXc9EMIvaAjiiAkAASAtQElgEYNsCaiwKDAgAEK_U4sSDgICgLxIcCMu4g5cJEgcIBBWkvhzCMIvSAjgAQAFYBJABC4ABRo0BVoz9E6UBZz7s5K0B5hXbDbUBf_lOXbgB6eGb9gfAAROAAgSIAsgBGIWYttsMUAxYAqABhZi22wygAYO-tM8DoAGj4cPOA6ABhaiHwAOgAY6k2PoJoAHg-Pm6Dw"
		item1                     *Item
		tooltipParams2            = "item/CoUBCNaYpeALEgcIBBVWt3pQHXKOHSEdmwYAyx1mIwZQHY9Y1h4d83AHrTCL1gI4qANAAFASWARgqgNqLAoMCAAQq57e9oOAgOAHEhwImsXD2gsSBwgEFYm-HMIwi9ICOABAAVgEkAELgAFGpQFmIwZQrQHQ7mgxtQGb2dnvuAHUxsPHDcABFRis3dWHCg"
		item2                     *Item
		compareItemsAttributesRaw map[string]float64
		err                       error
	)

	item1, err = d.GetItem(tooltipParams1)
	if err != nil {
		t.Fail()
	}

	item2, err = d.GetItem(tooltipParams2)
	if err != nil {
		t.Fail()
	}

	compareItemsAttributesRaw = CompareItemsAttributesRaw(*item1, *item2)

	fmt.Println(compareItemsAttributesRaw)
}

func TestGetFollower(t *testing.T) {
	var (
		followerJSON *[]byte
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

	fmt.Println(string(*followerJSON))
	fmt.Println(*follower)
}

func TestGetArtisan(t *testing.T) {
	var (
		artisanJSON *[]byte
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

	fmt.Println(string(*artisanJSON))
	fmt.Println(*artisan)
}
