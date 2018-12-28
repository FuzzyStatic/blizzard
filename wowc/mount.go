package wowc

// MountMasterList structure
type MountMasterList struct {
	Mounts []struct {
		Name       string `json:"name"`
		SpellID    int    `json:"spellId"`
		CreatureID int    `json:"creatureId"`
		ItemID     int    `json:"itemId"`
		QualityID  int    `json:"qualityId"`
		Icon       string `json:"icon,omitempty"`
		IsGround   bool   `json:"isGround"`
		IsFlying   bool   `json:"isFlying"`
		IsAquatic  bool   `json:"isAquatic"`
		IsJumping  bool   `json:"isJumping"`
	} `json:"mounts"`
}
