/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-17 19:43:46
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 17:03:17
 */

package starcraft2

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard"
	"github.com/spf13/viper"
)

var s *Starcraft2

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

	s = New(
		blizzard.New(
			blizzard.Auth{
				AccessToken: accessToken,
				APIKey:      apiKey,
			},
			blizzard.US),
	)
}

func TestGetProfile(t *testing.T) {
	var (
		profileJSON []byte
		profile     *Profile
		err         error
	)

	profileJSON, err = s.GetProfileJSON(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	profile, err = s.GetProfile(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(profileJSON))
	fmt.Println(profile)
}

func TestGetProfileLadders(t *testing.T) {
	var (
		profileLaddersJSON []byte
		profileLadders     *ProfileLadders
		err                error
	)

	profileLaddersJSON, err = s.GetProfileLaddersJSON(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	profileLadders, err = s.GetProfileLadders(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(profileLaddersJSON))
	fmt.Println(profileLadders)
}

func TestGetProfileMatches(t *testing.T) {
	var (
		profileMatchesJSON []byte
		profileMatches     *ProfileMatches
		err                error
	)

	profileMatchesJSON, err = s.GetProfileMatchesJSON(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	profileMatches, err = s.GetProfileMatches(2537456, 1, "Neeb")
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(profileMatchesJSON))
	fmt.Println(profileMatches)
}

func TestGetLadder(t *testing.T) {
	var (
		ladderJSON []byte
		ladder     *Ladder
		err        error
	)

	ladderJSON, err = s.GetLadderJSON(263049)
	if err != nil {
		t.Fail()
	}

	ladder, err = s.GetLadder(263049)
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(ladderJSON))
	fmt.Println(ladder)
}

func TestGetAchievements(t *testing.T) {
	var (
		achievementsJSON []byte
		achievements     *Achievements
		err              error
	)

	achievementsJSON, err = s.GetAchievementsJSON()
	if err != nil {
		t.Fail()
	}

	achievements, err = s.GetAchievements()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(achievementsJSON))
	fmt.Println(achievements)
}

func TestGetRewards(t *testing.T) {
	var (
		rewardsJSON []byte
		rewards     *Rewards
		err         error
	)

	rewardsJSON, err = s.GetRewardsJSON()
	if err != nil {
		t.Fail()
	}

	rewards, err = s.GetRewards()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(rewardsJSON))
	fmt.Println(rewards)
}

func TestGetProfileUser(t *testing.T) {
	var (
		profileUserJSON []byte
		profileUser     *ProfileUser
		err             error
	)

	profileUserJSON, err = s.GetProfileUserJSON()
	if err != nil {
		t.Fail()
	}

	profileUser, err = s.GetProfileUser()
	if err != nil {
		t.Fail()
	}

	fmt.Println(string(profileUserJSON))
	fmt.Println(profileUser)
}
