/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:22
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 17:21:15
 */

// Package diablo3 is a client library to use Blizzard Diablo 3 API calls.
package diablo3

import (
	"errors"
	"strconv"
	"strings"

	"github.com/FuzzyStatic/blizzard"
)

// Diablo3 regional API URLs, locale, access token, api key
type Diablo3 struct {
	Auth         blizzard.Auth
	Locale       string
	DataURL      string
	CommunityURL string
}

// New create new Diablo3 structure
func New(auth blizzard.Auth, region blizzard.Region) *Diablo3 {
	var d = Diablo3{Auth: auth}

	switch region {
	case blizzard.CN:
		d.Locale = "zh_CN"
		d.DataURL = "https://api.battle.net.cn/data/d3"
		d.CommunityURL = "https://api.battle.net.cn/d3"
	case blizzard.EU:
		d.Locale = "en_GB"
		d.DataURL = "https://eu.api.battle.net/data/d3"
		d.CommunityURL = "https://eu.api.battle.net/d3"
	case blizzard.KR:
		d.Locale = "ko_KR"
		d.DataURL = "https://kr.api.battle.net/data/d3"
		d.CommunityURL = "https://kr.api.battle.net/d3"
	case blizzard.TW:
		d.Locale = "zh_TW"
		d.DataURL = "https://tw.api.battle.net/data/d3"
		d.CommunityURL = "https://tw.api.battle.net/d3"
	case blizzard.US:
		d.Locale = "en_US"
		d.DataURL = "https://us.api.battle.net/data/d3"
		d.CommunityURL = "https://us.api.battle.net/d3"
	default: // USA! USA!
		d.Locale = "en_US"
		d.DataURL = "https://us.api.battle.net/data/d3"
		d.CommunityURL = "https://us.api.battle.net/d3"
	}

	return &d
}

func (d *Diablo3) getArtisanJSON(artisanNamePath string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.CommunityURL + dataPath + artisanPath + artisanNamePath + "?" + localeQuery +
		d.Locale + "&" + apiKeyQuery + d.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

func (d *Diablo3) getArtisan(artisanNamePath string) (*Artisan, error) {
	var (
		artisan Artisan
		json    *[]byte
		err     error
	)

	json, err = d.getArtisanJSON(artisanNamePath)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &artisan)
	if err != nil {
		return nil, err
	}

	return &artisan, nil
}

// GetBlacksmithJSON gets JSON for Blacksmith artisan
func (d *Diablo3) GetBlacksmithJSON() (*[]byte, error) {
	return d.getArtisanJSON(blacksmithPath)
}

// GetJewelerJSON gets JSON for Jeweler artisan
func (d *Diablo3) GetJewelerJSON() (*[]byte, error) {
	return d.getArtisanJSON(jewelerPath)
}

// GetMysticJSON gets JSON for Mystic artisan
func (d *Diablo3) GetMysticJSON() (*[]byte, error) {
	return d.getArtisanJSON(mysticPath)
}

// GetBlacksmith puts information into Artisan structure Blacksmith
func (d *Diablo3) GetBlacksmith() (*Artisan, error) {
	return d.getArtisan(blacksmithPath)
}

// GetJeweler puts information into Artisan structure Jeweler
func (d *Diablo3) GetJeweler() (*Artisan, error) {
	return d.getArtisan(jewelerPath)
}

// GetMystic puts information into Artisan structure Mystic
func (d *Diablo3) GetMystic() (*Artisan, error) {
	return d.getArtisan(mysticPath)
}

// GetEraJSON gets specified eraID JSON information
func (d *Diablo3) GetEraJSON(eraID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + eraIndexPath + "/" + strconv.Itoa(eraID) + "?" + accessTokenQuery +
		d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetEra puts specified era information into Era structure
func (d *Diablo3) GetEra(eraID int) (*Era, error) {
	var (
		era  Era
		json *[]byte
		err  error
	)

	json, err = d.GetEraJSON(eraID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &era)
	if err != nil {
		return nil, err
	}

	return &era, nil
}

// GetEraIndexJSON gets season index JSON information
func (d *Diablo3) GetEraIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + eraIndexPath + "/?" + accessTokenQuery + d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetEraIndex puts era index information into EraIndex structure
func (d *Diablo3) GetEraIndex() (*EraIndex, error) {
	var (
		eraIndex EraIndex
		json     *[]byte
		err      error
	)

	json, err = d.GetEraIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &eraIndex)
	if err != nil {
		return nil, err
	}

	return &eraIndex, nil
}

