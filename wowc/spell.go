package wowc

// Spell structure
type Spell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Range       string `json:"range"`
	CastTime    string `json:"castTime"`
}
