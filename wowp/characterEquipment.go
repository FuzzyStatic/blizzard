package wowp

// CharacterEquipmentSummary structure
type CharacterEquipmentSummary struct {
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
	EquippedItems []struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"item"`
		Slot struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"slot"`
		Quantity  int   `json:"quantity"`
		Context   int   `json:"context"`
		BonusList []int `json:"bonus_list,omitempty"`
		Quality   struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Name                 string `json:"name"`
		ModifiedAppearanceID int    `json:"modified_appearance_id,omitempty"`
		AzeriteDetails       struct {
			PercentageToNextLevel float64 `json:"percentage_to_next_level"`
			Level                 struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"level"`
			SelectedPowers []struct {
				ID           int `json:"id"`
				Tier         int `json:"tier"`
				SpellTooltip struct {
					Spell struct {
						Key struct {
							Href string `json:"href"`
						} `json:"key"`
						Name string `json:"name"`
						ID   int    `json:"id"`
					} `json:"spell"`
					Description string `json:"description"`
					CastTime    string `json:"cast_time"`
				} `json:"spell_tooltip,omitempty"`
			} `json:"selected_powers"`
			SelectedPowersString string `json:"selected_powers_string"`
		} `json:"azerite_details,omitempty"`
		NameDescription string `json:"name_description,omitempty"`
		Media           struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"media"`
		ItemClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_class"`
		ItemSubclass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"item_subclass"`
		InventoryType struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"inventory_type"`
		Binding struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"binding"`
		Armor struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"armor,omitempty"`
		Stats []struct {
			Type struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"type"`
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
			IsNegated     bool   `json:"is_negated,omitempty"`
		} `json:"stats"`
		Level struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"level"`
		Transmog struct {
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item"`
			DisplayString            string `json:"display_string"`
			ItemModifiedAppearanceID int    `json:"item_modified_appearance_id"`
		} `json:"transmog,omitempty"`
		Durability struct {
			Value         int    `json:"value"`
			DisplayString string `json:"display_string"`
		} `json:"durability,omitempty"`
		UniqueEquipped string `json:"unique_equipped,omitempty"`
		Spells         []struct {
			Spell struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"spell"`
			Description string `json:"description"`
		} `json:"spells,omitempty"`
		Description      string `json:"description,omitempty"`
		IsSubclassHidden bool   `json:"is_subclass_hidden,omitempty"`
		SellPrice        struct {
			Value          int `json:"value"`
			DisplayStrings struct {
				Header string `json:"header"`
				Gold   string `json:"gold"`
				Silver string `json:"silver"`
				Copper string `json:"copper"`
			} `json:"display_strings"`
		} `json:"sell_price,omitempty"`
		Enchantments []struct {
			DisplayString string `json:"display_string"`
			SourceItem    struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"source_item"`
			EnchantmentID   int `json:"enchantment_id"`
			EnchantmentSlot struct {
				ID   int    `json:"id"`
				Type string `json:"type"`
			} `json:"enchantment_slot"`
		} `json:"enchantments,omitempty"`
		LimitCategory string `json:"limit_category,omitempty"`
		Requirements  struct {
			Level struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"level"`
			Skill struct {
				Profession struct {
					Name string `json:"name"`
					ID   int    `json:"id"`
				} `json:"profession"`
				Level         int    `json:"level"`
				DisplayString string `json:"display_string"`
			} `json:"skill"`
		} `json:"requirements,omitempty"`
		Sockets []struct {
			SocketType struct {
				Type string `json:"type"`
				Name string `json:"name"`
			} `json:"socket_type"`
			Item struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"item"`
			DisplayString string `json:"display_string"`
			Media         struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				ID int `json:"id"`
			} `json:"media"`
		} `json:"sockets,omitempty"`
		Weapon struct {
			Damage struct {
				MinValue      int    `json:"min_value"`
				MaxValue      int    `json:"max_value"`
				DisplayString string `json:"display_string"`
				DamageClass   struct {
					Type string `json:"type"`
					Name string `json:"name"`
				} `json:"damage_class"`
			} `json:"damage"`
			AttackSpeed struct {
				Value         int    `json:"value"`
				DisplayString string `json:"display_string"`
			} `json:"attack_speed"`
			Dps struct {
				Value         float64 `json:"value"`
				DisplayString string  `json:"display_string"`
			} `json:"dps"`
		} `json:"weapon,omitempty"`
	} `json:"equipped_items"`
}
