/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:37:16
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:39:38
 */

package worldofwarcraft

import "encoding/json"

// GuildPerks structure
type GuildPerks struct {
	Perks []struct {
		GuildLevel int `json:"guildLevel"`
		Spell      struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Subtext     string `json:"subtext"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			CastTime    string `json:"castTime"`
		} `json:"spell"`
	} `json:"perks"`
}

// JSON2Struct creates GuildPerks structure from JSON byte array
func (g *GuildPerks) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, g)
}
