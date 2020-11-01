package sc2gd

// QueueID IDs for different league queues
type QueueID int

// QueueID constants (1=WoL 1v1, 2=WoL 2v2, 3=WoL 3v3, 4=WoL 4v4) DO NOT REARRANGE
const (
	_ QueueID = iota
	WoL1v1
	WoL2v2
	WoL3v3
	WoL4v4
)

// QueueID constants (101=HotS 1v1, 102=HotS 2v2, 103=HotS 3v3, 104=HotS 4v4) DO NOT REARRANGE
const (
	_ QueueID = 100 + iota
	HotS1v1
	HotS2v2
	HotS3v3
	HotS4v4
)

// QueueID constants (201=LotV 1v1, 202=LotV 2v2, 203=LotV 3v3, 204=LotV 4v4, 206=LotV Archon) DO NOT REARRANGE
const (
	_ QueueID = 200 + iota
	LotV1v1
	LotV2v2
	LotV3v3
	LotV4v4
	_
	LotVArchon
)

// TeamType different team types
type TeamType int

// TeamType constants (0=arranged, 1=random) DO NOT REARRANGE
const (
	Arranged TeamType = iota
	Random
)

// LeagueID available leagues
type LeagueID int

// LeagueID constants ( 0=Bronze, 1=Silver, 2=Gold, 3=Platinum, 4=Diamond, 5=Master, 6=Grandmaster) DO NOT REARRANGE
const (
	Bronze LeagueID = iota
	Silver
	Gold
	Platinum
	Diamond
	Master
	Grandmaster
)

// League structure
type League struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Key struct {
		LeagueID int `json:"league_id"`
		SeasonID int `json:"season_id"`
		QueueID  int `json:"queue_id"`
		TeamType int `json:"team_type"`
	} `json:"key"`
	Tier []struct {
		ID        int `json:"id"`
		MinRating int `json:"min_rating"`
		MaxRating int `json:"max_rating"`
		Division  []struct {
			ID          int `json:"id"`
			LadderID    int `json:"ladder_id"`
			MemberCount int `json:"member_count"`
		} `json:"division"`
	} `json:"tier"`
}
