package owl

// SummaryData structure
type SummaryData struct {
	Players []Player `json:"players"`
	Teams   []Team   `json:"teams"`
}

// Player structure
type Player struct {
	Number        int    `json:"number"`
	PreferredSlot int    `json:"preferredSlot"`
	GivenName     string `json:"givenName"`
	Teams         []struct {
		ID            interface{} `json:"id"`
		EarliestMatch int64       `json:"earliestMatch,omitempty"`
		Stats         Stats       `json:"stats,omitempty"`
		Heroes        Heroes      `json:"heroes,omitempty"`
	} `json:"teams"`
	Name         string   `json:"name"`
	FamilyName   string   `json:"familyName"`
	Competitions []string `json:"competitions"`
	Role         string   `json:"role"`
	ID           int      `json:"id"`
	HeadshotURL  string   `json:"headshotUrl"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	CurrentTeam int `json:"currentTeam"`
}

type Team struct {
	ID           int      `json:"id"`
	Competitions []string `json:"competitions"`
	Name         string   `json:"name"`
	Roster       []int    `json:"roster"`
	Code         string   `json:"code"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
}
