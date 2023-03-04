package owl

// TeamsAPI structure
type TeamsAPI struct {
	ID           int      `json:"id"`
	Competitions []string `json:"competitions"`
	Name         string   `json:"name"`
	Roster       []int    `json:"roster"`
	Code         string   `json:"code"`
	Segments     struct {
		Owl22022KickoffClashQualifiers     TeamSegment `json:"owl2-2022-kickoff-clash-qualifiers"`
		Owl22022MidseasonMadnessQualifiers TeamSegment `json:"owl2-2022-midseason-madness-qualifiers"`
		Owl22022SummerShowdownQualifiers   TeamSegment `json:"owl2-2022-summer-showdown-qualifiers"`
		Owl22022CountdownCupQualifiers     TeamSegment `json:"owl2-2022-countdown-cup-qualifiers"`
		Owl22022Regular                    TeamSegment `json:"owl2-2022-regular"`
		OverwatchTest2022Ow2Playtest2      TeamSegment `json:"overwatch-test-2022-ow2-playtest2"`
		Owl22022MidseasonMadnessTournament TeamSegment `json:"owl2-2022-midseason-madness-tournament"`
		Owl22022SummerShowdownTournament   TeamSegment `json:"owl2-2022-summer-showdown-tournament"`
		Owl22022PostSeason                 TeamSegment `json:"owl2-2022-post-season"`
		Owl22022RegularTournaments         TeamSegment `json:"owl2-2022-regular-tournaments"`
		Owl22022RegularTournamentsPlayoffs TeamSegment `json:"owl2-2022-regular-tournaments-playoffs"`
	} `json:"segments"`
	TeamStats    TeamStats `json:"teamStats"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
}

// TeamStats structure
type TeamStats struct {
	HeroDamageDone int `json:"heroDamageDone"`
	HealingDone    int `json:"healingDone"`
	DamageTaken    int `json:"damageTaken"`
	FinalBlows     int `json:"finalBlows"`
	Eliminations   int `json:"eliminations"`
	Deaths         int `json:"deaths"`
}

// TeamSegment structure
type TeamSegment struct {
	ID             int       `json:"id"`
	TeamStats      TeamStats `json:"teamStats"`
	Logo           string    `json:"logo"`
	Icon           string    `json:"icon"`
	PrimaryColor   string    `json:"primaryColor"`
	SecondaryColor string    `json:"secondaryColor"`
	Name           string    `json:"name"`
}
