// Package wowgd contains types for the World of Warcraft Game Data APIs
package wowgd

// Bracket type
type Bracket string

// Bracket field for PVP API calls
const (
	Bracket2v2 Bracket = "2v2"
	Bracket3v3 Bracket = "3v3"
	BracketRBG Bracket = "rbg"
)

// TimeLeft string
type TimeLeft string

// TimeLeft field for Auction House structure
const (
	TimeLeftShort    TimeLeft = "SHORT"
	TimeLeftMedium   TimeLeft = "MEDIUM"
	TimeLeftLong     TimeLeft = "LONG"
	TimeLeftVeryLong TimeLeft = "VERY_LONG"
)
