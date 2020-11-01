package d3c

// Profile structure
type Profile struct {
	BattleTag                  string `json:"battleTag"`
	ParagonLevel               int    `json:"paragonLevel"`
	ParagonLevelHardcore       int    `json:"paragonLevelHardcore"`
	ParagonLevelSeason         int    `json:"paragonLevelSeason"`
	ParagonLevelSeasonHardcore int    `json:"paragonLevelSeasonHardcore"`
	GuildName                  string `json:"guildName"`
	Heroes                     []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Class     string `json:"class"`
		ClassSlug string `json:"classSlug"`
		Gender    int    `json:"gender"`
		Level     int    `json:"level"`
		Kills     struct {
			Elites int `json:"elites"`
		} `json:"kills"`
		ParagonLevel int  `json:"paragonLevel"`
		Hardcore     bool `json:"hardcore"`
		Seasonal     bool `json:"seasonal"`
		Dead         bool `json:"dead"`
		LastUpdated  int  `json:"last-updated"`
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
		DemonHunter float64 `json:"demon-hunter"`
		Barbarian   float64 `json:"barbarian"`
		WitchDoctor float64 `json:"witch-doctor"`
		Wizard      float64 `json:"wizard"`
		Monk        float64 `json:"monk"`
		Crusader    float64 `json:"crusader"`
	} `json:"timePlayed"`
	Progression struct {
		Act1 bool `json:"act1"`
		Act3 bool `json:"act3"`
		Act2 bool `json:"act2"`
		Act5 bool `json:"act5"`
		Act4 bool `json:"act4"`
	} `json:"progression"`
	SeasonalProfiles struct {
		Season13 struct {
			SeasonID             int `json:"seasonId"`
			ParagonLevel         int `json:"paragonLevel"`
			ParagonLevelHardcore int `json:"paragonLevelHardcore"`
			Kills                struct {
				Monsters         int `json:"monsters"`
				Elites           int `json:"elites"`
				HardcoreMonsters int `json:"hardcoreMonsters"`
			} `json:"kills"`
			TimePlayed struct {
				DemonHunter float64 `json:"demon-hunter"`
				Barbarian   float64 `json:"barbarian"`
				WitchDoctor float64 `json:"witch-doctor"`
				Necromancer float64 `json:"necromancer"`
				Wizard      float64 `json:"wizard"`
				Monk        float64 `json:"monk"`
				Crusader    float64 `json:"crusader"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
		} `json:"season13"`
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
				DemonHunter float64 `json:"demon-hunter"`
				Barbarian   float64 `json:"barbarian"`
				WitchDoctor float64 `json:"witch-doctor"`
				Necromancer float64 `json:"necromancer"`
				Wizard      float64 `json:"wizard"`
				Monk        float64 `json:"monk"`
				Crusader    float64 `json:"crusader"`
			} `json:"timePlayed"`
			HighestHardcoreLevel int `json:"highestHardcoreLevel"`
		} `json:"season0"`
	} `json:"seasonalProfiles"`
	JewelerSeasonHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"jewelerSeasonHardcore"`
	MysticSeasonHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"mysticSeasonHardcore"`
	BlacksmithSeason struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"blacksmithSeason"`
	JewelerSeason struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"jewelerSeason"`
	BlacksmithHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"blacksmithHardcore"`
	JewelerHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"jewelerHardcore"`
	MysticHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"mysticHardcore"`
	BlacksmithSeasonHardcore struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"blacksmithSeasonHardcore"`
	MysticSeason struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"mysticSeason"`
	Blacksmith struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"blacksmith"`
	Jeweler struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"jeweler"`
	Mystic struct {
		Slug  string `json:"slug"`
		Level int    `json:"level"`
	} `json:"mystic"`
}

// ProfileHero structure
type ProfileHero struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Gender       int    `json:"gender"`
	Level        int    `json:"level"`
	ParagonLevel int    `json:"paragonLevel"`
	Kills        struct {
		Elites int `json:"elites"`
	} `json:"kills"`
	Hardcore      bool `json:"hardcore"`
	Seasonal      bool `json:"seasonal"`
	SeasonCreated int  `json:"seasonCreated"`
	Skills        struct {
		Active []struct {
			Skill struct {
				Slug            string `json:"slug"`
				Name            string `json:"name"`
				Icon            string `json:"icon"`
				Level           int    `json:"level"`
				TooltipURL      string `json:"tooltipUrl"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
			} `json:"skill"`
			Rune struct {
				Slug            string `json:"slug"`
				Type            string `json:"type"`
				Name            string `json:"name"`
				Level           int    `json:"level"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
			} `json:"rune"`
		} `json:"active"`
		Passive []struct {
			Skill struct {
				Slug            string `json:"slug"`
				Name            string `json:"name"`
				Icon            string `json:"icon"`
				Level           int    `json:"level"`
				TooltipURL      string `json:"tooltipUrl"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
				FlavorText      string `json:"flavorText"`
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
			TransmogItem  struct {
				ID            string `json:"id"`
				Name          string `json:"name"`
				Icon          string `json:"icon"`
				DisplayColor  string `json:"displayColor"`
				TooltipParams string `json:"tooltipParams"`
			} `json:"transmogItem"`
		} `json:"head"`
		Neck struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"neck"`
		Torso struct {
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
		} `json:"torso"`
		Shoulders struct {
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
		} `json:"shoulders"`
		Legs struct {
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
		} `json:"legs"`
		Waist struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"waist"`
		Hands struct {
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
		} `json:"hands"`
		Bracers struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"bracers"`
		Feet struct {
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
		} `json:"feet"`
		LeftFinger struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"leftFinger"`
		RightFinger struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"rightFinger"`
		MainHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
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
	} `json:"items"`
	Followers struct {
		Templar struct {
			Slug  string `json:"slug"`
			Level int    `json:"level"`
			Items struct {
				Neck struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"neck"`
				LeftFinger struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"leftFinger"`
				RightFinger struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"rightFinger"`
				MainHand struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"mainHand"`
				OffHand struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"offHand"`
			} `json:"items"`
			Stats struct {
				GoldFind        float64 `json:"goldFind"`
				MagicFind       float64 `json:"magicFind"`
				ExperienceBonus float64 `json:"experienceBonus"`
			} `json:"stats"`
			Skills []struct {
				Slug            string `json:"slug"`
				Name            string `json:"name"`
				Icon            string `json:"icon"`
				Level           int    `json:"level"`
				TooltipURL      string `json:"tooltipUrl"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
			} `json:"skills"`
		} `json:"templar"`
		Scoundrel struct {
			Slug  string `json:"slug"`
			Level int    `json:"level"`
			Items struct {
				MainHand struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"mainHand"`
			} `json:"items"`
			Stats struct {
				GoldFind        float64 `json:"goldFind"`
				MagicFind       float64 `json:"magicFind"`
				ExperienceBonus float64 `json:"experienceBonus"`
			} `json:"stats"`
			Skills []struct {
				Slug            string `json:"slug"`
				Name            string `json:"name"`
				Icon            string `json:"icon"`
				Level           int    `json:"level"`
				TooltipURL      string `json:"tooltipUrl"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
			} `json:"skills"`
		} `json:"scoundrel"`
		Enchantress struct {
			Slug  string `json:"slug"`
			Level int    `json:"level"`
			Items struct {
				MainHand struct {
					ID            string `json:"id"`
					Name          string `json:"name"`
					Icon          string `json:"icon"`
					DisplayColor  string `json:"displayColor"`
					TooltipParams string `json:"tooltipParams"`
				} `json:"mainHand"`
			} `json:"items"`
			Stats struct {
				GoldFind        float64 `json:"goldFind"`
				MagicFind       float64 `json:"magicFind"`
				ExperienceBonus float64 `json:"experienceBonus"`
			} `json:"stats"`
			Skills []struct {
				Slug            string `json:"slug"`
				Name            string `json:"name"`
				Icon            string `json:"icon"`
				Level           int    `json:"level"`
				TooltipURL      string `json:"tooltipUrl"`
				Description     string `json:"description"`
				DescriptionHTML string `json:"descriptionHtml"`
			} `json:"skills"`
		} `json:"enchantress"`
	} `json:"followers"`
	LegendaryPowers []struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
	} `json:"legendaryPowers"`
	Progression struct {
		Act1 struct {
			Completed       bool `json:"completed"`
			CompletedQuests []struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"completedQuests"`
		} `json:"act1"`
		Act2 struct {
			Completed       bool `json:"completed"`
			CompletedQuests []struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"completedQuests"`
		} `json:"act2"`
		Act3 struct {
			Completed       bool `json:"completed"`
			CompletedQuests []struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"completedQuests"`
		} `json:"act3"`
		Act4 struct {
			Completed       bool `json:"completed"`
			CompletedQuests []struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"completedQuests"`
		} `json:"act4"`
		Act5 struct {
			Completed       bool `json:"completed"`
			CompletedQuests []struct {
				Slug string `json:"slug"`
				Name string `json:"name"`
			} `json:"completedQuests"`
		} `json:"act5"`
	} `json:"progression"`
	Alive                    bool `json:"alive"`
	LastUpdated              int  `json:"lastUpdated"`
	HighestSoloRiftCompleted int  `json:"highestSoloRiftCompleted"`
	Stats                    struct {
		Life              float64 `json:"life"`
		Damage            float64 `json:"damage"`
		Toughness         float64 `json:"toughness"`
		Healing           float64 `json:"healing"`
		AttackSpeed       float64 `json:"attackSpeed"`
		Armor             float64 `json:"armor"`
		Strength          float64 `json:"strength"`
		Dexterity         float64 `json:"dexterity"`
		Vitality          float64 `json:"vitality"`
		Intelligence      float64 `json:"intelligence"`
		PhysicalResist    float64 `json:"physicalResist"`
		FireResist        float64 `json:"fireResist"`
		ColdResist        float64 `json:"coldResist"`
		LightningResist   float64 `json:"lightningResist"`
		PoisonResist      float64 `json:"poisonResist"`
		ArcaneResist      float64 `json:"arcaneResist"`
		BlockChance       float64 `json:"blockChance"`
		BlockAmountMin    float64 `json:"blockAmountMin"`
		BlockAmountMax    float64 `json:"blockAmountMax"`
		GoldFind          float64 `json:"goldFind"`
		CritChance        float64 `json:"critChance"`
		Thorns            float64 `json:"thorns"`
		LifeSteal         float64 `json:"lifeSteal"`
		LifePerKill       float64 `json:"lifePerKill"`
		LifeOnHit         float64 `json:"lifeOnHit"`
		PrimaryResource   float64 `json:"primaryResource"`
		SecondaryResource float64 `json:"secondaryResource"`
	} `json:"stats"`
}

// ProfileHeroItems structure
type ProfileHeroItems struct {
	Head struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"head"`
	Neck struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			JewelRank                int      `json:"jewelRank"`
			JewelSecondaryUnlockRank int      `json:"jewelSecondaryUnlockRank"`
			Attributes               []string `json:"attributes"`
			IsGem                    bool     `json:"isGem"`
			IsJewel                  bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"neck"`
	Torso struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop int `json:"seasonRequiredToDrop"`
		Transmog             struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
		} `json:"transmog"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"torso"`
	Shoulders struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"shoulders"`
	Legs struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		Set struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"legs"`
	Waist struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets            int  `json:"openSockets"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"waist"`
	Hands struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"hands"`
	Bracers struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		ElementalType    string  `json:"elementalType"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets            int  `json:"openSockets"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"bracers"`
	Feet struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"feet"`
	LeftFinger struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			JewelRank                int      `json:"jewelRank"`
			JewelSecondaryUnlockRank int      `json:"jewelSecondaryUnlockRank"`
			Attributes               []string `json:"attributes"`
			IsGem                    bool     `json:"isGem"`
			IsJewel                  bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"leftFinger"`
	RightFinger struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			JewelRank                int      `json:"jewelRank"`
			JewelSecondaryUnlockRank int      `json:"jewelSecondaryUnlockRank"`
			Attributes               []string `json:"attributes"`
			IsGem                    bool     `json:"isGem"`
			IsJewel                  bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"rightFinger"`
	MainHand struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		Damage           string  `json:"damage"`
		Dps              string  `json:"dps"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		ElementalType    string  `json:"elementalType"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Gems        []struct {
			Item struct {
				ID   string `json:"id"`
				Slug string `json:"slug"`
				Name string `json:"name"`
				Icon string `json:"icon"`
				Path string `json:"path"`
			} `json:"item"`
			Attributes []string `json:"attributes"`
			IsGem      bool     `json:"isGem"`
			IsJewel    bool     `json:"isJewel"`
		} `json:"gems"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"mainHand"`
	OffHand struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		Icon          string `json:"icon"`
		DisplayColor  string `json:"displayColor"`
		TooltipParams string `json:"tooltipParams"`
		RequiredLevel int    `json:"requiredLevel"`
		ItemLevel     int    `json:"itemLevel"`
		StackSizeMax  int    `json:"stackSizeMax"`
		AccountBound  bool   `json:"accountBound"`
		FlavorText    string `json:"flavorText"`
		TypeName      string `json:"typeName"`
		Type          struct {
			TwoHanded bool   `json:"twoHanded"`
			ID        string `json:"id"`
		} `json:"type"`
		Armor            float64 `json:"armor"`
		AttacksPerSecond float64 `json:"attacksPerSecond"`
		MinDamage        float64 `json:"minDamage"`
		MaxDamage        float64 `json:"maxDamage"`
		Slots            string  `json:"slots"`
		Augmentation     string  `json:"augmentation"`
		Attributes       struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributes"`
		AttributesHTML struct {
			Primary   []string `json:"primary"`
			Secondary []string `json:"secondary"`
		} `json:"attributesHtml"`
		OpenSockets int `json:"openSockets"`
		Set         struct {
			Name            string `json:"name"`
			Slug            string `json:"slug"`
			Description     string `json:"description"`
			DescriptionHTML string `json:"descriptionHtml"`
		} `json:"set"`
		SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
		IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
	} `json:"offHand"`
}

