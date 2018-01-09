/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-08 21:50:21
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-08 21:50:53
 */

package worldofwarcraft

import "encoding/json"

// ChallengeRegionLeaderboard structure
type ChallengeRegionLeaderboard struct {
	Challenge []struct {
		Groups []interface{} `json:"groups"`
	} `json:"challenge"`
}

// JSON2Struct creates ChallengeRegionLeaderboard structure from JSON byte array
func (c *ChallengeRegionLeaderboard) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
