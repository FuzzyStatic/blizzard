package wowp

// ChildCriteria struct used for representing nested Achievement criteria
type ChildCriteria []struct {
	ID            int           `json:"id"`
	Amount        float64       `json:"amount"`
	IsCompleted   bool          `json:"is_completed"`
	ChildCriteria ChildCriteria `json:"child_criteria,omitempty"`
}

// CharacterAchievementsSummary structure
type CharacterAchievementsSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	TotalQuantity int `json:"total_quantity"`
	TotalPoints   int `json:"total_points"`
	Achievements  []struct {
		ID          int `json:"id,omitempty"`
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		Criteria struct {
			ID            int           `json:"id"`
			Amount        float64       `json:"amount"`
			IsCompleted   bool          `json:"is_completed"`
			ChildCriteria ChildCriteria `json:"child_criteria,omitempty"`
		} `json:"criteria,omitempty"`
		CompleteTimestamp int64 `json:"completed_timestamp"`
	} `json:"achievements"`
	CategoryProgress []struct {
		Category struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"category"`
		Quantity int `json:"quantity"`
		Points   int `json:"points"`
	} `json:"category_progress"`
	RecentEvents []struct {
		Achievement struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"achievement"`
		Timestamp int64 `json:"timestamp"`
	} `json:"recent_events"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
}

// Statistics struct used for representing nested Statistics
type Statistics []struct {
	ID                   int     `json:"id"`
	Name                 string  `json:"name"`
	LastUpdatedTimestamp int64   `json:"last_updated_timestamp"`
	Quantity             float64 `json:"quantity"`
	Description          string  `json:"description,omitempty"`
}

// CharacterAchievementsStatistics structure
type CharacterAchievementsStatistics struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Character struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Realm struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
			Slug string `json:"slug"`
		} `json:"realm"`
	} `json:"character"`
	Categories []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SubCategories []struct {
			ID         int        `json:"id"`
			Name       string     `json:"name"`
			Statistics Statistics `json:"statistics"`
		} `json:"sub_categories,omitempty"`
		Statistics Statistics `json:"statistics,omitempty"`
	} `json:"categories"`
}
