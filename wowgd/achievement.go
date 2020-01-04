package wowgd

// AchievementCategoriesIndex structure
type AchievementCategoriesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Categories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"categories"`
	RootCategories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"root_categories"`
	GuildCategories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"guild_categories"`
}

// AchievementCategory structure
type AchievementCategory struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Achievements []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"achievements"`
	Subcategories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"subcategories"`
	IsGuildCategory     bool `json:"is_guild_category"`
	AggregatesByFaction struct {
		Alliance struct {
			Quantity int `json:"quantity"`
			Points   int `json:"points"`
		} `json:"alliance"`
		Horde struct {
			Quantity int `json:"quantity"`
			Points   int `json:"points"`
		} `json:"horde"`
	} `json:"aggregates_by_faction"`
	DisplayOrder int `json:"display_order"`
}

// AchievementIndex structure
type AchievementIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Achievements []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"achievements"`
}

// Achievement structure
type Achievement struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID       int `json:"id"`
	Category struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"category"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Points        int    `json:"points"`
	IsAccountWide bool   `json:"is_account_wide"`
	Criteria      struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		Amount      int    `json:"amount"`
	} `json:"criteria"`
	NextAchievement struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"next_achievement"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
	DisplayOrder int `json:"display_order"`
}

// AchievementMedia structure
type AchievementMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"assets"`
}
