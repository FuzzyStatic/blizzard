package d3c

// Hero structure
type Hero struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	MaleName        string `json:"maleName"`
	FemaleName      string `json:"femaleName"`
	Icon            string `json:"icon"`
	SkillCategories []struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"skillCategories"`
	Skills struct {
		Active []struct {
			Slug            string `json:"slug"`
			Name            string `json:"name"`
			Icon            string `json:"icon"`
			Level           int    `json:"level"`
			TooltipURL      string `json:"tooltipUrl"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"active"`
		Passive []struct {
			Slug            string `json:"slug"`
			Name            string `json:"name"`
			Icon            string `json:"icon"`
			Level           int    `json:"level"`
			TooltipURL      string `json:"tooltipUrl"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
			FlavorText      string `json:"flavorText"`
		} `json:"passive"`
	} `json:"skills"`
}

// Skill structure
type Skill struct {
	Skill struct {
		Slug            string `json:"slug"`
		Name            string `json:"name"`
		Icon            string `json:"icon"`
		Level           int    `json:"level"`
		TooltipURL      string `json:"tooltipUrl"`
		Description     string `json:"description"`
		DescriptionHTML string `json:"descriptionHtml"`
	} `json:"skill"`
	Runes []struct {
		Slug            string `json:"slug"`
		Type            string `json:"type"`
		Name            string `json:"name"`
		Level           int    `json:"level"`
		Description     string `json:"description"`
		DescriptionHTML string `json:"descriptionHtml"`
	} `json:"runes"`
}
