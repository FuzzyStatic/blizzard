/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:50
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:38:50
 */

package diablo3

import (
	"encoding/json"
)

// Profile structure
type Profile struct {
	BattleTag                  string `json:"battleTag"`
	ParagonLevel               int    `json:"paragonLevel"`
	ParagonLevelHardcore       int    `json:"paragonLevelHardcore"`
	ParagonLevelSeason         int    `json:"paragonLevelSeason"`
	ParagonLevelSeasonHardcore int    `json:"paragonLevelSeasonHardcore"`
	GuildName                  string `json:"guildName"`
	Heroes                     []struct {
		Heroes
	} `json:"heroes"`
	LastHeroPlayed int `json:"lastHeroPlayed"`
	LastUpdated    int `json:"lastUpdated"`
	Kills          struct {
		Monsters         int `json:"monsters"`
		Elites           int `json:"elites"`
		HardcoreMonsters int `json:"hardcoreMonsters"`
	} `json:"kills"`
	HighestHardcoreLevel int `json:"highestHardcoreLevel"`
	TimePlayed           struct {
		Barbarian   float64 `json:"barbarian"`
		Crusader    float64 `json:"crusader"`
		DemonHunter float64 `json:"demon-hunter"`
		Monk        float64 `json:"monk"`
		WitchDoctor float64 `json:"witch-doctor"`
		Wizard      float64 `json:"wizard"`
	} `json:"timePlayed"`
	Progression struct {
		Act1 bool `json:"act1"`
		Act2 bool `json:"act2"`
		Act3 bool `json:"act3"`
		Act4 bool `json:"act4"`
		Act5 bool `json:"act5"`
	} `json:"progression"`
	FallenHeroes []struct {
		Name  string `json:"name"`
		Level int    `json:"level"`
		Stats struct {
			Life            int     `json:"life"`
			Damage          float64 `json:"damage"`
			Toughness       float64 `json:"toughness"`
			Healing         float64 `json:"healing"`
			Armor           int     `json:"armor"`
			Strength        int     `json:"strength"`
			Dexterity       int     `json:"dexterity"`
			Vitality        int     `json:"vitality"`
			Intelligence    int     `json:"intelligence"`
			PhysicalResist  int     `json:"physicalResist"`
			FireResist      int     `json:"fireResist"`
			ColdResist      int     `json:"coldResist"`
			LightningResist int     `json:"lightningResist"`
			PoisonResist    int     `json:"poisonResist"`
			ArcaneResist    int     `json:"arcaneResist"`
			DamageIncrease  float64 `json:"damageIncrease"`
			CritChance      float64 `json:"critChance"`
			DamageReduction float64 `json:"damageReduction"`
		} `json:"stats"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
				SetItemsEquipped []string `json:"setItemsEquipped"`
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
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"transmogItem"`
			} `json:"mainHand"`
			OffHand struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
				TransmogItem  struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"transmogItem"`
			} `json:"offHand"`
			Waist struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"waist"`
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
			Neck struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"neck"`
		} `json:"items"`
		Kills struct {
			Elites int `json:"elites"`
		} `json:"kills"`
		Hardcore bool `json:"hardcore"`
		Death    struct {
			Killer   int `json:"killer"`
			Time     int `json:"time"`
			Location int `json:"location"`
		} `json:"death"`
		HeroID int    `json:"heroId"`
		Gender int    `json:"gender"`
		Class  string `json:"class"`
	} `json:"fallenHeroes"`
	Blacksmith struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"blacksmith"`
	Jeweler struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"jeweler"`
	Mystic struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"mystic"`
	BlacksmithHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"blacksmithHardcore"`
	JewelerHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"jewelerHardcore"`
	MysticHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"mysticHardcore"`
	BlacksmithSeason struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"blacksmithSeason"`
	JewelerSeason struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"jewelerSeason"`
	MysticSeason struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"mysticSeason"`
	BlacksmithSeasonHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"blacksmithSeasonHardcore"`
	JewelerSeasonHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"jewelerSeasonHardcore"`
	MysticSeasonHardcore struct {
		Slug        string `json:"slug"`
		Level       int    `json:"level"`
		StepCurrent int    `json:"stepCurrent"`
		StepMax     int    `json:"stepMax"`
	} `json:"mysticSeasonHardcore"`
	SeasonalProfiles struct {
		Season4 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season4"`
		Season3 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season3"`
		Season6 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season6"`
		Season5 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season5"`
		Season8 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season8"`
		Season7 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season7"`
		Season9 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season9"`
		Season0 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season0"`
		Season1 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season1"`
		Season2 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				Barbarian   float64 `json:"barbarian"`
				Crusader    float64 `json:"crusader"`
				DemonHunter float64 `json:"demon-hunter"`
				Monk        float64 `json:"monk"`
				WitchDoctor float64 `json:"witch-doctor"`
				Wizard      float64 `json:"wizard"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
			Progression          struct {
				Act1 bool `json:"act1"`
				Act2 bool `json:"act2"`
				Act3 bool `json:"act3"`
				Act4 bool `json:"act4"`
				Act5 bool `json:"act5"`
			} `json:"progression"`
		} `json:"season2"`
	} `json:"seasonalProfiles"`
}

// Heroes structure contains all heroes for profile
type Heroes struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Class  string `json:"class"`
	Gender int    `json:"gender"`
	Level  int    `json:"level"`
	Kills  struct {
		Elites int `json:"elites"`
	} `json:"kills"`
	ParagonLevel int  `json:"paragonLevel"`
	Hardcore     bool `json:"hardcore"`
	Seasonal     bool `json:"seasonal"`
	LastUpdated  int  `json:"last-updated"`
	Dead         bool `json:"dead"`
}

// JSON2Struct creates Profile structure from JSON byte array
func (p *Profile) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, p)
}

// GetHeroNamesFromProfile get all hero names from profile
func (p *Profile) GetHeroNamesFromProfile() *[]string {
	var heroes []string

	for _, v := range (*p).Heroes {
		heroes = append(heroes, v.Name)
	}

	return &heroes
}

// GetHeroIDsAndNamesFromProfile get all hero ids and names from profile
func (p *Profile) GetHeroIDsAndNamesFromProfile() *map[int]string {
	m := make(map[int]string)

	for _, v := range (*p).Heroes {
		m[v.ID] = v.Name
	}

	return &m
}
