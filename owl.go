package blizzard

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/FuzzyStatic/blizzard/v3/owl"
)

// TempData is used because the OWL structure is key by just ID numbers without a proper field key.
type TempData struct {
	Players map[string]interface{} `json:"players"`
	Teams   map[string]interface{} `json:"teams"`
}

// OWLSummaryData Returns a summary of OWL stats where you can get entity IDs (playerId, matchId, segmentId, and teamId).
func (c *Client) OWLSummaryData(ctx context.Context) (*owl.SummaryData, *Header, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.apiHost+"/owl/v1/owl2", nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json")

	q := req.URL.Query()
	q.Set("locale", c.locale.String())
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, nil, errors.New(res.Status)
	}

	header := getHeader(res.Header)
	if err != nil {
		return nil, nil, err
	}

	var td TempData
	err = json.Unmarshal(body, &td)
	if err != nil {
		return nil, nil, err
	}

	var dat owl.SummaryData
	for _, v := range td.Players {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}

		var player owl.Player
		json.Unmarshal(b, &player)
		if err != nil {
			return nil, nil, err
		}

		dat.Players = append(dat.Players, player)
	}

	for _, v := range td.Teams {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}

		var team owl.Team
		err = json.Unmarshal(b, &team)
		if err != nil {
			return nil, nil, err
		}

		dat.Teams = append(dat.Teams, team)
	}

	return &dat, header, nil
}

// OWLPlayersAPI returns stats for a player.
func (c *Client) OWLPlayersAPI(ctx context.Context, playerId int) (*owl.PlayersAPI, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/owl/v1/players/%d", playerId),
		&owl.PlayersAPI{},
	)
	return dat.(*owl.PlayersAPI), header, err
}

// OWLMatchesAPI returns stats for a match.
func (c *Client) OWLMatchesAPI(ctx context.Context, matchId int) (*owl.MatchesAPI, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/owl/v1/matches/%d", matchId),
		&owl.MatchesAPI{},
	)
	return dat.(*owl.MatchesAPI), header, err
}

// OWLSegmentsAPI returns stats for a segment.
func (c *Client) OWLSegmentsAPI(ctx context.Context, segmentId string) (*owl.SegmentsAPI, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/owl/v1/segments/%s", segmentId),
		&owl.SegmentsAPI{},
	)
	return dat.(*owl.SegmentsAPI), header, err
}

// OWLTeamsAPI returns stats for a team.
func (c *Client) OWLTeamsAPI(ctx context.Context, teamId int) (*owl.TeamsAPI, *Header, error) {
	dat, header, err := c.getStructDataNoNamespace(ctx,
		fmt.Sprintf("/owl/v1/teams/%d", teamId),
		&owl.TeamsAPI{},
	)
	return dat.(*owl.TeamsAPI), header, err
}
