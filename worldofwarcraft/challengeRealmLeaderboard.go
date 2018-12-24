/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 22:38:46
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 22:47:52
 */

package worldofwarcraft

import (
	"encoding/json"
	"time"
)

// ChallengeRealmLeaderboard structure
type ChallengeRealmLeaderboard struct {
	Challenge []struct {
		Realm struct {
			Name            string   `json:"name"`
			Slug            string   `json:"slug"`
			Battlegroup     string   `json:"battlegroup"`
			Locale          string   `json:"locale"`
			Timezone        string   `json:"timezone"`
			ConnectedRealms []string `json:"connected_realms"`
		} `json:"realm"`
		Map struct {
			ID               int    `json:"id"`
			Name             string `json:"name"`
			Slug             string `json:"slug"`
			HasChallengeMode bool   `json:"hasChallengeMode"`
			BronzeCriteria   struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"bronzeCriteria"`
			SilverCriteria struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"silverCriteria"`
			GoldCriteria struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"goldCriteria"`
		} `json:"map"`
		Groups []struct {
			Ranking int `json:"ranking"`
			Time    struct {
				Time         int  `json:"time"`
				Hours        int  `json:"hours"`
				Minutes      int  `json:"minutes"`
				Seconds      int  `json:"seconds"`
				Milliseconds int  `json:"milliseconds"`
				IsPositive   bool `json:"isPositive"`
			} `json:"time"`
			Date        time.Time `json:"date"`
			Medal       string    `json:"medal"`
			Faction     string    `json:"faction"`
			IsRecurring bool      `json:"isRecurring"`
			Members     []struct {
				Character struct {
					Name              string `json:"name"`
					Realm             string `json:"realm"`
					Battlegroup       string `json:"battlegroup"`
					Class             int    `json:"class"`
					Race              int    `json:"race"`
					Gender            int    `json:"gender"`
					Level             int    `json:"level"`
					AchievementPoints int    `json:"achievementPoints"`
					Thumbnail         string `json:"thumbnail"`
					Spec              struct {
						Name            string `json:"name"`
						Role            string `json:"role"`
						BackgroundImage string `json:"backgroundImage"`
						Icon            string `json:"icon"`
						Description     string `json:"description"`
						Order           int    `json:"order"`
					} `json:"spec"`
					Guild        string `json:"guild"`
					GuildRealm   string `json:"guildRealm"`
					LastModified int    `json:"lastModified"`
				} `json:"character,omitempty"`
				Spec struct {
					Name            string `json:"name"`
					Role            string `json:"role"`
					BackgroundImage string `json:"backgroundImage"`
					Icon            string `json:"icon"`
					Description     string `json:"description"`
					Order           int    `json:"order"`
				} `json:"spec"`
			} `json:"members"`
		} `json:"groups"`
	} `json:"challenge"`
}

// JSON2Struct creates ChallengeRealmLeaderboard structure from JSON byte array
func (c *ChallengeRealmLeaderboard) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, c)
}
