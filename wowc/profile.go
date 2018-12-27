package wowc

// Profile structure
type Profile struct {
	Characters []struct {
		Name              string `json:"name"`
		Realm             string `json:"realm"`
		Battlegroup       string `json:"battlegroup"`
		Class             int    `json:"class"`
		Race              int    `json:"race"`
		Gender            int    `json:"gender"`
		Level             int    `json:"level"`
		AchievementPoints int    `json:"achievementPoints"`
		Thumbnail         string `json:"thumbnail"`
		Guild             string `json:"guild,omitempty"`
		GuildRealm        string `json:"guildRealm,omitempty"`
		LastModified      int64  `json:"lastModified"`
		Spec              struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"spec,omitempty"`
	} `json:"characters"`
}