// GetEraLeaderboardJSON gets specified season information JSON
func (d *Diablo3) GetEraLeaderboardJSON(eraID int, groupPath string, hardcore bool) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + eraIndexPath + "/" + strconv.Itoa(eraID) + leaderboardPath +
		riftPath

	if hardcore {
		url = url + hardcorePath
	}

	url = url + groupPath + "?access_token=" + d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetEraLeaderboard puts specified era leaderboard information into EraLeaderboard structure
func (d *Diablo3) GetEraLeaderboard(eraID int, groupPath string, hardcore bool) (*EraLeaderboard, error) {
	var (
		eraLeaderboard EraLeaderboard
		json           *[]byte
		err            error
	)

	json, err = d.GetEraLeaderboardJSON(eraID, groupPath, hardcore)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &eraLeaderboard)
	if err != nil {
		return nil, err
	}

	return &eraLeaderboard, nil
}

func (d *Diablo3) getFollowerJSON(followerNamePath string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.CommunityURL + dataPath + followerPath + followerNamePath + "?" + localeQuery +
		d.Locale + "&" + apiKeyQuery + d.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

func (d *Diablo3) getFollower(followerNamePath string) (*Follower, error) {
	var (
		follower Follower
		json     *[]byte
		err      error
	)

	json, err = d.getFollowerJSON(followerNamePath)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &follower)
	if err != nil {
		return nil, err
	}

	return &follower, nil
}

// GetEnchantressJSON gets enchantress follower information JSON
func (d *Diablo3) GetEnchantressJSON() (*[]byte, error) {
	return d.getFollowerJSON(enchantressPath)
}

// GetScoundrelJSON gets scoundrel follower information JSON
func (d *Diablo3) GetScoundrelJSON() (*[]byte, error) {
	return d.getFollowerJSON(scoundrelPath)
}

// GetTemplarJSON gets templar follower information JSON
func (d *Diablo3) GetTemplarJSON() (*[]byte, error) {
	return d.getFollowerJSON(templarPath)
}

// GetEnchantress puts enchantress information into Follower structure
func (d *Diablo3) GetEnchantress() (*Follower, error) {
	return d.getFollower(enchantressPath)
}

// GetScoundrel puts scoundrel follower information into Follower structure
func (d *Diablo3) GetScoundrel() (*Follower, error) {
	return d.getFollower(scoundrelPath)
}

// GetTemplar puts templar follower information into Follower structure
func (d *Diablo3) GetTemplar() (*Follower, error) {
	return d.getFollower(templarPath)
}

// GetHeroJSON gets specified battleTag hero information JSON
// Formats accepted: Player-1234 or Player#1234
func (d *Diablo3) GetHeroJSON(battleTag string, heroID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	battleTag = strings.Replace(battleTag, "#", "-", -1)

	url = d.CommunityURL + profilePath + "/" + battleTag + heroPath + "/" + strconv.Itoa(heroID) +
		"?" + localeQuery + d.Locale + "&" + apiKeyQuery + d.Auth.APIKey

	if err = blizzard.GetURLBody(url, &json); err != nil {
		return nil, err
	}

	return &json, nil
}

// GetHero puts specified follower information into Hero structure
func (d *Diablo3) GetHero(battleTag string, heroID int) (*Hero, error) {
	var (
		hero Hero
		json *[]byte
		err  error
	)

	json, err = d.GetHeroJSON(battleTag, heroID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &hero)
	if err != nil {
		return nil, err
	}

	if hero.Code == "NOTFOUND" {
		return nil, errors.New(hero.Reason)
	}

	return &hero, nil
}

