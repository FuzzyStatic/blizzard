package blizzard

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/FuzzyStatic/blizzard/owl"
)

// TempData is used because the OWL structure is key by just ID numbers without a proper field key.
type TempData struct {
	Players map[string]interface{} `json:"players"`
	Teams   map[string]interface{} `json:"teams"`
}

// OWLSummaryData Returns a summary of OWL stats where you can get entity IDs (playerId, matchId, segmentId, and teamId)
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
	for v := range td.Players {
		var player owl.Player
		b, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}

		err = json.Unmarshal(b, &player)
		if err != nil {
			return nil, nil, err
		}

		dat.Players = append(dat.Players, player)
	}

	for v := range td.Teams {
		var team owl.Team
		b, err := json.Marshal(v)
		if err != nil {
			return nil, nil, err
		}

		err = json.Unmarshal(b, &team)
		if err != nil {
			return nil, nil, err
		}

		dat.Teams = append(dat.Teams, team)
	}

	return &dat, header, nil
}
