package blizzard

import (
	"encoding/json"
	"fmt"

	"github.com/FuzzyStatic/blizzard/sc2gd"
)

// SC2LeagueData returns all SC2 league data from for seasonID, queue ID, team type, and league ID
func (c *Client) SC2LeagueData(seasonID int, queueID sc2gd.QueueID, teamType sc2gd.TeamType, leagueID sc2gd.LeagueID) (*sc2gd.League, []byte, error) {
	var (
		dat sc2gd.League
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+fmt.Sprintf("/data/sc2/league/%d/%d/%d/%d?locale=%s", seasonID, queueID, teamType, leagueID, c.locale), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
