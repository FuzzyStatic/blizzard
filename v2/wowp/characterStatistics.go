package wowp

// CharacterStatisticsSummary structure
type CharacterStatisticsSummary struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Health    int `json:"health"`
	Power     int `json:"power"`
	PowerType struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"power_type"`
	Speed struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
	} `json:"speed"`
	Strength struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"strength"`
	Agility struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"agility"`
	Intellect struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"intellect"`
	Stamina struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"stamina"`
	MeleeCrit struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"melee_crit"`
	MeleeHaste struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"melee_haste"`
	Mastery struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"mastery"`
	BonusArmor int `json:"bonus_armor"`
	Lifesteal  struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"lifesteal"`
	Versatility                 float64 `json:"versatility"`
	VersatilityDamageDoneBonus  float64 `json:"versatility_damage_done_bonus"`
	VersatilityHealingDoneBonus float64 `json:"versatility_healing_done_bonus"`
	VersatilityDamageTakenBonus float64 `json:"versatility_damage_taken_bonus"`
	Avoidance                   struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
	} `json:"avoidance"`
	AttackPower       int     `json:"attack_power"`
	MainHandDamageMin float64 `json:"main_hand_damage_min"`
	MainHandDamageMax float64 `json:"main_hand_damage_max"`
	MainHandSpeed     float64 `json:"main_hand_speed"`
	MainHandDps       float64 `json:"main_hand_dps"`
	OffHandDamageMin  float64 `json:"off_hand_damage_min"`
	OffHandDamageMax  float64 `json:"off_hand_damage_max"`
	OffHandSpeed      float64 `json:"off_hand_speed"`
	OffHandDps        float64 `json:"off_hand_dps"`
	SpellPower        int     `json:"spell_power"`
	SpellPenetration  int     `json:"spell_penetration"`
	SpellCrit         struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"spell_crit"`
	ManaRegen       float64 `json:"mana_regen"`
	ManaRegenCombat float64 `json:"mana_regen_combat"`
	Armor           struct {
		Base      int `json:"base"`
		Effective int `json:"effective"`
	} `json:"armor"`
	Dodge struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"dodge"`
	Parry struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"parry"`
	Block struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"block"`
	RangedCrit struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"ranged_crit"`
	RangedHaste struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"ranged_haste"`
	SpellHaste struct {
		Rating      int     `json:"rating"`
		RatingBonus float64 `json:"rating_bonus"`
		Value       float64 `json:"value"`
	} `json:"spell_haste"`
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
