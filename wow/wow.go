// Package wowp contains types for the World of Warcraft Profile APIs
package wow

// Bracket type
type Bracket string

// Bracket field for PVP API calls
const (
	Bracket2v2 Bracket = "2v2"
	Bracket3v3 Bracket = "3v3"
	BracketRBG Bracket = "rbg"
)
