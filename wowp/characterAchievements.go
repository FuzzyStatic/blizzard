package wowp

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
			ID            int     `json:"id"`
			Amount        float64 `json:"amount"`
			IsCompleted   bool    `json:"is_completed"`
			ChildCriteria []struct {
				ID          int     `json:"id"`
				Amount      float64 `json:"amount"`
				IsCompleted bool    `json:"is_completed"`
			} `json:"child_criteria"`
		} `json:"criteria,omitempty"`
		CompleteTimestamp int64 `json:"completed_timestamp"`
	} `json:"achievements"`
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
