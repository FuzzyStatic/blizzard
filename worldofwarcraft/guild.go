/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 19:58:00
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 20:05:40
 */

package worldofwarcraft

import "encoding/json"

// Guild structure
type Guild struct {
	LastModified      int64  `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Level             int    `json:"level"`
	Side              int    `json:"side"`
	AchievementPoints int    `json:"achievementPoints"`
	Achievements      struct {
		AchievementsCompleted          []int         `json:"achievementsCompleted"`
		AchievementsCompletedTimestamp []int64       `json:"achievementsCompletedTimestamp"`
		Criteria                       []int         `json:"criteria"`
		CriteriaQuantity               []interface{} `json:"criteriaQuantity"`
		CriteriaTimestamp              []int64       `json:"criteriaTimestamp"`
		CriteriaCreated                []int64       `json:"criteriaCreated"`
	} `json:"achievements"`
	Members []struct {
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
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
	Emblem struct {
		Icon              int    `json:"icon"`
		IconColor         string `json:"iconColor"`
		IconColorID       int    `json:"iconColorId"`
		Border            int    `json:"border"`
		BorderColor       string `json:"borderColor"`
		BorderColorID     int    `json:"borderColorId"`
		BackgroundColor   string `json:"backgroundColor"`
		BackgroundColorID int    `json:"backgroundColorId"`
	} `json:"emblem"`
	News []struct {
		Type        string `json:"type"`
		Character   string `json:"character"`
		Timestamp   int64  `json:"timestamp"`
		ItemID      int    `json:"itemId,omitempty"`
		Context     string `json:"context"`
		BonusLists  []int  `json:"bonusLists"`
		Achievement struct {
			ID          int           `json:"id"`
			Title       string        `json:"title"`
			Points      int           `json:"points"`
			Description string        `json:"description"`
			Reward      string        `json:"reward"`
			RewardItems []interface{} `json:"rewardItems"`
			Icon        string        `json:"icon"`
			Criteria    []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool `json:"accountWide"`
			FactionID   int  `json:"factionId"`
		} `json:"achievement,omitempty"`
	} `json:"news"`
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
		Groups []interface{} `json:"groups"`
	} `json:"challenge"`
}

// JSON2Struct creates Guild structure from JSON byte array
func (g *Guild) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, g)
}
