/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 21:45:21
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:57:01
 */

package worldofwarcraft

import "encoding/json"

// BossIndex structure
type BossIndex struct {
	Bosses []struct {
		Boss
		Location struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"location,omitempty"`
	} `json:"bosses"`
}

// JSON2Struct creates BossIndex structure from JSON byte array
func (boss *BossIndex) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, boss)
}