// GetItemJSON gets specified item information JSON
// Expected format: item/<string>
func (d *Diablo3) GetItemJSON(tooltipParams string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	if tooltipParams != "" {
		url = d.CommunityURL + dataPath + "/" + tooltipParams + "?" + localeQuery + d.Locale +
			"&" + apiKeyQuery + d.Auth.APIKey

		err = blizzard.GetURLBody(url, &json)
		if err != nil {
			return nil, err
		}
	}

	return &json, nil
}

// GetItem puts specified follower information into Item structure
func (d *Diablo3) GetItem(tooltipParams string) (*Item, error) {
	var (
		item Item
		json *[]byte
		err  error
	)

	json, err = d.GetItemJSON(tooltipParams)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetAllItems returns a map with all the item structures from a hero
func (d *Diablo3) GetAllItems(hero *Hero) map[string]*Item {
	m := make(map[string]*Item)
	m["Bracers"], _ = d.GetItem(hero.Items.Bracers.TooltipParams)
	m["Feet"], _ = d.GetItem(hero.Items.Feet.TooltipParams)
	m["Hands"], _ = d.GetItem(hero.Items.Hands.TooltipParams)
	m["Head"], _ = d.GetItem(hero.Items.Head.TooltipParams)
	m["LeftFinger"], _ = d.GetItem(hero.Items.LeftFinger.TooltipParams)
	m["Legs"], _ = d.GetItem(hero.Items.Legs.TooltipParams)
	m["MainHand"], _ = d.GetItem(hero.Items.MainHand.TooltipParams)
	m["OffHand"], _ = d.GetItem(hero.Items.OffHand.TooltipParams)
	m["RightFinger"], _ = d.GetItem(hero.Items.RightFinger.TooltipParams)
	m["Shoulders"], _ = d.GetItem(hero.Items.Shoulders.TooltipParams)
	m["Torso"], _ = d.GetItem(hero.Items.Torso.TooltipParams)
	return m
}

// CompareAllItemsBetweenHeroes returns the difference values of all attributes for
// item1 compared to item2 for all items between hero1 and hero2
func (d *Diablo3) CompareAllItemsBetweenHeroes(hero1, hero2 *Hero) map[string]map[string]float64 {
	m := make(map[string]map[string]float64)

	mItems1 := hero1.GetAllItemTooltipParams()
	mItems2 := hero2.GetAllItemTooltipParams()

	for i, v := range mItems1 {
		for j, w := range mItems2 {
			if i == j {
				item1, _ := d.GetItem(v)
				item2, _ := d.GetItem(w)
				m[i] = CompareItemsAttributesRaw(*item1, *item2)
				delete(mItems2, j)
			}
		}
	}

	return m
}

// GetProfileJSON gets specified battleTag profile JSON information
// Formats accepted: Player-1234 or Player#1234
func (d *Diablo3) GetProfileJSON(battleTag string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	battleTag = strings.Replace(battleTag, "#", "-", -1)

	url = d.CommunityURL + profilePath + "/" + battleTag + "/?" + localeQuery + d.Locale +
		"&" + apiKeyQuery + d.Auth.APIKey

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetProfile puts specified battleTag information in Profile structure
// Formats accepted: Player-1234 or Player#1234
func (d *Diablo3) GetProfile(battleTag string) (*Profile, error) {
	var (
		profile Profile
		json    *[]byte
		err     error
	)

	json, err = d.GetProfileJSON(battleTag)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// GetSeasonJSON gets specified season information JSON
func (d *Diablo3) GetSeasonJSON(seasonID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + seasonIndexPath + "/" + strconv.Itoa(seasonID) + "?" + accessTokenQuery +
		d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetSeason puts specified season information into Season structure
func (d *Diablo3) GetSeason(seasonID int) (*Season, error) {
	var (
		season Season
		json   *[]byte
		err    error
	)

	json, err = d.GetSeasonJSON(seasonID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &season)
	if err != nil {
		return nil, err
	}

	return &season, nil
}

// GetSeasonIndexJSON gets season index information JSON
func (d *Diablo3) GetSeasonIndexJSON() (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + seasonIndexPath + "/?" + accessTokenQuery + d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetSeasonIndex puts season index information into SeasonIndex structure
func (d *Diablo3) GetSeasonIndex() (*SeasonIndex, error) {
	var (
		seasonIndex SeasonIndex
		json        *[]byte
		err         error
	)

	json, err = d.GetSeasonIndexJSON()
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &seasonIndex)
	if err != nil {
		return nil, err
	}

	return &seasonIndex, nil
}

// GetSeasonAchievementPointsJSON puts specified season information JSON
func (d *Diablo3) GetSeasonAchievementPointsJSON(seasonID int) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + seasonIndexPath + "/" + strconv.Itoa(seasonID) + leaderboardPath +
		achievementPointsPath + "?access_token=" + d.Auth.AccessToken

	err = blizzard.GetURLBody(url, &json)
	if err != nil {
		return nil, err
	}

	return &json, nil
}

// GetSeasonAchievementPoints puts specified season information into SeasonAchievementPoints structure
func (d *Diablo3) GetSeasonAchievementPoints(seasonID int) (*SeasonAchievementPoints, error) {
	var (
		seasonAchievementPoints SeasonAchievementPoints
		json                    *[]byte
		err                     error
	)

	json, err = d.GetSeasonAchievementPointsJSON(seasonID)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &seasonAchievementPoints)
	if err != nil {
		return nil, err
	}

	return &seasonAchievementPoints, nil
}

func (d *Diablo3) getRiftJSON(seasonID int, hardcore bool, groupPath string) (*[]byte, error) {
	var (
		url  string
		json []byte
		err  error
	)

	url = d.DataURL + seasonIndexPath + "/" + strconv.Itoa(seasonID) +
		leaderboardPath + riftPath

	if hardcore {
		url = url + hardcorePath
	}

	url = url + groupPath + "?access_token=" + d.Auth.AccessToken

	if err = blizzard.GetURLBody(url, &json); err != nil {
		return nil, err
	}

	return &json, nil
}

func (d *Diablo3) getRift(seasonID int, hardcore bool, groupPath string) (*SeasonRift, error) {
	var (
		seasonRift SeasonRift
		json       *[]byte
		err        error
	)

	json, err = d.getRiftJSON(seasonID, hardcore, groupPath)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &seasonRift)
	if err != nil {
		return nil, err
	}

	return &seasonRift, nil
}

// GetBarbarianRiftJSON gets specified season information JSON for Barbarian
func (d *Diablo3) GetBarbarianRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, barbarianPath)
}

