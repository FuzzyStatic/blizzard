package wowc

// CharacterProfile structure
type CharacterProfile struct {
	LastModified        int64  `json:"lastModified"`
	Name                string `json:"name"`
	Realm               string `json:"realm"`
	Battlegroup         string `json:"battlegroup"`
	Class               int    `json:"class"`
	Race                int    `json:"race"`
	Gender              int    `json:"gender"`
	Level               int    `json:"level"`
	AchievementPoints   int    `json:"achievementPoints"`
	Thumbnail           string `json:"thumbnail"`
	CalcClass           string `json:"calcClass"`
	Faction             int    `json:"faction"`
	TotalHonorableKills int    `json:"totalHonorableKills"`
}
