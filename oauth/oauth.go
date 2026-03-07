// Package oauth contains types for the OAuth APIs
package oauth

// Profile type
type Profile string

// Profile field authorize scopes
const (
	ProfileD3  Profile = "d3.profile"
	ProfileSC2 Profile = "sc2.profile"
	ProfileWoW Profile = "wow.profile"
	// Streaming provider service scopes
	ProfileStreamingTitles Profile = "streaming.titles"
	ProfileOpenID          Profile = "openid"
)
