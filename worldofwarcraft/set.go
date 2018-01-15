/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-14 20:44:15
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-14 20:44:59
 */

package worldofwarcraft

import "encoding/json"

// Set structure
type Set struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SetBonuses []struct {
		Description string `json:"description"`
		Threshold   int    `json:"threshold"`
	} `json:"setBonuses"`
	Items []int `json:"items"`
}

// JSON2Struct creates Set structure from JSON byte array
func (s *Set) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, s)
}
