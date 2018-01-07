/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:28
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:38:28
 */

package worldofwarcraft

import "encoding/json"

// ConnectedRealmIndex structure
type ConnectedRealmIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ConnectedRealms []struct {
		Href string `json:"href"`
	} `json:"connected_realms"`
}

// JSON2Struct creates ConnectedRealmIndex structure from JSON byte array
func (c *ConnectedRealmIndex) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
