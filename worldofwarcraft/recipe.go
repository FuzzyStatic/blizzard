/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:27:50
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:28:11
 */

package worldofwarcraft

import "encoding/json"

// Recipe structure
type Recipe struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Icon       string `json:"icon"`
}

// JSON2Struct creates Recipe structure from JSON byte array
func (r *Recipe) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, r)
}
