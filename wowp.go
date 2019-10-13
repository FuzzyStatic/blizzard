package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/wowp"
)

const (
	profileWowCharacterPath   = "/profile/wow/character"
	mythicKeystoneProfilePath = "/mythic-keystone-profile"
	seasonPath                = "/season"
)

// WoWCharacterMythicKeystoneProfile returns a Mythic Keystone Profile index for a character
func (c *Client) WoWCharacterMythicKeystoneProfile(realmSlug, characterName string) (*wowp.CharacterMythicKeystoneProfile, []byte, error) {
	var (
		dat wowp.CharacterMythicKeystoneProfile
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+profileWowCharacterPath+"/"+realmSlug+"/"+characterName+"/"+mythicKeystoneProfilePath+"?"+localeQuery+c.locale.String(), c.profileNamespace)
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

	b, err = c.getURLBody(c.apiURL+profileWowCharacterPath+"/"+realmSlug+"/"+characterName+"/"+mythicKeystoneProfilePath+seasonPath+"/"+strconv.Itoa(seasonID)+"?"+localeQuery+c.locale.String(), c.profileNamespace)
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
