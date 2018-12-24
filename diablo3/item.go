/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:55
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:38:55
 */

package diablo3

import (
	"encoding/json"
	"reflect"
)

// Item structure
type Item struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	DisplayColor    string `json:"displayColor"`
	TooltipParams   string `json:"tooltipParams"`
	RequiredLevel   int    `json:"requiredLevel"`
	ItemLevel       int    `json:"itemLevel"`
	StackSizeMax    int    `json:"stackSizeMax"`
	BonusAffixes    int    `json:"bonusAffixes"`
	BonusAffixesMax int    `json:"bonusAffixesMax"`
	AccountBound    bool   `json:"accountBound"`
	FlavorText      string `json:"flavorText"`
	TypeName        string `json:"typeName"`
	Type            struct {
		TwoHanded bool   `json:"twoHanded"`
		ID        string `json:"id"`
	} `json:"type"`
	DamageRange string `json:"damageRange"`
	Armor       struct {
		MinMax
	} `json:"armor"`
	Slots      []string `json:"slots"`
	Attributes struct {
		Primary []struct {
			Text      string `json:"text"`
			Color     string `json:"color"`
			AffixType string `json:"affixType"`
		} `json:"primary"`
		Secondary []struct {
			Text      string `json:"text"`
			Color     string `json:"color"`
			AffixType string `json:"affixType"`
		} `json:"secondary"`
		Passive []interface{} `json:"passive"`
	} `json:"attributes"`
	AttributesRaw struct {
		StrengthItem struct {
			MinMax
		} `json:"Strength_Item"`
		DexterityItem struct {
			MinMax
		} `json:"Dexterity_Item"`
		IntelligenceItem struct {
			MinMax
		} `json:"Intelligence_Item"`
		VitalityItem struct {
			MinMax
		} `json:"Vitality_Item"`
		AttacksPerSecondPercent struct {
			MinMax
		} `json:"Attacks_Per_Second_Percent"`
		CritPercentBonusCapped struct {
			MinMax
		} `json:"Crit_Percent_Bonus_Capped"`
		CritDamagePercent struct {
			MinMax
		} `json:"Crit_Damage_Percent"`
		SplashDamageEffectPercent struct {
			MinMax
		} `json:"Splash_Damage_Effect_Percent"`
		DamageDealtPercentBonusPhysical struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Physical"`
		DamageDealtPercentBonusCold struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Cold"`
		DamageDealtPercentBonusFire struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Fire"`
		DamageDealtPercentBonusLightning struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Lightning"`
		DamageDealtPercentBonusPoison struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Poison"`
		DamageDealtPercentBonusArcane struct {
			MinMax
		} `json:"Damage_Dealt_Percent_Bonus#Arcane"`
		PowerDamagePercentBonusBarbarianAncientSpear struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_AncientSpear"`
		PowerDamagePercentBonusBarbarianAvalanche struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Avalanche"`
		PowerDamagePercentBonusBarbarianBash struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Bash"`
		PowerDamagePercentBonusBarbarianCallOfTheAncients struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_CallOfTheAncients"`
		PowerDamagePercentBonusBarbarianCleave struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Cleave"`
		PowerDamagePercentBonusBarbarianEarthquake struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Earthquake"`
		PowerDamagePercentBonusBarbarianFrenzy struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Frenzy"`
		PowerDamagePercentBonusBarbarianFuriousCharge struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_FuriousCharge"`
		PowerDamagePercentBonusBarbarianHammerOfTheAncients struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_HammerOfTheAncients"`
		PowerDamagePercentBonusBarbarianOverpower struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Overpower"`
		PowerDamagePercentBonusBarbarianRend struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Rend"`
		PowerDamagePercentBonusBarbarianRevenge struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Revenge"`
		PowerDamagePercentBonusBarbarianSeismicSlam struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_SeismicSlam"`
		PowerDamagePercentBonusBarbarianWeaponThrow struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_WeaponThrow"`
		PowerDamagePercentBonusBarbarianWhirlwind struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Barbarian_Whirlwind"`
		PowerDamagePercentBonusCrusaderBlessedHammer struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_BlessedHammer"`
		PowerDamagePercentBonusCrusaderBlessedShield struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_BlessedShield"`
		PowerDamagePercentBonusCrusaderBombardment struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Bombardment"`
		PowerDamagePercentBonusCrusaderCodemn struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Codemn"`
		PowerDamagePercentBonusCrusaderFallingSword struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_FallingSword"`
		PowerDamagePercentBonusCrusaderFistOfTheHeavens struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_FistOfTheHeavens"`
		PowerDamagePercentBonusCrusaderHeavensFury struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_HeavensFury"`
		PowerDamagePercentBonusCrusaderJustice struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Justice"`
		PowerDamagePercentBonusCrusaderPhalanx struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Phalanx"`
		PowerDamagePercentBonusCrusaderPunish struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Punish"`
		PowerDamagePercentBonusCrusaderShieldBash struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_ShieldBash"`
		PowerDamagePercentBonusCrusaderSlash struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Slash"`
		PowerDamagePercentBonusCrusaderSmite struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_Smite"`
		PowerDamagePercentBonusCrusaderSweepAttack struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Crusader_SweepAttack"`
		PowerDamagePercentBonusDemonhunterBolas struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Bolas"`
		PowerDamagePercentBonusDemonhunterChakram struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Chakram"`
		PowerDamagePercentBonusDemonhunterClusterArrow struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_ClusterArrow"`
		PowerDamagePercentBonusDemonhunterCompanion struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Companion"`
		PowerDamagePercentBonusDemonhunterElemenentalArrow struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_ElemenentalArrow"`
		PowerDamagePercentBonusDemonhunterEntanglingShot struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_EntanglingShot"`
		PowerDamagePercentBonusDemonhunterEvasiveFire struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_EvasiveFire"`
		PowerDamagePercentBonusDemonhunterFanOfKnives struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_FanOfKnives"`
		PowerDamagePercentBonusDemonhunterGrenade struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Grenade"`
		PowerDamagePercentBonusDemonhunterHungeringArrow struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_HungeringArrow"`
		PowerDamagePercentBonusDemonhunterMultishot struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Multishot"`
		PowerDamagePercentBonusDemonhunterImpale struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Impale"`
		PowerDamagePercentBonusDemonhunterRainOfVengeance struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_RainOfVengeance"`
		PowerDamagePercentBonusDemonhunterRapidFire struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_RapidFire"`
		PowerDamagePercentBonusDemonhunterSentry struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Sentry"`
		PowerDamagePercentBonusDemonhunterSpikeTrap struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_SpikeTrap"`
		PowerDamagePercentBonusDemonhunterStrafe struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Demonhunter_Strafe"`
		PowerDamagePercentBonusMonkCripplingWave struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_CripplingWave"`
		PowerDamagePercentBonusMonkCycloneStrike struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_CycloneStrike"`
		PowerDamagePercentBonusMonkDashingStrike struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_DashingStrike"`
		PowerDamagePercentBonusMonkDeadlyReach struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_DeadlyReach"`
		PowerDamagePercentBonusMonkExplodingPalm struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_ExplodingPalm"`
		PowerDamagePercentBonusMonkFistsOfThunder struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_FistsOfThunder"`
		PowerDamagePercentBonusMonkLashingTailKick struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_LashingTailKick"`
		PowerDamagePercentBonusMonkMysticAlly struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_MysticAlly"`
		PowerDamagePercentBonusMonkSevenSidedStrike struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_SevenSidedStrike"`
		PowerDamagePercentBonusMonkSweepingWind struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_SweepingWind"`
		PowerDamagePercentBonusMonkTempestRush struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_TempestRush"`
		PowerDamagePercentBonusMonkWaveOfLight struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_WaveOfLight"`
		PowerDamagePercentBonusMonkWayOfTheHundredFists struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Monk_WayOfTheHundredFists"`
		PowerDamagePercentBonusP6NecroBoneNova struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_BoneNova"`
		PowerDamagePercentBonusP6NecroBoneArmor struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_BoneArmor"`
		PowerDamagePercentBonusP6NecroBoneSpear struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_BoneSpear"`
		PowerDamagePercentBonusP6NecroBoneSpikes struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_BoneSpikes"`
		PowerDamagePercentBonusP6NecroBoneSpirit struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_BoneSpirit"`
		PowerDamagePercentBonusP6NecroCorpseExplosion struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_CorpseExplosion"`
		PowerDamagePercentBonusP6NecroCorpseLance struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_CorpseLance"`
		PowerDamagePercentBonusP6NecroGrimScythe struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_GrimScythe"`
		PowerDamagePercentBonusP6NecroRaiseDead struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_RaiseDead"`
		PowerDamagePercentBonusP6NecroRevive struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_Revive"`
		PowerDamagePercentBonusP6NecroSiphonBlood struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_SiphonBlood"`
		PowerDamagePercentBonusP6NecroTraitGolemSpawner struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_Trait_GolemSpawner"`
		PowerDamagePercentBonusP6NecroTraitSkeletonSpawner struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#P6_Necro_Trait_SkeletonSpawner"`
		PowerDamagePercentBonusWitchdoctorAcidCloud struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_AcidCloud"`
		PowerDamagePercentBonusWitchdoctorCorpseSpiders struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_CorpseSpiders"`
		PowerDamagePercentBonusWitchdoctorFetishArmy struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_FetishArmy"`
		PowerDamagePercentBonusWitchdoctorFirebats struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Firebats"`
		PowerDamagePercentBonusWitchdoctorFirebomb struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Firebomb"`
		PowerDamagePercentBonusWitchdoctorGargantuan struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Gargantuan"`
		PowerDamagePercentBonusWitchdoctorGraspOfTheDead struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_GraspOfTheDead"`
		PowerDamagePercentBonusWitchdoctorHaunt struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Haunt"`
		PowerDamagePercentBonusWitchdoctorLocustSwarm struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_LocustSwarm"`
		PowerDamagePercentBonusWitchdoctorPiranhas struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Piranhas"`
		PowerDamagePercentBonusWitchdoctorPlagueOfToads struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_PlagueOfToads"`
		PowerDamagePercentBonusWitchdoctorPoisonDart struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_PoisonDart"`
		PowerDamagePercentBonusWitchdoctorSacrifice struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_Sacrifice"`
		PowerDamagePercentBonusWitchdoctorSpiritBarrage struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_SpiritBarrage"`
		PowerDamagePercentBonusWitchdoctorWallOfDeath struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_WallOfDeath"`
		PowerDamagePercentBonusWitchdoctorZombieCharger struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_ZombieCharger"`
		PowerDamagePercentBonusWitchdoctorZombieDogSpawnerPassive struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Witchdoctor_ZombieDogSpawner_Passive"`
		PowerDamagePercentBonusWizardArcaneOrb struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_ArcaneOrb"`
		PowerDamagePercentBonusWizardArcaneTorrent struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_ArcaneTorrent"`
		PowerDamagePercentBonusWizardBlackHole struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_BlackHole"`
		PowerDamagePercentBonusWizardBlizzard struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Blizzard"`
		PowerDamagePercentBonusWizardDisintergrate struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Disintergrate"`
		PowerDamagePercentBonusWizardElectrocute struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Electrocute"`
		PowerDamagePercentBonusWizardEnergyTwister struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_EnergyTwister"`
		PowerDamagePercentBonusWizardExplosiveBlast struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_ExplosiveBlast"`
		PowerDamagePercentBonusWizardFamiliar struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Familiar"`
		PowerDamagePercentBonusWizardHydra struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Hydra"`
		PowerDamagePercentBonusWizardMagicMissle struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_MagicMissle"`
		PowerDamagePercentBonusWizardMeteor struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_Meteor"`
		PowerDamagePercentBonusWizardRayOfFrost struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_RayOfFrost"`
		PowerDamagePercentBonusWizardShockPulse struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_ShockPulse"`
		PowerDamagePercentBonusWizardSpectralBlade struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_SpectralBlade"`
		PowerDamagePercentBonusWizardWaveOfForce struct {
			MinMax
		} `json:"Power_Damage_Percent_Bonus#Wizard_WaveOfForce"`
		ArmorItem struct {
			MinMax
		} `json:"Armor_Item"`
		ArmorBonusItem struct {
			MinMax
		} `json:"Armor_Bonus_Item"`
		ResistancePhysical struct {
			MinMax
		} `json:"Resistance#Physical"`
		ResistanceAll struct {
			MinMax
		} `json:"ResistanceAll"`
		ResistanceCold struct {
			MinMax
		} `json:"Resistance#Cold"`
		ResistanceFire struct {
			MinMax
		} `json:"Resistance#Fire"`
		ResistanceLightning struct {
			MinMax
		} `json:"Resistance#Lightning"`
		ResistancePoison struct {
			MinMax
		} `json:"Resistance#Poison"`
		ResistanceArcane struct {
			MinMax
		} `json:"Resistance#Arcane"`
		DamagePercentReductionFromMelee struct {
			MinMax
		} `json:"Damage_Percent_Reduction_From_Melee"`
		HitpointsRegenPerSecond struct {
			MinMax
		} `json:"Hitpoints_Regen_Per_Second"`
		HealthGlobeBonusHealth struct {
			MinMax
		} `json:"Health_Globe_Bonus_Health"`
		GoldPickUpRadius struct {
			MinMax
		} `json:"Gold_PickUp_Radius"`
		ResourceMaxBonusArcanum struct {
			MinMax
		} `json:"Resource_Max_Bonus#Arcanum"`
		ResourceCostReductionPercentAll struct {
			MinMax
		} `json:"Resource_Cost_Reduction_Percent_All"`
		ExperienceBonus struct {
			MinMax
		} `json:"Experience_Bonus"`
		DamageMinPhysical struct {
			MinMax
		} `json:"Damage_Min#Physical"`
		DamageDeltaPhysical struct {
			MinMax
		} `json:"Damage_Delta#Physical"`
		Post212Drop struct {
			MinMax
		} `json:"Post_2_1_2_Drop"`
		ItemLegendaryItemLevelOverride struct {
			MinMax
		} `json:"Item_LegendaryItem_Level_Override"`
		ItemBindingLevelOverride struct {
			MinMax
		} `json:"Item_Binding_Level_Override"`
		DurabilityCur struct {
			MinMax
		} `json:"Durability_Cur"`
		ItemLegendaryItemBaseItem struct {
			MinMax
		} `json:"Item_Legendary_Item_Base_Item"`
		Sockets struct {
			MinMax
		} `json:"Sockets"`
		Loot20Drop struct {
			MinMax
		} `json:"Loot_2_0_Drop"`
		DamageWeaponDeltaPhysical struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Physical"`
		DamageWeaponDeltaCold struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Cold"`
		DamageWeaponDeltaFire struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Fire"`
		DamageWeaponDeltaLightning struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Lightning"`
		DamageWeaponDeltaPoison struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Poison"`
		DamageWeaponDeltaArcane struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Arcane"`
		DamageWeaponDeltaHoly struct {
			MinMax
		} `json:"Damage_Weapon_Delta#Holy"`
		DamageWeaponMinPhysical struct {
			MinMax
		} `json:"Damage_Weapon_Min#Physical"`
		DamageWeaponMinCold struct {
			MinMax
		} `json:"Damage_Weapon_Min#Cold"`
		DamageWeaponMinFire struct {
			MinMax
		} `json:"Damage_Weapon_Min#Fire"`
		DamageWeaponMinLightning struct {
			MinMax
		} `json:"Damage_Weapon_Min#Lightning"`
		DamageWeaponMinPoison struct {
			MinMax
		} `json:"Damage_Weapon_Min#Poison"`
		DamageWeaponMinArcane struct {
			MinMax
		} `json:"Damage_Weapon_Min#Arcane"`
		DamageWeaponMinHoly struct {
			MinMax
		} `json:"Damage_Weapon_Min#Holy"`
		DurabilityMax struct {
			MinMax
		} `json:"Durability_Max"`
		OnHitStunProcChance struct {
			MinMax
		} `json:"On_Hit_Stun_Proc_Chance"`
		Quiver struct {
			MinMax
		} `json:"Quiver"`
		ItemPowerPassiveP6ItemPassiveUniqueRing001 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_001"`
		ItemPowerPassiveP6ItemPassiveUniqueRing002 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_002"`
		ItemPowerPassiveP6ItemPassiveUniqueRing003 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_003"`
		ItemPowerPassiveP6ItemPassiveUniqueRing004 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_004"`
		ItemPowerPassiveP6ItemPassiveUniqueRing005 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_005"`
		ItemPowerPassiveP6ItemPassiveUniqueRing006 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_006"`
		ItemPowerPassiveP6ItemPassiveUniqueRing007 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_007"`
		ItemPowerPassiveP6ItemPassiveUniqueRing008 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_008"`
		ItemPowerPassiveP6ItemPassiveUniqueRing009 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_009"`
		ItemPowerPassiveP6ItemPassiveUniqueRing010 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_010"`
		ItemPowerPassiveP6ItemPassiveUniqueRing011 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_011"`
		ItemPowerPassiveP6ItemPassiveUniqueRing012 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_012"`
		ItemPowerPassiveP6ItemPassiveUniqueRing037 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_037"`
		ItemPowerPassiveP6ItemPassiveUniqueRing038 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_038"`
		ItemPowerPassiveP6ItemPassiveUniqueRing039 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_039"`
		ItemPowerPassiveP6ItemPassiveUniqueRing040 struct {
			MinMax
		} `json:"Item_Power_Passive#P6_ItemPassive_Unique_Ring_040"`
		ItemPowerPassiveItemPassiveUniqueRing653x1 struct {
			MinMax
		} `json:"Item_Power_Passive#ItemPassive_Unique_Ring_653_x1"`
		Season struct {
			MinMax
		} `json:"Season"`
		AncientRank struct {
			MinMax
		} `json:"Ancient_Rank"`
	} `json:"attributesRaw"`
	RandomAffixes []interface{} `json:"randomAffixes"`
	Gems          []struct {
		Item struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"item"`
		IsGem      bool `json:"isGem"`
		IsJewel    bool `json:"isJewel"`
		Attributes struct {
			Primary []struct {
				Text      string `json:"text"`
				Color     string `json:"color"`
				AffixType string `json:"affixType"`
			} `json:"primary"`
			Secondary []interface{} `json:"secondary"`
			Passive   []interface{} `json:"passive"`
		} `json:"attributes"`
		AttributesRaw struct {
			PowerCooldownReductionPercentAll struct {
				MinMax
			} `json:"Power_Cooldown_Reduction_Percent_All"`
		} `json:"attributesRaw"`
	} `json:"gems"`
	SocketEffects []interface{} `json:"socketEffects"`
	Set           struct {
		Name  string `json:"name"`
		Items []struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"items"`
		Slug  string `json:"slug"`
		Ranks []struct {
			Required   int `json:"required"`
			Attributes struct {
				Primary   []interface{} `json:"primary"`
				Secondary []interface{} `json:"secondary"`
				Passive   []struct {
					Text      string `json:"text"`
					Color     string `json:"color"`
					AffixType string `json:"affixType"`
				} `json:"passive"`
			} `json:"attributes"`
			AttributesRaw struct {
				ItemPowerPassiveItemPassiveUniqueRing727X1 struct {
					MinMax
				} `json:"Item_Power_Passive#ItemPassive_Unique_Ring_727_x1"`
			} `json:"attributesRaw"`
		} `json:"ranks"`
	} `json:"set"`
	SetItemsEquipped       []string      `json:"setItemsEquipped"`
	CraftedBy              []interface{} `json:"craftedBy"`
	SeasonRequiredToDrop   int           `json:"seasonRequiredToDrop"`
	IsSeasonRequiredToDrop bool          `json:"isSeasonRequiredToDrop"`
	Description            interface{}   `json:"description"`
	BlockChance            string        `json:"blockChance"`
}

// MinMax min/max values of attributes
type MinMax struct {
	Min float64 `json:"min,omitempty"`
	Max float64 `json:"max,omitempty"`
}

// JSON2Struct creates Item structure from JSON byte array
func (i *Item) JSON2Struct(b []byte) error {
	return json.Unmarshal(b, i)
}

// GetItemAttributesRaw returns the raw attributes of an item in a map
func GetItemAttributesRaw(item *Item) map[string]float64 {
	m := make(map[string]float64)
	v := reflect.ValueOf(item.AttributesRaw)

	for i := 0; i < v.NumField(); i++ {
		minMax := v.Field(i).Interface().(struct{ MinMax })
		if minMax.Max != 0.0 {
			m[v.Type().Field(i).Name] = minMax.Max
		}
	}

	return m
}

// CompareItemsAttributesRaw returns the difference values of all attributes for item1
// compared to item2
func CompareItemsAttributesRaw(item1, item2 *Item) map[string]float64 {
	m := make(map[string]float64)

	mItem1 := GetItemAttributesRaw(item1)
	mItem2 := GetItemAttributesRaw(item2)

	for i := range mItem1 {
		m[i] = mItem1[i] - mItem2[i]
	}

	return m
}
