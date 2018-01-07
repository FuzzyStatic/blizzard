/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 15:51:15
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 15:51:36
 */

package worldofwarcraft

import "encoding/json"

// MythicLeaderboard structure
type MythicLeaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Map struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"map"`
	Period         int `json:"period"`
	PeriodStart    int `json:"period_start"`
	PeriodEnd      int `json:"period_end"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	LeadingGroups []struct {
		Ranking            int `json:"ranking"`
		Duration           int `json:"duration"`
		CompletedTimestamp int `json:"completed_timestamp"`
		KeystoneLevel      int `json:"keystone_level"`
		KeystoneAffixes    []struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"keystone_affixes"`
		Members []struct {
			Profile struct {
				Name  string `json:"name"`
				ID    int    `json:"id"`
				Realm struct {
					Key struct {
						Href string `json:"href"`
					} `json:"key"`
					Name string `json:"name"`
					ID   int    `json:"id"`
					Slug string `json:"slug"`
				} `json:"realm"`
			} `json:"profile"`
			Faction struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"faction"`
			Specialization struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"specialization"`
		} `json:"members"`
	} `json:"leading_groups"`
	KeystoneAffixes []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"keystone_affixes"`
	MapChallengeModeID int    `json:"map_challenge_mode_id"`
	Name               string `json:"name"`
}

// JSON2Struct creates MythicLeaderboard structure from JSON byte array
func (m *MythicLeaderboard) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, m)
}
