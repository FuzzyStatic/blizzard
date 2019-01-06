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
	wowPath                          = "/wow"
	wowUserPath                      = wowPath + "/user"
	wowUserCharactersPath            = wowUserPath + "/characters"
	wowAchievementPath               = wowPath + "/achievement"
	wowAuctionDataPath               = wowPath + "/auction" + dataPath
	wowBossPath                      = wowPath + "/boss"
	wowChallengePath                 = wowPath + "/challenge"
	wowChallengeRegionPath           = wowChallengePath + "/region"
	wowCharacterPath                 = wowPath + "/character"
	wowGuildPath                     = wowPath + "/guild"
	wowItemPath                      = wowPath + "/item"
	wowItemSetPath                   = wowItemPath + "/set"
	wowMountPath                     = wowPath + "/mount"
	wowPetPath                       = wowPath + "/pet"
	wowPetAbilityPath                = wowPetPath + "/ability"
	wowPetSpeciesPath                = wowPetPath + "/species"
	wowPetStatsPath                  = wowPetPath + "/stats"
	wowPetsLevelField                = "level="
	wowPetsBreedIDField              = "breedId="
	wowPetsQualityIDField            = "qualityId="
	wowPVPLeaderboardPath            = wowPath + "/leaderboard"
	wowQuestPath                     = wowPath + "/quest"
	wowRealmStatusPath               = wowPath + "/realm/status"
	wowRecipePath                    = wowPath + "/recipe"
	wowSpellPath                     = wowPath + "/spell"
	wowZonePath                      = wowPath + "/zone"
	wowDataPath                      = wowPath + "/data"
	wowDataBattlegroupsPath          = wowDataPath + "/battlegroups"
	wowDataCharacterPath             = wowDataPath + "/character"
	wowDataCharacterRacesPath        = wowDataCharacterPath + "/races"
	wowDataCharacterClassesPath      = wowDataCharacterPath + "/classes"
	wowDataCharacterAchievementsPath = wowDataCharacterPath + "/achievements"
	wowDataGuildPath                 = wowDataPath + "/guild"
	wowDataGuildRewardsPath          = wowDataGuildPath + "/rewards"
	wowDataGuildPerksPath            = wowDataGuildPath + "/perks"
	wowDataGuildAchievementsPath     = wowDataGuildPath + "/achievements"
	wowDataItemClassesPath           = wowDataPath + "/item/classes"
	wowDataTalentsPath               = wowDataPath + "/talents"
	wowDataPetTypesPath              = wowDataPath + "/pet/types"
)

