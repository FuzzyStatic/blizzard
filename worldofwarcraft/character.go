/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-11 22:25:18
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-11 23:11:47
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
		AchievementsCompleted          []int         `json:"achievementsCompleted"`
		AchievementsCompletedTimestamp []int64       `json:"achievementsCompletedTimestamp"`
		Criteria                       []int         `json:"criteria"`
		CriteriaQuantity               []interface{} `json:"criteriaQuantity"`
		CriteriaTimestamp              []int64       `json:"criteriaTimestamp"`
		CriteriaCreated                []interface{} `json:"criteriaCreated"`
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
	TotalHonorableKills int `json:"totalHonorableKills"`
}

// JSON2Struct creates Character structure from JSON byte array
func (c *Character) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, c)
}
