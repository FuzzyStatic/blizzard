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
	wowBattlegroupsPath              = wowDataPath + "/battlegroups"
	wowDataCharacterPath             = wowDataPath + "/character"
	wowDataCharacterRacesPath        = wowDataCharacterPath + "/races"
	wowDataCharacterClassesPath      = wowDataCharacterPath + "/classes"
	wowDataCharacterAchievementsPath = wowDataCharacterPath + "/achievements"
	wowDataGuildPath                 = wowDataPath + "/guild"
	wowDataGuildRewardsPath          = wowDataGuildPath + "/rewards"
	wowDataGuildPerksPath            = wowDataGuildPath + "/perks"
	wowDataGuildAchievementsPath     = wowDataGuildPath + "/achievements"
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

// WoWBoss provides information about bosses. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
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

// WoWItem provides detailed item information, including item set information
func (c *Config) WoWItem(itemID int) (*wowc.Item, error) {
	var (
		dat wowc.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowItemPath + "/" + strconv.Itoa(itemID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWItemSet provides detailed item information, including item set information
func (c *Config) WoWItemSet(setID int) (*wowc.ItemSet, error) {
	var (
		dat wowc.ItemSet
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowItemSetPath + "/" + strconv.Itoa(setID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWMountMasterList returns a list of all supported mounts
func (c *Config) WoWMountMasterList() (*wowc.MountMasterList, error) {
	var (
		dat wowc.MountMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowMountPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWPetMasterList returns a list of all supported battle and vanity pets
func (c *Config) WoWPetMasterList() (*wowc.PetMasterList, error) {
	var (
		dat wowc.PetMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowPetPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWPetAbility returns data about a individual battle pet ability ID. This resource does not provide ability tooltips
func (c *Config) WoWPetAbility(abilityID int) (*wowc.PetAbility, error) {
	var (
		dat wowc.PetAbility
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowPetAbilityPath + "/" + strconv.Itoa(abilityID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWPetSpecies returns data about an individual pet species. Use pets as the field value in a character profile request to get species IDs. Each species also has data about its six abilities
func (c *Config) WoWPetSpecies(speciesID int) (*wowc.PetSpecies, error) {
	var (
		dat wowc.PetSpecies
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowPetSpeciesPath + "/" + strconv.Itoa(speciesID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWPetStats returns detailed information about a given species of pet
func (c *Config) WoWPetStats(speciesID, level, breedID, qualityID int) (*wowc.PetStats, error) {
	var (
		dat wowc.PetStats
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowPetStatsPath + "/" + strconv.Itoa(speciesID) + "?" + wowPetsLevelField + strconv.Itoa(level) + "&" + wowPetsBreedIDField + strconv.Itoa(breedID) + "&" + wowPetsQualityIDField + strconv.Itoa(qualityID) + "&" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWPVPLeaderboard provides leaderboard information for the 2v2, 3v3, 5v5, and Rated Battleground leaderboards
func (c *Config) WoWPVPLeaderboard(bracket string) (*wowc.PVPLeaderboard, error) {
	var (
		dat wowc.PVPLeaderboard
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowPVPLeaderboardPath + "/" + bracket + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWQuest returns metadata for a specified quest
func (c *Config) WoWQuest(questID int) (*wowc.Quest, error) {
	var (
		dat wowc.Quest
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowQuestPath + "/" + strconv.Itoa(questID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWRealmStatus returns metadata for a specified quest
func (c *Config) WoWRealmStatus() (*wowc.RealmStatus, error) {
	var (
		dat wowc.RealmStatus
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowRealmStatusPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWRecipe returns basic recipe information
func (c *Config) WoWRecipe(recipeID int) (*wowc.Recipe, error) {
	var (
		dat wowc.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowRecipePath + "/" + strconv.Itoa(recipeID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWSpell returns information about spells
func (c *Config) WoWSpell(spellID int) (*wowc.Spell, error) {
	var (
		dat wowc.Spell
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowSpellPath + "/" + strconv.Itoa(spellID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWZoneMasterList returns a list of all supported zones and their bosses. A "zone" in this context should be considered a dungeon or a raid, not a world zone. A "boss" in this context should be considered a boss encounter, which may include more than one NPC
func (c *Config) WoWZoneMasterList() (*wowc.ZoneMasterList, error) {
	var (
		dat wowc.ZoneMasterList
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowZonePath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWZone returns information about zone
func (c *Config) WoWZone(zoneID int) (*wowc.Zone, error) {
	var (
		dat wowc.Zone
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowZonePath + "/" + strconv.Itoa(zoneID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWRegionBattlegroups returns a list of battlegroups for the specified region. Note the trailing / on this request path
func (c *Config) WoWRegionBattlegroups() (*wowc.RegionBattlegroups, error) {
	var (
		dat wowc.RegionBattlegroups
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowBattlegroupsPath + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWCharacterRaces returns a list of races and their associated faction, name, unique ID, and skin
func (c *Config) WoWCharacterRaces() (*wowc.CharacterRaces, error) {
	var (
		dat wowc.CharacterRaces
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataCharacterRacesPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWCharacterClasses returns a list of character classes
func (c *Config) WoWCharacterClasses() (*wowc.CharacterClasses, error) {
	var (
		dat wowc.CharacterClasses
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataCharacterClassesPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWCharacterAchievements returns a list of all achievements that characters can earn as well as the category structure and hierarchy
func (c *Config) WoWCharacterAchievements() (*wowc.CharacterAchievements, error) {
	var (
		dat wowc.CharacterAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataCharacterAchievementsPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWGuildRewards provides a list of all guild rewards
func (c *Config) WoWGuildRewards() (*wowc.GuildRewards, error) {
	var (
		dat wowc.GuildRewards
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataGuildRewardsPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWGuildPerks returns a list of all guild achievements as well as the category structure and hierarchy
func (c *Config) WoWGuildPerks() (*wowc.GuildPerks, error) {
	var (
		dat wowc.GuildPerks
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataGuildPerksPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// WoWGuildAchievements returns a list of all guild achievements as well as the category structure and hierarchy
func (c *Config) WoWGuildAchievements() (*wowc.GuildAchievements, error) {
	var (
		dat wowc.GuildAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowDataGuildAchievementsPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
