package blizzard

import (
	"encoding/json"

	"github.com/FuzzyStatic/blizzard/wowgd"
)

const (
	dataWowPath                       = dataPath + wowPath
	indexPath                         = "/index"
	wowConnectedRealm                 = dataWowPath + "/connected-realm"
	wowConnectedRealmIndex            = wowConnectedRealm + indexPath
	wowKeystoneAffinity               = dataWowPath + "/keystone-affinity"
	wowKeystoneAffinityIndex          = wowKeystoneAffinity + indexPath
	wowLeaderboardHallOfFame          = dataWowPath + "/leaderboard/hall-of-fame"
	wowMythicKeystonePath             = dataWowPath + "/mythic-keystone"
	wowMythicKeystoneDungeonPath      = wowMythicKeystonePath + "/dungeon"
	wowMythicKeystoneDungeonIndexPath = wowMythicKeystoneDungeonPath + indexPath
)

// WoWConnectedRealmIndex returns
func (c *Config) WoWConnectedRealmIndex() (*wowgd.ConnectedRealmIndex, error) {
	var (
		dat wowgd.ConnectedRealmIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + wowConnectedRealmIndex + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
