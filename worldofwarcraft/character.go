/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-11 22:25:18
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-13 22:21:43
 */

package worldofwarcraft

import "encoding/json"

// Character structure
type Character struct {
	LastModified      int64  `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Class             int    `json:"class"`
	Race              int    `json:"race"`
	Gender            int    `json:"gender"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievementPoints"`
	Thumbnail         string `json:"thumbnail"`
	CalcClass         string `json:"calcClass"`
	Faction           int    `json:"faction"`
	Achievements      struct {
		AchievementsCompleted          []int   `json:"achievementsCompleted"`
		AchievementsCompletedTimestamp []int64 `json:"achievementsCompletedTimestamp"`
		Criteria                       []int   `json:"criteria"`
		CriteriaQuantity               []int64 `json:"criteriaQuantity"`
		CriteriaTimestamp              []int64 `json:"criteriaTimestamp"`
		CriteriaCreated                []int64 `json:"criteriaCreated"`
	} `json:"achievements"`
	Appearance struct {
		FaceVariation        int   `json:"faceVariation"`
		SkinColor            int   `json:"skinColor"`
		HairVariation        int   `json:"hairVariation"`
		HairColor            int   `json:"hairColor"`
		FeatureVariation     int   `json:"featureVariation"`
		ShowHelm             bool  `json:"showHelm"`
		ShowCloak            bool  `json:"showCloak"`
		CustomDisplayOptions []int `json:"customDisplayOptions"`
	} `json:"appearance"`
	Feed []struct {
		Type        string `json:"type"`
		Timestamp   int64  `json:"timestamp"`
		Achievement struct {
			ID          int           `json:"id"`
			Title       string        `json:"title"`
			Points      int           `json:"points"`
			Description string        `json:"description"`
			RewardItems []interface{} `json:"rewardItems"`
			Icon        string        `json:"icon"`
			Criteria    []struct {
				ID          int    `json:"id"`
				Description string `json:"description"`
				OrderIndex  int    `json:"orderIndex"`
				Max         int    `json:"max"`
			} `json:"criteria"`
			AccountWide bool `json:"accountWide"`
			FactionID   int  `json:"factionId"`
		} `json:"achievement,omitempty"`
		FeatOfStrength bool   `json:"featOfStrength,omitempty"`
		ItemID         int    `json:"itemId,omitempty"`
		Context        string `json:"context,omitempty"`
		BonusLists     []int  `json:"bonusLists,omitempty"`
		Criteria       struct {
			ID          int    `json:"id"`
			Description string `json:"description"`
			OrderIndex  int    `json:"orderIndex"`
			Max         int    `json:"max"`
		} `json:"criteria,omitempty"`
		Quantity int    `json:"quantity,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"feed"`
	Guild struct {
		Name              string `json:"name"`
		Realm             string `json:"realm"`
		Battlegroup       string `json:"battlegroup"`
		Members           int    `json:"members"`
		AchievementPoints int    `json:"achievementPoints"`
		Emblem            struct {
			Icon              int    `json:"icon"`
			IconColor         string `json:"iconColor"`
			IconColorID       int    `json:"iconColorId"`
			Border            int    `json:"border"`
			BorderColor       string `json:"borderColor"`
			BorderColorID     int    `json:"borderColorId"`
			BackgroundColor   string `json:"backgroundColor"`
			BackgroundColorID int    `json:"backgroundColorId"`
		} `json:"emblem"`
	} `json:"guild"`
	HunterPets []struct {
		Name     string `json:"name"`
		Creature int    `json:"creature"`
		Selected bool   `json:"selected"`
		Slot     int    `json:"slot"`
		Spec     struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"spec"`
		CalcSpec   string `json:"calcSpec"`
		FamilyID   int    `json:"familyId"`
		FamilyName string `json:"familyName"`
	} `json:"hunterPets"`
	Items struct {
		AverageItemLevel         int `json:"averageItemLevel"`
		AverageItemLevelEquipped int `json:"averageItemLevelEquipped"`
		Head                     struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"head"`
		Neck struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"neck"`
		Shoulder struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"shoulder"`
		Back struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"back"`
		Chest struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"chest"`
		Shirt struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats                []interface{} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"shirt"`
		Wrist struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"wrist"`
		Hands struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"hands"`
		Waist struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Tinker          int `json:"tinker"`
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"waist"`
		Legs struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"legs"`
		Feet struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TransmogItem    int `json:"transmogItem"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemID                      int `json:"itemId"`
				TransmogItemAppearanceModID int `json:"transmogItemAppearanceModId"`
			} `json:"appearance"`
		} `json:"feet"`
		Finger1 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"finger1"`
		Finger2 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"finger2"`
		Trinket1 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"trinket1"`
		Trinket2 struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor                int           `json:"armor"`
			Context              string        `json:"context"`
			BonusLists           []int         `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
			} `json:"appearance"`
		} `json:"trinket2"`
		MainHand struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				Gem0            int `json:"gem0"`
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor      int `json:"armor"`
			WeaponInfo struct {
				Damage struct {
					Min      int     `json:"min"`
					Max      int     `json:"max"`
					ExactMin float64 `json:"exactMin"`
					ExactMax float64 `json:"exactMax"`
				} `json:"damage"`
				WeaponSpeed float64 `json:"weaponSpeed"`
				Dps         float64 `json:"dps"`
			} `json:"weaponInfo"`
			Context              string `json:"context"`
			BonusLists           []int  `json:"bonusLists"`
			ArtifactID           int    `json:"artifactId"`
			DisplayInfoID        int    `json:"displayInfoId"`
			ArtifactAppearanceID int    `json:"artifactAppearanceId"`
			ArtifactTraits       []struct {
				ID   int `json:"id"`
				Rank int `json:"rank"`
			} `json:"artifactTraits"`
			Relics []struct {
				Socket     int           `json:"socket"`
				ItemID     int           `json:"itemId"`
				Context    int           `json:"context"`
				BonusLists []interface{} `json:"bonusLists"`
			} `json:"relics"`
			Appearance struct {
				ItemAppearanceModID int `json:"itemAppearanceModId"`
			} `json:"appearance"`
		} `json:"mainHand"`
		OffHand struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			Quality       int    `json:"quality"`
			ItemLevel     int    `json:"itemLevel"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
			Stats []struct {
				Stat   int `json:"stat"`
				Amount int `json:"amount"`
			} `json:"stats"`
			Armor      int `json:"armor"`
			WeaponInfo struct {
				Damage struct {
					Min      int     `json:"min"`
					Max      int     `json:"max"`
					ExactMin float64 `json:"exactMin"`
					ExactMax float64 `json:"exactMax"`
				} `json:"damage"`
				WeaponSpeed float64 `json:"weaponSpeed"`
				Dps         float64 `json:"dps"`
			} `json:"weaponInfo"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			ArtifactID           int           `json:"artifactId"`
			DisplayInfoID        int           `json:"displayInfoId"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			Relics               []interface{} `json:"relics"`
			Appearance           struct {
				ItemAppearanceModID int `json:"itemAppearanceModId"`
			} `json:"appearance"`
		} `json:"offHand"`
	} `json:"items"`
	Mounts struct {
		NumCollected    int `json:"numCollected"`
		NumNotCollected int `json:"numNotCollected"`
		Collected       []struct {
			Name       string `json:"name"`
			SpellID    int    `json:"spellId"`
			CreatureID int    `json:"creatureId"`
			ItemID     int    `json:"itemId"`
			QualityID  int    `json:"qualityId"`
			Icon       string `json:"icon"`
			IsGround   bool   `json:"isGround"`
			IsFlying   bool   `json:"isFlying"`
			IsAquatic  bool   `json:"isAquatic"`
			IsJumping  bool   `json:"isJumping"`
		} `json:"collected"`
	} `json:"mounts"`
	Pets struct {
		NumCollected    int `json:"numCollected"`
		NumNotCollected int `json:"numNotCollected"`
		Collected       []struct {
			Name       string `json:"name"`
			SpellID    int    `json:"spellId"`
			CreatureID int    `json:"creatureId"`
			ItemID     int    `json:"itemId"`
			QualityID  int    `json:"qualityId"`
			Icon       string `json:"icon"`
			Stats      struct {
				SpeciesID    int `json:"speciesId"`
				BreedID      int `json:"breedId"`
				PetQualityID int `json:"petQualityId"`
				Level        int `json:"level"`
				Health       int `json:"health"`
				Power        int `json:"power"`
				Speed        int `json:"speed"`
			} `json:"stats"`
			BattlePetGUID               string `json:"battlePetGuid"`
			IsFavorite                  bool   `json:"isFavorite"`
			IsFirstAbilitySlotSelected  bool   `json:"isFirstAbilitySlotSelected"`
			IsSecondAbilitySlotSelected bool   `json:"isSecondAbilitySlotSelected"`
			IsThirdAbilitySlotSelected  bool   `json:"isThirdAbilitySlotSelected"`
			CreatureName                string `json:"creatureName"`
			CanBattle                   bool   `json:"canBattle"`
		} `json:"collected"`
	} `json:"pets"`
	PetSlots []struct {
		Slot          int    `json:"slot"`
		BattlePetGUID string `json:"battlePetGuid"`
		IsEmpty       bool   `json:"isEmpty"`
		IsLocked      bool   `json:"isLocked"`
		Abilities     []int  `json:"abilities"`
	} `json:"petSlots"`
	Professions struct {
		Primary []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Icon    string `json:"icon"`
			Rank    int    `json:"rank"`
			Max     int    `json:"max"`
			Recipes []int  `json:"recipes"`
		} `json:"primary"`
		Secondary []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Icon    string `json:"icon"`
			Rank    int    `json:"rank"`
			Max     int    `json:"max"`
			Recipes []int  `json:"recipes"`
		} `json:"secondary"`
	} `json:"professions"`
	Progression struct {
		Raids []struct {
			Name   string `json:"name"`
			Lfr    int    `json:"lfr"`
			Normal int    `json:"normal"`
			Heroic int    `json:"heroic"`
			Mythic int    `json:"mythic"`
			ID     int    `json:"id"`
			Bosses []struct {
				ID              int    `json:"id"`
				Name            string `json:"name"`
				NormalKills     int    `json:"normalKills"`
				NormalTimestamp int64  `json:"normalTimestamp"`
			} `json:"bosses"`
		} `json:"raids"`
	} `json:"progression"`
	PVP struct {
		Brackets struct {
			ARENABRACKET2V2 struct {
				Slug         string `json:"slug"`
				Rating       int    `json:"rating"`
				WeeklyPlayed int    `json:"weeklyPlayed"`
				WeeklyWon    int    `json:"weeklyWon"`
				WeeklyLost   int    `json:"weeklyLost"`
				SeasonPlayed int    `json:"seasonPlayed"`
				SeasonWon    int    `json:"seasonWon"`
				SeasonLost   int    `json:"seasonLost"`
			} `json:"ARENA_BRACKET_2v2"`
			ARENABRACKET3V3 struct {
				Slug         string `json:"slug"`
				Rating       int    `json:"rating"`
				WeeklyPlayed int    `json:"weeklyPlayed"`
				WeeklyWon    int    `json:"weeklyWon"`
				WeeklyLost   int    `json:"weeklyLost"`
				SeasonPlayed int    `json:"seasonPlayed"`
				SeasonWon    int    `json:"seasonWon"`
				SeasonLost   int    `json:"seasonLost"`
			} `json:"ARENA_BRACKET_3v3"`
			ARENABRACKETRBG struct {
				Slug         string `json:"slug"`
				Rating       int    `json:"rating"`
				WeeklyPlayed int    `json:"weeklyPlayed"`
				WeeklyWon    int    `json:"weeklyWon"`
				WeeklyLost   int    `json:"weeklyLost"`
				SeasonPlayed int    `json:"seasonPlayed"`
				SeasonWon    int    `json:"seasonWon"`
				SeasonLost   int    `json:"seasonLost"`
			} `json:"ARENA_BRACKET_RBG"`
			ARENABRACKET2V2SKIRMISH struct {
				Slug         string `json:"slug"`
				Rating       int    `json:"rating"`
				WeeklyPlayed int    `json:"weeklyPlayed"`
				WeeklyWon    int    `json:"weeklyWon"`
				WeeklyLost   int    `json:"weeklyLost"`
				SeasonPlayed int    `json:"seasonPlayed"`
				SeasonWon    int    `json:"seasonWon"`
				SeasonLost   int    `json:"seasonLost"`
			} `json:"ARENA_BRACKET_2v2_SKIRMISH"`
			UNKNOWN struct {
				Slug         string `json:"slug"`
				Rating       int    `json:"rating"`
				WeeklyPlayed int    `json:"weeklyPlayed"`
				WeeklyWon    int    `json:"weeklyWon"`
				WeeklyLost   int    `json:"weeklyLost"`
				SeasonPlayed int    `json:"seasonPlayed"`
				SeasonWon    int    `json:"seasonWon"`
				SeasonLost   int    `json:"seasonLost"`
			} `json:"UNKNOWN"`
		} `json:"brackets"`
	} `json:"pvp"`
	Quests     []int `json:"quests"`
	Reputation []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Standing int    `json:"standing"`
		Value    int    `json:"value"`
		Max      int    `json:"max"`
	} `json:"reputation"`
	Statistics struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		SubCategories []struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Statistics []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Quantity    int    `json:"quantity"`
				LastUpdated int64  `json:"lastUpdated"`
				Money       bool   `json:"money"`
			} `json:"statistics"`
			SubCategories []struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Statistics []struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Quantity    int    `json:"quantity"`
					LastUpdated int64  `json:"lastUpdated"`
					Money       bool   `json:"money"`
					Highest     string `json:"highest,omitempty"`
				} `json:"statistics"`
			} `json:"subCategories,omitempty"`
		} `json:"subCategories"`
	} `json:"statistics"`
	Stats struct {
		Health                      int     `json:"health"`
		PowerType                   string  `json:"powerType"`
		Power                       int     `json:"power"`
		Str                         int     `json:"str"`
		Agi                         int     `json:"agi"`
		Int                         int     `json:"int"`
		Sta                         int     `json:"sta"`
		SpeedRating                 float64 `json:"speedRating"`
		SpeedRatingBonus            float64 `json:"speedRatingBonus"`
		Crit                        float64 `json:"crit"`
		CritRating                  int     `json:"critRating"`
		Haste                       float64 `json:"haste"`
		HasteRating                 int     `json:"hasteRating"`
		HasteRatingPercent          float64 `json:"hasteRatingPercent"`
		Mastery                     float64 `json:"mastery"`
		MasteryRating               int     `json:"masteryRating"`
		Leech                       float64 `json:"leech"`
		LeechRating                 float64 `json:"leechRating"`
		LeechRatingBonus            float64 `json:"leechRatingBonus"`
		Versatility                 int     `json:"versatility"`
		VersatilityDamageDoneBonus  float64 `json:"versatilityDamageDoneBonus"`
		VersatilityHealingDoneBonus float64 `json:"versatilityHealingDoneBonus"`
		VersatilityDamageTakenBonus float64 `json:"versatilityDamageTakenBonus"`
		AvoidanceRating             float64 `json:"avoidanceRating"`
		AvoidanceRatingBonus        float64 `json:"avoidanceRatingBonus"`
		SpellPen                    int     `json:"spellPen"`
		SpellCrit                   float64 `json:"spellCrit"`
		SpellCritRating             int     `json:"spellCritRating"`
		Mana5                       float64 `json:"mana5"`
		Mana5Combat                 float64 `json:"mana5Combat"`
		Armor                       int     `json:"armor"`
		Dodge                       float64 `json:"dodge"`
		DodgeRating                 int     `json:"dodgeRating"`
		Parry                       float64 `json:"parry"`
		ParryRating                 int     `json:"parryRating"`
		Block                       float64 `json:"block"`
		BlockRating                 int     `json:"blockRating"`
		MainHandDmgMin              float64 `json:"mainHandDmgMin"`
		MainHandDmgMax              float64 `json:"mainHandDmgMax"`
		MainHandSpeed               float64 `json:"mainHandSpeed"`
		MainHandDps                 float64 `json:"mainHandDps"`
		OffHandDmgMin               float64 `json:"offHandDmgMin"`
		OffHandDmgMax               float64 `json:"offHandDmgMax"`
		OffHandSpeed                float64 `json:"offHandSpeed"`
		OffHandDps                  float64 `json:"offHandDps"`
		RangedDmgMin                float64 `json:"rangedDmgMin"`
		RangedDmgMax                float64 `json:"rangedDmgMax"`
		RangedSpeed                 float64 `json:"rangedSpeed"`
		RangedDps                   float64 `json:"rangedDps"`
	} `json:"stats"`
	Talents []struct {
		Selected bool `json:"selected,omitempty"`
		Talents  []struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				Range       string `json:"range"`
				PowerCost   string `json:"powerCost"`
				CastTime    string `json:"castTime"`
				Cooldown    string `json:"cooldown"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec,omitempty"`
		} `json:"talents"`
		Spec struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"spec,omitempty"`
		CalcTalent string `json:"calcTalent"`
		CalcSpec   string `json:"calcSpec"`
	} `json:"talents"`
	Titles []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"titles"`
	Audit struct {
		NumberOfIssues int `json:"numberOfIssues"`
		Slots          struct {
			Num2  int `json:"2"`
			Num4  int `json:"4"`
			Num5  int `json:"5"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
		} `json:"slots"`
		EmptyGlyphSlots     int  `json:"emptyGlyphSlots"`
		UnspentTalentPoints int  `json:"unspentTalentPoints"`
		NoSpec              bool `json:"noSpec"`
		UnenchantedItems    struct {
			Num2  int `json:"2"`
			Num4  int `json:"4"`
			Num6  int `json:"6"`
			Num7  int `json:"7"`
			Num8  int `json:"8"`
			Num9  int `json:"9"`
			Num14 int `json:"14"`
			Num15 int `json:"15"`
			Num16 int `json:"16"`
		} `json:"unenchantedItems"`
		EmptySockets          int `json:"emptySockets"`
		ItemsWithEmptySockets struct {
			Num15 int `json:"15"`
		} `json:"itemsWithEmptySockets"`
		AppropriateArmorType   int `json:"appropriateArmorType"`
		InappropriateArmorType struct {
		} `json:"inappropriateArmorType"`
		LowLevelItems struct {
		} `json:"lowLevelItems"`
		LowLevelThreshold   int `json:"lowLevelThreshold"`
		MissingExtraSockets struct {
			Num5 int `json:"5"`
		} `json:"missingExtraSockets"`
		RecommendedBeltBuckle struct {
			ID          int           `json:"id"`
			Description string        `json:"description"`
			Name        string        `json:"name"`
			Icon        string        `json:"icon"`
			Stackable   int           `json:"stackable"`
			ItemBind    int           `json:"itemBind"`
			BonusStats  []interface{} `json:"bonusStats"`
			ItemSpells  []struct {
				SpellID int `json:"spellId"`
				Spell   struct {
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Icon        string `json:"icon"`
					Description string `json:"description"`
					CastTime    string `json:"castTime"`
				} `json:"spell"`
				NCharges   int    `json:"nCharges"`
				Consumable bool   `json:"consumable"`
				CategoryID int    `json:"categoryId"`
				Trigger    string `json:"trigger"`
			} `json:"itemSpells"`
			BuyPrice          int  `json:"buyPrice"`
			ItemClass         int  `json:"itemClass"`
			ItemSubClass      int  `json:"itemSubClass"`
			ContainerSlots    int  `json:"containerSlots"`
			InventoryType     int  `json:"inventoryType"`
			Equippable        bool `json:"equippable"`
			ItemLevel         int  `json:"itemLevel"`
			MaxCount          int  `json:"maxCount"`
			MaxDurability     int  `json:"maxDurability"`
			MinFactionID      int  `json:"minFactionId"`
			MinReputation     int  `json:"minReputation"`
			Quality           int  `json:"quality"`
			SellPrice         int  `json:"sellPrice"`
			RequiredSkill     int  `json:"requiredSkill"`
			RequiredLevel     int  `json:"requiredLevel"`
			RequiredSkillRank int  `json:"requiredSkillRank"`
			ItemSource        struct {
				SourceID   int    `json:"sourceId"`
				SourceType string `json:"sourceType"`
			} `json:"itemSource"`
			BaseArmor            int           `json:"baseArmor"`
			HasSockets           bool          `json:"hasSockets"`
			IsAuctionable        bool          `json:"isAuctionable"`
			Armor                int           `json:"armor"`
			DisplayInfoID        int           `json:"displayInfoId"`
			NameDescription      string        `json:"nameDescription"`
			NameDescriptionColor string        `json:"nameDescriptionColor"`
			Upgradable           bool          `json:"upgradable"`
			HeroicTooltip        bool          `json:"heroicTooltip"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			AvailableContexts    []string      `json:"availableContexts"`
			BonusSummary         struct {
				DefaultBonusLists []interface{} `json:"defaultBonusLists"`
				ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
				BonusChances      []interface{} `json:"bonusChances"`
			} `json:"bonusSummary"`
			ArtifactID int `json:"artifactId"`
		} `json:"recommendedBeltBuckle"`
		MissingBlacksmithSockets struct {
		} `json:"missingBlacksmithSockets"`
		MissingEnchanterEnchants struct {
		} `json:"missingEnchanterEnchants"`
		MissingEngineerEnchants struct {
			Num9  int `json:"9"`
			Num14 int `json:"14"`
		} `json:"missingEngineerEnchants"`
		MissingScribeEnchants struct {
		} `json:"missingScribeEnchants"`
		NMissingJewelcrafterGems   int `json:"nMissingJewelcrafterGems"`
		RecommendedJewelcrafterGem struct {
			ID                int           `json:"id"`
			Description       string        `json:"description"`
			Name              string        `json:"name"`
			Icon              string        `json:"icon"`
			Stackable         int           `json:"stackable"`
			ItemBind          int           `json:"itemBind"`
			BonusStats        []interface{} `json:"bonusStats"`
			ItemSpells        []interface{} `json:"itemSpells"`
			BuyPrice          int           `json:"buyPrice"`
			ItemClass         int           `json:"itemClass"`
			ItemSubClass      int           `json:"itemSubClass"`
			ContainerSlots    int           `json:"containerSlots"`
			InventoryType     int           `json:"inventoryType"`
			Equippable        bool          `json:"equippable"`
			ItemLevel         int           `json:"itemLevel"`
			MaxCount          int           `json:"maxCount"`
			MaxDurability     int           `json:"maxDurability"`
			MinFactionID      int           `json:"minFactionId"`
			MinReputation     int           `json:"minReputation"`
			Quality           int           `json:"quality"`
			SellPrice         int           `json:"sellPrice"`
			RequiredSkill     int           `json:"requiredSkill"`
			RequiredLevel     int           `json:"requiredLevel"`
			RequiredSkillRank int           `json:"requiredSkillRank"`
			ItemSource        struct {
				SourceID   int    `json:"sourceId"`
				SourceType string `json:"sourceType"`
			} `json:"itemSource"`
			BaseArmor            int           `json:"baseArmor"`
			HasSockets           bool          `json:"hasSockets"`
			IsAuctionable        bool          `json:"isAuctionable"`
			Armor                int           `json:"armor"`
			DisplayInfoID        int           `json:"displayInfoId"`
			NameDescription      string        `json:"nameDescription"`
			NameDescriptionColor string        `json:"nameDescriptionColor"`
			Upgradable           bool          `json:"upgradable"`
			HeroicTooltip        bool          `json:"heroicTooltip"`
			Context              string        `json:"context"`
			BonusLists           []interface{} `json:"bonusLists"`
			AvailableContexts    []string      `json:"availableContexts"`
			BonusSummary         struct {
				DefaultBonusLists []interface{} `json:"defaultBonusLists"`
				ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
				BonusChances      []interface{} `json:"bonusChances"`
			} `json:"bonusSummary"`
			ArtifactID int `json:"artifactId"`
		} `json:"recommendedJewelcrafterGem"`
		MissingLeatherworkerEnchants struct {
		} `json:"missingLeatherworkerEnchants"`
	} `json:"audit"`
	TotalHonorableKills int `json:"totalHonorableKills"`
}

// JSON2Struct creates Character structure from JSON byte array
func (c *Character) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
