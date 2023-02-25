package blizzard

import (
	"context"
	"fmt"
	"strings"

	"github.com/FuzzyStatic/blizzard/wow"
	"github.com/FuzzyStatic/blizzard/wowp"
	"golang.org/x/oauth2"
)

// WoWAccountProfileSummary Returns a profile summary for an account.
func (c *Client) WoWAccountProfileSummary(ctx context.Context, token *oauth2.Token) (*wowp.AccountProfileSummary, *Header, error) {
	dat, header, err := c.getStructDataOAuth(ctx,
		"/profile/user/wow",
		c.GetProfileNamespace(),
		token,
		&wowp.AccountProfileSummary{},
	)
	return dat.(*wowp.AccountProfileSummary), header, err
}

// WoWCharacterAchievementsSummary returns a summary of the achievements a character has completed.
func (c *Client) WoWCharacterAchievementsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterAchievementsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/achievements", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterAchievementsSummary{},
	)
	return dat.(*wowp.CharacterAchievementsSummary), header, err
}

// WoWCharacterAchievementsStatistics returns a character's statistics as they pertain to achievements.
func (c *Client) WoWCharacterAchievementsStatistics(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterAchievementsStatistics, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/achievements/statistics", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterAchievementsStatistics{},
	)
	return dat.(*wowp.CharacterAchievementsStatistics), header, err
}

// WoWCharacterAppearanceSummary returns a summary of a character's appearance settings.
func (c *Client) WoWCharacterAppearanceSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterAppearanceSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/appearance", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterAppearanceSummary{},
	)
	return dat.(*wowp.CharacterAppearanceSummary), header, err
}

// WoWCharacterCollectionsIndex returns an index of collection types for a character.
func (c *Client) WoWCharacterCollectionsIndex(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterCollectionsIndex, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/collections", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterCollectionsIndex{},
	)
	return dat.(*wowp.CharacterCollectionsIndex), header, err
}

// WoWCharacterMountsCollectionSummary returns a summary of the mounts a character has obtained.
func (c *Client) WoWCharacterMountsCollectionSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterMountsCollectionSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/mounts", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterMountsCollectionSummary{},
	)
	return dat.(*wowp.CharacterMountsCollectionSummary), header, err
}

// WoWCharacterPetsCollectionSummary returns a summary of the mounts a character has obtained.
func (c *Client) WoWCharacterPetsCollectionSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterPetsCollectionSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/pets", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterPetsCollectionSummary{},
	)
	return dat.(*wowp.CharacterPetsCollectionSummary), header, err
}

// WoWCharacterEncountersSummary returns a summary of a character's encounters.
func (c *Client) WoWCharacterEncountersSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterEncountersSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterEncountersSummary{},
	)
	return dat.(*wowp.CharacterEncountersSummary), header, err
}

// WoWCharacterDungeons returns a summary of a character's completed dungeons.
func (c *Client) WoWCharacterDungeons(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterDungeons, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters/dungeons", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterDungeons{},
	)
	return dat.(*wowp.CharacterDungeons), header, err
}

// WoWCharacterRaids returns a summary of a character's completed raids.
func (c *Client) WoWCharacterRaids(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterRaids, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/encounters/raids", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterRaids{},
	)
	return dat.(*wowp.CharacterRaids), header, err
}

// WoWCharacterEquipmentSummary returns a summary of the items equipped by a character.
func (c *Client) WoWCharacterEquipmentSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterEquipmentSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/equipment", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterEquipmentSummary{},
	)
	return dat.(*wowp.CharacterEquipmentSummary), header, err
}

// WoWCharacterHunterPetsSummary if the character is a hunter, returns a summary of the character's hunter pets. Otherwise, returns an HTTP 404 Not Found error.
func (c *Client) WoWCharacterHunterPetsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterHunterPetsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/hunter-pets", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterHunterPetsSummary{},
	)
	return dat.(*wowp.CharacterHunterPetsSummary), header, err
}

