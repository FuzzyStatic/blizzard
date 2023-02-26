package owl

// SummaryData structure
type SummaryData struct {
	Players []Player `json:"players"`
	Teams   []Team   `json:"teams"`
}

// Player structure
type Player struct {
	Number        int    `json:"number"`
	PreferredSlot int    `json:"preferredSlot"`
	GivenName     string `json:"givenName"`
	Teams         []struct {
		ID            interface{} `json:"id"`
		EarliestMatch int64       `json:"earliestMatch,omitempty"`
		Stats         struct {
			HeroDamageDone  interface{} `json:"heroDamageDone"`
			HealingDone     interface{} `json:"healingDone"`
			DamageTaken     interface{} `json:"damageTaken"`
			FinalBlows      interface{} `json:"finalBlows"`
			Eliminations    interface{} `json:"eliminations"`
			Deaths          interface{} `json:"deaths"`
			TimeSpentOnFire interface{} `json:"timeSpentOnFire"`
			SoloKills       interface{} `json:"soloKills"`
			UltsUsed        interface{} `json:"ultsUsed"`
			UltsEarned      interface{} `json:"ultsEarned"`
			TimePlayed      interface{} `json:"timePlayed"`
		} `json:"stats,omitempty"`
		Heroes struct {
			Ana struct {
				DamageTaken        int `json:"damageTaken"`
				EnemiesSlept       int `json:"enemiesSlept"`
				ShotsHit           int `json:"shotsHit"`
				TimePlayed         int `json:"timePlayed"`
				Eliminations       int `json:"eliminations"`
				UltsEarned         int `json:"ultsEarned"`
				UltsUsed           int `json:"ultsUsed"`
				Deaths             int `json:"deaths"`
				ScopedHits         int `json:"scopedHits"`
				BioticGrenadeKills int `json:"bioticGrenadeKills"`
				TimeSpentOnFire    int `json:"timeSpentOnFire"`
				HealingDone        int `json:"healingDone"`
				FinalBlows         int `json:"finalBlows"`
				HeroDamageDone     int `json:"heroDamageDone"`
			} `json:"ana"`
			Ashe struct {
				CriticalHits           int `json:"criticalHits"`
				DamageTaken            int `json:"damageTaken"`
				ShotsHit               int `json:"shotsHit"`
				ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
				TimePlayed             int `json:"timePlayed"`
				Eliminations           int `json:"eliminations"`
				UltsEarned             int `json:"ultsEarned"`
				KnockbackKills         int `json:"knockbackKills"`
				UltsUsed               int `json:"ultsUsed"`
				Deaths                 int `json:"deaths"`
				ScopedHits             int `json:"scopedHits"`
				TimeSpentOnFire        int `json:"timeSpentOnFire"`
				FinalBlows             int `json:"finalBlows"`
				ScopedCriticalHits     int `json:"scopedCriticalHits"`
				SoloKills              int `json:"soloKills"`
				HeroDamageDone         int `json:"heroDamageDone"`
				BobKills               int `json:"bobKills"`
			} `json:"ashe"`
			Baptiste struct {
				CriticalHits                    int `json:"criticalHits"`
				DamageTaken                     int `json:"damageTaken"`
				AmplificationMatrixAssists      int `json:"amplificationMatrixAssists"`
				ShotsHit                        int `json:"shotsHit"`
				TimePlayed                      int `json:"timePlayed"`
				Eliminations                    int `json:"eliminations"`
				UltsEarned                      int `json:"ultsEarned"`
				UltsUsed                        int `json:"ultsUsed"`
				Deaths                          int `json:"deaths"`
				TimeSpentOnFire                 int `json:"timeSpentOnFire"`
				ImmortalityFieldDeathsPrevented int `json:"immortalityFieldDeathsPrevented"`
				DamageAmplified                 int `json:"damageAmplified"`
				HealingDone                     int `json:"healingDone"`
				FinalBlows                      int `json:"finalBlows"`
				SoloKills                       int `json:"soloKills"`
				HeroDamageDone                  int `json:"heroDamageDone"`
			} `json:"baptiste"`
			Doomfist struct {
				HeroDamageDone    int `json:"heroDamageDone"`
				DamageTaken       int `json:"damageTaken"`
				TimePlayed        int `json:"timePlayed"`
				ShotsHit          int `json:"shotsHit"`
				AbilityDamageDone int `json:"abilityDamageDone"`
			} `json:"doomfist"`
			DVa struct {
				CriticalHits   int `json:"criticalHits"`
				DamageTaken    int `json:"damageTaken"`
				ShotsHit       int `json:"shotsHit"`
				TimePlayed     int `json:"timePlayed"`
				Eliminations   int `json:"eliminations"`
				UltsEarned     int `json:"ultsEarned"`
				KnockbackKills int `json:"knockbackKills"`
				UltsUsed       int `json:"ultsUsed"`
				Deaths         int `json:"deaths"`
				FinalBlows     int `json:"finalBlows"`
				HeroDamageDone int `json:"heroDamageDone"`
			} `json:"d-va"`
			Echo struct {
				CriticalHits      int `json:"criticalHits"`
				FocusingBeamKills int `json:"focusingBeamKills"`
				DamageTaken       int `json:"damageTaken"`
				StickyBombsKills  int `json:"stickyBombsKills"`
				ShotsHit          int `json:"shotsHit"`
				TimePlayed        int `json:"timePlayed"`
				Eliminations      int `json:"eliminations"`
				UltsEarned        int `json:"ultsEarned"`
				UltsUsed          int `json:"ultsUsed"`
				Deaths            int `json:"deaths"`
				HealingDone       int `json:"healingDone"`
				FinalBlows        int `json:"finalBlows"`
				SoloKills         int `json:"soloKills"`
				HeroDamageDone    int `json:"heroDamageDone"`
			} `json:"echo"`
			Genji struct {
				CriticalHits     int `json:"criticalHits"`
				DamageTaken      int `json:"damageTaken"`
				DragonbladeKills int `json:"dragonbladeKills"`
				ShotsHit         int `json:"shotsHit"`
				TimePlayed       int `json:"timePlayed"`
				Eliminations     int `json:"eliminations"`
				UltsEarned       int `json:"ultsEarned"`
				UltsUsed         int `json:"ultsUsed"`
				Deaths           int `json:"deaths"`
				TimeSpentOnFire  int `json:"timeSpentOnFire"`
				FinalBlows       int `json:"finalBlows"`
				HeroDamageDone   int `json:"heroDamageDone"`
			} `json:"genji"`
			Lucio struct {
				CriticalHits   int `json:"criticalHits"`
				DamageTaken    int `json:"damageTaken"`
				ShotsHit       int `json:"shotsHit"`
				TimePlayed     int `json:"timePlayed"`
				Eliminations   int `json:"eliminations"`
				UltsEarned     int `json:"ultsEarned"`
				KnockbackKills int `json:"knockbackKills"`
				UltsUsed       int `json:"ultsUsed"`
				Deaths         int `json:"deaths"`
				HealingDone    int `json:"healingDone"`
				FinalBlows     int `json:"finalBlows"`
				HeroDamageDone int `json:"heroDamageDone"`
			} `json:"lucio"`
			Reaper struct {
				HeroDamageDone int `json:"heroDamageDone"`
				HealingDone    int `json:"healingDone"`
				DamageTaken    int `json:"damageTaken"`
				Deaths         int `json:"deaths"`
				UltsUsed       int `json:"ultsUsed"`
				UltsEarned     int `json:"ultsEarned"`
				TimePlayed     int `json:"timePlayed"`
				CriticalHits   int `json:"criticalHits"`
				ShotsHit       int `json:"shotsHit"`
			} `json:"reaper"`
			Roadhog struct {
				HeroDamageDone    int `json:"heroDamageDone"`
				DamageTaken       int `json:"damageTaken"`
				FinalBlows        int `json:"finalBlows"`
				Eliminations      int `json:"eliminations"`
				Deaths            int `json:"deaths"`
				TimeSpentOnFire   int `json:"timeSpentOnFire"`
				SoloKills         int `json:"soloKills"`
				UltsUsed          int `json:"ultsUsed"`
				UltsEarned        int `json:"ultsEarned"`
				TimePlayed        int `json:"timePlayed"`
				CriticalHits      int `json:"criticalHits"`
				ShotsHit          int `json:"shotsHit"`
				AccretionKills    int `json:"accretionKills"`
				GraviticFluxKills int `json:"graviticFluxKills"`
				KnockbackKills    int `json:"knockbackKills"`
			} `json:"roadhog"`
			Sigma struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				SoloKills                  int `json:"soloKills"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				CriticalHits               int `json:"criticalHits"`
				ShotsHit                   int `json:"shotsHit"`
				AccretionKills             int `json:"accretionKills"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
				KnockbackKills             int `json:"knockbackKills"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
			} `json:"sigma"`
			Soldier76 struct {
				CriticalHits     int `json:"criticalHits"`
				DamageTaken      int `json:"damageTaken"`
				HelixRocketKills int `json:"helixRocketKills"`
				ShotsHit         int `json:"shotsHit"`
				TimePlayed       int `json:"timePlayed"`
				Eliminations     int `json:"eliminations"`
				UltsEarned       int `json:"ultsEarned"`
				UltsUsed         int `json:"ultsUsed"`
				Deaths           int `json:"deaths"`
				HealingDone      int `json:"healingDone"`
				FinalBlows       int `json:"finalBlows"`
				HeroDamageDone   int `json:"heroDamageDone"`
			} `json:"soldier-76"`
			Tracer struct {
				HeroDamageDone  int `json:"heroDamageDone"`
				DamageTaken     int `json:"damageTaken"`
				FinalBlows      int `json:"finalBlows"`
				Eliminations    int `json:"eliminations"`
				TimeSpentOnFire int `json:"timeSpentOnFire"`
				TimePlayed      int `json:"timePlayed"`
				CriticalHits    int `json:"criticalHits"`
				ShotsHit        int `json:"shotsHit"`
			} `json:"tracer"`
			Widowmaker struct {
				CriticalHits           int `json:"criticalHits"`
				DamageTaken            int `json:"damageTaken"`
				ShotsHit               int `json:"shotsHit"`
				ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
				TimePlayed             int `json:"timePlayed"`
				Eliminations           int `json:"eliminations"`
				UltsEarned             int `json:"ultsEarned"`
				UltsUsed               int `json:"ultsUsed"`
				Deaths                 int `json:"deaths"`
				ScopedHits             int `json:"scopedHits"`
				FinalBlows             int `json:"finalBlows"`
				ScopedCriticalHits     int `json:"scopedCriticalHits"`
				SoloKills              int `json:"soloKills"`
				HeroDamageDone         int `json:"heroDamageDone"`
			} `json:"widowmaker"`
			Winston struct {
				DamageTaken     int `json:"damageTaken"`
				PrimalRageKills int `json:"primalRageKills"`
				JumpPackKills   int `json:"jumpPackKills"`
				ShotsHit        int `json:"shotsHit"`
				TimePlayed      int `json:"timePlayed"`
				Eliminations    int `json:"eliminations"`
				UltsEarned      int `json:"ultsEarned"`
				KnockbackKills  int `json:"knockbackKills"`
				UltsUsed        int `json:"ultsUsed"`
				Deaths          int `json:"deaths"`
				TimeSpentOnFire int `json:"timeSpentOnFire"`
				FinalBlows      int `json:"finalBlows"`
				HeroDamageDone  int `json:"heroDamageDone"`
			} `json:"winston"`
			WreckingBall struct {
				HeroDamageDone    int `json:"heroDamageDone"`
				DamageTaken       int `json:"damageTaken"`
				FinalBlows        int `json:"finalBlows"`
				Eliminations      int `json:"eliminations"`
				Deaths            int `json:"deaths"`
				TimeSpentOnFire   int `json:"timeSpentOnFire"`
				UltsUsed          int `json:"ultsUsed"`
				UltsEarned        int `json:"ultsEarned"`
				TimePlayed        int `json:"timePlayed"`
				AccretionKills    int `json:"accretionKills"`
				GraviticFluxKills int `json:"graviticFluxKills"`
			} `json:"wrecking-ball"`
			Zarya struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				TimePlayed                 int `json:"timePlayed"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				ShotsHit                   int `json:"shotsHit"`
			} `json:"zarya"`
		} `json:"heroes,omitempty"`
	} `json:"teams"`
	Name         string   `json:"name"`
	FamilyName   string   `json:"familyName"`
	Competitions []string `json:"competitions"`
	Role         string   `json:"role"`
	ID           int      `json:"id"`
	HeadshotURL  string   `json:"headshotUrl"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	CurrentTeam int `json:"currentTeam"`
}

type Team struct {
	ID           int      `json:"id"`
	Competitions []string `json:"competitions"`
	Name         string   `json:"name"`
	Roster       []int    `json:"roster"`
	Code         string   `json:"code"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
}

type AutoGenerated struct {
	Echo struct {
		CriticalHits      int `json:"criticalHits"`
		FocusingBeamKills int `json:"focusingBeamKills"`
		DamageTaken       int `json:"damageTaken"`
		StickyBombsKills  int `json:"stickyBombsKills"`
		ShotsHit          int `json:"shotsHit"`
		TimePlayed        int `json:"timePlayed"`
		Eliminations      int `json:"eliminations"`
		UltsEarned        int `json:"ultsEarned"`
		UltsUsed          int `json:"ultsUsed"`
		Deaths            int `json:"deaths"`
		HealingDone       int `json:"healingDone"`
		FinalBlows        int `json:"finalBlows"`
		SoloKills         int `json:"soloKills"`
		HeroDamageDone    int `json:"heroDamageDone"`
	} `json:"echo"`
}