// ProfileHeroFollowerItems structure
type ProfileHeroFollowerItems struct {
	Templar struct {
		MainHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor                  float64 `json:"armor"`
			Damage                 string  `json:"damage"`
			Dps                    string  `json:"dps"`
			AttacksPerSecond       float64 `json:"attacksPerSecond"`
			MinDamage              float64 `json:"minDamage"`
			MaxDamage              float64 `json:"maxDamage"`
			Slots                  string  `json:"slots"`
			OpenSockets            int     `json:"openSockets"`
			SeasonRequiredToDrop   int     `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool    `json:"isSeasonRequiredToDrop"`
		} `json:"mainHand"`
		OffHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor                  float64 `json:"armor"`
			AttacksPerSecond       float64 `json:"attacksPerSecond"`
			MinDamage              float64 `json:"minDamage"`
			MaxDamage              float64 `json:"maxDamage"`
			Slots                  string  `json:"slots"`
			OpenSockets            int     `json:"openSockets"`
			SeasonRequiredToDrop   int     `json:"seasonRequiredToDrop"`
			BlockChance            string  `json:"blockChance"`
			IsSeasonRequiredToDrop bool    `json:"isSeasonRequiredToDrop"`
		} `json:"offHand"`
	} `json:"templar"`
	Scoundrel struct {
		MainHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			Damage           string  `json:"damage"`
			Dps              string  `json:"dps"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary []string `json:"primary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary []string `json:"primary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"mainHand"`
	} `json:"scoundrel"`
	Enchantress struct {
		Neck struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			FlavorText    string `json:"flavorText"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"neck"`
		LeftFinger struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			FlavorText    string `json:"flavorText"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"leftFinger"`
		RightFinger struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			FlavorText    string `json:"flavorText"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"rightFinger"`
		MainHand struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			FlavorText    string `json:"flavorText"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			Damage           string  `json:"damage"`
			Dps              string  `json:"dps"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			ElementalType    string  `json:"elementalType"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"mainHand"`
		Special struct {
			ID            string `json:"id"`
			Name          string `json:"name"`
			Icon          string `json:"icon"`
			DisplayColor  string `json:"displayColor"`
			TooltipParams string `json:"tooltipParams"`
			RequiredLevel int    `json:"requiredLevel"`
			ItemLevel     int    `json:"itemLevel"`
			StackSizeMax  int    `json:"stackSizeMax"`
			AccountBound  bool   `json:"accountBound"`
			FlavorText    string `json:"flavorText"`
			TypeName      string `json:"typeName"`
			Type          struct {
				TwoHanded bool   `json:"twoHanded"`
				ID        string `json:"id"`
			} `json:"type"`
			Armor            float64 `json:"armor"`
			AttacksPerSecond float64 `json:"attacksPerSecond"`
			MinDamage        float64 `json:"minDamage"`
			MaxDamage        float64 `json:"maxDamage"`
			Slots            string  `json:"slots"`
			Attributes       struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributes"`
			AttributesHTML struct {
				Primary   []string `json:"primary"`
				Secondary []string `json:"secondary"`
			} `json:"attributesHtml"`
			OpenSockets            int  `json:"openSockets"`
			SeasonRequiredToDrop   int  `json:"seasonRequiredToDrop"`
			IsSeasonRequiredToDrop bool `json:"isSeasonRequiredToDrop"`
		} `json:"special"`
	} `json:"enchantress"`
}
