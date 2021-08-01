package wowgd

// ModifiedCraftingIndex structure
type ModifiedCraftingIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Categories struct {
		Href string `json:"href"`
	} `json:"categories"`
	SlotTypes struct {
		Href string `json:"href"`
	} `json:"slot_types"`
}

// ModifiedCraftingCategoryIndex structure
type ModifiedCraftingCategoryIndex struct {
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
}

// ModifiedCraftingCategory structure
type ModifiedCraftingCategory struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ModifiedCraftingReagentSlotTypeIndex structure
type ModifiedCraftingReagentSlotTypeIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	SlotTypes []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name,omitempty"`
		ID   int    `json:"id"`
	} `json:"slot_types"`
}

// ModifiedCraftingReagentSlotType structure
type ModifiedCraftingReagentSlotType struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID                   int    `json:"id"`
	Description          string `json:"description"`
	CompatibleCategories []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"compatible_categories"`
}
