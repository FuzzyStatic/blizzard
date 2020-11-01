package wowgd

// Token structure
type Token struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
	Price                int   `json:"price"`
}
