package d3c

// Follower structure
type Follower struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	RealName string `json:"realName"`
	Portrait string `json:"portrait"`
	Skills   []struct {
		Slug            string `json:"slug"`
		Name            string `json:"name"`
		Icon            string `json:"icon"`
		Level           int    `json:"level"`
		TooltipURL      string `json:"tooltipUrl"`
		Description     string `json:"description"`
		DescriptionHTML string `json:"descriptionHtml"`
	} `json:"skills"`
}
