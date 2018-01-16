/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:22:45
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:23:24
 */

package worldofwarcraft

import "encoding/json"

// CharacterClasses structure
type CharacterClasses struct {
	Classes []struct {
		ID        int    `json:"id"`
		Mask      int    `json:"mask"`
		PowerType string `json:"powerType"`
		Name      string `json:"name"`
	} `json:"classes"`
}

// JSON2Struct creates CharacterClasses structure from JSON byte array
func (c *CharacterClasses) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
