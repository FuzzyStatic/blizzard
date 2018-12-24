/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:13
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:00:56
 */

package diablo3

import "encoding/json"

// EraIndex structure
type EraIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Era []struct {
		Href string `json:"href"`
	} `json:"era"`
	CurrentEra     int    `json:"current_era"`
	LastUpdateTime string `json:"last_update_time"`
	GeneratedBy    string `json:"generated_by"`
}

// JSON2Struct creates EraIndex structure from JSON byte array
func (e *EraIndex) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, e)
}
