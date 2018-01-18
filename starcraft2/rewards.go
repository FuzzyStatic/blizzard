/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-18 16:52:55
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-18 16:53:31
 */

package starcraft2

import "encoding/json"

// Rewards structure
type Rewards struct {
	Portraits []struct {
		Title string `json:"title"`
		ID    int64  `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"portraits"`
	TerranDecals []struct {
		Title string `json:"title"`
		ID    int64  `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"terranDecals"`
	ZergDecals []struct {
		Title string `json:"title"`
		ID    int    `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"zergDecals"`
	ProtossDecals []struct {
		Title string `json:"title"`
		ID    int    `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"protossDecals"`
	Skins []struct {
		Name  string `json:"name"`
		Title string `json:"title"`
		ID    int64  `json:"id"`
		Icon  struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"skins"`
	Animations []struct {
		Title   string `json:"title"`
		Command string `json:"command"`
		ID      int64  `json:"id"`
		Icon    struct {
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
			Offset int    `json:"offset"`
			URL    string `json:"url"`
		} `json:"icon"`
		AchievementID int `json:"achievementId"`
	} `json:"animations"`
}

// JSON2Struct creates Rewards structure from JSON byte array
func (r *Rewards) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, r)
}
