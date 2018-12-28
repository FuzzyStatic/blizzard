package wowc

// Bracket field for PVP API calls
const (
	Bracket2v2 = "2v2"
	Bracket3v3 = "3v3"
	Bracket5v5 = "5v5"
	BracketRBG = "rbg"
)

// PVPLeaderboard structure
type PVPLeaderboard struct {
	Rows []struct {
		Ranking      int    `json:"ranking"`
		Rating       int    `json:"rating"`
		Name         string `json:"name"`
		RealmID      int    `json:"realmId"`
		RealmName    string `json:"realmName"`
		RealmSlug    string `json:"realmSlug"`
		RaceID       int    `json:"raceId"`
		ClassID      int    `json:"classId"`
		SpecID       int    `json:"specId"`
		FactionID    int    `json:"factionId"`
		GenderID     int    `json:"genderId"`
		SeasonWins   int    `json:"seasonWins"`
		SeasonLosses int    `json:"seasonLosses"`
		WeeklyWins   int    `json:"weeklyWins"`
		WeeklyLosses int    `json:"weeklyLosses"`
		Tier         string `json:"tier"`
	} `json:"rows"`
}
