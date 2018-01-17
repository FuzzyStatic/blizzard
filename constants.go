/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:40:25
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-16 19:29:59
 */

package blizzard

// Region type
type Region int

// Region constants
const (
	CN Region = iota
	EU
	KR
	SEA
	TW
	US
)

// Path constants
const (
	accountPath      = "/account"
	userPath         = "/user"
	accessTokenQuery = "access_token="
)
