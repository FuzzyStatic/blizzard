/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-18 16:48:09
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 16:49:26
 */

package starcraft2

import "encoding/json"

// Achievements structure
type Achievements struct {
	Achievements []struct {
		Title         string `json:"title"`
		Description   string `json:"description"`
		AchievementID int64  `json:"achievementId"`
		CategoryID    int    `json:"categoryId"`
		Points        int    `json:"points"`
		Icon          struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
	} `json:"achievements"`
	Categories []struct {
		Title                 string `json:"title"`
		CategoryID            int    `json:"categoryId"`
		FeaturedAchievementID int64  `json:"featuredAchievementId"`
		Children              []struct {
			Title                 string `json:"title"`
			CategoryID            int    `json:"categoryId"`
			FeaturedAchievementID int    `json:"featuredAchievementId"`
		} `json:"children,omitempty"`
	} `json:"categories"`
}

// JSON2Struct creates Achievements structure from JSON byte array
func (a *Achievements) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, a)
}
