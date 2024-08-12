package wowgd

// ItemClassesIndex structure
type ItemClassesIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ItemClasses []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item_classes"`
}

// ItemClass structure
type ItemClass struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ClassID        int    `json:"class_id"`
	Name           string `json:"name"`
	ItemSubclasses []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item_subclasses"`
}

// ItemSetsIndex structure
type ItemSetsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ItemSets []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item_sets"`
}

// ItemSet structure
type ItemSet struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"items"`
	Effects []struct {
		DisplayString string `json:"display_string"`
		RequiredCount int    `json:"required_count"`
	} `json:"effects"`
	IsEffectActive bool `json:"is_effect_active"`
}

// ItemSubclass structure
type ItemSubclass struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ClassID                int    `json:"class_id"`
	SubclassID             int    `json:"subclass_id"`
	DisplayName            string `json:"display_name"`
	HideSubclassInTooltips bool   `json:"hide_subclass_in_tooltips"`
}

// Item structure
type Item struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Quality struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"quality"`
	Level         int `json:"level"`
	RequiredLevel int `json:"required_level"`
	Media         struct {
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
	PurchasePrice int  `json:"purchase_price"`
	SellPrice     int  `json:"sell_price"`
	MaxCount      int  `json:"max_count"`
	IsEquippable  bool `json:"is_equippable"`
	IsStackable   bool `json:"is_stackable"`
	PreviewItem   struct {
		Item struct {
			Key struct {
				Href string `json:"href"`
			} `json:"key"`
			ID int `json:"id"`
		} `json:"item"`
		Quality struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"quality"`
		Name  string `json:"name"`
		Media struct {
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
		UniqueEquipped string `json:"unique_equipped"`
		Spells         []struct {
			Spell struct {
				Key struct {
					Href string `json:"href"`
				} `json:"key"`
				Name string `json:"name"`
				ID   int    `json:"id"`
			} `json:"spell"`
			Description string `json:"description"`
		} `json:"spells"`
		IsSubclassHidden bool `json:"is_subclass_hidden"`
		NameDescription  *struct {
			DisplayString string `json:"display_string"`
		} `json:"name_description"`
	} `json:"preview_item"`
	PurchaseQuantity int `json:"purchase_quantity"`
}

// ItemMedia structure
type ItemMedia struct {
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

// ItemSearch structure
type ItemSearch struct {
	Page              int  `json:"page"`
	PageSize          int  `json:"pageSize"`
	MaxPageSize       int  `json:"maxPageSize"`
	PageCount         int  `json:"pageCount"`
	ResultCountCapped bool `json:"resultCountCapped"`
	Results           []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			Level         int `json:"level"`
			RequiredLevel int `json:"required_level"`
			SellPrice     int `json:"sell_price"`
			ItemSubclass  struct {
				Name struct {
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"item_subclass"`
			IsEquippable     bool `json:"is_equippable"`
			PurchaseQuantity int  `json:"purchase_quantity"`
			Media            struct {
				ID int `json:"id"`
			} `json:"media"`
			ItemClass struct {
				Name struct {
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				ID int `json:"id"`
			} `json:"item_class"`
			Name struct {
				RuRU string `json:"ru_RU"`
				EnGB string `json:"en_GB"`
				ZhTW string `json:"zh_TW"`
				KoKR string `json:"ko_KR"`
				EnUS string `json:"en_US"`
				EsMX string `json:"es_MX"`
				PtBR string `json:"pt_BR"`
				EsES string `json:"es_ES"`
				ZhCN string `json:"zh_CN"`
				FrFR string `json:"fr_FR"`
				DeDE string `json:"de_DE"`
			} `json:"name"`
			Quality struct {
				Type string `json:"type"`
				Name struct {
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
			} `json:"quality"`
			InventoryType struct {
				Type string `json:"type"`
				Name struct {
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
			} `json:"inventory_type"`
			PurchasePrice int  `json:"purchase_price"`
			MaxCount      int  `json:"max_count"`
			IsStackable   bool `json:"is_stackable"`
			ID            int  `json:"id"`
		} `json:"data"`
	} `json:"results"`
}

type ItemAppearance struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int `json:"id"`
	Slot struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"slot"`
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
	ItemDisplayInfoID int `json:"item_display_info_id"`
	Items             []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"items"`
	Media struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}

type ItemAppearanceSearch struct {
	Page              int  `json:"page"`
	PageSize          int  `json:"pageSize"`
	MaxPageSize       int  `json:"maxPageSize"`
	PageCount         int  `json:"pageCount"`
	ResultCountCapped bool `json:"resultCountCapped"`
	Results           []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Data struct {
			ItemDisplayInfoID int `json:"item_display_info_id"`
			ID                int `json:"id"`
			Slot              struct {
				Name struct {
					ItIT string `json:"it_IT"`
					RuRU string `json:"ru_RU"`
					EnGB string `json:"en_GB"`
					ZhTW string `json:"zh_TW"`
					KoKR string `json:"ko_KR"`
					EnUS string `json:"en_US"`
					EsMX string `json:"es_MX"`
					PtBR string `json:"pt_BR"`
					EsES string `json:"es_ES"`
					ZhCN string `json:"zh_CN"`
					FrFR string `json:"fr_FR"`
					DeDE string `json:"de_DE"`
				} `json:"name"`
				Type string `json:"type"`
			} `json:"slot"`
		} `json:"data"`
	} `json:"results"`
}

type ItemAppearanceSetsIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	AppearanceSets []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"appearance_sets"`
}

type ItemAppearanceSet struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID          int    `json:"id"`
	SetName     string `json:"set_name"`
	Appearances []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"appearances"`
}

type ItemAppearanceSlotIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Slots []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
	} `json:"slots"`
}

type ItemAppearanceSlot struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Appearances []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"appearances"`
}
