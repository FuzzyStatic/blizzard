/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 12:12:24
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:14:37
 */

package worldofwarcraft

import "encoding/json"

// Quest structure
type Quest struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	ReqLevel              int    `json:"reqLevel"`
	SuggestedPartyMembers int    `json:"suggestedPartyMembers"`
	Category              string `json:"category"`
	Level                 int    `json:"level"`
}

// JSON2Struct creates Quest structure from JSON byte array
func (q *Quest) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, q)
}
