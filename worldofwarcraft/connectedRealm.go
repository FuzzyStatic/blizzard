/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:58:41
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:00:17
 */

package worldofwarcraft

import "encoding/json"

// ConnectedRealm structure
type ConnectedRealm struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID       int  `json:"id"`
	HasQueue bool `json:"has_queue"`
	Status   struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"status"`
	Population struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"population"`
	Realms []struct {
		ID     int `json:"id"`
		Region struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"region"`
		ConnectedRealm struct {
			Href string `json:"href"`
		} `json:"connected_realm"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Locale   string `json:"locale"`
		Timezone string `json:"timezone"`
		Type     struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"type"`
		IsTournament bool   `json:"is_tournament"`
		Slug         string `json:"slug"`
	} `json:"realms"`
	MythicLeaderboards struct {
		Href string `json:"href"`
	} `json:"mythic_leaderboards"`
}

// JSON2Struct creates ConnectedRealm structure from JSON byte array
func (c *ConnectedRealm) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, c)
}
