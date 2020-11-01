package sc2c

// Player structure
type Player []struct {
	Name       string `json:"name"`
	ProfileURL string `json:"profileUrl"`
	AvatarURL  string `json:"avatarUrl"`
	ProfileID  string `json:"profileId"`
	RegionID   int    `json:"regionId"`
	RealmID    int    `json:"realmId"`
}
