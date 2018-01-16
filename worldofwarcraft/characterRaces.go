/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:10:53
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:17:49
 */

package worldofwarcraft

import "encoding/json"

// CharacterRaces structure
type CharacterRaces struct {
	Races []struct {
		ID   int    `json:"id"`
		Mask int    `json:"mask"`
		Side string `json:"side"`
		Name string `json:"name"`
	} `json:"races"`
}

// JSON2Struct creates CharacterRaces structure from JSON byte array
func (c *CharacterRaces) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
