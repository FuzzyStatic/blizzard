package oauth

// UserInfo structure
type UserInfo struct {
	Sub       string `json:"sub"`
	ID        int    `json:"id"`
	Battletag string `json:"battletag"`
}
