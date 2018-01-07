/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:59
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:01:21
 */

package diablo3

import "encoding/json"

// Hero structure
type Hero struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Class  string `json:"class"`
	Gender int    `json:"gender"`
	Level  int    `json:"level"`
	Kills  struct {
		Elites int `json:"elites"`
	} `json:"kills"`
	ParagonLevel  int  `json:"paragonLevel"`
	Hardcore      bool `json:"hardcore"`
	Seasonal      bool `json:"seasonal"`
	SeasonCreated int  `json:"seasonCreated"`
	Skills        struct {
		Active []struct {
			Skill struct {
				Slug              string `json:"slug"`
				Name              string `json:"name"`
				Icon              string `json:"icon"`
				Level             int    `json:"level"`
				CategorySlug      string `json:"categorySlug"`
				TooltipURL        string `json:"tooltipUrl"`
				Description       string `json:"description"`
				SimpleDescription string `json:"simpleDescription"`
				SkillCalcID       string `json:"skillCalcId"`
			} `json:"skill"`
			Rune struct {
				Slug              string `json:"slug"`
				Type              string `json:"type"`
				Name              string `json:"name"`
				Level             int    `json:"level"`
				Description       string `json:"description"`
				SimpleDescription string `json:"simpleDescription"`
				TooltipParams     string `json:"tooltipParams"`
				SkillCalcID       string `json:"skillCalcId"`
				Order             int    `json:"order"`
			} `json:"rune"`
		} `json:"active"`
		Passive []struct {
			Skill struct {
				Slug        string `json:"slug"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Level       int    `json:"level"`
				TooltipURL  string `json:"tooltipUrl"`
				Description string `json:"description"`
				Flavor      string `json:"flavor"`
				SkillCalcID string `json:"skillCalcId"`
			} `json:"skill"`
		} `json:"passive"`
	} `json:"skills"`
	Items struct {
		Head struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"head"`
		Torso struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"torso"`
		Feet struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"feet"`
		Hands struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"hands"`
		Shoulders struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"shoulders"`
		Legs struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			DyeColor      struct {
				Item struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"item"`
			} `json:"dyeColor"`
			TransmogItem struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"legs"`
		Bracers struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"bracers"`
		MainHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			TransmogItem  struct {
				ID               string   `json:"id"`
				Name             string   `json:"name"`
				Icon             string   `json:"icon"`
				DisplayColor     string   `json:"displayColor"`
				TooltipParams    string   `json:"tooltipParams"`
				SetItemsEquipped []string `json:"setItemsEquipped"`
			} `json:"transmogItem"`
		} `json:"mainHand"`
		OffHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			TransmogItem  struct {
				ID               string   `json:"id"`
				Name             string   `json:"name"`
				Icon             string   `json:"icon"`
				DisplayColor     string   `json:"displayColor"`
				TooltipParams    string   `json:"tooltipParams"`
				SetItemsEquipped []string `json:"setItemsEquipped"`
			} `json:"transmogItem"`
		} `json:"offHand"`
		RightFinger struct {
			ID               string   `json:"id"`
			Name             string   `json:"name"`
			Icon             string   `json:"icon"`
			DisplayColor     string   `json:"displayColor"`
			TooltipParams    string   `json:"tooltipParams"`
			SetItemsEquipped []string `json:"setItemsEquipped"`
		} `json:"rightFinger"`
		LeftFinger struct {
			ID               string   `json:"id"`
			Name             string   `json:"name"`
			Icon             string   `json:"icon"`
			DisplayColor     string   `json:"displayColor"`
			TooltipParams    string   `json:"tooltipParams"`
			SetItemsEquipped []string `json:"setItemsEquipped"`
		} `json:"leftFinger"`
	} `json:"items"`
	Followers struct {
	} `json:"followers"`
	LegendaryPowers []struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
	} `json:"legendaryPowers"`
	Stats struct {
		Life              int     `json:"life"`
		Damage            float64 `json:"damage"`
		Toughness         float64 `json:"toughness"`
		Healing           float64 `json:"healing"`
		AttackSpeed       float64 `json:"attackSpeed"`
		Armor             int     `json:"armor"`
		Strength          int     `json:"strength"`
		Dexterity         int     `json:"dexterity"`
		Vitality          int     `json:"vitality"`
		Intelligence      int     `json:"intelligence"`
		PhysicalResist    int     `json:"physicalResist"`
		FireResist        int     `json:"fireResist"`
		ColdResist        int     `json:"coldResist"`
		LightningResist   int     `json:"lightningResist"`
		PoisonResist      int     `json:"poisonResist"`
		ArcaneResist      int     `json:"arcaneResist"`
		CritDamage        float64 `json:"critDamage"`
		BlockChance       float64 `json:"blockChance"`
		BlockAmountMin    int     `json:"blockAmountMin"`
		BlockAmountMax    int     `json:"blockAmountMax"`
		DamageIncrease    float64 `json:"damageIncrease"`
		CritChance        float64 `json:"critChance"`
		DamageReduction   float64 `json:"damageReduction"`
		Thorns            float64 `json:"thorns"`
		LifeSteal         float64 `json:"lifeSteal"`
		LifePerKill       float64 `json:"lifePerKill"`
		GoldFind          float64 `json:"goldFind"`
		MagicFind         float64 `json:"magicFind"`
		LifeOnHit         float64 `json:"lifeOnHit"`
		PrimaryResource   int     `json:"primaryResource"`
		SecondaryResource int     `json:"secondaryResource"`
	} `json:"stats"`
	Progression struct {
		Act1 struct {
			Completed       bool          `json:"completed"`
			CompletedQuests []interface{} `json:"completedQuests"`
		} `json:"act1"`
		Act2 struct {
			Completed       bool          `json:"completed"`
			CompletedQuests []interface{} `json:"completedQuests"`
		} `json:"act2"`
		Act3 struct {
			Completed       bool          `json:"completed"`
			CompletedQuests []interface{} `json:"completedQuests"`
		} `json:"act3"`
		Act4 struct {
			Completed       bool          `json:"completed"`
			CompletedQuests []interface{} `json:"completedQuests"`
		} `json:"act4"`
		Act5 struct {
			Completed       bool          `json:"completed"`
			CompletedQuests []interface{} `json:"completedQuests"`
		} `json:"act5"`
	} `json:"progression"`
	Dead        bool   `json:"dead"`
	LastUpdated int    `json:"last-updated"`
	Code        string `json:"code"`
	Reason      string `json:"reason"`
}

// JSON2Struct creates Hero structure from JSON byte array
func (h *Hero) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, h)
}

// GetAllItemTooltipParams returns a map with all the item tooltipParams from a hero
func (h *Hero) GetAllItemTooltipParams() map[string]string {
	return map[string]string{
		"Bracers":     h.Items.Bracers.TooltipParams,
		"Feet":        h.Items.Feet.TooltipParams,
		"Hands":       h.Items.Hands.TooltipParams,
		"Head":        h.Items.Head.TooltipParams,
		"LeftFinger":  h.Items.LeftFinger.TooltipParams,
		"Legs":        h.Items.Legs.TooltipParams,
		"MainHand":    h.Items.MainHand.TooltipParams,
		"OffHand":     h.Items.OffHand.TooltipParams,
		"RightFinger": h.Items.RightFinger.TooltipParams,
		"Shoulders":   h.Items.Shoulders.TooltipParams,
		"Torso":       h.Items.Torso.TooltipParams,
	}
}
