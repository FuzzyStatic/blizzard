/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:30:18
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:33:14
 */

package worldofwarcraft

import "encoding/json"

// Spell structure
type Spell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Subtext     string `json:"subtext"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
}

// JSON2Struct creates Spell structure from JSON byte array
func (s *Spell) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, s)
}