// WoWUserCharacters returns all characters for user's Access Token
func (c *Client) WoWUserCharacters(accessToken string) (*wowc.Profile, []byte, error) {
	var (
		dat wowc.Profile
		req *http.Request
		res *http.Response
		b   []byte
		err error
	)

	err = c.updateAccessTokenIfExp()
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("GET", c.apiURL+wowUserCharactersPath+"?access_token="+accessToken, nil)
	if err != nil {
		return nil, nil, err
	}

	//req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")

	res, err = c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		err = res.Body.Close()
		if err != nil {
			return
		}
	}()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, nil, errors.New(res.Status)
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWAchievement returns data about an individual achievement
func (c *Client) WoWAchievement(achievementID int) (*wowc.Achievement, []byte, error) {
	var (
		dat wowc.Achievement
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowAchievementPath+"/"+strconv.Itoa(achievementID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWAuctionFiles returns list of auction URLs containing auction data
func (c *Client) WoWAuctionFiles(realm string) (*wowc.AuctionFiles, []byte, error) {
	var (
		dat wowc.AuctionFiles
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowAuctionDataPath+"/"+realm+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWAuctionData returns auction data for realm
func (c *Client) WoWAuctionData(realm string) ([]*wowc.AuctionData, error) {
	var (
		af   *wowc.AuctionFiles
		adad []*wowc.AuctionData
		err  error
	)

	af, _, err = c.WoWAuctionFiles(realm)
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
func (c *Client) WoWBossMasterList() (*wowc.BossMasterList, []byte, error) {
	var (
		dat wowc.BossMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowBossPath+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWBoss provides information about bosses. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
func (c *Client) WoWBoss(bossID int) (*wowc.Boss, []byte, error) {
	var (
		dat wowc.Boss
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowBossPath+"/"+strconv.Itoa(bossID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWChallengeModeRealmLeaderboard returns data for all nine challenge mode maps (currently). The map field includes the current medal times for each dungeon. Each ladder provides data about each character that was part of each run. The character data includes the current cached specialization of the character while the member field includes the specialization of the character during the challenge mode run
func (c *Client) WoWChallengeModeRealmLeaderboard(realm string) (*wowc.ChallengeModeRealmLeaderboard, []byte, error) {
	var (
		dat wowc.ChallengeModeRealmLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowChallengePath+"/"+realm+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWChallengeModeRegionLeaderboard has the exact same data format as the realm leaderboards except there is no realm field. Instead, the response has the top 100 results gathered for each map for all of the available realm leaderboards in a region
func (c *Client) WoWChallengeModeRegionLeaderboard() (*wowc.ChallengeModeRegionLeaderboard, []byte, error) {
	var (
		dat wowc.ChallengeModeRegionLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowChallengeRegionPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWCharacterProfile is the primary way to access character information. This API can be used to fetch a single character at a time through an HTTP GET request to a URL describing the character profile resource. By default, these requests return a basic dataset, and each request can return zero or more additional fields. To access this API, craft a resource URL pointing to the desired character for which to retrieve information
// Optional field constants are prefixed with the word "FieldCharacter"
func (c *Client) WoWCharacterProfile(realm, characterName string, optionalFields []string) (*wowc.CharacterProfile, []byte, error) {
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

	b, err = c.getURLBody(c.apiURL+wowCharacterPath+"/"+realm+"/"+characterName+"?"+fieldStr+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWGuildProfile  is the primary way to access guild information. This API can fetch a single guild at a time through an HTTP GET request to a URL describing the guild profile resource. By default, these requests return a basic dataset and each request can retrieve zero or more additional fields.
// Although this endpoint has no required query string parameters, requests can optionally pass the fields query string parameter to indicate that one or more of the optional datasets is to be retrieved. Those additional fields are listed in the method titled "Optional Fields"
// Optional field constants are prefixed with the word "FieldGuild"
func (c *Client) WoWGuildProfile(realm, guildName string, optionalFields []string) (*wowc.GuildProfile, []byte, error) {
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

	b, err = c.getURLBody(c.apiURL+wowGuildPath+"/"+realm+"/"+guildName+"?"+fieldStr+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(string(b))

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWItem provides detailed item information, including item set information
func (c *Client) WoWItem(itemID int) (*wowc.Item, []byte, error) {
	var (
		dat wowc.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowItemPath+"/"+strconv.Itoa(itemID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWItemSet provides detailed item information, including item set information
func (c *Client) WoWItemSet(setID int) (*wowc.ItemSet, []byte, error) {
	var (
		dat wowc.ItemSet
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowItemSetPath+"/"+strconv.Itoa(setID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWMountMasterList returns a list of all supported mounts
func (c *Client) WoWMountMasterList() (*wowc.MountMasterList, []byte, error) {
	var (
		dat wowc.MountMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowMountPath+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPetMasterList returns a list of all supported battle and vanity pets
func (c *Client) WoWPetMasterList() (*wowc.PetMasterList, []byte, error) {
	var (
		dat wowc.PetMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPetPath+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPetAbility returns data about a individual battle pet ability ID. This resource does not provide ability tooltips
func (c *Client) WoWPetAbility(abilityID int) (*wowc.PetAbility, []byte, error) {
	var (
		dat wowc.PetAbility
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPetAbilityPath+"/"+strconv.Itoa(abilityID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPetSpecies returns data about an individual pet species. Use pets as the field value in a character profile request to get species IDs. Each species also has data about its six abilities
func (c *Client) WoWPetSpecies(speciesID int) (*wowc.PetSpecies, []byte, error) {
	var (
		dat wowc.PetSpecies
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPetSpeciesPath+"/"+strconv.Itoa(speciesID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPetStats returns detailed information about a given species of pet
func (c *Client) WoWPetStats(speciesID, level, breedID, qualityID int) (*wowc.PetStats, []byte, error) {
	var (
		dat wowc.PetStats
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPetStatsPath+"/"+strconv.Itoa(speciesID)+"?"+wowPetsLevelField+strconv.Itoa(level)+"&"+wowPetsBreedIDField+strconv.Itoa(breedID)+"&"+wowPetsQualityIDField+strconv.Itoa(qualityID)+"&"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPVPLeaderboard provides leaderboard information for the 2v2, 3v3, 5v5, and Rated Battleground leaderboards
func (c *Client) WoWPVPLeaderboard(bracket string) (*wowc.PVPLeaderboard, []byte, error) {
	var (
		dat wowc.PVPLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowPVPLeaderboardPath+"/"+bracket+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWQuest returns metadata for a specified quest
func (c *Client) WoWQuest(questID int) (*wowc.Quest, []byte, error) {
	var (
		dat wowc.Quest
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowQuestPath+"/"+strconv.Itoa(questID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRealmStatus returns metadata for a specified quest
func (c *Client) WoWRealmStatus() (*wowc.RealmStatus, []byte, error) {
	var (
		dat wowc.RealmStatus
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRealmStatusPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRecipe returns basic recipe information
func (c *Client) WoWRecipe(recipeID int) (*wowc.Recipe, []byte, error) {
	var (
		dat wowc.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowRecipePath+"/"+strconv.Itoa(recipeID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWSpell returns information about spells
func (c *Client) WoWSpell(spellID int) (*wowc.Spell, []byte, error) {
	var (
		dat wowc.Spell
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowSpellPath+"/"+strconv.Itoa(spellID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWZoneMasterList returns a list of all supported zones and their bosses. A "zone" in this context should be considered a dungeon or a raid, not a world zone. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
func (c *Client) WoWZoneMasterList() (*wowc.ZoneMasterList, []byte, error) {
	var (
		dat wowc.ZoneMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowZonePath+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWZone returns information about zone
func (c *Client) WoWZone(zoneID int) (*wowc.Zone, []byte, error) {
	var (
		dat wowc.Zone
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowZonePath+"/"+strconv.Itoa(zoneID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWRegionBattlegroups returns a list of battlegroups for the specified region. Note the trailing / on this request path
func (c *Client) WoWRegionBattlegroups() (*wowc.RegionBattlegroups, []byte, error) {
	var (
		dat wowc.RegionBattlegroups
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataBattlegroupsPath+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWCharacterRaces returns a list of races and their associated faction, name, unique ID, and skin
func (c *Client) WoWCharacterRaces() (*wowc.CharacterRaces, []byte, error) {
	var (
		dat wowc.CharacterRaces
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataCharacterRacesPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWCharacterClasses returns a list of character classes
func (c *Client) WoWCharacterClasses() (*wowc.CharacterClasses, []byte, error) {
	var (
		dat wowc.CharacterClasses
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataCharacterClassesPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWCharacterAchievements returns a list of all achievements that characters can earn as well as the category structure and hierarchy
func (c *Client) WoWCharacterAchievements() (*wowc.CharacterAchievements, []byte, error) {
	var (
		dat wowc.CharacterAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataCharacterAchievementsPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWGuildRewards provides a list of all guild rewards
func (c *Client) WoWGuildRewards() (*wowc.GuildRewards, []byte, error) {
	var (
		dat wowc.GuildRewards
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataGuildRewardsPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWGuildPerks returns a list of all guild achievements as well as the category structure and hierarchy
func (c *Client) WoWGuildPerks() (*wowc.GuildPerks, []byte, error) {
	var (
		dat wowc.GuildPerks
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataGuildPerksPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWGuildAchievements returns a list of all guild achievements as well as the category structure and hierarchy
func (c *Client) WoWGuildAchievements() (*wowc.GuildAchievements, []byte, error) {
	var (
		dat wowc.GuildAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataGuildAchievementsPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWItemClasses returns a list of item classes
func (c *Client) WoWItemClasses() (*wowc.ItemClasses, []byte, error) {
	var (
		dat wowc.ItemClasses
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataItemClassesPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWTalents returns a list of talents, specs, and glyphs for each class
func (c *Client) WoWTalents() (*wowc.Talents, []byte, error) {
	var (
		dat wowc.Talents
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataTalentsPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}

// WoWPetTypes returns a list of the different battle pet types, including what they are strong and weak against
func (c *Client) WoWPetTypes() (*wowc.PetTypes, []byte, error) {
	var (
		dat wowc.PetTypes
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wowDataPetTypesPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, nil, err
	}

	return &dat, b, nil
}
