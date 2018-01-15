/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 10:35:54
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 12:00:42
 */

// Package starcraft2 is a client library to use Blizzard Starcraft 2 API calls.
package starcraft2

import "github.com/FuzzyStatic/blizzard"

// Starcraft2 regional API URLs, locale, access token, api key
type Starcraft2 struct {
	Auth         blizzard.Auth
	Locale       string
	DataURL      string
	CommunityURL string
}

// New create new Starcraft2 structure
func New(auth blizzard.Auth, region blizzard.Region) *Starcraft2 {
	var s = Starcraft2{Auth: auth}

	switch region {
	case blizzard.EU:
		s.Locale = "en_GB"
		s.DataURL = "https://eu.api.battle.net/data/wow"
		s.CommunityURL = "https://eu.api.battle.net/wow"
	case blizzard.KR:
		s.Locale = "ko_KR"
		s.DataURL = "https://kr.api.battle.net/data/wow"
		s.CommunityURL = "https://kr.api.battle.net/wow"
	case blizzard.SEA:
		s.Locale = "en_US"
		s.DataURL = "https://sea.api.battle.net/data/wow"
		s.CommunityURL = "https://sea.api.battle.net/wow"
	case blizzard.TW:
		s.Locale = "zh_TW"
		s.DataURL = "https://tw.api.battle.net/data/wow"
		s.CommunityURL = "https://tw.api.battle.net/wow"
	case blizzard.US:
		s.Locale = "en_US"
		s.DataURL = "https://us.api.battle.net/data/wow"
		s.CommunityURL = "https://us.api.battle.net/wow"
	default: // USA! USA!
		s.Locale = "en_US"
		s.DataURL = "https://us.api.battle.net/data/wow"
		s.CommunityURL = "https://us.api.battle.net/wow"
	}

	return &s
}
