/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:21:44
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:22:14
 */

package worldofwarcraft

import "encoding/json"

// RealmStatus structure
type RealmStatus struct {
	Realms []struct {
		Type            string   `json:"type"`
		Population      string   `json:"population"`
		Queue           bool     `json:"queue"`
		Status          bool     `json:"status"`
		Name            string   `json:"name"`
		Slug            string   `json:"slug"`
		Battlegroup     string   `json:"battlegroup"`
		Locale          string   `json:"locale"`
		Timezone        string   `json:"timezone"`
		ConnectedRealms []string `json:"connected_realms"`
	} `json:"realms"`
}

// JSON2Struct creates RealmStatus structure from JSON byte array
func (r *RealmStatus) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, r)
}
