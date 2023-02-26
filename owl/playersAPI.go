package owl

// PlayersAPI structure
type PlayersAPI struct {
	Number        int    `json:"number"`
	PreferredSlot int    `json:"preferredSlot"`
	GivenName     string `json:"givenName"`
	Teams         []struct {
		ID            int   `json:"id"`
		EarliestMatch int64 `json:"earliestMatch,omitempty"`
		Stats         struct {
			HeroDamageDone             int `json:"heroDamageDone"`
			DamageTaken                int `json:"damageTaken"`
			FinalBlows                 int `json:"finalBlows"`
			Eliminations               int `json:"eliminations"`
			Deaths                     int `json:"deaths"`
			UltsUsed                   int `json:"ultsUsed"`
			UltsEarned                 int `json:"ultsEarned"`
			TimePlayed                 int `json:"timePlayed"`
			GravitonSurgeKills         int `json:"gravitonSurgeKills"`
			EnergyMax                  int `json:"energyMax"`
			LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
			CriticalHits               int `json:"criticalHits"`
			ShotsHit                   int `json:"shotsHit"`
			TimeSpentOnFire            int `json:"timeSpentOnFire"`
			SoloKills                  int `json:"soloKills"`
			AccretionKills             int `json:"accretionKills"`
			GraviticFluxKills          int `json:"graviticFluxKills"`
			KnockbackKills             int `json:"knockbackKills"`
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
		Logo           string `json:"logo,omitempty"`
		Icon           string `json:"icon,omitempty"`
		PrimaryColor   string `json:"primaryColor,omitempty"`
		SecondaryColor string `json:"secondaryColor,omitempty"`
	} `json:"teams"`
	CurrentTeams []int    `json:"currentTeams"`
	Name         string   `json:"name"`
	FamilyName   string   `json:"familyName"`
	Competitions []string `json:"competitions"`
	Role         string   `json:"role"`
	ID           int      `json:"id"`
	HeadshotURL  string   `json:"headshotUrl"`
	Stats        struct {
		HeroDamageDone             int `json:"heroDamageDone"`
		DamageTaken                int `json:"damageTaken"`
		FinalBlows                 int `json:"finalBlows"`
		Eliminations               int `json:"eliminations"`
		Deaths                     int `json:"deaths"`
		UltsUsed                   int `json:"ultsUsed"`
		UltsEarned                 int `json:"ultsEarned"`
		TimePlayed                 int `json:"timePlayed"`
		GravitonSurgeKills         int `json:"gravitonSurgeKills"`
		EnergyMax                  int `json:"energyMax"`
		LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
		CriticalHits               int `json:"criticalHits"`
		ShotsHit                   int `json:"shotsHit"`
		TimeSpentOnFire            int `json:"timeSpentOnFire"`
		SoloKills                  int `json:"soloKills"`
		AccretionKills             int `json:"accretionKills"`
		GraviticFluxKills          int `json:"graviticFluxKills"`
		KnockbackKills             int `json:"knockbackKills"`
	} `json:"stats"`
	SegmentStats struct {
		Owl22022KickoffClashQualifiers struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					HealingDone                int `json:"healingDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					TimeSpentOnFire            int `json:"timeSpentOnFire"`
					SoloKills                  int `json:"soloKills"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					CriticalHits               int `json:"criticalHits"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"stats"`
				Heroes struct {
					DVa struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						HealingDone     int `json:"healingDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						SoloKills       int `json:"soloKills"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
						CriticalHits    int `json:"criticalHits"`
						ShotsHit        int `json:"shotsHit"`
					} `json:"d-va"`
					Zarya struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						HealingDone                int `json:"healingDone"`
						DamageTaken                int `json:"damageTaken"`
						FinalBlows                 int `json:"finalBlows"`
						Eliminations               int `json:"eliminations"`
						Deaths                     int `json:"deaths"`
						TimeSpentOnFire            int `json:"timeSpentOnFire"`
						SoloKills                  int `json:"soloKills"`
						UltsUsed                   int `json:"ultsUsed"`
						UltsEarned                 int `json:"ultsEarned"`
						TimePlayed                 int `json:"timePlayed"`
						GravitonSurgeKills         int `json:"gravitonSurgeKills"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						ShotsHit                   int `json:"shotsHit"`
					} `json:"zarya"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				HealingDone                int `json:"healingDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				SoloKills                  int `json:"soloKills"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				GravitonSurgeKills         int `json:"gravitonSurgeKills"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				CriticalHits               int `json:"criticalHits"`
				ShotsHit                   int `json:"shotsHit"`
			} `json:"stats"`
			Heroes struct {
				DVa struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					HealingDone     int `json:"healingDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
					CriticalHits    int `json:"criticalHits"`
					ShotsHit        int `json:"shotsHit"`
				} `json:"d-va"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					HealingDone                int `json:"healingDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					TimeSpentOnFire            int `json:"timeSpentOnFire"`
					SoloKills                  int `json:"soloKills"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"zarya"`
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-kickoff-clash-qualifiers"`
		Owl22022MidseasonMadnessQualifiers struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					HealingDone       int `json:"healingDone"`
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
				} `json:"stats"`
				Heroes struct {
					DVa struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						HealingDone     int `json:"healingDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						SoloKills       int `json:"soloKills"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
						CriticalHits    int `json:"criticalHits"`
						ShotsHit        int `json:"shotsHit"`
						KnockbackKills  int `json:"knockbackKills"`
					} `json:"d-va"`
					Sigma struct {
						HeroDamageDone    int `json:"heroDamageDone"`
						HealingDone       int `json:"healingDone"`
						DamageTaken       int `json:"damageTaken"`
						FinalBlows        int `json:"finalBlows"`
						Eliminations      int `json:"eliminations"`
						Deaths            int `json:"deaths"`
						TimeSpentOnFire   int `json:"timeSpentOnFire"`
						SoloKills         int `json:"soloKills"`
						UltsUsed          int `json:"ultsUsed"`
						UltsEarned        int `json:"ultsEarned"`
						TimePlayed        int `json:"timePlayed"`
						AccretionKills    int `json:"accretionKills"`
						GraviticFluxKills int `json:"graviticFluxKills"`
					} `json:"sigma"`
					Roadhog struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						HealingDone     int `json:"healingDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						SoloKills       int `json:"soloKills"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
					} `json:"roadhog"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone    int `json:"heroDamageDone"`
				HealingDone       int `json:"healingDone"`
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
			} `json:"stats"`
			Heroes struct {
				DVa struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					HealingDone     int `json:"healingDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
					CriticalHits    int `json:"criticalHits"`
					ShotsHit        int `json:"shotsHit"`
					KnockbackKills  int `json:"knockbackKills"`
				} `json:"d-va"`
				Sigma struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					HealingDone       int `json:"healingDone"`
					DamageTaken       int `json:"damageTaken"`
					FinalBlows        int `json:"finalBlows"`
					Eliminations      int `json:"eliminations"`
					Deaths            int `json:"deaths"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					SoloKills         int `json:"soloKills"`
					UltsUsed          int `json:"ultsUsed"`
					UltsEarned        int `json:"ultsEarned"`
					TimePlayed        int `json:"timePlayed"`
					AccretionKills    int `json:"accretionKills"`
					GraviticFluxKills int `json:"graviticFluxKills"`
				} `json:"sigma"`
				Roadhog struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					HealingDone     int `json:"healingDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
				} `json:"roadhog"`
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-midseason-madness-qualifiers"`
		Owl22022SummerShowdownQualifiers struct {
			ID    int `json:"id"`
			Teams []struct {
				ID    int `json:"id"`
				Stats struct {
				} `json:"stats"`
				Heroes struct {
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
			} `json:"stats"`
			Heroes struct {
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-summer-showdown-qualifiers"`
		Owl22022CountdownCupQualifiers struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
					AccretionKills             int `json:"accretionKills"`
					SoloKills                  int `json:"soloKills"`
					GraviticFluxKills          int `json:"graviticFluxKills"`
				} `json:"stats"`
				Heroes struct {
					Sigma struct {
						HeroDamageDone    int `json:"heroDamageDone"`
						DamageTaken       int `json:"damageTaken"`
						FinalBlows        int `json:"finalBlows"`
						Eliminations      int `json:"eliminations"`
						Deaths            int `json:"deaths"`
						UltsUsed          int `json:"ultsUsed"`
						UltsEarned        int `json:"ultsEarned"`
						TimePlayed        int `json:"timePlayed"`
						AccretionKills    int `json:"accretionKills"`
						SoloKills         int `json:"soloKills"`
						GraviticFluxKills int `json:"graviticFluxKills"`
					} `json:"sigma"`
					Zarya struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						DamageTaken                int `json:"damageTaken"`
						FinalBlows                 int `json:"finalBlows"`
						Eliminations               int `json:"eliminations"`
						Deaths                     int `json:"deaths"`
						UltsUsed                   int `json:"ultsUsed"`
						UltsEarned                 int `json:"ultsEarned"`
						TimePlayed                 int `json:"timePlayed"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						ShotsHit                   int `json:"shotsHit"`
					} `json:"zarya"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				ShotsHit                   int `json:"shotsHit"`
				AccretionKills             int `json:"accretionKills"`
				SoloKills                  int `json:"soloKills"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
			} `json:"stats"`
			Heroes struct {
				Sigma struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					DamageTaken       int `json:"damageTaken"`
					FinalBlows        int `json:"finalBlows"`
					Eliminations      int `json:"eliminations"`
					Deaths            int `json:"deaths"`
					UltsUsed          int `json:"ultsUsed"`
					UltsEarned        int `json:"ultsEarned"`
					TimePlayed        int `json:"timePlayed"`
					AccretionKills    int `json:"accretionKills"`
					SoloKills         int `json:"soloKills"`
					GraviticFluxKills int `json:"graviticFluxKills"`
				} `json:"sigma"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"zarya"`
			} `json:"heroes"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-countdown-cup-qualifiers"`
		Owl22022Regular struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
					AccretionKills             int `json:"accretionKills"`
					SoloKills                  int `json:"soloKills"`
					GraviticFluxKills          int `json:"graviticFluxKills"`
				} `json:"stats"`
				Heroes struct {
					Sigma struct {
						HeroDamageDone    int `json:"heroDamageDone"`
						DamageTaken       int `json:"damageTaken"`
						FinalBlows        int `json:"finalBlows"`
						Eliminations      int `json:"eliminations"`
						Deaths            int `json:"deaths"`
						UltsUsed          int `json:"ultsUsed"`
						UltsEarned        int `json:"ultsEarned"`
						TimePlayed        int `json:"timePlayed"`
						AccretionKills    int `json:"accretionKills"`
						SoloKills         int `json:"soloKills"`
						GraviticFluxKills int `json:"graviticFluxKills"`
					} `json:"sigma"`
					Zarya struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						DamageTaken                int `json:"damageTaken"`
						FinalBlows                 int `json:"finalBlows"`
						Eliminations               int `json:"eliminations"`
						Deaths                     int `json:"deaths"`
						UltsUsed                   int `json:"ultsUsed"`
						UltsEarned                 int `json:"ultsEarned"`
						TimePlayed                 int `json:"timePlayed"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						ShotsHit                   int `json:"shotsHit"`
					} `json:"zarya"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				ShotsHit                   int `json:"shotsHit"`
				AccretionKills             int `json:"accretionKills"`
				SoloKills                  int `json:"soloKills"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
			} `json:"stats"`
			Heroes struct {
				Sigma struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					DamageTaken       int `json:"damageTaken"`
					FinalBlows        int `json:"finalBlows"`
					Eliminations      int `json:"eliminations"`
					Deaths            int `json:"deaths"`
					UltsUsed          int `json:"ultsUsed"`
					UltsEarned        int `json:"ultsEarned"`
					TimePlayed        int `json:"timePlayed"`
					AccretionKills    int `json:"accretionKills"`
					SoloKills         int `json:"soloKills"`
					GraviticFluxKills int `json:"graviticFluxKills"`
				} `json:"sigma"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"zarya"`
			} `json:"heroes"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-regular"`
		OverwatchTest2022Ow2Playtest2 struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					HealingDone     int `json:"healingDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
				} `json:"stats"`
				Heroes struct {
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone  int `json:"heroDamageDone"`
				HealingDone     int `json:"healingDone"`
				DamageTaken     int `json:"damageTaken"`
				FinalBlows      int `json:"finalBlows"`
				Eliminations    int `json:"eliminations"`
				Deaths          int `json:"deaths"`
				TimeSpentOnFire int `json:"timeSpentOnFire"`
				SoloKills       int `json:"soloKills"`
				UltsUsed        int `json:"ultsUsed"`
				UltsEarned      int `json:"ultsEarned"`
				TimePlayed      int `json:"timePlayed"`
			} `json:"stats"`
			Heroes struct {
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"overwatch-test-2022-ow2-playtest2"`
		Owl22022MidseasonMadnessTournament struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					HealingDone       int `json:"healingDone"`
					DamageTaken       int `json:"damageTaken"`
					FinalBlows        int `json:"finalBlows"`
					Eliminations      int `json:"eliminations"`
					Deaths            int `json:"deaths"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					SoloKills         int `json:"soloKills"`
					UltsUsed          int `json:"ultsUsed"`
					UltsEarned        int `json:"ultsEarned"`
					TimePlayed        int `json:"timePlayed"`
					AccretionKills    int `json:"accretionKills"`
					GraviticFluxKills int `json:"graviticFluxKills"`
				} `json:"stats"`
				Heroes struct {
					Sigma struct {
						HeroDamageDone    int `json:"heroDamageDone"`
						HealingDone       int `json:"healingDone"`
						DamageTaken       int `json:"damageTaken"`
						FinalBlows        int `json:"finalBlows"`
						Eliminations      int `json:"eliminations"`
						Deaths            int `json:"deaths"`
						TimeSpentOnFire   int `json:"timeSpentOnFire"`
						SoloKills         int `json:"soloKills"`
						UltsUsed          int `json:"ultsUsed"`
						UltsEarned        int `json:"ultsEarned"`
						TimePlayed        int `json:"timePlayed"`
						AccretionKills    int `json:"accretionKills"`
						GraviticFluxKills int `json:"graviticFluxKills"`
					} `json:"sigma"`
					WreckingBall struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						HealingDone     int `json:"healingDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						SoloKills       int `json:"soloKills"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
					} `json:"wrecking-ball"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone    int `json:"heroDamageDone"`
				HealingDone       int `json:"healingDone"`
				DamageTaken       int `json:"damageTaken"`
				FinalBlows        int `json:"finalBlows"`
				Eliminations      int `json:"eliminations"`
				Deaths            int `json:"deaths"`
				TimeSpentOnFire   int `json:"timeSpentOnFire"`
				SoloKills         int `json:"soloKills"`
				UltsUsed          int `json:"ultsUsed"`
				UltsEarned        int `json:"ultsEarned"`
				TimePlayed        int `json:"timePlayed"`
				AccretionKills    int `json:"accretionKills"`
				GraviticFluxKills int `json:"graviticFluxKills"`
			} `json:"stats"`
			Heroes struct {
				Sigma struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					HealingDone       int `json:"healingDone"`
					DamageTaken       int `json:"damageTaken"`
					FinalBlows        int `json:"finalBlows"`
					Eliminations      int `json:"eliminations"`
					Deaths            int `json:"deaths"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					SoloKills         int `json:"soloKills"`
					UltsUsed          int `json:"ultsUsed"`
					UltsEarned        int `json:"ultsEarned"`
					TimePlayed        int `json:"timePlayed"`
					AccretionKills    int `json:"accretionKills"`
					GraviticFluxKills int `json:"graviticFluxKills"`
				} `json:"sigma"`
				WreckingBall struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					HealingDone     int `json:"healingDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
				} `json:"wrecking-ball"`
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-midseason-madness-tournament"`
		Owl22022SummerShowdownTournament struct {
			ID    int `json:"id"`
			Teams []struct {
				ID    int `json:"id"`
				Stats struct {
				} `json:"stats"`
				Heroes struct {
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
			} `json:"stats"`
			Heroes struct {
			} `json:"heroes"`
			Logo               string `json:"logo"`
			Icon               string `json:"icon"`
			PrimaryColor       string `json:"primaryColor"`
			SecondaryColor     string `json:"secondaryColor"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-summer-showdown-tournament"`
		Owl22022PostSeason struct {
			ID    int `json:"id"`
			Teams []struct {
				ID    int `json:"id"`
				Stats struct {
				} `json:"stats"`
				Heroes struct {
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
			} `json:"stats"`
			Heroes struct {
			} `json:"heroes"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-post-season"`
		Owl22022RegularTournaments struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					TimeSpentOnFire            int `json:"timeSpentOnFire"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					AccretionKills             int `json:"accretionKills"`
					GraviticFluxKills          int `json:"graviticFluxKills"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					CriticalHits               int `json:"criticalHits"`
					ShotsHit                   int `json:"shotsHit"`
					KnockbackKills             int `json:"knockbackKills"`
					SoloKills                  int `json:"soloKills"`
				} `json:"stats"`
				Heroes struct {
					Sigma struct {
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
						SoloKills         int `json:"soloKills"`
					} `json:"sigma"`
					WreckingBall struct {
						TimePlayed int `json:"timePlayed"`
					} `json:"wrecking-ball"`
					DVa struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
						CriticalHits    int `json:"criticalHits"`
						ShotsHit        int `json:"shotsHit"`
						KnockbackKills  int `json:"knockbackKills"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
					} `json:"d-va"`
					Zarya struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						DamageTaken                int `json:"damageTaken"`
						FinalBlows                 int `json:"finalBlows"`
						Eliminations               int `json:"eliminations"`
						Deaths                     int `json:"deaths"`
						UltsUsed                   int `json:"ultsUsed"`
						UltsEarned                 int `json:"ultsEarned"`
						TimePlayed                 int `json:"timePlayed"`
						GravitonSurgeKills         int `json:"gravitonSurgeKills"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						ShotsHit                   int `json:"shotsHit"`
					} `json:"zarya"`
					Roadhog struct {
						DamageTaken int `json:"damageTaken"`
						TimePlayed  int `json:"timePlayed"`
					} `json:"roadhog"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				AccretionKills             int `json:"accretionKills"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
				GravitonSurgeKills         int `json:"gravitonSurgeKills"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				CriticalHits               int `json:"criticalHits"`
				ShotsHit                   int `json:"shotsHit"`
				KnockbackKills             int `json:"knockbackKills"`
				SoloKills                  int `json:"soloKills"`
			} `json:"stats"`
			Heroes struct {
				Sigma struct {
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
					SoloKills         int `json:"soloKills"`
				} `json:"sigma"`
				WreckingBall struct {
					TimePlayed int `json:"timePlayed"`
				} `json:"wrecking-ball"`
				DVa struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
					CriticalHits    int `json:"criticalHits"`
					ShotsHit        int `json:"shotsHit"`
					KnockbackKills  int `json:"knockbackKills"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
				} `json:"d-va"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"zarya"`
				Roadhog struct {
					DamageTaken int `json:"damageTaken"`
					TimePlayed  int `json:"timePlayed"`
				} `json:"roadhog"`
			} `json:"heroes"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-regular-tournaments"`
		Owl22022RegularTournamentsPlayoffs struct {
			ID    int `json:"id"`
			Teams []struct {
				ID            int   `json:"id"`
				EarliestMatch int64 `json:"earliestMatch"`
				Stats         struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					TimeSpentOnFire            int `json:"timeSpentOnFire"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					AccretionKills             int `json:"accretionKills"`
					GraviticFluxKills          int `json:"graviticFluxKills"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					CriticalHits               int `json:"criticalHits"`
					ShotsHit                   int `json:"shotsHit"`
					KnockbackKills             int `json:"knockbackKills"`
					SoloKills                  int `json:"soloKills"`
				} `json:"stats"`
				Heroes struct {
					Sigma struct {
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
						SoloKills         int `json:"soloKills"`
					} `json:"sigma"`
					WreckingBall struct {
						TimePlayed int `json:"timePlayed"`
					} `json:"wrecking-ball"`
					DVa struct {
						HeroDamageDone  int `json:"heroDamageDone"`
						DamageTaken     int `json:"damageTaken"`
						FinalBlows      int `json:"finalBlows"`
						Eliminations    int `json:"eliminations"`
						Deaths          int `json:"deaths"`
						UltsUsed        int `json:"ultsUsed"`
						UltsEarned      int `json:"ultsEarned"`
						TimePlayed      int `json:"timePlayed"`
						CriticalHits    int `json:"criticalHits"`
						ShotsHit        int `json:"shotsHit"`
						KnockbackKills  int `json:"knockbackKills"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
					} `json:"d-va"`
					Zarya struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						DamageTaken                int `json:"damageTaken"`
						FinalBlows                 int `json:"finalBlows"`
						Eliminations               int `json:"eliminations"`
						Deaths                     int `json:"deaths"`
						UltsUsed                   int `json:"ultsUsed"`
						UltsEarned                 int `json:"ultsEarned"`
						TimePlayed                 int `json:"timePlayed"`
						GravitonSurgeKills         int `json:"gravitonSurgeKills"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						ShotsHit                   int `json:"shotsHit"`
					} `json:"zarya"`
					Roadhog struct {
						DamageTaken int `json:"damageTaken"`
						TimePlayed  int `json:"timePlayed"`
					} `json:"roadhog"`
				} `json:"heroes"`
			} `json:"teams"`
			Stats struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				FinalBlows                 int `json:"finalBlows"`
				Eliminations               int `json:"eliminations"`
				Deaths                     int `json:"deaths"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				UltsUsed                   int `json:"ultsUsed"`
				UltsEarned                 int `json:"ultsEarned"`
				TimePlayed                 int `json:"timePlayed"`
				AccretionKills             int `json:"accretionKills"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
				GravitonSurgeKills         int `json:"gravitonSurgeKills"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				CriticalHits               int `json:"criticalHits"`
				ShotsHit                   int `json:"shotsHit"`
				KnockbackKills             int `json:"knockbackKills"`
				SoloKills                  int `json:"soloKills"`
			} `json:"stats"`
			Heroes struct {
				Sigma struct {
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
					SoloKills         int `json:"soloKills"`
				} `json:"sigma"`
				WreckingBall struct {
					TimePlayed int `json:"timePlayed"`
				} `json:"wrecking-ball"`
				DVa struct {
					HeroDamageDone  int `json:"heroDamageDone"`
					DamageTaken     int `json:"damageTaken"`
					FinalBlows      int `json:"finalBlows"`
					Eliminations    int `json:"eliminations"`
					Deaths          int `json:"deaths"`
					UltsUsed        int `json:"ultsUsed"`
					UltsEarned      int `json:"ultsEarned"`
					TimePlayed      int `json:"timePlayed"`
					CriticalHits    int `json:"criticalHits"`
					ShotsHit        int `json:"shotsHit"`
					KnockbackKills  int `json:"knockbackKills"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
				} `json:"d-va"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					FinalBlows                 int `json:"finalBlows"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					UltsUsed                   int `json:"ultsUsed"`
					UltsEarned                 int `json:"ultsEarned"`
					TimePlayed                 int `json:"timePlayed"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
				} `json:"zarya"`
				Roadhog struct {
					DamageTaken int `json:"damageTaken"`
					TimePlayed  int `json:"timePlayed"`
				} `json:"roadhog"`
			} `json:"heroes"`
			TeamLogo           string `json:"teamLogo"`
			TeamIcon           string `json:"teamIcon"`
			TeamPrimaryColor   string `json:"teamPrimaryColor"`
			TeamSecondaryColor string `json:"teamSecondaryColor"`
			Name               string `json:"name"`
			FamilyName         string `json:"familyName"`
			GivenName          string `json:"givenName"`
			HeadshotURL        string `json:"headshotUrl"`
			Role               string `json:"role"`
		} `json:"owl2-2022-regular-tournaments-playoffs"`
	} `json:"segmentStats"`
	Heroes struct {
		DVa struct {
			HeroDamageDone  int `json:"heroDamageDone"`
			DamageTaken     int `json:"damageTaken"`
			FinalBlows      int `json:"finalBlows"`
			Eliminations    int `json:"eliminations"`
			Deaths          int `json:"deaths"`
			UltsUsed        int `json:"ultsUsed"`
			UltsEarned      int `json:"ultsEarned"`
			TimePlayed      int `json:"timePlayed"`
			CriticalHits    int `json:"criticalHits"`
			ShotsHit        int `json:"shotsHit"`
			TimeSpentOnFire int `json:"timeSpentOnFire"`
			KnockbackKills  int `json:"knockbackKills"`
		} `json:"d-va"`
		Zarya struct {
			HeroDamageDone             int `json:"heroDamageDone"`
			DamageTaken                int `json:"damageTaken"`
			FinalBlows                 int `json:"finalBlows"`
			Eliminations               int `json:"eliminations"`
			Deaths                     int `json:"deaths"`
			UltsUsed                   int `json:"ultsUsed"`
			UltsEarned                 int `json:"ultsEarned"`
			TimePlayed                 int `json:"timePlayed"`
			GravitonSurgeKills         int `json:"gravitonSurgeKills"`
			EnergyMax                  int `json:"energyMax"`
			LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
			ShotsHit                   int `json:"shotsHit"`
		} `json:"zarya"`
		Sigma struct {
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
			AccretionKills    int `json:"accretionKills"`
			GraviticFluxKills int `json:"graviticFluxKills"`
		} `json:"sigma"`
		Roadhog struct {
			DamageTaken int `json:"damageTaken"`
			TimePlayed  int `json:"timePlayed"`
		} `json:"roadhog"`
		WreckingBall struct {
			TimePlayed int `json:"timePlayed"`
		} `json:"wrecking-ball"`
	} `json:"heroes"`
	AlternateIds []struct {
		Competitions []string `json:"competitions"`
		ID           int      `json:"id"`
	} `json:"alternateIds"`
}
