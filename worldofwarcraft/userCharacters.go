/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-18 17:09:25
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 17:09:59
 */

package worldofwarcraft

import "encoding/json"

// UserCharacters structure
type UserCharacters struct {
	Characters []struct {
		Name              string `json:"name"`
		Realm             string `json:"realm"`
		Battlegroup       string `json:"battlegroup"`
		Class             int    `json:"class"`
		Race              int    `json:"race"`
		Gender            int    `json:"gender"`
		Level             int    `json:"level"`
		AchievementPoints int    `json:"achievementPoints"`
		Thumbnail         string `json:"thumbnail"`
		Guild             string `json:"guild,omitempty"`
		GuildRealm        string `json:"guildRealm,omitempty"`
		LastModified      int64  `json:"lastModified"`
		Spec              struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"spec,omitempty"`
	} `json:"characters"`
}

// JSON2Struct creates UserCharacters structure from JSON byte array
func (u *UserCharacters) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, u)
}
