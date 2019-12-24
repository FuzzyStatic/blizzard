package blizzard

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/FuzzyStatic/blizzard/wowp"
)

// WoWCharacterAchievementsSummary returns a summary of the achievements a character has completed.
func (c *Client) WoWCharacterAchievementsSummary(realmSlug, characterName string) (*wowp.CharacterAchievementsSummary, []byte, error) {
	var (
		dat wowp.CharacterAchievementsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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

// WoWCharacterAppearanceSummary returns a summary of a character's appearance settings.
func (c *Client) WoWCharacterAppearanceSummary(realmSlug, characterName string) (*wowp.CharacterAppearanceSummary, []byte, error) {
	var (
		dat wowp.CharacterAppearanceSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterCollectionsIndex(realmSlug, characterName string) (*wowp.CharacterCollectionsIndex, []byte, error) {
	var (
		dat wowp.CharacterCollectionsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterMountsCollectionSummary(realmSlug, characterName string) (*wowp.CharacterMountsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterMountsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterPetsCollectionSummary(realmSlug, characterName string) (*wowp.CharacterPetsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterPetsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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

// WoWCharacterEquipmentSummary returns a summary of the items equipped by a character.
func (c *Client) WoWCharacterEquipmentSummary(realmSlug, characterName string) (*wowp.CharacterEquipmentSummary, []byte, error) {
	var (
		dat wowp.CharacterEquipmentSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterHunterPetsSummary(realmSlug, characterName string) (*wowp.CharacterHunterPetsSummary, []byte, error) {
	var (
		dat wowp.CharacterHunterPetsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterMediaSummary(realmSlug, characterName string) (*wowp.CharacterMediaSummary, []byte, error) {
	var (
		dat wowp.CharacterMediaSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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

// WoWCharacterProfileSummary returns a profile summary for a character.
func (c *Client) WoWCharacterProfileSummary(realmSlug, characterName string) (*wowp.CharacterProfileSummary, []byte, error) {
	var (
		dat wowp.CharacterProfileSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterProfileStatus(realmSlug, characterName string) (*wowp.CharacterProfileStatus, []byte, error) {
	var (
		dat wowp.CharacterProfileStatus
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterMythicKeystoneProfile(realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterMythicKeystoneProfileSeason(realmSlug, characterName string, seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfileSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterPvPBracketStatistics(realmSlug, characterName string, pvpBracket wowp.Bracket) (*wowp.CharacterPvPBracketStatistics, []byte, error) {
	var (
		dat wowp.CharacterPvPBracketStatistics
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterPvPSummary(realmSlug, characterName string) (*wowp.CharacterPvPSummary, []byte, error) {
	var (
		dat wowp.CharacterPvPSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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

// WoWCharacterReputationsSummary returns a summary of a character's reputations.
func (c *Client) WoWCharacterReputationsSummary(realmSlug, characterName string) (*wowp.CharacterReputationsSummary, []byte, error) {
	var (
		dat wowp.CharacterReputationsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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

// WoWCharacterSpecializationsSummary returns a summary of a character's specializations.
func (c *Client) WoWCharacterSpecializationsSummary(realmSlug, characterName string) (*wowp.CharacterSpecializationsSummary, []byte, error) {
	var (
		dat wowp.CharacterSpecializationsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterStatisticsSummary(realmSlug, characterName string) (*wowp.CharacterStatisticsSummary, []byte, error) {
	var (
		dat wowp.CharacterStatisticsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWCharacterTitlesSummary(realmSlug, characterName string) (*wowp.CharacterTitlesSummary, []byte, error) {
	var (
		dat wowp.CharacterTitlesSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWGuild(realmSlug, nameSlug string) (*wowp.Guild, []byte, error) {
	var (
		dat wowp.Guild
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWGuildAchievements(realmSlug, nameSlug string) (*wowp.GuildAchievements, []byte, error) {
	var (
		dat wowp.GuildAchievements
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
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
func (c *Client) WoWGuildRoster(realmSlug, nameSlug string) (*wowp.GuildRoster, []byte, error) {
	var (
		dat wowp.GuildRoster
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/data/wow/guild/%s/%s/roster?locale=%s", realmSlug, strings.Replace(strings.ToLower(nameSlug), " ", "-", -1), c.locale),
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
