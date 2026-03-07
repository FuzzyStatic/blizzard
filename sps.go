package blizzard

import (
	"context"

	"github.com/FuzzyStatic/blizzard/v3/sps"
	"golang.org/x/oauth2"
)

// SPSPlayableTitles returns a list of playable titles for the account associated
// with the provided OAuth2 token.  The caller must request the
// "streaming.titles" and "openid" scopes during authorization.
func (c *Client) SPSPlayableTitles(ctx context.Context, token *oauth2.Token) (*sps.PlayableTitles, *Header, error) {
	dat, header, err := c.getStructDataOAuth(ctx,
		"/streaming/titles",
		"",
		token,
		&sps.PlayableTitles{},
	)
	return dat.(*sps.PlayableTitles), header, err
}
