package blizzard

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/FuzzyStatic/blizzard/d3c"
)

const (
	d3Path                          = "/d3"
	d3DataPath                      = d3Path + "/data"
	actPath                         = d3DataPath + "/act"
	artisanPath                     = d3DataPath + "/artisan"
	blacksmithPath                  = artisanPath + "/blacksmith"
	jewelerPath                     = artisanPath + "/jeweler"
	mysticPath                      = artisanPath + "/mystic"
	blacksmithRecipePath            = blacksmithPath + "/recipe"
	jewelerRecipePath               = jewelerPath + "/recipe"
	followerPath                    = d3DataPath + "/follower"
	enchantressPath                 = followerPath + "/enchantress"
	scoundrelPath                   = followerPath + "/scoundrel"
	templarPath                     = followerPath + "/templar"
	heroPath                        = d3DataPath + "/hero"
	barbarianPath                   = heroPath + "/barbarian"
	crusaderPath                    = heroPath + "/crusader"
	demonHunterPath                 = heroPath + "/demon-hunter"
	monkPath                        = heroPath + "/monk"
	necromancerPath                 = heroPath + "/necromancer"
	wizardPath                      = heroPath + "/wizard"
	witchDoctorPath                 = heroPath + "/witch-doctor"
	skillPath                       = "/skill"
	barbarianSkillPath              = barbarianPath + skillPath
	crusaderSkillPath               = crusaderPath + skillPath
	demonHunterSkillPath            = demonHunterPath + skillPath
	monkSkillPath                   = monkPath + skillPath
	necromancerSkillPath            = necromancerPath + skillPath
	wizardSkillPath                 = wizardPath + skillPath
	witchDoctorSkillPath            = witchDoctorPath + skillPath
	itemTypePath                    = d3DataPath + "/item-type"
	itemPath                        = d3DataPath + "/item"
	d3ProfilePath                   = d3Path + profilePath
	profileHeroPartialPath          = "/hero"
	profileItemsPartialPath         = "/items"
	profileFollowerItemsPartialPath = "/follower-items"
)

// D3ActIndex returns an index of acts
func (c *Config) D3ActIndex() (*d3c.ActIndex, error) {
	var (
		dat d3c.ActIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+actPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Act returns information about act based on ID
func (c *Config) D3Act(id int) (*d3c.Act, error) {
	var (
		dat d3c.Act
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+actPath+"/"+strconv.Itoa(id)+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Blacksmith returns blacksmith data
func (c *Config) D3Blacksmith() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+blacksmithPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Jeweler returns jeweler data
func (c *Config) D3Jeweler() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+jewelerPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Mystic returns mystic data
func (c *Config) D3Mystic() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+mysticPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3BlacksmithRecipe returns blacksmith recipe data
func (c *Config) D3BlacksmithRecipe(recipeSlug string) (*d3c.Recipe, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+blacksmithRecipePath+"/"+recipeSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3JewelerRecipe returns jeweler recipe data
func (c *Config) D3JewelerRecipe(recipeSlug string) (*d3c.Recipe, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+jewelerRecipePath+"/"+recipeSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Enchantress returns enchantress data
func (c *Config) D3Enchantress() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+enchantressPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Scoundrel returns scoundrel data
func (c *Config) D3Scoundrel() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+scoundrelPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Templar returns templar data
func (c *Config) D3Templar() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+templarPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Barbarian returns barbarian data
func (c *Config) D3Barbarian() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+barbarianPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Crusader returns crusader data
func (c *Config) D3Crusader() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+crusaderPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3DemonHunter returns demon hunter data
func (c *Config) D3DemonHunter() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+demonHunterPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Monk returns monk data
func (c *Config) D3Monk() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+monkPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Necromancer returns necromancer data
func (c *Config) D3Necromancer() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+necromancerPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Wizard returns wizard data
func (c *Config) D3Wizard() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wizardPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3WitchDoctor returns witch doctor data
func (c *Config) D3WitchDoctor() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+witchDoctorPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3BarbarianSkill returns barbarian skill data
func (c *Config) D3BarbarianSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+barbarianSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3CrusaderSkill returns crusader skill data
func (c *Config) D3CrusaderSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+crusaderSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3DemonHunterSkill returns demon hunter skill data
func (c *Config) D3DemonHunterSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+demonHunterSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3MonkSkill returns monk skill data
func (c *Config) D3MonkSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+monkSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3NecromancerSkill returns necromancer skill data
func (c *Config) D3NecromancerSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+necromancerSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3WizardSkill returns wizard skill data
func (c *Config) D3WizardSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wizardSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3WitchDoctorSkill returns witch doctor skill data
func (c *Config) D3WitchDoctorSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+witchDoctorSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3ItemTypeIndex returns an index of item types
func (c *Config) D3ItemTypeIndex() (*d3c.ItemTypeIndex, error) {
	var (
		dat d3c.ItemTypeIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemTypePath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3ItemType returns item type data
func (c *Config) D3ItemType(itemTypeSlug string) (*d3c.ItemType, error) {
	var (
		dat d3c.ItemType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemTypePath+"/"+itemTypeSlug+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Item returns item data
func (c *Config) D3Item(itemSlug, itemID string) (*d3c.Item, error) {
	var (
		dat d3c.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemPath+"/"+itemSlug+"-"+itemID+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3Profile returns profile data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3Profile(account string) (*d3c.Profile, error) {
	var (
		dat d3c.Profile
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+"/?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3ProfileHero returns profile's hero data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3ProfileHero(account string, heroID int) (*d3c.ProfileHero, error) {
	var (
		dat d3c.ProfileHero
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3ProfileHeroItems returns profile's hero's item data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3ProfileHeroItems(account string, heroID int) (*d3c.ProfileHeroItems, error) {
	var (
		dat d3c.ProfileHeroItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+profileItemsPartialPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3ProfileHeroFollowerItems returns profile's hero's follower item data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3ProfileHeroFollowerItems(account string, heroID int) (*d3c.ProfileHeroFollowerItems, error) {
	var (
		dat d3c.ProfileHeroFollowerItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+profileFollowerItemsPartialPath+"?"+localeQuery+c.locale, "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
