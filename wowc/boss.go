package wowc

// BossMasterList structure
type BossMasterList struct {
	Bosses []Boss `json:"bosses"`
}

// Boss structure
type Boss struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	URLSlug               string `json:"urlSlug"`
	Description           string `json:"description"`
	ZoneID                int    `json:"zoneId"`
	AvailableInNormalMode bool   `json:"availableInNormalMode"`
	AvailableInHeroicMode bool   `json:"availableInHeroicMode"`
	Health                int    `json:"health"`
	HeroicHealth          int    `json:"heroicHealth"`
	Level                 int    `json:"level"`
	HeroicLevel           int    `json:"heroicLevel"`
	JournalID             int    `json:"journalId"`
	Npcs                  []struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		URLSlug           string `json:"urlSlug"`
		CreatureDisplayID int    `json:"creatureDisplayId"`
	} `json:"npcs"`
	EncounterFaction int `json:"encounterFaction"`
}
