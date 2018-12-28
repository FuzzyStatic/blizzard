package blizzard

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/FuzzyStatic/blizzard/wowc"
)

const (
	wowPath                = "/wow"
	wowUserPath            = wowPath + "/user"
	wowUserCharactersPath  = wowUserPath + "/characters"
	wowAchievementPath     = wowPath + "/achievement"
	wowAuctionDataPath     = wowPath + "/auction" + dataPath
	wowBossPath            = wowPath + "/boss"
	wowChallengePath       = wowPath + "/challenge"
	wowChallengeRegionPath = wowChallengePath + "/region"
	wowCharacterPath       = wowPath + "/character"
	wowGuildPath           = wowPath + "/guild"
)

// Fields for Character/Guild API calls
const (
	FieldCharacterAchievements = "achievements"
	FieldCharacterAppearance   = "appearance"
	FieldCharacterFeed         = "feed"
	FieldCharacterGuild        = "guild"
	FieldCharacterItems        = "items"
	FieldCharacterMounts       = "mounts"
	FieldCharacterPets         = "pets"
	FieldCharacterPetSlots     = "petSlots"
	FieldCharacterProfessions  = "professions"
	FieldCharacterProgression  = "progression"
	FieldCharacterPVP          = "pvp"
	FieldCharacterQuests       = "quests"
	FieldCharacterReputation   = "reputation"
	FieldCharacterStatistics   = "statistics"
	FieldCharacterStats        = "stats"
	FieldCharacterTalents      = "talents"
	FieldCharacterTitle        = "title"
	FieldCharacterAudit        = "audit"
	FieldGuildMembers          = "members"
	FieldGuildAchievements     = "achievements"
	FieldGuildNews             = "news"
	FieldGuildChallenge        = "challenge"
)

// WoWUserCharacters returns all characters for user's Access Token
func (c *Config) WoWUserCharacters(accessToken string) (*wowc.Profile, error) {
	var (
		dat wowc.Profile
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	err = c.updateAccessTokenIfExp()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", c.apiURL+wowUserCharactersPath+"?access_token="+accessToken, nil)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	res, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New(res.Status)
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWAchievement returns data about an individual achievement
func (c *Config) WoWAchievement(achievementID int) (*wowc.Achievement, error) {
	var (
		dat wowc.Achievement
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowAchievementPath + "/" + strconv.Itoa(achievementID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWAuctionFiles returns list of auction URLs containing auction data
func (c *Config) WoWAuctionFiles(realm string) (*wowc.AuctionFiles, error) {
	var (
		dat wowc.AuctionFiles
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowAuctionDataPath + "/" + realm + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWAuctionData returns auction data for realm
func (c *Config) WoWAuctionData(realm string) ([]*wowc.AuctionData, error) {
	var (
		af   *wowc.AuctionFiles
		adad []*wowc.AuctionData
		err  error
	)

	af, err = c.WoWAuctionFiles(realm)
	if err != nil {
		return nil, err
	}

	for _, file := range af.Files {
		var (
			dat wowc.AuctionData
			req *http.Request
			res *http.Response
			b   []byte
			err error
		)

		req, err = http.NewRequest("GET", file.URL, nil)
		if err != nil {
			return nil, err
		}

		res, err = c.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer func() {
			err = res.Body.Close()
			if err != nil {
				return
			}
		}()

		b, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		switch res.StatusCode {
		case http.StatusNotFound:
			return nil, errors.New(res.Status)
		}

		err = json.Unmarshal(b, &dat)
		if err != nil {
			return nil, err
		}

		adad = append(adad, &dat)
	}

	return adad, nil
}

// WoWBossMasterList returns a list of all supported bosses. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
func (c *Config) WoWBossMasterList() (*wowc.BossMasterList, error) {
	var (
		dat wowc.BossMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowBossPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWBoss  provides information about bosses. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
func (c *Config) WoWBoss(bossID int) (*wowc.Boss, error) {
	var (
		dat wowc.Boss
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowBossPath + "/" + strconv.Itoa(bossID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWChallengeModeRealmLeaderboard returns data for all nine challenge mode maps (currently). The map field includes the current medal times for each dungeon. Each ladder provides data about each character that was part of each run. The character data includes the current cached specialization of the character while the member field includes the specialization of the character during the challenge mode run
func (c *Config) WoWChallengeModeRealmLeaderboard(realm string) (*wowc.ChallengeModeRealmLeaderboard, error) {
	var (
		dat wowc.ChallengeModeRealmLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowChallengePath + "/" + realm + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWChallengeModeRegionLeaderboard has the exact same data format as the realm leaderboards except there is no realm field. Instead, the response has the top 100 results gathered for each map for all of the available realm leaderboards in a region
func (c *Config) WoWChallengeModeRegionLeaderboard() (*wowc.ChallengeModeRegionLeaderboard, error) {
	var (
		dat wowc.ChallengeModeRegionLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowChallengeRegionPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWCharacterProfile is the primary way to access character information. This API can be used to fetch a single character at a time through an HTTP GET request to a URL describing the character profile resource. By default, these requests return a basic dataset, and each request can return zero or more additional fields. To access this API, craft a resource URL pointing to the desired character for which to retrieve information
// Optional field constants are prefixed with the word "FieldCharacter"
func (c *Config) WoWCharacterProfile(realm, characterName string, optionalFields []string) (*wowc.CharacterProfile, error) {
	var (
		dat      wowc.CharacterProfile
		fieldStr string
		b        []byte
		err      error
	)

	if optionalFields != nil {
		fieldStr = "fields="

		for i := 0; i < len(optionalFields); i++ {
			fieldStr = fieldStr + optionalFields[i]

			if i < len(optionalFields)-1 {
				fieldStr = fieldStr + ","
			}
		}

		fieldStr = fieldStr + "&"
	}

	b, err = c.getURLBody(c.apiURL + wowCharacterPath + "/" + realm + "/" + characterName + "?" + fieldStr + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWGuildProfile  is the primary way to access guild information. This API can fetch a single guild at a time through an HTTP GET request to a URL describing the guild profile resource. By default, these requests return a basic dataset and each request can retrieve zero or more additional fields.
// Although this endpoint has no required query string parameters, requests can optionally pass the fields query string parameter to indicate that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields"
// Optional field constants are prefixed with the word "FieldGuild"
func (c *Config) WoWGuildProfile(realm, guildName string, optionalFields []string) (*wowc.GuildProfile, error) {
	var (
		dat      wowc.GuildProfile
		fieldStr string
		b        []byte
		err      error
	)

	if optionalFields != nil {
		fieldStr = "fields="

		for i := 0; i < len(optionalFields); i++ {
			fieldStr = fieldStr + optionalFields[i]

			if i < len(optionalFields)-1 {
				fieldStr = fieldStr + ","
			}
		}

		fieldStr = fieldStr + "&"
	}

	b, err = c.getURLBody(c.apiURL + wowGuildPath + "/" + realm + "/" + guildName + "?" + fieldStr + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