// WoWCharacterMediaSummary returns a summary of the media assets available for a character (such as an avatar render).
func (c *Client) WoWCharacterMediaSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterMediaSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/character-media", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterMediaSummary{},
	)
	return dat.(*wowp.CharacterMediaSummary), header, err
}

// WoWMythicKeystoneProfileIndex returns the Mythic Keystone season details for a character.
// Returns a 404 Not Found for characters that have not yet completed a Mythic Keystone dungeon for the specified season.
func (c *Client) WoWMythicKeystoneProfileIndex(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile",
			realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterMythicKeystoneProfile{},
	)
	return dat.(*wowp.CharacterMythicKeystoneProfile), header, err
}

// WoWMythicKeystoneSeasonDetails returns a summary of the media assets available for a character (such as an avatar render).
func (c *Client) WoWMythicKeystoneSeasonDetails(ctx context.Context, realmSlug, characterName string,
	seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile/season/%d",
			realmSlug, strings.ToLower(characterName), seasonID),
		c.GetProfileNamespace(),
		&wowp.CharacterMythicKeystoneProfileSeason{},
	)
	return dat.(*wowp.CharacterMythicKeystoneProfileSeason), header, err
}

// WoWCharacterProfessionsSummary returns a summary of professions for a character.
func (c *Client) WoWCharacterProfessionsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterProfessionsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/professions", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterProfessionsSummary{},
	)
	return dat.(*wowp.CharacterProfessionsSummary), header, err
}

// WoWCharacterProfileSummary returns a profile summary for a character.
func (c *Client) WoWCharacterProfileSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterProfileSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterProfileSummary{},
	)
	return dat.(*wowp.CharacterProfileSummary), header, err
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
func (c *Client) WoWCharacterProfileStatus(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterProfileStatus, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/status", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterProfileStatus{},
	)
	return dat.(*wowp.CharacterProfileStatus), header, err
}

// WoWCharacterMythicKeystoneProfile returns a Mythic Keystone Profile index for a character.
func (c *Client) WoWCharacterMythicKeystoneProfile(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterMythicKeystoneProfile{},
	)
	return dat.(*wowp.CharacterMythicKeystoneProfile), header, err
}

// WoWCharacterMythicKeystoneProfileSeason returns a Mythic Keystone Profile index for a character.
// Note: this request returns a 404 for characters that have not completed a Mythic Keystone dungeon.
func (c *Client) WoWCharacterMythicKeystoneProfileSeason(ctx context.Context, realmSlug, characterName string,
	seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile/season/%d", realmSlug, strings.ToLower(characterName), seasonID),
		c.GetProfileNamespace(),
		&wowp.CharacterMythicKeystoneProfileSeason{},
	)
	return dat.(*wowp.CharacterMythicKeystoneProfileSeason), header, err
}

// WoWCharacterPvPBracketStatistics returns the PvP bracket statistics for a character.
func (c *Client) WoWCharacterPvPBracketStatistics(ctx context.Context,
	realmSlug, characterName string, pvpBracket wow.Bracket) (*wowp.CharacterPvPBracketStatistics, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/pvp-bracket/%s", realmSlug, strings.ToLower(characterName), pvpBracket),
		c.GetProfileNamespace(),
		&wowp.CharacterPvPBracketStatistics{},
	)
	return dat.(*wowp.CharacterPvPBracketStatistics), header, err
}

// WoWCharacterPvPSummary returns a PvP summary for a character.
func (c *Client) WoWCharacterPvPSummary(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterPvPSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/pvp-summary", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterPvPSummary{},
	)
	return dat.(*wowp.CharacterPvPSummary), header, err
}

// WoWCharacterQuests returns a character's active quests as well as a link to the character's completed quests.
func (c *Client) WoWCharacterQuests(ctx context.Context, realmSlug, characterName string) (*wowp.CharacterQuests, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/quests", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterQuests{},
	)
	return dat.(*wowp.CharacterQuests), header, err
}

