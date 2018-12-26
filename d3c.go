package blizzard

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/FuzzyStatic/blizzard/d3c"
)

const (
	d3Path                          = "/d3"
	dataPath                        = d3Path + "/data"
	actPath                         = dataPath + "/act"
	artisanPath                     = dataPath + "/artisan"
	blacksmithPath                  = artisanPath + "/blacksmith"
	jewelerPath                     = artisanPath + "/jeweler"
	mysticPath                      = artisanPath + "/mystic"
	blacksmithRecipePath            = blacksmithPath + "/recipe"
	jewelerRecipePath               = jewelerPath + "/recipe"
	followerPath                    = dataPath + "/follower"
	enchantressPath                 = followerPath + "/enchantress"
	scoundrelPath                   = followerPath + "/scoundrel"
	templarPath                     = followerPath + "/templar"
	heroPath                        = dataPath + "/hero"
	barbarianPath                   = heroPath + "/barbarian"
	crusaderPath                    = heroPath + "/crusader"
	demonHunterPath                 = heroPath + "/demon-hunter"
	necromancerPath                 = heroPath + "/necromancer"
	wizardPath                      = heroPath + "/wizard"
	witchDoctorPath                 = heroPath + "/witch-doctor"
	skillPath                       = "/skill"
	barbarianSkillPath              = barbarianPath + skillPath
	crusaderSkillPath               = crusaderPath + skillPath
	demonHunterSkillPath            = demonHunterPath + skillPath
	necromancerSkillPath            = necromancerPath + skillPath
	wizardSkillPath                 = wizardPath + skillPath
	witchDoctorSkillPath            = witchDoctorPath + skillPath
	itemTypePath                    = dataPath + "/item-type"
	itemPath                        = dataPath + "/item"
	profilePath                     = d3Path + "/profile"
	profileHeroPartialPath          = "/hero"
	profileItemsPartialPath         = "/items"
	profileFollowerItemsPartialPath = "/follower-items"
)

// D3GetActIndex returns an index of acts
func (c *Config) D3GetActIndex() (*d3c.ActIndex, error) {
	var (
		dat d3c.ActIndex
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + actPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetAct returns information about act based on ID
func (c *Config) D3GetAct(id int) (*d3c.Act, error) {
	var (
		dat d3c.Act
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + actPath + "/" + strconv.Itoa(id) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetBlacksmith returns blacksmith data
func (c *Config) D3GetBlacksmith() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + blacksmithPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetJeweler returns jeweler data
func (c *Config) D3GetJeweler() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + jewelerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetMystic returns mystic data
func (c *Config) D3GetMystic() (*d3c.Artisan, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + mysticPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetBlacksmithRecipe returns blacksmith recipe data
func (c *Config) D3GetBlacksmithRecipe(recipeSlug string) (*d3c.Recipe, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + blacksmithRecipePath + "/" + recipeSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetJewelerRecipe returns jeweler recipe data
func (c *Config) D3GetJewelerRecipe(recipeSlug string) (*d3c.Recipe, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + jewelerRecipePath + "/" + recipeSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetEnchantress returns enchantress data
func (c *Config) D3GetEnchantress() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + enchantressPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetScoundrel returns scoundrel data
func (c *Config) D3GetScoundrel() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + scoundrelPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetTemplar returns templar data
func (c *Config) D3GetTemplar() (*d3c.Follower, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + templarPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetBarbarian returns barbarian data
func (c *Config) D3GetBarbarian() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + barbarianPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetCrusader returns crusader data
func (c *Config) D3GetCrusader() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + crusaderPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetDemonHunter returns demon hunter data
func (c *Config) D3GetDemonHunter() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + demonHunterPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetNecromancer returns necromancer data
func (c *Config) D3GetNecromancer() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + necromancerPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetWizard returns wizard data
func (c *Config) D3GetWizard() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + wizardPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetWitchDoctor returns witch doctor data
func (c *Config) D3GetWitchDoctor() (*d3c.Hero, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + witchDoctorPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetBarbarianSkill returns barbarian skill data
func (c *Config) D3GetBarbarianSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + barbarianSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetCrusaderSkill returns crusader skill data
func (c *Config) D3GetCrusaderSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + crusaderSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetDemonHunterSkill returns demon hunter skill data
func (c *Config) D3GetDemonHunterSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + demonHunterSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetNecromancerSkill returns necromancer skill data
func (c *Config) D3GetNecromancerSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + necromancerSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetWizardSkill returns wizard skill data
func (c *Config) D3GetWizardSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + wizardSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetWitchDoctorSkill returns witch doctor skill data
func (c *Config) D3GetWitchDoctorSkill(skillSlug string) (*d3c.Skill, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + witchDoctorSkillPath + "/" + skillSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetItemTypeIndex returns an index of item types
func (c *Config) D3GetItemTypeIndex() (*d3c.ItemTypeIndex, error) {
	var (
		dat d3c.ItemTypeIndex
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + itemTypePath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetItemType returns item type data
func (c *Config) D3GetItemType(itemTypeSlug string) (*d3c.ItemType, error) {
	var (
		dat d3c.ItemType
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + itemTypePath + "/" + itemTypeSlug + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetItem returns item data
func (c *Config) D3GetItem(itemSlug, itemID string) (*d3c.Item, error) {
	var (
		dat d3c.Item
		b   []byte
		err error
	)

	b, err = c.GetURLBody(c.apiURL + itemPath + "/" + itemSlug + "-" + itemID + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetProfile returns profile data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3GetProfile(account string) (*d3c.Profile, error) {
	var (
		dat d3c.Profile
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.GetURLBody(c.apiURL + profilePath + "/" + account + "/?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetProfileHero returns profile's hero data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3GetProfileHero(account string, heroID int) (*d3c.ProfileHero, error) {
	var (
		dat d3c.ProfileHero
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.GetURLBody(c.apiURL + profilePath + "/" + account + profileHeroPartialPath + "/" + strconv.Itoa(heroID) + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetProfileHeroItems returns profile's hero's item data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3GetProfileHeroItems(account string, heroID int) (*d3c.ProfileHeroItems, error) {
	var (
		dat d3c.ProfileHeroItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.GetURLBody(c.apiURL + profilePath + "/" + account + profileHeroPartialPath + "/" + strconv.Itoa(heroID) + profileItemsPartialPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}

// D3GetProfileHeroFollowerItems returns profile's hero's follower item data
// Formats accepted: Player-1234 or Player#1234
func (c *Config) D3GetProfileHeroFollowerItems(account string, heroID int) (*d3c.ProfileHeroFollowerItems, error) {
	var (
		dat d3c.ProfileHeroFollowerItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.GetURLBody(c.apiURL + profilePath + "/" + account + profileHeroPartialPath + "/" + strconv.Itoa(heroID) + profileFollowerItemsPartialPath + "?" + localeQuery + c.locale)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return nil, err
	}

	return &dat, nil
}
