package blizzard

import (
	"encoding/json"
	"strconv"

	"github.com/FuzzyStatic/blizzard/sc2gd"
)

const (
	sc2LeaguePath = dataPath + sc2Path + "/league"
)

// SC2GetLeagueData returns all SC2 league data queue ID, team type, and league ID
func (c *Config) SC2GetLeagueData(seasonID int, queueID sc2gd.QueueID, teamType sc2gd.TeamType, leagueID sc2gd.LeagueID) (*sc2gd.League, error) {
	var (
		dat sc2gd.League
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL + sc2LeaguePath + "/" + strconv.Itoa(seasonID) + "/" + strconv.Itoa(int(queueID)) + "/" + strconv.Itoa(int(teamType)) + "/" + strconv.Itoa(int(leagueID)) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
