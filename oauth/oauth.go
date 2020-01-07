// Package oauth contains types for the OAuth APIs
package oauth

// Profile type
type Profile string

// PRofile field authorize scopes
const (
	ProfileD3  Profile = "d3.profile"
	ProfileSC2 Profile = "sc2.profile"
	ProfileWoW Profile = "wow.profile"
)
