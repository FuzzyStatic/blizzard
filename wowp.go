package blizzard

import (
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/wowp"
)

// WoWCharacterAchievementsSummary returns a summary of the achievements a character has completed
func (c *Client) WoWCharacterAchievementsSummary(realmSlug, characterName string) (*wowp.CharacterAchievementsSummary, []byte, error) {
	var (
		dat wowp.CharacterAchievementsSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/achievements?locale=%s", realmSlug, characterName, c.locale),
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

// WoWCharacterAppearanceSummary returns a summary of a character's appearance settings
func (c *Client) WoWCharacterAppearanceSummary(realmSlug, characterName string) (*wowp.CharacterAppearanceSummary, []byte, error) {
	var (
		dat wowp.CharacterAppearanceSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/appearance?locale=%s", realmSlug, characterName, c.locale),
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

// WoWCharacterCollectionsIndex returns an index of collection types for a character
func (c *Client) WoWCharacterCollectionsIndex(realmSlug, characterName string) (*wowp.CharacterCollectionsIndex, []byte, error) {
	var (
		dat wowp.CharacterCollectionsIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections?locale=%s", realmSlug, characterName, c.locale),
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

// WoWCharacterMountsCollectionSummary returns a summary of the mounts a character has obtained
func (c *Client) WoWCharacterMountsCollectionSummary(realmSlug, characterName string) (*wowp.CharacterMountsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterMountsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/mounts?locale=%s", realmSlug, characterName, c.locale),
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

// WoWCharacterPetsCollectionSummary returns a summary of the mounts a character has obtained
func (c *Client) WoWCharacterPetsCollectionSummary(realmSlug, characterName string) (*wowp.CharacterPetsCollectionSummary, []byte, error) {
	var (
		dat wowp.CharacterPetsCollectionSummary
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/collections/pets?locale=%s", realmSlug, characterName, c.locale),
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

// WoWCharacterMythicKeystoneProfile returns a Mythic Keystone Profile index for a character
func (c *Client) WoWCharacterMythicKeystoneProfile(realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile?locale=%s", realmSlug, characterName, c.locale),
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
// Note: this request returns a 404 for characters that have not completed a Mythic Keystone dungeon
func (c *Client) WoWCharacterMythicKeystoneProfileSeason(realmSlug, characterName string, seasonID int) (*wowp.CharacterMythicKeystoneProfileSeason, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfileSeason
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+
		fmt.Sprintf("/profile/wow/character/%s/%s/mythic-keystone-profile/season/%d?locale=%s", realmSlug, characterName, seasonID, c.locale),
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
