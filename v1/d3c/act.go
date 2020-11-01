package d3c

// ActIndex contains slice of Acts
type ActIndex struct {
	Acts []Act `json:"acts"`
}

// Act contains act data
type Act struct {
	Slug   string `json:"slug"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Quests []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"quests"`
}
