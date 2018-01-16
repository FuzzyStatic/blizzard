/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:46:09
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:46:59
 */

package worldofwarcraft

import "encoding/json"

// ItemClasses structure
type ItemClasses struct {
	Classes []struct {
		Class      int    `json:"class"`
		Name       string `json:"name"`
		Subclasses []struct {
			Subclass int    `json:"subclass"`
			Name     string `json:"name"`
		} `json:"subclasses"`
	} `json:"classes"`
}

// JSON2Struct creates ItemClasses structure from JSON byte array
func (i *ItemClasses) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, i)
}
