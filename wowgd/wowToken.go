package wowgd

// WowToken structure
type WowToken struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
	Price                int   `json:"price"`
}
