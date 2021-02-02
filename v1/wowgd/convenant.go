package wowgd

// CovenantsIndex structure
type CovenantsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Covenants []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"covenants"`
}

// Covenant structure
type Covenant struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	SignatureAbility struct {
		ID           int `json:"id"`
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
			Cooldown    string `json:"cooldown"`
		} `json:"spell_tooltip"`
	} `json:"signature_ability"`
	ClassAbilities []struct {
		ID            int `json:"id"`
		PlayableClass struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"playable_class"`
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
			PowerCost   string `json:"power_cost"`
			Range       string `json:"range"`
			Cooldown    string `json:"cooldown"`
		} `json:"spell_tooltip,omitempty"`
	} `json:"class_abilities"`
	Soulbinds []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"soulbinds"`
	RenownRewards []struct {
		Level  int `json:"level"`
		Reward struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"reward"`
	} `json:"renown_rewards"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

// CovenantMedia structure
type CovenantMedia struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Assets []struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		FileDataID int    `json:"file_data_id"`
	} `json:"assets"`
}

// CovenantSoulbindsIndex structure
type CovenantSoulbindsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Soulbinds []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"soulbinds"`
}

// CovenantSoulbind structure
type CovenantSoulbind struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Covenant struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"covenant"`
	Creature struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"creature"`
	Follower struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"follower"`
	TalentTree struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"talent_tree"`
}

// CovenantConduitsIndex structure
type CovenantConduitsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Conduits []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"conduits"`
}

// CovenantConduit structure
type CovenantConduit struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Item struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item"`
	SocketType struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"socket_type"`
	Ranks []struct {
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
		} `json:"spell_tooltip"`
	} `json:"ranks"`
}
