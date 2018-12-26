package d3c

// Artisan structure
type Artisan struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Portrait string `json:"portrait"`
	Training struct {
		Tiers []struct {
			Tier           int `json:"tier"`
			TrainedRecipes []struct {
				ID       string `json:"id"`
				Slug     string `json:"slug"`
				Name     string `json:"name"`
				Cost     int    `json:"cost"`
				Reagents []struct {
					Quantity int `json:"quantity"`
					Item     struct {
						ID   string `json:"id"`
						Slug string `json:"slug"`
						Name string `json:"name"`
						Icon string `json:"icon"`
						Path string `json:"path"`
					} `json:"item"`
				} `json:"reagents"`
				ItemProduced struct {
					ID   string `json:"id"`
					Slug string `json:"slug"`
					Name string `json:"name"`
					Icon string `json:"icon"`
					Path string `json:"path"`
				} `json:"itemProduced"`
			} `json:"trainedRecipes"`
			TaughtRecipes []struct {
				ID       string `json:"id"`
				Slug     string `json:"slug"`
				Name     string `json:"name"`
				Cost     int    `json:"cost"`
				Reagents []struct {
					Quantity int `json:"quantity"`
					Item     struct {
						ID   string `json:"id"`
						Slug string `json:"slug"`
						Name string `json:"name"`
						Icon string `json:"icon"`
						Path string `json:"path"`
					} `json:"item"`
				} `json:"reagents"`
				ItemProduced struct {
					ID   string `json:"id"`
					Slug string `json:"slug"`
					Name string `json:"name"`
					Icon string `json:"icon"`
					Path string `json:"path"`
				} `json:"itemProduced"`
			} `json:"taughtRecipes"`
		} `json:"tiers"`
	} `json:"training"`
}

// Recipe artisan recipe structure
type Recipe struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Reagents []struct {
		Quantity int `json:"quantity"`
		Item     struct {
			ID   string `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
			Icon string `json:"icon"`
			Path string `json:"path"`
		} `json:"item"`
	} `json:"reagents"`
	ItemProduced struct {
		ID   string `json:"id"`
		Slug string `json:"slug"`
		Name string `json:"name"`
		Icon string `json:"icon"`
		Path string `json:"path"`
	} `json:"itemProduced"`
}
