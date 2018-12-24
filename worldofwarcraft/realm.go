/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 16:17:15
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 16:21:47
 */

package worldofwarcraft

import "encoding/json"

// Realm structure
type Realm struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID     int `json:"id"`
	Region struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"region"`
	ConnectedRealm struct {
		Href string `json:"href"`
	} `json:"connected_realm"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	Type     struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"type"`
	IsTournament bool   `json:"is_tournament"`
	Slug         string `json:"slug"`
}

// JSON2Struct creates Realm structure from JSON byte array
func (r *Realm) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, r)
}
