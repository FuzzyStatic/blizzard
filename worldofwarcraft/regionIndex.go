/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 17:31:27
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 17:32:00
 */

package worldofwarcraft

import "encoding/json"

// RegionIndex structure
type RegionIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Regions []struct {
		Href string `json:"href"`
	} `json:"regions"`
}

// JSON2Struct creates RegionIndex structure from JSON byte array
func (r *RegionIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, r)
}
