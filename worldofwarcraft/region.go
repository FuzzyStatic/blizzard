/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 17:36:11
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 17:41:23
 */

package worldofwarcraft

import "encoding/json"

// Region structure
type Region struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

// JSON2Struct creates Region structure from JSON byte array
func (r *Region) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, r)
}