// GetCrusaderRiftJSON gets specified season information JSON for Crusader
func (d *Diablo3) GetCrusaderRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, crusaderPath)
}

// GetDemonHunterRiftJSON gets specified season information JSON for Demon Hunter
func (d *Diablo3) GetDemonHunterRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, demonHunterPath)
}

// GetNecromancerRiftJSON gets specified season information JSON for Necromancer
func (d *Diablo3) GetNecromancerRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, necromancerPath)
}

// GetWizardRiftJSON gets specified season information JSON for Wizard
func (d *Diablo3) GetWizardRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, wizardPath)

}

// GetWitchDoctorRiftJSON gets specified season information JSON for Witch Doctor
func (d *Diablo3) GetWitchDoctorRiftJSON(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, witchDoctorPath)
}

// GetBarbarianRift gets specified season information JSON for Barbarian
func (d *Diablo3) GetBarbarianRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, barbarianPath)
}

// GetCrusaderRift gets specified season information JSON for Crusader
func (d *Diablo3) GetCrusaderRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, crusaderPath)
}

// GetDemonHunterRift gets specified season information JSON for Demon Hunter
func (d *Diablo3) GetDemonHunterRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, demonHunterPath)
}

// GetNecromancerRift gets specified season information JSON for Necromancer
func (d *Diablo3) GetNecromancerRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, necromancerPath)
}

// GetWizardRift gets specified season information JSON for Wizard
func (d *Diablo3) GetWizardRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, wizardPath)

}

