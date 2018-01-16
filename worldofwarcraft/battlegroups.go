/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 17:59:24
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:01:36
 */

package worldofwarcraft

import "encoding/json"

// Battlegroups structure
type Battlegroups struct {
	Battlegroups []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"battlegroups"`
}

// JSON2Struct creates Battlegroups structure from JSON byte array
func (bg *Battlegroups) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, bg)
}
