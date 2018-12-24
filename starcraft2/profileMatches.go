/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-17 20:05:22
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-17 20:05:48
 */

package starcraft2

import "encoding/json"

// ProfileMatches structure
type ProfileMatches struct {
	Matches []struct {
		Map      string `json:"map"`
		Type     string `json:"type"`
		Decision string `json:"decision"`
		Speed    string `json:"speed"`
		Date     int    `json:"date"`
	} `json:"matches"`
}

// JSON2Struct creates ProfileMatches structure from JSON byte array
func (p *ProfileMatches) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, p)
}