// GetWitchDoctorRift gets specified season information JSON for Witch Doctor
func (d *Diablo3) GetWitchDoctorRift(seasonID int, hardcore bool) (*[]byte, error) {
	return d.getRiftJSON(seasonID, hardcore, witchDoctorPath)
}

func (d *Diablo3) getCurrentSeasonRiftJSON(hardcore bool, groupPath string) (*[]byte, error) {
	var (
		seasonIndex *SeasonIndex
		url         string
		json        []byte
		err         error
	)

	seasonIndex, err = d.GetSeasonIndex()
	if err != nil {
		return nil, err
	}

	url = d.DataURL + seasonIndexPath + "/" + strconv.Itoa(seasonIndex.CurrentSeason) +
		leaderboardPath + riftPath

	if hardcore {
		url = url + hardcorePath
	}

	url = url + groupPath + "?access_token=" + d.Auth.AccessToken

	if err = blizzard.GetURLBody(url, &json); err != nil {
		return nil, err
	}

	return &json, nil
}

func (d *Diablo3) getCurrentSeasonRift(hardcore bool, groupPath string) (*SeasonRift, error) {
	var (
		seasonRift SeasonRift
		json       *[]byte
		err        error
	)

	json, err = d.getCurrentSeasonRiftJSON(hardcore, groupPath)
	if err != nil {
		return nil, err
	}

	err = blizzard.GetStruct(json, &seasonRift)
	if err != nil {
		return nil, err
	}

	return &seasonRift, nil
}

// GetCurrentSeasonBarbarianRiftJSON gets current season information JSON for Barbarian
func (d *Diablo3) GetCurrentSeasonBarbarianRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, barbarianPath)
}

// GetCurrentSeasonCrusaderRiftJSON gets current season information JSON for Crusader
func (d *Diablo3) GetCurrentSeasonCrusaderRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, crusaderPath)
}

// GetCurrentSeasonDemonHunterRiftJSON gets current season information JSON for Demon Hunter
func (d *Diablo3) GetCurrentSeasonDemonHunterRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, demonHunterPath)
}

// GetCurrentSeasonNecromancerRiftJSON gets current season information JSON for Necromancer
func (d *Diablo3) GetCurrentSeasonNecromancerRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, necromancerPath)
}

// GetCurrentSeasonWizardRiftJSON gets current season information JSON for Wizard
func (d *Diablo3) GetCurrentSeasonWizardRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, wizardPath)
}

// GetCurrentSeasonWitchDoctorRiftJSON gets current season information JSON for Witch Doctor
func (d *Diablo3) GetCurrentSeasonWitchDoctorRiftJSON(hardcore bool) (*[]byte, error) {
	return d.getCurrentSeasonRiftJSON(hardcore, wizardPath)
}

// GetCurrentSeasonBarbarianRift puts current season information into SeasonRift structure for Barbarian
func (d *Diablo3) GetCurrentSeasonBarbarianRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, barbarianPath)
}

// GetCurrentSeasonCrusaderRift puts current season information into SeasonRift structure for Crusader
func (d *Diablo3) GetCurrentSeasonCrusaderRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, crusaderPath)
}

// GetCurrentSeasonDemonHunterRift puts current season information into SeasonRift structure Demon Hunter
func (d *Diablo3) GetCurrentSeasonDemonHunterRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, demonHunterPath)
}

// GetCurrentSeasonNecromancerRift puts current season information into SeasonRift structure Necromancer
func (d *Diablo3) GetCurrentSeasonNecromancerRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, necromancerPath)
}

// GetCurrentSeasonWizardRift puts current season information into SeasonRift structure Wizard
func (d *Diablo3) GetCurrentSeasonWizardRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, wizardPath)
}

// GetCurrentSeasonWitchDoctorRift puts current season information into SeasonRift structure Witch Doctor
func (d *Diablo3) GetCurrentSeasonWitchDoctorRift(hardcore bool) (*SeasonRift, error) {
	return d.getCurrentSeasonRift(hardcore, witchDoctorPath)
}
