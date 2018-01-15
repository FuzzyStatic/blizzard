/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:44:24
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 13:04:38
 */

package worldofwarcraft

import "encoding/json"

// ZoneIndex structure
type ZoneIndex struct {
	Zones []struct {
		Zone
	} `json:"zones"`
}

// JSON2Struct creates ZoneIndex structure from JSON byte array
func (z *ZoneIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, z)
}
