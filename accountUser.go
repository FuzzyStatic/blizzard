/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-16 19:31:01
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-16 19:31:34
 */

package blizzard

import "encoding/json"

// AccountUser structure
type AccountUser struct {
	Battletag string `json:"battletag"`
	ID        int    `json:"id"`
}

// JSON2Struct creates AccountUser structure from JSON byte array
func (a *AccountUser) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, a)
}
