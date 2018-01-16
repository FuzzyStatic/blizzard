/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 17:45:41
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:55:13
 */

package worldofwarcraft

import "encoding/json"

// TokenIndex structure
type TokenIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	LastUpdated int `json:"last_updated"`
	Price       int `json:"price"`
}

// JSON2Struct creates TokenIndex structure from JSON byte array
func (t *TokenIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, t)
}
