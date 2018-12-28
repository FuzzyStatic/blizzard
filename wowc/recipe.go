package wowc

// Recipe structure
type Recipe struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Icon       string `json:"icon"`
}
