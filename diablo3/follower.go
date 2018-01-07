/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:04
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:01:14
 */

package diablo3

import "encoding/json"

// Follower structure
type Follower struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	RealName string `json:"realName"`
	Portrait string `json:"portrait"`
	Skills   struct {
		Active []struct {
			Slug              string `json:"slug"`
			Name              string `json:"name"`
			Icon              string `json:"icon"`
			Level             int    `json:"level"`
			TooltipURL        string `json:"tooltipUrl"`
			Description       string `json:"description"`
			SimpleDescription string `json:"simpleDescription"`
			SkillCalcID       string `json:"skillCalcId"`
		} `json:"active"`
		Passive []interface{} `json:"passive"`
	} `json:"skills"`
}

// JSON2Struct creates Follower structure from JSON byte array
func (f *Follower) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, f)
}
