/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 16:08:22
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 16:10:32
 */

package worldofwarcraft

import "encoding/json"

// RealmIndex structure
type RealmIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Realms []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
		Slug string `json:"slug"`
	} `json:"realms"`
}

// JSON2Struct creates RealmIndex structure from JSON byte array
func (r *RealmIndex) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, r)
}