// WoWCharacterCompletedQuests returns a list of quests that a character has completed.
func (c *Client) WoWCharacterCompletedQuests(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterCompletedQuests, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/quests/completed", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterCompletedQuests{},
	)
	return dat.(*wowp.CharacterCompletedQuests), header, err
}

// WoWCharacterReputationsSummary returns a summary of a character's reputations.
func (c *Client) WoWCharacterReputationsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterReputationsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/reputations", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterReputationsSummary{},
	)
	return dat.(*wowp.CharacterReputationsSummary), header, err
}

// WoWCharacterSoulbinds returns a character's soulbinds.
func (c *Client) WoWCharacterSoulbinds(ctx context.Context, realmSlug,
	characterName string) (*wowp.CharacterSoulbinds, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/soulbinds", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterSoulbinds{},
	)
	return dat.(*wowp.CharacterSoulbinds), header, err
}

// WoWCharacterSpecializationsSummary returns a summary of a character's specializations.
func (c *Client) WoWCharacterSpecializationsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterSpecializationsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/specializations", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterSpecializationsSummary{},
	)
	return dat.(*wowp.CharacterSpecializationsSummary), header, err
}

// WoWCharacterStatisticsSummary returns a statistics summary for a character.
func (c *Client) WoWCharacterStatisticsSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterStatisticsSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/statistics", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterStatisticsSummary{},
	)
	return dat.(*wowp.CharacterStatisticsSummary), header, err
}

// WoWCharacterTitlesSummary returns a summary of titles a character has obtained.
func (c *Client) WoWCharacterTitlesSummary(ctx context.Context,
	realmSlug, characterName string) (*wowp.CharacterTitlesSummary, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/profile/wow/character/%s/%s/titles", realmSlug, strings.ToLower(characterName)),
		c.GetProfileNamespace(),
		&wowp.CharacterTitlesSummary{},
	)
	return dat.(*wowp.CharacterTitlesSummary), header, err
}

// WoWGuild returns a single guild by its name and realm.
func (c *Client) WoWGuild(ctx context.Context, realmSlug, nameSlug string) (*wowp.Guild, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/guild/%s/%s", realmSlug, formatNameSlug(nameSlug)),
		c.GetProfileNamespace(),
		&wowp.Guild{},
	)
	return dat.(*wowp.Guild), header, err
}

// WoWGuildActivity returns a single guild's activity by name and realm.
func (c *Client) WoWGuildActivity(ctx context.Context, realmSlug, nameSlug string) (*wowp.GuildActivity, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/guild/%s/%s/activity", realmSlug, formatNameSlug(nameSlug)),
		c.GetProfileNamespace(),
		&wowp.GuildActivity{},
	)
	return dat.(*wowp.GuildActivity), header, err
}

// WoWGuildAchievements returns a single guild's achievements by name and realm.
func (c *Client) WoWGuildAchievements(ctx context.Context, realmSlug, nameSlug string) (*wowp.GuildAchievements, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/guild/%s/%s/achievements", realmSlug, formatNameSlug(nameSlug)),
		c.GetProfileNamespace(),
		&wowp.GuildAchievements{},
	)
	return dat.(*wowp.GuildAchievements), header, err
}

// WoWGuildRoster returns a single guild's roster by its name and realm.
func (c *Client) WoWGuildRoster(ctx context.Context, realmSlug, nameSlug string) (*wowp.GuildRoster, *Header, error) {
	dat, header, err := c.getStructData(ctx,
		fmt.Sprintf("/data/wow/guild/%s/%s/roster", realmSlug, formatNameSlug(nameSlug)),
		c.GetProfileNamespace(),
		&wowp.GuildRoster{},
	)
	return dat.(*wowp.GuildRoster), header, err
}

func formatNameSlug(nameSlug string) string {
	return strings.Replace(strings.ToLower(nameSlug), " ", "-", -1)
}
