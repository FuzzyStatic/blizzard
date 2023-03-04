package owl

type MatchesAPI struct {
	CompetitionID string `json:"competitionId"`
	Conclusion    string `json:"conclusion"`
	EndTimestamp  int64  `json:"endTimestamp"`
	Games         struct {
		Num60411 struct {
			Teams struct {
				Num15188 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
					} `json:"teamStats"`
				} `json:"15188"`
				Num15206 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
					} `json:"teamStats"`
				} `json:"15206"`
			} `json:"teams"`
			ID      int    `json:"id"`
			Map     string `json:"map"`
			MapType string `json:"mapType"`
			Number  int    `json:"number"`
			Players struct {
				Num15480 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits      int `json:"criticalHits"`
						DamageTaken       int `json:"damageTaken"`
						HelixRocketKills  int `json:"helixRocketKills"`
						ShotsHit          int `json:"shotsHit"`
						TimePlayed        int `json:"timePlayed"`
						Eliminations      int `json:"eliminations"`
						UltsEarned        int `json:"ultsEarned"`
						UltsUsed          int `json:"ultsUsed"`
						Deaths            int `json:"deaths"`
						TimeSpentOnFire   int `json:"timeSpentOnFire"`
						HealingDone       int `json:"healingDone"`
						FinalBlows        int `json:"finalBlows"`
						HeroDamageDone    int `json:"heroDamageDone"`
						DeathBlossomKills int `json:"deathBlossomKills"`
					} `json:"stats"`
					Heroes struct {
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
							TimeSpentOnFire  int `json:"timeSpentOnFire"`
							HealingDone      int `json:"healingDone"`
							FinalBlows       int `json:"finalBlows"`
							HeroDamageDone   int `json:"heroDamageDone"`
						} `json:"soldier-76"`
						Reaper struct {
							CriticalHits      int `json:"criticalHits"`
							DamageTaken       int `json:"damageTaken"`
							ShotsHit          int `json:"shotsHit"`
							TimePlayed        int `json:"timePlayed"`
							Eliminations      int `json:"eliminations"`
							UltsEarned        int `json:"ultsEarned"`
							UltsUsed          int `json:"ultsUsed"`
							Deaths            int `json:"deaths"`
							TimeSpentOnFire   int `json:"timeSpentOnFire"`
							HealingDone       int `json:"healingDone"`
							FinalBlows        int `json:"finalBlows"`
							DeathBlossomKills int `json:"deathBlossomKills"`
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"reaper"`
					} `json:"heroes"`
				} `json:"15480"`
				Num15622 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15622"`
				Num15643 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits       int `json:"criticalHits"`
						DamageTaken        int `json:"damageTaken"`
						HelixRocketKills   int `json:"helixRocketKills"`
						ShotsHit           int `json:"shotsHit"`
						TimePlayed         int `json:"timePlayed"`
						Eliminations       int `json:"eliminations"`
						UltsEarned         int `json:"ultsEarned"`
						UltsUsed           int `json:"ultsUsed"`
						TacticalVisorKills int `json:"tacticalVisorKills"`
						Deaths             int `json:"deaths"`
						HealingDone        int `json:"healingDone"`
						FinalBlows         int `json:"finalBlows"`
						SoloKills          int `json:"soloKills"`
						HeroDamageDone     int `json:"heroDamageDone"`
						DeathBlossomKills  int `json:"deathBlossomKills"`
					} `json:"stats"`
					Heroes struct {
						Soldier76 struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							HelixRocketKills   int `json:"helixRocketKills"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							TacticalVisorKills int `json:"tacticalVisorKills"`
							Deaths             int `json:"deaths"`
							HealingDone        int `json:"healingDone"`
							FinalBlows         int `json:"finalBlows"`
							SoloKills          int `json:"soloKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"soldier-76"`
						Reaper struct {
							CriticalHits      int `json:"criticalHits"`
							DamageTaken       int `json:"damageTaken"`
							ShotsHit          int `json:"shotsHit"`
							TimePlayed        int `json:"timePlayed"`
							Eliminations      int `json:"eliminations"`
							UltsEarned        int `json:"ultsEarned"`
							UltsUsed          int `json:"ultsUsed"`
							Deaths            int `json:"deaths"`
							HealingDone       int `json:"healingDone"`
							FinalBlows        int `json:"finalBlows"`
							DeathBlossomKills int `json:"deathBlossomKills"`
							SoloKills         int `json:"soloKills"`
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"reaper"`
					} `json:"heroes"`
				} `json:"15643"`
				Num15787 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15787"`
				Num15937 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						HeroDamageDone    int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
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
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"echo"`
					} `json:"heroes"`
				} `json:"15937"`
				Num16205 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits    int `json:"criticalHits"`
						DamageTaken     int `json:"damageTaken"`
						ShotsHit        int `json:"shotsHit"`
						TimePlayed      int `json:"timePlayed"`
						Eliminations    int `json:"eliminations"`
						UltsEarned      int `json:"ultsEarned"`
						KnockbackKills  int `json:"knockbackKills"`
						UltsUsed        int `json:"ultsUsed"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						HealingDone     int `json:"healingDone"`
						FinalBlows      int `json:"finalBlows"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							CriticalHits    int `json:"criticalHits"`
							DamageTaken     int `json:"damageTaken"`
							ShotsHit        int `json:"shotsHit"`
							TimePlayed      int `json:"timePlayed"`
							Eliminations    int `json:"eliminations"`
							UltsEarned      int `json:"ultsEarned"`
							KnockbackKills  int `json:"knockbackKills"`
							UltsUsed        int `json:"ultsUsed"`
							Deaths          int `json:"deaths"`
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							HealingDone     int `json:"healingDone"`
							FinalBlows      int `json:"finalBlows"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"16205"`
				Num16232 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"16232"`
				Num16240 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						DamageTaken                int `json:"damageTaken"`
						JumpPackKills              int `json:"jumpPackKills"`
						ShotsHit                   int `json:"shotsHit"`
						TimePlayed                 int `json:"timePlayed"`
						Eliminations               int `json:"eliminations"`
						UltsEarned                 int `json:"ultsEarned"`
						KnockbackKills             int `json:"knockbackKills"`
						UltsUsed                   int `json:"ultsUsed"`
						Deaths                     int `json:"deaths"`
						TimeSpentOnFire            int `json:"timeSpentOnFire"`
						FinalBlows                 int `json:"finalBlows"`
						SoloKills                  int `json:"soloKills"`
						HeroDamageDone             int `json:"heroDamageDone"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					} `json:"stats"`
					Heroes struct {
						Winston struct {
							DamageTaken     int `json:"damageTaken"`
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
							SoloKills       int `json:"soloKills"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"winston"`
						Zarya struct {
							DamageTaken                int `json:"damageTaken"`
							EnergyMax                  int `json:"energyMax"`
							ShotsHit                   int `json:"shotsHit"`
							TimePlayed                 int `json:"timePlayed"`
							Eliminations               int `json:"eliminations"`
							LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
							UltsEarned                 int `json:"ultsEarned"`
							UltsUsed                   int `json:"ultsUsed"`
							Deaths                     int `json:"deaths"`
							TimeSpentOnFire            int `json:"timeSpentOnFire"`
							GravitonSurgeKills         int `json:"gravitonSurgeKills"`
							FinalBlows                 int `json:"finalBlows"`
							HeroDamageDone             int `json:"heroDamageDone"`
						} `json:"zarya"`
					} `json:"heroes"`
				} `json:"16240"`
				Num16241 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						DamageTaken       int `json:"damageTaken"`
						PrimalRageKills   int `json:"primalRageKills"`
						JumpPackKills     int `json:"jumpPackKills"`
						ShotsHit          int `json:"shotsHit"`
						TimePlayed        int `json:"timePlayed"`
						Eliminations      int `json:"eliminations"`
						UltsEarned        int `json:"ultsEarned"`
						KnockbackKills    int `json:"knockbackKills"`
						UltsUsed          int `json:"ultsUsed"`
						Deaths            int `json:"deaths"`
						FinalBlows        int `json:"finalBlows"`
						SoloKills         int `json:"soloKills"`
						HeroDamageDone    int `json:"heroDamageDone"`
						CriticalHits      int `json:"criticalHits"`
						AbilityDamageDone int `json:"abilityDamageDone"`
					} `json:"stats"`
					Heroes struct {
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
							FinalBlows      int `json:"finalBlows"`
							SoloKills       int `json:"soloKills"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"winston"`
						Doomfist struct {
							HeroDamageDone    int `json:"heroDamageDone"`
							DamageTaken       int `json:"damageTaken"`
							Deaths            int `json:"deaths"`
							TimePlayed        int `json:"timePlayed"`
							CriticalHits      int `json:"criticalHits"`
							ShotsHit          int `json:"shotsHit"`
							AbilityDamageDone int `json:"abilityDamageDone"`
						} `json:"doomfist"`
					} `json:"heroes"`
				} `json:"16241"`
				Num16554 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						HeroDamageDone   int `json:"heroDamageDone"`
						DamageTaken      int `json:"damageTaken"`
						FinalBlows       int `json:"finalBlows"`
						Eliminations     int `json:"eliminations"`
						TimeSpentOnFire  int `json:"timeSpentOnFire"`
						TimePlayed       int `json:"timePlayed"`
						CriticalHits     int `json:"criticalHits"`
						ShotsHit         int `json:"shotsHit"`
						DragonbladeKills int `json:"dragonbladeKills"`
						UltsEarned       int `json:"ultsEarned"`
						UltsUsed         int `json:"ultsUsed"`
						Deaths           int `json:"deaths"`
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"16554"`
			} `json:"players"`
			State string `json:"state"`
		} `json:"60411"`
		Num60413 struct {
			Teams struct {
				Num15188 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
					} `json:"teamStats"`
				} `json:"15188"`
				Num15206 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						HeroDamageDone int `json:"heroDamageDone"`
						FinalBlows     int `json:"finalBlows"`
					} `json:"teamStats"`
				} `json:"15206"`
			} `json:"teams"`
			ID      int    `json:"id"`
			Map     string `json:"map"`
			MapType string `json:"mapType"`
			Number  int    `json:"number"`
			Players struct {
				Num15480 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
							CriticalHits   int `json:"criticalHits"`
							DamageTaken    int `json:"damageTaken"`
							ShotsHit       int `json:"shotsHit"`
							TimePlayed     int `json:"timePlayed"`
							Eliminations   int `json:"eliminations"`
							UltsEarned     int `json:"ultsEarned"`
							UltsUsed       int `json:"ultsUsed"`
							Deaths         int `json:"deaths"`
							FinalBlows     int `json:"finalBlows"`
							HeroDamageDone int `json:"heroDamageDone"`
						} `json:"tracer"`
					} `json:"heroes"`
				} `json:"15480"`
				Num15622 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15622"`
				Num15643 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits       int `json:"criticalHits"`
						DamageTaken        int `json:"damageTaken"`
						HelixRocketKills   int `json:"helixRocketKills"`
						ShotsHit           int `json:"shotsHit"`
						TimePlayed         int `json:"timePlayed"`
						Eliminations       int `json:"eliminations"`
						UltsEarned         int `json:"ultsEarned"`
						UltsUsed           int `json:"ultsUsed"`
						TacticalVisorKills int `json:"tacticalVisorKills"`
						Deaths             int `json:"deaths"`
						HealingDone        int `json:"healingDone"`
						FinalBlows         int `json:"finalBlows"`
						SoloKills          int `json:"soloKills"`
						HeroDamageDone     int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Soldier76 struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							HelixRocketKills   int `json:"helixRocketKills"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							TacticalVisorKills int `json:"tacticalVisorKills"`
							Deaths             int `json:"deaths"`
							HealingDone        int `json:"healingDone"`
							FinalBlows         int `json:"finalBlows"`
							SoloKills          int `json:"soloKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"soldier-76"`
					} `json:"heroes"`
				} `json:"15643"`
				Num15787 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						HealingDone        int `json:"healingDone"`
						HeroDamageDone     int `json:"heroDamageDone"`
					} `json:"stats"`
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
							HealingDone        int `json:"healingDone"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"ana"`
					} `json:"heroes"`
				} `json:"15787"`
				Num15937 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"15937"`
				Num16205 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"16205"`
				Num16232 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits    int `json:"criticalHits"`
						DamageTaken     int `json:"damageTaken"`
						ShotsHit        int `json:"shotsHit"`
						TimePlayed      int `json:"timePlayed"`
						Eliminations    int `json:"eliminations"`
						UltsEarned      int `json:"ultsEarned"`
						KnockbackKills  int `json:"knockbackKills"`
						UltsUsed        int `json:"ultsUsed"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						HealingDone     int `json:"healingDone"`
						FinalBlows      int `json:"finalBlows"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							CriticalHits    int `json:"criticalHits"`
							DamageTaken     int `json:"damageTaken"`
							ShotsHit        int `json:"shotsHit"`
							TimePlayed      int `json:"timePlayed"`
							Eliminations    int `json:"eliminations"`
							UltsEarned      int `json:"ultsEarned"`
							KnockbackKills  int `json:"knockbackKills"`
							UltsUsed        int `json:"ultsUsed"`
							Deaths          int `json:"deaths"`
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							HealingDone     int `json:"healingDone"`
							FinalBlows      int `json:"finalBlows"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"16232"`
				Num16240 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						SoloKills       int `json:"soloKills"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
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
							SoloKills       int `json:"soloKills"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"winston"`
					} `json:"heroes"`
				} `json:"16240"`
				Num16241 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"16241"`
				Num16554 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits       int `json:"criticalHits"`
						DamageTaken        int `json:"damageTaken"`
						ShotsHit           int `json:"shotsHit"`
						TimePlayed         int `json:"timePlayed"`
						Eliminations       int `json:"eliminations"`
						UltsEarned         int `json:"ultsEarned"`
						UltsUsed           int `json:"ultsUsed"`
						PulseBombsAttached int `json:"pulseBombsAttached"`
						Deaths             int `json:"deaths"`
						TimeSpentOnFire    int `json:"timeSpentOnFire"`
						FinalBlows         int `json:"finalBlows"`
						PulseBombKills     int `json:"pulseBombKills"`
						SoloKills          int `json:"soloKills"`
						HeroDamageDone     int `json:"heroDamageDone"`
						DragonbladeKills   int `json:"dragonbladeKills"`
					} `json:"stats"`
					Heroes struct {
						Tracer struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							PulseBombsAttached int `json:"pulseBombsAttached"`
							Deaths             int `json:"deaths"`
							TimeSpentOnFire    int `json:"timeSpentOnFire"`
							FinalBlows         int `json:"finalBlows"`
							PulseBombKills     int `json:"pulseBombKills"`
							SoloKills          int `json:"soloKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"tracer"`
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
							FinalBlows       int `json:"finalBlows"`
							SoloKills        int `json:"soloKills"`
							HeroDamageDone   int `json:"heroDamageDone"`
						} `json:"genji"`
					} `json:"heroes"`
				} `json:"16554"`
			} `json:"players"`
			State string `json:"state"`
		} `json:"60413"`
		Num60414 struct {
			ID     int    `json:"id"`
			Number int    `json:"number"`
			State  string `json:"state"`
			Teams  struct {
				Num15188 struct {
					ID        int `json:"id"`
					TeamStats struct {
					} `json:"teamStats"`
				} `json:"15188"`
				Num15206 struct {
					ID        int `json:"id"`
					TeamStats struct {
					} `json:"teamStats"`
				} `json:"15206"`
			} `json:"teams"`
			Map     string `json:"map"`
			MapType string `json:"mapType"`
			Players struct {
			} `json:"players"`
		} `json:"60414"`
		Num60432 struct {
			Teams struct {
				Num15188 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
						HealingDone    int `json:"healingDone"`
					} `json:"teamStats"`
				} `json:"15188"`
				Num15206 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
					} `json:"teamStats"`
				} `json:"15206"`
			} `json:"teams"`
			ID      int    `json:"id"`
			Map     string `json:"map"`
			MapType string `json:"mapType"`
			Number  int    `json:"number"`
			Players struct {
				Num15480 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						HealingDone            int `json:"healingDone"`
					} `json:"stats"`
					Heroes struct {
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
						Soldier76 struct {
							CriticalHits   int `json:"criticalHits"`
							DamageTaken    int `json:"damageTaken"`
							ShotsHit       int `json:"shotsHit"`
							TimePlayed     int `json:"timePlayed"`
							Eliminations   int `json:"eliminations"`
							UltsEarned     int `json:"ultsEarned"`
							UltsUsed       int `json:"ultsUsed"`
							Deaths         int `json:"deaths"`
							HealingDone    int `json:"healingDone"`
							FinalBlows     int `json:"finalBlows"`
							HeroDamageDone int `json:"heroDamageDone"`
						} `json:"soldier-76"`
						Reaper struct {
							HeroDamageDone int `json:"heroDamageDone"`
							HealingDone    int `json:"healingDone"`
							DamageTaken    int `json:"damageTaken"`
							Deaths         int `json:"deaths"`
							TimePlayed     int `json:"timePlayed"`
							CriticalHits   int `json:"criticalHits"`
							ShotsHit       int `json:"shotsHit"`
						} `json:"reaper"`
					} `json:"heroes"`
				} `json:"15480"`
				Num15622 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15622"`
				Num15643 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits       int `json:"criticalHits"`
						DamageTaken        int `json:"damageTaken"`
						HelixRocketKills   int `json:"helixRocketKills"`
						ShotsHit           int `json:"shotsHit"`
						TimePlayed         int `json:"timePlayed"`
						Eliminations       int `json:"eliminations"`
						UltsEarned         int `json:"ultsEarned"`
						UltsUsed           int `json:"ultsUsed"`
						TacticalVisorKills int `json:"tacticalVisorKills"`
						Deaths             int `json:"deaths"`
						TimeSpentOnFire    int `json:"timeSpentOnFire"`
						HealingDone        int `json:"healingDone"`
						FinalBlows         int `json:"finalBlows"`
						SoloKills          int `json:"soloKills"`
						HeroDamageDone     int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Soldier76 struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							HelixRocketKills   int `json:"helixRocketKills"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							TacticalVisorKills int `json:"tacticalVisorKills"`
							Deaths             int `json:"deaths"`
							TimeSpentOnFire    int `json:"timeSpentOnFire"`
							HealingDone        int `json:"healingDone"`
							FinalBlows         int `json:"finalBlows"`
							SoloKills          int `json:"soloKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"soldier-76"`
					} `json:"heroes"`
				} `json:"15643"`
				Num15787 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15787"`
				Num15937 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"15937"`
				Num16205 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits    int `json:"criticalHits"`
						DamageTaken     int `json:"damageTaken"`
						ShotsHit        int `json:"shotsHit"`
						TimePlayed      int `json:"timePlayed"`
						Eliminations    int `json:"eliminations"`
						UltsEarned      int `json:"ultsEarned"`
						KnockbackKills  int `json:"knockbackKills"`
						UltsUsed        int `json:"ultsUsed"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						HealingDone     int `json:"healingDone"`
						FinalBlows      int `json:"finalBlows"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							CriticalHits    int `json:"criticalHits"`
							DamageTaken     int `json:"damageTaken"`
							ShotsHit        int `json:"shotsHit"`
							TimePlayed      int `json:"timePlayed"`
							Eliminations    int `json:"eliminations"`
							UltsEarned      int `json:"ultsEarned"`
							KnockbackKills  int `json:"knockbackKills"`
							UltsUsed        int `json:"ultsUsed"`
							Deaths          int `json:"deaths"`
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							HealingDone     int `json:"healingDone"`
							FinalBlows      int `json:"finalBlows"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"16205"`
				Num16232 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits    int `json:"criticalHits"`
						DamageTaken     int `json:"damageTaken"`
						ShotsHit        int `json:"shotsHit"`
						TimePlayed      int `json:"timePlayed"`
						Eliminations    int `json:"eliminations"`
						UltsEarned      int `json:"ultsEarned"`
						KnockbackKills  int `json:"knockbackKills"`
						UltsUsed        int `json:"ultsUsed"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						HealingDone     int `json:"healingDone"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							CriticalHits    int `json:"criticalHits"`
							DamageTaken     int `json:"damageTaken"`
							ShotsHit        int `json:"shotsHit"`
							TimePlayed      int `json:"timePlayed"`
							Eliminations    int `json:"eliminations"`
							UltsEarned      int `json:"ultsEarned"`
							KnockbackKills  int `json:"knockbackKills"`
							UltsUsed        int `json:"ultsUsed"`
							Deaths          int `json:"deaths"`
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							HealingDone     int `json:"healingDone"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"16232"`
				Num16240 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits               int `json:"criticalHits"`
						DamageTaken                int `json:"damageTaken"`
						ShotsHit                   int `json:"shotsHit"`
						TimePlayed                 int `json:"timePlayed"`
						Eliminations               int `json:"eliminations"`
						UltsEarned                 int `json:"ultsEarned"`
						KnockbackKills             int `json:"knockbackKills"`
						UltsUsed                   int `json:"ultsUsed"`
						Deaths                     int `json:"deaths"`
						FinalBlows                 int `json:"finalBlows"`
						HeroDamageDone             int `json:"heroDamageDone"`
						JumpPackKills              int `json:"jumpPackKills"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					} `json:"stats"`
					Heroes struct {
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
						Winston struct {
							DamageTaken    int `json:"damageTaken"`
							JumpPackKills  int `json:"jumpPackKills"`
							ShotsHit       int `json:"shotsHit"`
							TimePlayed     int `json:"timePlayed"`
							Eliminations   int `json:"eliminations"`
							UltsEarned     int `json:"ultsEarned"`
							UltsUsed       int `json:"ultsUsed"`
							Deaths         int `json:"deaths"`
							FinalBlows     int `json:"finalBlows"`
							HeroDamageDone int `json:"heroDamageDone"`
						} `json:"winston"`
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
					} `json:"heroes"`
				} `json:"16240"`
				Num16241 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						DamageTaken       int `json:"damageTaken"`
						PrimalRageKills   int `json:"primalRageKills"`
						JumpPackKills     int `json:"jumpPackKills"`
						ShotsHit          int `json:"shotsHit"`
						TimePlayed        int `json:"timePlayed"`
						Eliminations      int `json:"eliminations"`
						UltsEarned        int `json:"ultsEarned"`
						KnockbackKills    int `json:"knockbackKills"`
						UltsUsed          int `json:"ultsUsed"`
						Deaths            int `json:"deaths"`
						TimeSpentOnFire   int `json:"timeSpentOnFire"`
						FinalBlows        int `json:"finalBlows"`
						HeroDamageDone    int `json:"heroDamageDone"`
						AbilityDamageDone int `json:"abilityDamageDone"`
					} `json:"stats"`
					Heroes struct {
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
						Doomfist struct {
							HeroDamageDone    int `json:"heroDamageDone"`
							DamageTaken       int `json:"damageTaken"`
							TimePlayed        int `json:"timePlayed"`
							ShotsHit          int `json:"shotsHit"`
							AbilityDamageDone int `json:"abilityDamageDone"`
						} `json:"doomfist"`
					} `json:"heroes"`
				} `json:"16241"`
				Num16554 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits           int `json:"criticalHits"`
						DamageTaken            int `json:"damageTaken"`
						ShotsHit               int `json:"shotsHit"`
						ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
						TimePlayed             int `json:"timePlayed"`
						Eliminations           int `json:"eliminations"`
						Deaths                 int `json:"deaths"`
						ScopedHits             int `json:"scopedHits"`
						FinalBlows             int `json:"finalBlows"`
						ScopedCriticalHits     int `json:"scopedCriticalHits"`
						HeroDamageDone         int `json:"heroDamageDone"`
						UltsEarned             int `json:"ultsEarned"`
						UltsUsed               int `json:"ultsUsed"`
						SoloKills              int `json:"soloKills"`
						DragonbladeKills       int `json:"dragonbladeKills"`
						TimeSpentOnFire        int `json:"timeSpentOnFire"`
					} `json:"stats"`
					Heroes struct {
						Widowmaker struct {
							CriticalHits           int `json:"criticalHits"`
							DamageTaken            int `json:"damageTaken"`
							ShotsHit               int `json:"shotsHit"`
							ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
							TimePlayed             int `json:"timePlayed"`
							Eliminations           int `json:"eliminations"`
							Deaths                 int `json:"deaths"`
							ScopedHits             int `json:"scopedHits"`
							FinalBlows             int `json:"finalBlows"`
							ScopedCriticalHits     int `json:"scopedCriticalHits"`
							HeroDamageDone         int `json:"heroDamageDone"`
						} `json:"widowmaker"`
						Tracer struct {
							CriticalHits   int `json:"criticalHits"`
							DamageTaken    int `json:"damageTaken"`
							ShotsHit       int `json:"shotsHit"`
							TimePlayed     int `json:"timePlayed"`
							Eliminations   int `json:"eliminations"`
							UltsEarned     int `json:"ultsEarned"`
							UltsUsed       int `json:"ultsUsed"`
							Deaths         int `json:"deaths"`
							FinalBlows     int `json:"finalBlows"`
							SoloKills      int `json:"soloKills"`
							HeroDamageDone int `json:"heroDamageDone"`
						} `json:"tracer"`
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
					} `json:"heroes"`
				} `json:"16554"`
			} `json:"players"`
			State string `json:"state"`
		} `json:"60432"`
		Num60433 struct {
			Teams struct {
				Num15188 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
						HealingDone    int `json:"healingDone"`
					} `json:"teamStats"`
				} `json:"15188"`
				Num15206 struct {
					ID        int `json:"id"`
					Score     int `json:"score"`
					TeamStats struct {
						DamageTaken    int `json:"damageTaken"`
						Eliminations   int `json:"eliminations"`
						Deaths         int `json:"deaths"`
						HealingDone    int `json:"healingDone"`
						FinalBlows     int `json:"finalBlows"`
						HeroDamageDone int `json:"heroDamageDone"`
					} `json:"teamStats"`
				} `json:"15206"`
			} `json:"teams"`
			ID      int    `json:"id"`
			Map     string `json:"map"`
			MapType string `json:"mapType"`
			Number  int    `json:"number"`
			Players struct {
				Num15480 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						HealingDone            int `json:"healingDone"`
						PulseBombsAttached     int `json:"pulseBombsAttached"`
						TimeSpentOnFire        int `json:"timeSpentOnFire"`
						PulseBombKills         int `json:"pulseBombKills"`
					} `json:"stats"`
					Heroes struct {
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
						Reaper struct {
							CriticalHits   int `json:"criticalHits"`
							DamageTaken    int `json:"damageTaken"`
							ShotsHit       int `json:"shotsHit"`
							TimePlayed     int `json:"timePlayed"`
							Eliminations   int `json:"eliminations"`
							Deaths         int `json:"deaths"`
							HealingDone    int `json:"healingDone"`
							FinalBlows     int `json:"finalBlows"`
							SoloKills      int `json:"soloKills"`
							HeroDamageDone int `json:"heroDamageDone"`
						} `json:"reaper"`
						Tracer struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							PulseBombsAttached int `json:"pulseBombsAttached"`
							Deaths             int `json:"deaths"`
							TimeSpentOnFire    int `json:"timeSpentOnFire"`
							FinalBlows         int `json:"finalBlows"`
							PulseBombKills     int `json:"pulseBombKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"tracer"`
					} `json:"heroes"`
				} `json:"15480"`
				Num15622 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
					Heroes struct {
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
						Lucio struct {
							TimePlayed int `json:"timePlayed"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"15622"`
				Num15643 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits       int `json:"criticalHits"`
						DamageTaken        int `json:"damageTaken"`
						ShotsHit           int `json:"shotsHit"`
						TimePlayed         int `json:"timePlayed"`
						Eliminations       int `json:"eliminations"`
						UltsEarned         int `json:"ultsEarned"`
						UltsUsed           int `json:"ultsUsed"`
						PulseBombsAttached int `json:"pulseBombsAttached"`
						Deaths             int `json:"deaths"`
						FinalBlows         int `json:"finalBlows"`
						PulseBombKills     int `json:"pulseBombKills"`
						SoloKills          int `json:"soloKills"`
						HeroDamageDone     int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Tracer struct {
							CriticalHits       int `json:"criticalHits"`
							DamageTaken        int `json:"damageTaken"`
							ShotsHit           int `json:"shotsHit"`
							TimePlayed         int `json:"timePlayed"`
							Eliminations       int `json:"eliminations"`
							UltsEarned         int `json:"ultsEarned"`
							UltsUsed           int `json:"ultsUsed"`
							PulseBombsAttached int `json:"pulseBombsAttached"`
							Deaths             int `json:"deaths"`
							FinalBlows         int `json:"finalBlows"`
							PulseBombKills     int `json:"pulseBombKills"`
							SoloKills          int `json:"soloKills"`
							HeroDamageDone     int `json:"heroDamageDone"`
						} `json:"tracer"`
					} `json:"heroes"`
				} `json:"15643"`
				Num15648 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits           int `json:"criticalHits"`
						DamageTaken            int `json:"damageTaken"`
						HelixRocketKills       int `json:"helixRocketKills"`
						ShotsHit               int `json:"shotsHit"`
						TimePlayed             int `json:"timePlayed"`
						Eliminations           int `json:"eliminations"`
						UltsEarned             int `json:"ultsEarned"`
						UltsUsed               int `json:"ultsUsed"`
						Deaths                 int `json:"deaths"`
						HealingDone            int `json:"healingDone"`
						FinalBlows             int `json:"finalBlows"`
						HeroDamageDone         int `json:"heroDamageDone"`
						ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
						ScopedHits             int `json:"scopedHits"`
						ScopedCriticalHits     int `json:"scopedCriticalHits"`
						SoloKills              int `json:"soloKills"`
					} `json:"stats"`
					Heroes struct {
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
					} `json:"heroes"`
				} `json:"15648"`
				Num15787 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
					} `json:"stats"`
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
					} `json:"heroes"`
				} `json:"15787"`
				Num15937 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
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
						TimeSpentOnFire   int `json:"timeSpentOnFire"`
						FinalBlows        int `json:"finalBlows"`
						SoloKills         int `json:"soloKills"`
						HeroDamageDone    int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
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
							TimeSpentOnFire   int `json:"timeSpentOnFire"`
							FinalBlows        int `json:"finalBlows"`
							SoloKills         int `json:"soloKills"`
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"echo"`
					} `json:"heroes"`
				} `json:"15937"`
				Num16232 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits    int `json:"criticalHits"`
						DamageTaken     int `json:"damageTaken"`
						ShotsHit        int `json:"shotsHit"`
						TimePlayed      int `json:"timePlayed"`
						Eliminations    int `json:"eliminations"`
						UltsEarned      int `json:"ultsEarned"`
						KnockbackKills  int `json:"knockbackKills"`
						UltsUsed        int `json:"ultsUsed"`
						Deaths          int `json:"deaths"`
						TimeSpentOnFire int `json:"timeSpentOnFire"`
						HealingDone     int `json:"healingDone"`
						FinalBlows      int `json:"finalBlows"`
						HeroDamageDone  int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							CriticalHits    int `json:"criticalHits"`
							DamageTaken     int `json:"damageTaken"`
							ShotsHit        int `json:"shotsHit"`
							TimePlayed      int `json:"timePlayed"`
							Eliminations    int `json:"eliminations"`
							UltsEarned      int `json:"ultsEarned"`
							KnockbackKills  int `json:"knockbackKills"`
							UltsUsed        int `json:"ultsUsed"`
							Deaths          int `json:"deaths"`
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							HealingDone     int `json:"healingDone"`
							FinalBlows      int `json:"finalBlows"`
							HeroDamageDone  int `json:"heroDamageDone"`
						} `json:"lucio"`
					} `json:"heroes"`
				} `json:"16232"`
				Num16240 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						CriticalHits               int `json:"criticalHits"`
						DamageTaken                int `json:"damageTaken"`
						ShotsHit                   int `json:"shotsHit"`
						TimePlayed                 int `json:"timePlayed"`
						Eliminations               int `json:"eliminations"`
						KnockbackKills             int `json:"knockbackKills"`
						TimeSpentOnFire            int `json:"timeSpentOnFire"`
						FinalBlows                 int `json:"finalBlows"`
						AbilityDamageDone          int `json:"abilityDamageDone"`
						HeroDamageDone             int `json:"heroDamageDone"`
						EnergyMax                  int `json:"energyMax"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						UltsEarned                 int `json:"ultsEarned"`
						UltsUsed                   int `json:"ultsUsed"`
						Deaths                     int `json:"deaths"`
						GravitonSurgeKills         int `json:"gravitonSurgeKills"`
						SoloKills                  int `json:"soloKills"`
					} `json:"stats"`
					Heroes struct {
						Doomfist struct {
							CriticalHits      int `json:"criticalHits"`
							DamageTaken       int `json:"damageTaken"`
							ShotsHit          int `json:"shotsHit"`
							TimePlayed        int `json:"timePlayed"`
							Eliminations      int `json:"eliminations"`
							KnockbackKills    int `json:"knockbackKills"`
							TimeSpentOnFire   int `json:"timeSpentOnFire"`
							FinalBlows        int `json:"finalBlows"`
							AbilityDamageDone int `json:"abilityDamageDone"`
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"doomfist"`
						Winston struct {
							TimeSpentOnFire int `json:"timeSpentOnFire"`
							TimePlayed      int `json:"timePlayed"`
						} `json:"winston"`
						Zarya struct {
							DamageTaken                int `json:"damageTaken"`
							EnergyMax                  int `json:"energyMax"`
							ShotsHit                   int `json:"shotsHit"`
							TimePlayed                 int `json:"timePlayed"`
							Eliminations               int `json:"eliminations"`
							LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
							UltsEarned                 int `json:"ultsEarned"`
							UltsUsed                   int `json:"ultsUsed"`
							Deaths                     int `json:"deaths"`
							TimeSpentOnFire            int `json:"timeSpentOnFire"`
							GravitonSurgeKills         int `json:"gravitonSurgeKills"`
							FinalBlows                 int `json:"finalBlows"`
							SoloKills                  int `json:"soloKills"`
							HeroDamageDone             int `json:"heroDamageDone"`
						} `json:"zarya"`
					} `json:"heroes"`
				} `json:"16240"`
				Num16288 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						HeroDamageDone             int `json:"heroDamageDone"`
						DamageTaken                int `json:"damageTaken"`
						Deaths                     int `json:"deaths"`
						TimePlayed                 int `json:"timePlayed"`
						CriticalHits               int `json:"criticalHits"`
						ShotsHit                   int `json:"shotsHit"`
						EnergyMax                  int `json:"energyMax"`
						Eliminations               int `json:"eliminations"`
						LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
						UltsEarned                 int `json:"ultsEarned"`
						UltsUsed                   int `json:"ultsUsed"`
						GraviticFluxKills          int `json:"graviticFluxKills"`
						AccretionKills             int `json:"accretionKills"`
						TimeSpentOnFire            int `json:"timeSpentOnFire"`
						FinalBlows                 int `json:"finalBlows"`
					} `json:"stats"`
					Heroes struct {
						DVa struct {
							HeroDamageDone int `json:"heroDamageDone"`
							DamageTaken    int `json:"damageTaken"`
							Deaths         int `json:"deaths"`
							TimePlayed     int `json:"timePlayed"`
							CriticalHits   int `json:"criticalHits"`
							ShotsHit       int `json:"shotsHit"`
						} `json:"d-va"`
						Zarya struct {
							DamageTaken                int `json:"damageTaken"`
							EnergyMax                  int `json:"energyMax"`
							ShotsHit                   int `json:"shotsHit"`
							TimePlayed                 int `json:"timePlayed"`
							Eliminations               int `json:"eliminations"`
							LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
							UltsEarned                 int `json:"ultsEarned"`
							UltsUsed                   int `json:"ultsUsed"`
							Deaths                     int `json:"deaths"`
							HeroDamageDone             int `json:"heroDamageDone"`
						} `json:"zarya"`
						Sigma struct {
							DamageTaken       int `json:"damageTaken"`
							GraviticFluxKills int `json:"graviticFluxKills"`
							TimePlayed        int `json:"timePlayed"`
							Eliminations      int `json:"eliminations"`
							UltsEarned        int `json:"ultsEarned"`
							UltsUsed          int `json:"ultsUsed"`
							AccretionKills    int `json:"accretionKills"`
							Deaths            int `json:"deaths"`
							TimeSpentOnFire   int `json:"timeSpentOnFire"`
							FinalBlows        int `json:"finalBlows"`
							HeroDamageDone    int `json:"heroDamageDone"`
						} `json:"sigma"`
					} `json:"heroes"`
				} `json:"16288"`
				Num16530 struct {
					ID     int `json:"id"`
					TeamID int `json:"teamId"`
					Stats  struct {
						TimePlayed       int `json:"timePlayed"`
						CriticalHits     int `json:"criticalHits"`
						DamageTaken      int `json:"damageTaken"`
						ShotsHit         int `json:"shotsHit"`
						Eliminations     int `json:"eliminations"`
						UltsEarned       int `json:"ultsEarned"`
						DamageDiscordOrb int `json:"damageDiscordOrb"`
						UltsUsed         int `json:"ultsUsed"`
						Deaths           int `json:"deaths"`
						HealingDone      int `json:"healingDone"`
						FinalBlows       int `json:"finalBlows"`
						HeroDamageDone   int `json:"heroDamageDone"`
					} `json:"stats"`
					Heroes struct {
						Lucio struct {
							TimePlayed int `json:"timePlayed"`
						} `json:"lucio"`
						Zenyatta struct {
							CriticalHits     int `json:"criticalHits"`
							DamageTaken      int `json:"damageTaken"`
							ShotsHit         int `json:"shotsHit"`
							TimePlayed       int `json:"timePlayed"`
							Eliminations     int `json:"eliminations"`
							UltsEarned       int `json:"ultsEarned"`
							DamageDiscordOrb int `json:"damageDiscordOrb"`
							UltsUsed         int `json:"ultsUsed"`
							Deaths           int `json:"deaths"`
							HealingDone      int `json:"healingDone"`
							FinalBlows       int `json:"finalBlows"`
							HeroDamageDone   int `json:"heroDamageDone"`
						} `json:"zenyatta"`
					} `json:"heroes"`
				} `json:"16530"`
			} `json:"players"`
			State string `json:"state"`
		} `json:"60433"`
	} `json:"games"`
	ID                   int    `json:"id"`
	LocalTimeZone        string `json:"localTimeZone"`
	LocalScheduledDate   string `json:"localScheduledDate"`
	SeasonID             string `json:"seasonId"`
	SegmentID            string `json:"segmentId"`
	StartTimestamp       int64  `json:"startTimestamp"`
	ActualStartTimestamp int64  `json:"actualStartTimestamp"`
	ActualEndTimestamp   int64  `json:"actualEndTimestamp"`
	State                string `json:"state"`
	Teams                struct {
		Num15188 struct {
			ID        int `json:"id"`
			Score     int `json:"score"`
			TeamStats struct {
				DamageTaken    int `json:"damageTaken"`
				Eliminations   int `json:"eliminations"`
				Deaths         int `json:"deaths"`
				FinalBlows     int `json:"finalBlows"`
				HeroDamageDone int `json:"heroDamageDone"`
				HealingDone    int `json:"healingDone"`
			} `json:"teamStats"`
		} `json:"15188"`
		Num15206 struct {
			ID        int `json:"id"`
			Score     int `json:"score"`
			TeamStats struct {
				DamageTaken    int `json:"damageTaken"`
				Eliminations   int `json:"eliminations"`
				Deaths         int `json:"deaths"`
				HealingDone    int `json:"healingDone"`
				FinalBlows     int `json:"finalBlows"`
				HeroDamageDone int `json:"heroDamageDone"`
			} `json:"teamStats"`
		} `json:"15206"`
	} `json:"teams"`
	Players struct {
		Num15480 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
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
				HealingDone            int `json:"healingDone"`
				HelixRocketKills       int `json:"helixRocketKills"`
				DeathBlossomKills      int `json:"deathBlossomKills"`
				PulseBombsAttached     int `json:"pulseBombsAttached"`
				PulseBombKills         int `json:"pulseBombKills"`
			} `json:"stats"`
			Heroes struct {
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
				Soldier76 struct {
					CriticalHits     int `json:"criticalHits"`
					DamageTaken      int `json:"damageTaken"`
					ShotsHit         int `json:"shotsHit"`
					TimePlayed       int `json:"timePlayed"`
					Eliminations     int `json:"eliminations"`
					UltsEarned       int `json:"ultsEarned"`
					UltsUsed         int `json:"ultsUsed"`
					Deaths           int `json:"deaths"`
					HealingDone      int `json:"healingDone"`
					FinalBlows       int `json:"finalBlows"`
					HeroDamageDone   int `json:"heroDamageDone"`
					HelixRocketKills int `json:"helixRocketKills"`
					TimeSpentOnFire  int `json:"timeSpentOnFire"`
				} `json:"soldier-76"`
				Reaper struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					HealingDone       int `json:"healingDone"`
					DamageTaken       int `json:"damageTaken"`
					Deaths            int `json:"deaths"`
					TimePlayed        int `json:"timePlayed"`
					CriticalHits      int `json:"criticalHits"`
					ShotsHit          int `json:"shotsHit"`
					Eliminations      int `json:"eliminations"`
					UltsEarned        int `json:"ultsEarned"`
					UltsUsed          int `json:"ultsUsed"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					FinalBlows        int `json:"finalBlows"`
					DeathBlossomKills int `json:"deathBlossomKills"`
					SoloKills         int `json:"soloKills"`
				} `json:"reaper"`
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
				Tracer struct {
					CriticalHits       int `json:"criticalHits"`
					DamageTaken        int `json:"damageTaken"`
					ShotsHit           int `json:"shotsHit"`
					TimePlayed         int `json:"timePlayed"`
					Eliminations       int `json:"eliminations"`
					UltsEarned         int `json:"ultsEarned"`
					UltsUsed           int `json:"ultsUsed"`
					PulseBombsAttached int `json:"pulseBombsAttached"`
					Deaths             int `json:"deaths"`
					TimeSpentOnFire    int `json:"timeSpentOnFire"`
					FinalBlows         int `json:"finalBlows"`
					PulseBombKills     int `json:"pulseBombKills"`
					HeroDamageDone     int `json:"heroDamageDone"`
				} `json:"tracer"`
			} `json:"heroes"`
		} `json:"15480"`
		Num15622 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				DamageTaken                     int `json:"damageTaken"`
				EnemiesSlept                    int `json:"enemiesSlept"`
				ShotsHit                        int `json:"shotsHit"`
				TimePlayed                      int `json:"timePlayed"`
				Eliminations                    int `json:"eliminations"`
				UltsEarned                      int `json:"ultsEarned"`
				UltsUsed                        int `json:"ultsUsed"`
				Deaths                          int `json:"deaths"`
				ScopedHits                      int `json:"scopedHits"`
				BioticGrenadeKills              int `json:"bioticGrenadeKills"`
				TimeSpentOnFire                 int `json:"timeSpentOnFire"`
				HealingDone                     int `json:"healingDone"`
				FinalBlows                      int `json:"finalBlows"`
				HeroDamageDone                  int `json:"heroDamageDone"`
				CriticalHits                    int `json:"criticalHits"`
				AmplificationMatrixAssists      int `json:"amplificationMatrixAssists"`
				ImmortalityFieldDeathsPrevented int `json:"immortalityFieldDeathsPrevented"`
				DamageAmplified                 int `json:"damageAmplified"`
				SoloKills                       int `json:"soloKills"`
			} `json:"stats"`
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
				Lucio struct {
					TimePlayed int `json:"timePlayed"`
				} `json:"lucio"`
			} `json:"heroes"`
		} `json:"15622"`
		Num15643 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits       int `json:"criticalHits"`
				DamageTaken        int `json:"damageTaken"`
				HelixRocketKills   int `json:"helixRocketKills"`
				ShotsHit           int `json:"shotsHit"`
				TimePlayed         int `json:"timePlayed"`
				Eliminations       int `json:"eliminations"`
				UltsEarned         int `json:"ultsEarned"`
				UltsUsed           int `json:"ultsUsed"`
				TacticalVisorKills int `json:"tacticalVisorKills"`
				Deaths             int `json:"deaths"`
				TimeSpentOnFire    int `json:"timeSpentOnFire"`
				HealingDone        int `json:"healingDone"`
				FinalBlows         int `json:"finalBlows"`
				SoloKills          int `json:"soloKills"`
				HeroDamageDone     int `json:"heroDamageDone"`
				DeathBlossomKills  int `json:"deathBlossomKills"`
				PulseBombsAttached int `json:"pulseBombsAttached"`
				PulseBombKills     int `json:"pulseBombKills"`
			} `json:"stats"`
			Heroes struct {
				Soldier76 struct {
					CriticalHits       int `json:"criticalHits"`
					DamageTaken        int `json:"damageTaken"`
					HelixRocketKills   int `json:"helixRocketKills"`
					ShotsHit           int `json:"shotsHit"`
					TimePlayed         int `json:"timePlayed"`
					Eliminations       int `json:"eliminations"`
					UltsEarned         int `json:"ultsEarned"`
					UltsUsed           int `json:"ultsUsed"`
					TacticalVisorKills int `json:"tacticalVisorKills"`
					Deaths             int `json:"deaths"`
					TimeSpentOnFire    int `json:"timeSpentOnFire"`
					HealingDone        int `json:"healingDone"`
					FinalBlows         int `json:"finalBlows"`
					SoloKills          int `json:"soloKills"`
					HeroDamageDone     int `json:"heroDamageDone"`
				} `json:"soldier-76"`
				Reaper struct {
					CriticalHits      int `json:"criticalHits"`
					DamageTaken       int `json:"damageTaken"`
					ShotsHit          int `json:"shotsHit"`
					TimePlayed        int `json:"timePlayed"`
					Eliminations      int `json:"eliminations"`
					UltsEarned        int `json:"ultsEarned"`
					UltsUsed          int `json:"ultsUsed"`
					Deaths            int `json:"deaths"`
					HealingDone       int `json:"healingDone"`
					FinalBlows        int `json:"finalBlows"`
					DeathBlossomKills int `json:"deathBlossomKills"`
					SoloKills         int `json:"soloKills"`
					HeroDamageDone    int `json:"heroDamageDone"`
				} `json:"reaper"`
				Tracer struct {
					CriticalHits       int `json:"criticalHits"`
					DamageTaken        int `json:"damageTaken"`
					ShotsHit           int `json:"shotsHit"`
					TimePlayed         int `json:"timePlayed"`
					Eliminations       int `json:"eliminations"`
					UltsEarned         int `json:"ultsEarned"`
					UltsUsed           int `json:"ultsUsed"`
					PulseBombsAttached int `json:"pulseBombsAttached"`
					Deaths             int `json:"deaths"`
					FinalBlows         int `json:"finalBlows"`
					PulseBombKills     int `json:"pulseBombKills"`
					SoloKills          int `json:"soloKills"`
					HeroDamageDone     int `json:"heroDamageDone"`
				} `json:"tracer"`
			} `json:"heroes"`
		} `json:"15643"`
		Num15648 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits           int `json:"criticalHits"`
				DamageTaken            int `json:"damageTaken"`
				HelixRocketKills       int `json:"helixRocketKills"`
				ShotsHit               int `json:"shotsHit"`
				TimePlayed             int `json:"timePlayed"`
				Eliminations           int `json:"eliminations"`
				UltsEarned             int `json:"ultsEarned"`
				UltsUsed               int `json:"ultsUsed"`
				Deaths                 int `json:"deaths"`
				HealingDone            int `json:"healingDone"`
				FinalBlows             int `json:"finalBlows"`
				HeroDamageDone         int `json:"heroDamageDone"`
				ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
				ScopedHits             int `json:"scopedHits"`
				ScopedCriticalHits     int `json:"scopedCriticalHits"`
				SoloKills              int `json:"soloKills"`
			} `json:"stats"`
			Heroes struct {
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
			} `json:"heroes"`
		} `json:"15648"`
		Num15787 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
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
			} `json:"stats"`
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
			} `json:"heroes"`
		} `json:"15787"`
		Num15937 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
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
				TimeSpentOnFire   int `json:"timeSpentOnFire"`
			} `json:"stats"`
			Heroes struct {
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
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
				} `json:"echo"`
			} `json:"heroes"`
		} `json:"15937"`
		Num16205 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits    int `json:"criticalHits"`
				DamageTaken     int `json:"damageTaken"`
				ShotsHit        int `json:"shotsHit"`
				TimePlayed      int `json:"timePlayed"`
				Eliminations    int `json:"eliminations"`
				UltsEarned      int `json:"ultsEarned"`
				KnockbackKills  int `json:"knockbackKills"`
				UltsUsed        int `json:"ultsUsed"`
				Deaths          int `json:"deaths"`
				TimeSpentOnFire int `json:"timeSpentOnFire"`
				HealingDone     int `json:"healingDone"`
				FinalBlows      int `json:"finalBlows"`
				HeroDamageDone  int `json:"heroDamageDone"`
			} `json:"stats"`
			Heroes struct {
				Lucio struct {
					CriticalHits    int `json:"criticalHits"`
					DamageTaken     int `json:"damageTaken"`
					ShotsHit        int `json:"shotsHit"`
					TimePlayed      int `json:"timePlayed"`
					Eliminations    int `json:"eliminations"`
					UltsEarned      int `json:"ultsEarned"`
					KnockbackKills  int `json:"knockbackKills"`
					UltsUsed        int `json:"ultsUsed"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					HealingDone     int `json:"healingDone"`
					FinalBlows      int `json:"finalBlows"`
					HeroDamageDone  int `json:"heroDamageDone"`
				} `json:"lucio"`
			} `json:"heroes"`
		} `json:"16205"`
		Num16232 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits    int `json:"criticalHits"`
				DamageTaken     int `json:"damageTaken"`
				ShotsHit        int `json:"shotsHit"`
				TimePlayed      int `json:"timePlayed"`
				Eliminations    int `json:"eliminations"`
				UltsEarned      int `json:"ultsEarned"`
				KnockbackKills  int `json:"knockbackKills"`
				UltsUsed        int `json:"ultsUsed"`
				Deaths          int `json:"deaths"`
				TimeSpentOnFire int `json:"timeSpentOnFire"`
				HealingDone     int `json:"healingDone"`
				HeroDamageDone  int `json:"heroDamageDone"`
				FinalBlows      int `json:"finalBlows"`
			} `json:"stats"`
			Heroes struct {
				Lucio struct {
					CriticalHits    int `json:"criticalHits"`
					DamageTaken     int `json:"damageTaken"`
					ShotsHit        int `json:"shotsHit"`
					TimePlayed      int `json:"timePlayed"`
					Eliminations    int `json:"eliminations"`
					UltsEarned      int `json:"ultsEarned"`
					KnockbackKills  int `json:"knockbackKills"`
					UltsUsed        int `json:"ultsUsed"`
					Deaths          int `json:"deaths"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					HealingDone     int `json:"healingDone"`
					HeroDamageDone  int `json:"heroDamageDone"`
					FinalBlows      int `json:"finalBlows"`
				} `json:"lucio"`
			} `json:"heroes"`
		} `json:"16232"`
		Num16240 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits               int `json:"criticalHits"`
				DamageTaken                int `json:"damageTaken"`
				ShotsHit                   int `json:"shotsHit"`
				TimePlayed                 int `json:"timePlayed"`
				Eliminations               int `json:"eliminations"`
				UltsEarned                 int `json:"ultsEarned"`
				KnockbackKills             int `json:"knockbackKills"`
				UltsUsed                   int `json:"ultsUsed"`
				Deaths                     int `json:"deaths"`
				FinalBlows                 int `json:"finalBlows"`
				HeroDamageDone             int `json:"heroDamageDone"`
				JumpPackKills              int `json:"jumpPackKills"`
				EnergyMax                  int `json:"energyMax"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				SoloKills                  int `json:"soloKills"`
				GravitonSurgeKills         int `json:"gravitonSurgeKills"`
				AbilityDamageDone          int `json:"abilityDamageDone"`
				PrimalRageKills            int `json:"primalRageKills"`
			} `json:"stats"`
			Heroes struct {
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
				Winston struct {
					DamageTaken     int `json:"damageTaken"`
					JumpPackKills   int `json:"jumpPackKills"`
					ShotsHit        int `json:"shotsHit"`
					TimePlayed      int `json:"timePlayed"`
					Eliminations    int `json:"eliminations"`
					UltsEarned      int `json:"ultsEarned"`
					UltsUsed        int `json:"ultsUsed"`
					Deaths          int `json:"deaths"`
					FinalBlows      int `json:"finalBlows"`
					HeroDamageDone  int `json:"heroDamageDone"`
					KnockbackKills  int `json:"knockbackKills"`
					TimeSpentOnFire int `json:"timeSpentOnFire"`
					SoloKills       int `json:"soloKills"`
					PrimalRageKills int `json:"primalRageKills"`
				} `json:"winston"`
				Zarya struct {
					HeroDamageDone             int `json:"heroDamageDone"`
					DamageTaken                int `json:"damageTaken"`
					Eliminations               int `json:"eliminations"`
					Deaths                     int `json:"deaths"`
					TimePlayed                 int `json:"timePlayed"`
					EnergyMax                  int `json:"energyMax"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					ShotsHit                   int `json:"shotsHit"`
					UltsEarned                 int `json:"ultsEarned"`
					UltsUsed                   int `json:"ultsUsed"`
					TimeSpentOnFire            int `json:"timeSpentOnFire"`
					GravitonSurgeKills         int `json:"gravitonSurgeKills"`
					FinalBlows                 int `json:"finalBlows"`
					SoloKills                  int `json:"soloKills"`
				} `json:"zarya"`
				Doomfist struct {
					CriticalHits      int `json:"criticalHits"`
					DamageTaken       int `json:"damageTaken"`
					ShotsHit          int `json:"shotsHit"`
					TimePlayed        int `json:"timePlayed"`
					Eliminations      int `json:"eliminations"`
					KnockbackKills    int `json:"knockbackKills"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					FinalBlows        int `json:"finalBlows"`
					AbilityDamageDone int `json:"abilityDamageDone"`
					HeroDamageDone    int `json:"heroDamageDone"`
				} `json:"doomfist"`
			} `json:"heroes"`
		} `json:"16240"`
		Num16241 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				DamageTaken       int `json:"damageTaken"`
				PrimalRageKills   int `json:"primalRageKills"`
				JumpPackKills     int `json:"jumpPackKills"`
				ShotsHit          int `json:"shotsHit"`
				TimePlayed        int `json:"timePlayed"`
				Eliminations      int `json:"eliminations"`
				UltsEarned        int `json:"ultsEarned"`
				KnockbackKills    int `json:"knockbackKills"`
				UltsUsed          int `json:"ultsUsed"`
				Deaths            int `json:"deaths"`
				TimeSpentOnFire   int `json:"timeSpentOnFire"`
				FinalBlows        int `json:"finalBlows"`
				HeroDamageDone    int `json:"heroDamageDone"`
				AbilityDamageDone int `json:"abilityDamageDone"`
				SoloKills         int `json:"soloKills"`
				CriticalHits      int `json:"criticalHits"`
			} `json:"stats"`
			Heroes struct {
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
					SoloKills       int `json:"soloKills"`
				} `json:"winston"`
				Doomfist struct {
					HeroDamageDone    int `json:"heroDamageDone"`
					DamageTaken       int `json:"damageTaken"`
					TimePlayed        int `json:"timePlayed"`
					ShotsHit          int `json:"shotsHit"`
					AbilityDamageDone int `json:"abilityDamageDone"`
					Deaths            int `json:"deaths"`
					CriticalHits      int `json:"criticalHits"`
				} `json:"doomfist"`
			} `json:"heroes"`
		} `json:"16241"`
		Num16288 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				HeroDamageDone             int `json:"heroDamageDone"`
				DamageTaken                int `json:"damageTaken"`
				Deaths                     int `json:"deaths"`
				TimePlayed                 int `json:"timePlayed"`
				CriticalHits               int `json:"criticalHits"`
				ShotsHit                   int `json:"shotsHit"`
				EnergyMax                  int `json:"energyMax"`
				Eliminations               int `json:"eliminations"`
				LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
				UltsEarned                 int `json:"ultsEarned"`
				UltsUsed                   int `json:"ultsUsed"`
				GraviticFluxKills          int `json:"graviticFluxKills"`
				AccretionKills             int `json:"accretionKills"`
				TimeSpentOnFire            int `json:"timeSpentOnFire"`
				FinalBlows                 int `json:"finalBlows"`
			} `json:"stats"`
			Heroes struct {
				DVa struct {
					HeroDamageDone int `json:"heroDamageDone"`
					DamageTaken    int `json:"damageTaken"`
					Deaths         int `json:"deaths"`
					TimePlayed     int `json:"timePlayed"`
					CriticalHits   int `json:"criticalHits"`
					ShotsHit       int `json:"shotsHit"`
				} `json:"d-va"`
				Zarya struct {
					DamageTaken                int `json:"damageTaken"`
					EnergyMax                  int `json:"energyMax"`
					ShotsHit                   int `json:"shotsHit"`
					TimePlayed                 int `json:"timePlayed"`
					Eliminations               int `json:"eliminations"`
					LifetimeEnergyAccumulation int `json:"lifetimeEnergyAccumulation"`
					UltsEarned                 int `json:"ultsEarned"`
					UltsUsed                   int `json:"ultsUsed"`
					Deaths                     int `json:"deaths"`
					HeroDamageDone             int `json:"heroDamageDone"`
				} `json:"zarya"`
				Sigma struct {
					DamageTaken       int `json:"damageTaken"`
					GraviticFluxKills int `json:"graviticFluxKills"`
					TimePlayed        int `json:"timePlayed"`
					Eliminations      int `json:"eliminations"`
					UltsEarned        int `json:"ultsEarned"`
					UltsUsed          int `json:"ultsUsed"`
					AccretionKills    int `json:"accretionKills"`
					Deaths            int `json:"deaths"`
					TimeSpentOnFire   int `json:"timeSpentOnFire"`
					FinalBlows        int `json:"finalBlows"`
					HeroDamageDone    int `json:"heroDamageDone"`
				} `json:"sigma"`
			} `json:"heroes"`
		} `json:"16288"`
		Num16530 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				TimePlayed       int `json:"timePlayed"`
				CriticalHits     int `json:"criticalHits"`
				DamageTaken      int `json:"damageTaken"`
				ShotsHit         int `json:"shotsHit"`
				Eliminations     int `json:"eliminations"`
				UltsEarned       int `json:"ultsEarned"`
				DamageDiscordOrb int `json:"damageDiscordOrb"`
				UltsUsed         int `json:"ultsUsed"`
				Deaths           int `json:"deaths"`
				HealingDone      int `json:"healingDone"`
				FinalBlows       int `json:"finalBlows"`
				HeroDamageDone   int `json:"heroDamageDone"`
			} `json:"stats"`
			Heroes struct {
				Lucio struct {
					TimePlayed int `json:"timePlayed"`
				} `json:"lucio"`
				Zenyatta struct {
					CriticalHits     int `json:"criticalHits"`
					DamageTaken      int `json:"damageTaken"`
					ShotsHit         int `json:"shotsHit"`
					TimePlayed       int `json:"timePlayed"`
					Eliminations     int `json:"eliminations"`
					UltsEarned       int `json:"ultsEarned"`
					DamageDiscordOrb int `json:"damageDiscordOrb"`
					UltsUsed         int `json:"ultsUsed"`
					Deaths           int `json:"deaths"`
					HealingDone      int `json:"healingDone"`
					FinalBlows       int `json:"finalBlows"`
					HeroDamageDone   int `json:"heroDamageDone"`
				} `json:"zenyatta"`
			} `json:"heroes"`
		} `json:"16530"`
		Num16554 struct {
			ID     int `json:"id"`
			TeamID int `json:"teamId"`
			Stats  struct {
				CriticalHits           int `json:"criticalHits"`
				DamageTaken            int `json:"damageTaken"`
				ShotsHit               int `json:"shotsHit"`
				ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
				TimePlayed             int `json:"timePlayed"`
				Eliminations           int `json:"eliminations"`
				Deaths                 int `json:"deaths"`
				ScopedHits             int `json:"scopedHits"`
				FinalBlows             int `json:"finalBlows"`
				ScopedCriticalHits     int `json:"scopedCriticalHits"`
				HeroDamageDone         int `json:"heroDamageDone"`
				UltsEarned             int `json:"ultsEarned"`
				UltsUsed               int `json:"ultsUsed"`
				SoloKills              int `json:"soloKills"`
				DragonbladeKills       int `json:"dragonbladeKills"`
				TimeSpentOnFire        int `json:"timeSpentOnFire"`
				PulseBombsAttached     int `json:"pulseBombsAttached"`
				PulseBombKills         int `json:"pulseBombKills"`
			} `json:"stats"`
			Heroes struct {
				Widowmaker struct {
					CriticalHits           int `json:"criticalHits"`
					DamageTaken            int `json:"damageTaken"`
					ShotsHit               int `json:"shotsHit"`
					ScopedCriticalHitKills int `json:"scopedCriticalHitKills"`
					TimePlayed             int `json:"timePlayed"`
					Eliminations           int `json:"eliminations"`
					Deaths                 int `json:"deaths"`
					ScopedHits             int `json:"scopedHits"`
					FinalBlows             int `json:"finalBlows"`
					ScopedCriticalHits     int `json:"scopedCriticalHits"`
					HeroDamageDone         int `json:"heroDamageDone"`
				} `json:"widowmaker"`
				Tracer struct {
					CriticalHits       int `json:"criticalHits"`
					DamageTaken        int `json:"damageTaken"`
					ShotsHit           int `json:"shotsHit"`
					TimePlayed         int `json:"timePlayed"`
					Eliminations       int `json:"eliminations"`
					UltsEarned         int `json:"ultsEarned"`
					UltsUsed           int `json:"ultsUsed"`
					Deaths             int `json:"deaths"`
					FinalBlows         int `json:"finalBlows"`
					SoloKills          int `json:"soloKills"`
					HeroDamageDone     int `json:"heroDamageDone"`
					TimeSpentOnFire    int `json:"timeSpentOnFire"`
					PulseBombsAttached int `json:"pulseBombsAttached"`
					PulseBombKills     int `json:"pulseBombKills"`
				} `json:"tracer"`
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
					SoloKills        int `json:"soloKills"`
				} `json:"genji"`
			} `json:"heroes"`
		} `json:"16554"`
	} `json:"players"`
	Winner int `json:"winner"`
}
