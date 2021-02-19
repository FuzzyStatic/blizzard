package blizzard

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/FuzzyStatic/blizzard/v2/wowp"
	"golang.org/x/oauth2"
)

// WoWAccountProfileSummary Returns a profile summary for an account.
func (c *Client) WoWAccountProfileSummary(ctx context.Context, token *oauth2.Token) (dat *wowp.AccountProfileSummary, b []byte, err error) {
	var (
		req *http.Request
		res *http.Response
	)

	req, err = http.NewRequest("GET", c.apiURL+
		fmt.Sprintf("/profile/user/wow?namespace=%s&locale=%s", c.profileNamespace, c.locale), nil)
	if err != nil {
		return
	}

	client := c.authorizedCfg.Client(ctx, token)
	if res, err = client.Do(req); err != nil {
		return
	}
	defer func() {
		if derr := res.Body.Close(); derr != nil {
			err = derr
		}
	}()

	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	err = json.Unmarshal(b, &dat)
	return
}

// WoWCharacterAchievementsSummary returns a summary of the achievements a character has completed.
func (c *Client) WoWCharacterAchievementsSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterAchievementsSummary, []byte, error) {
	var (
		dat wowp.CharacterAchievementsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/achievements?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterAchievementsStatistics returns a character's statistics as they pertain to achievements.
func (c *Client) WoWCharacterAchievementsStatistics(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterAchievementsStatistics, []byte, error) {
	var (
		dat wowp.CharacterAchievementsStatistics
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/achievements/statistics?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterAppearanceSummary returns a summary of a character's appearance settings.
func (c *Client) WoWCharacterAppearanceSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterAppearanceSummary, []byte, error) {
	var (
		dat wowp.CharacterAppearanceSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/appearance?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterCollectionsIndex returns an index of collection types for a character.
func (c *Client) WoWCharacterCollectionsIndex(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterCollectionsIndex, []byte, error) {
	var (
		dat wowp.CharacterCollectionsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterMountsCollectionSummary returns a summary of the mounts a character has obtained.
func (c *Client) WoWCharacterMountsCollectionSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterMountsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterMountsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/mounts?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterPetsCollectionSummary returns a summary of the mounts a character has obtained.
func (c *Client) WoWCharacterPetsCollectionSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterPetsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterPetsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/pets?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterEncountersSummary returns a summary of a character's encounters.
func (c *Client) WoWCharacterEncountersSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterEncountersSummary, []byte, error) {
	var (
		dat wowp.CharacterEncountersSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterDungeons returns a summary of a character's completed dungeons.
func (c *Client) WoWCharacterDungeons(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterDungeons, []byte, error) {
	var (
		dat wowp.CharacterDungeons
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters/dungeons?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterRaids returns a summary of a character's completed raids.
func (c *Client) WoWCharacterRaids(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterRaids, []byte, error) {
	var (
		dat wowp.CharacterRaids
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters/raids?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterEquipmentSummary returns a summary of the items equipped by a character.
func (c *Client) WoWCharacterEquipmentSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterEquipmentSummary, []byte, error) {
	var (
		dat wowp.CharacterEquipmentSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/equipment?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterHunterPetsSummary if the character is a hunter, returns a summary of the character's hunter pets. Otherwise, returns an HTTP 404 Not Found error.
func (c *Client) WoWCharacterHunterPetsSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterHunterPetsSummary, []byte, error) {
	var (
		dat wowp.CharacterHunterPetsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/hunter-pets?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterMediaSummary returns a summary of the media assets available for a character (such as an avatar render).
func (c *Client) WoWCharacterMediaSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterMediaSummary, []byte, error) {
	var (
		dat wowp.CharacterMediaSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/character-media?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneProfileIndex returns the Mythic Keystone season details for a character.
//
// Returns a 404 Not Found for characters that have not yet completed a Mythic Keystone dungeon for the
// specified season.
func (c *Client) WoWMythicKeystoneProfileIndex(ctx context.Context, realmSlug,
	characterName string) (*wowp.CharacterMythicKeystoneProfile, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile?locale=%s",
			realmSlug, strings.ToLower(characterName), c.locale), c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWMythicKeystoneSeasonDetails returns a summary of the media assets available for a character (such as an avatar render).
func (c *Client) WoWMythicKeystoneSeasonDetails(ctx context.Context, realmSlug,
	characterName string, seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfileSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile/season/%d?locale=%s",
			realmSlug, strings.ToLower(characterName), seasonID, c.locale), c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterProfessionsSummary returns a summary of professions for a character.
func (c *Client) WoWCharacterProfessionsSummary(ctx context.Context, realmSlug,
	characterName string) (*wowp.CharacterProfessionsSummary, []byte, error) {
	var (
		dat wowp.CharacterProfessionsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/professions?locale=%s",
			realmSlug, strings.ToLower(characterName), c.locale), c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterProfileSummary returns a profile summary for a character.
func (c *Client) WoWCharacterProfileSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterProfileSummary, []byte, error) {
	var (
		dat wowp.CharacterProfileSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterProfileStatus returns the status and a unique ID for a character.
// A client should delete information about a character from their application if any of the following conditions occur:
// * an HTTP 404 Not Found error is returned
// * the is_valid value is false
// * the returned character ID doesn't match the previously recorded value for the character
// The following example illustrates how to use this endpoint:
// 1. A client requests and stores information about a character, including its unique character ID and the timestamp of the request.
// 2. After 30 days, the client makes a request to the status endpoint to verify if the character information is still valid.
// 3. If character cannot be found, is not valid, or the characters IDs do not match, the client removes the information from their application.
// 4. If the character is valid and the character IDs match, the client retains the data for another 30 days.
func (c *Client) WoWCharacterProfileStatus(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterProfileStatus, []byte, error) {
	var (
		dat wowp.CharacterProfileStatus
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/status?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterMythicKeystoneProfile returns a Mythic Keystone Profile index for a character.
func (c *Client) WoWCharacterMythicKeystoneProfile(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterMythicKeystoneProfileSeason returns a Mythic Keystone Profile index for a character.
// Note: this request returns a 404 for characters that have not completed a Mythic Keystone dungeon.
func (c *Client) WoWCharacterMythicKeystoneProfileSeason(ctx context.Context, realmSlug, characterName string, seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfileSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile/season/%d?locale=%s", realmSlug, strings.ToLower(characterName), seasonID, c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterPvPBracketStatistics returns the PvP bracket statistics for a character.
func (c *Client) WoWCharacterPvPBracketStatistics(ctx context.Context, realmSlug, characterName string, pvpBracket wowp.Bracket) (*wowp.CharacterPvPBracketStatistics, []byte, error) {
	var (
		dat wowp.CharacterPvPBracketStatistics
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/pvp-bracket/%s?locale=%s", realmSlug, strings.ToLower(characterName), pvpBracket, c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterPvPSummary returns a PvP summary for a character.
func (c *Client) WoWCharacterPvPSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterPvPSummary, []byte, error) {
	var (
		dat wowp.CharacterPvPSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/pvp-summary?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterQuests returns a character's active quests as well as a link to the character's completed quests.
func (c *Client) WoWCharacterQuests(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterQuests, []byte, error) {
	var (
		dat wowp.CharacterQuests
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/quests?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterCompletedQuests returns a list of quests that a character has completed.
func (c *Client) WoWCharacterCompletedQuests(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterCompletedQuests, []byte, error) {
	var (
		dat wowp.CharacterCompletedQuests
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/quests/completed?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterReputationsSummary returns a summary of a character's reputations.
func (c *Client) WoWCharacterReputationsSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterReputationsSummary, []byte, error) {
	var (
		dat wowp.CharacterReputationsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/reputations?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterSoulbinds returns a character's soulbinds.
func (c *Client) WoWCharacterSoulbinds(ctx context.Context, realmSlug,
	characterName string) (*wowp.CharacterSoulbinds, []byte, error) {
	var (
		dat wowp.CharacterSoulbinds
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/soulbinds?locale=%s",
			realmSlug, strings.ToLower(characterName), c.locale), c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterSpecializationsSummary returns a summary of a character's specializations.
func (c *Client) WoWCharacterSpecializationsSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterSpecializationsSummary, []byte, error) {
	var (
		dat wowp.CharacterSpecializationsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/specializations?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterStatisticsSummary returns a statistics summary for a character.
func (c *Client) WoWCharacterStatisticsSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterStatisticsSummary, []byte, error) {
	var (
		dat wowp.CharacterStatisticsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/statistics?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWCharacterTitlesSummary returns a summary of titles a character has obtained.
func (c *Client) WoWCharacterTitlesSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterTitlesSummary, []byte, error) {
	var (
		dat wowp.CharacterTitlesSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/titles?locale=%s", realmSlug, strings.ToLower(characterName), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuild returns a single guild by its name and realm.
func (c *Client) WoWGuild(ctx context.Context, realmSlug, nameSlug string) (*wowp.Guild, []byte, error) {
	var (
		dat wowp.Guild
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/data/wow/guild/%s/%s?locale=%s", realmSlug, strings.Replace(strings.ToLower(nameSlug), " ", "-", -1), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuildAchievements returns a single guild's achievements by name and realm.
func (c *Client) WoWGuildAchievements(ctx context.Context, realmSlug, nameSlug string) (*wowp.GuildAchievements, []byte, error) {
	var (
		dat wowp.GuildAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/data/wow/guild/%s/%s/achievements?locale=%s", realmSlug, strings.Replace(strings.ToLower(nameSlug), " ", "-", -1), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// WoWGuildRoster returns a single guild's roster by its name and realm.
func (c *Client) WoWGuildRoster(ctx context.Context, realmSlug, nameSlug string) (*wowp.GuildRoster, []byte, error) {
	var (
		dat wowp.GuildRoster
		b   []byte
		err error
	)

	b, err = c.getURLBody(ctx, c.apiURL+
		fmt.Sprintf("/data/wow/guild/%s/%s/roster?locale=%s", realmSlug,
			strings.Replace(strings.ToLower(nameSlug), " ", "-", -1), c.locale),
		c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
