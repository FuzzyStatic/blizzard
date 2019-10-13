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
func (c *Client) D3ActIndex() (*d3c.ActIndex, []byte, error) {
	var (
		dat d3c.ActIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+actPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Act returns information about act based on ID
func (c *Client) D3Act(id int) (*d3c.Act, []byte, error) {
	var (
		dat d3c.Act
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+actPath+"/"+strconv.Itoa(id)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Blacksmith returns blacksmith data
func (c *Client) D3Blacksmith() (*d3c.Artisan, []byte, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+blacksmithPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Jeweler returns jeweler data
func (c *Client) D3Jeweler() (*d3c.Artisan, []byte, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+jewelerPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Mystic returns mystic data
func (c *Client) D3Mystic() (*d3c.Artisan, []byte, error) {
	var (
		dat d3c.Artisan
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+mysticPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3BlacksmithRecipe returns blacksmith recipe data
func (c *Client) D3BlacksmithRecipe(recipeSlug string) (*d3c.Recipe, []byte, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+blacksmithRecipePath+"/"+recipeSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3JewelerRecipe returns jeweler recipe data
func (c *Client) D3JewelerRecipe(recipeSlug string) (*d3c.Recipe, []byte, error) {
	var (
		dat d3c.Recipe
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+jewelerRecipePath+"/"+recipeSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Enchantress returns enchantress data
func (c *Client) D3Enchantress() (*d3c.Follower, []byte, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+enchantressPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Scoundrel returns scoundrel data
func (c *Client) D3Scoundrel() (*d3c.Follower, []byte, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+scoundrelPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Templar returns templar data
func (c *Client) D3Templar() (*d3c.Follower, []byte, error) {
	var (
		dat d3c.Follower
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+templarPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Barbarian returns barbarian data
func (c *Client) D3Barbarian() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+barbarianPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Crusader returns crusader data
func (c *Client) D3Crusader() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+crusaderPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3DemonHunter returns demon hunter data
func (c *Client) D3DemonHunter() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+demonHunterPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Monk returns monk data
func (c *Client) D3Monk() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+monkPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Necromancer returns necromancer data
func (c *Client) D3Necromancer() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+necromancerPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Wizard returns wizard data
func (c *Client) D3Wizard() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wizardPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3WitchDoctor returns witch doctor data
func (c *Client) D3WitchDoctor() (*d3c.Hero, []byte, error) {
	var (
		dat d3c.Hero
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+witchDoctorPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3BarbarianSkill returns barbarian skill data
func (c *Client) D3BarbarianSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+barbarianSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3CrusaderSkill returns crusader skill data
func (c *Client) D3CrusaderSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+crusaderSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3DemonHunterSkill returns demon hunter skill data
func (c *Client) D3DemonHunterSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+demonHunterSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3MonkSkill returns monk skill data
func (c *Client) D3MonkSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+monkSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3NecromancerSkill returns necromancer skill data
func (c *Client) D3NecromancerSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+necromancerSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3WizardSkill returns wizard skill data
func (c *Client) D3WizardSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+wizardSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3WitchDoctorSkill returns witch doctor skill data
func (c *Client) D3WitchDoctorSkill(skillSlug string) (*d3c.Skill, []byte, error) {
	var (
		dat d3c.Skill
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+witchDoctorSkillPath+"/"+skillSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3ItemTypeIndex returns an index of item types
func (c *Client) D3ItemTypeIndex() (*d3c.ItemTypeIndex, []byte, error) {
	var (
		dat d3c.ItemTypeIndex
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemTypePath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3ItemType returns item type data
func (c *Client) D3ItemType(itemTypeSlug string) (*d3c.ItemType, []byte, error) {
	var (
		dat d3c.ItemType
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemTypePath+"/"+itemTypeSlug+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Item returns item data
func (c *Client) D3Item(itemSlug, itemID string) (*d3c.Item, []byte, error) {
	var (
		dat d3c.Item
		b   []byte
		err error
	)

	b, err = c.getURLBody(c.apiURL+itemPath+"/"+itemSlug+"-"+itemID+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3Profile returns profile data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3Profile(account string) (*d3c.Profile, []byte, error) {
	var (
		dat d3c.Profile
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+"/?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3ProfileHero returns profile's hero data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHero(account string, heroID int) (*d3c.ProfileHero, []byte, error) {
	var (
		dat d3c.ProfileHero
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3ProfileHeroItems returns profile's hero's item data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHeroItems(account string, heroID int) (*d3c.ProfileHeroItems, []byte, error) {
	var (
		dat d3c.ProfileHeroItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+profileItemsPartialPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}

// D3ProfileHeroFollowerItems returns profile's hero's follower item data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHeroFollowerItems(account string, heroID int) (*d3c.ProfileHeroFollowerItems, []byte, error) {
	var (
		dat d3c.ProfileHeroFollowerItems
		b   []byte
		err error
	)

	account = strings.Replace(account, "#", "-", 1)

	b, err = c.getURLBody(c.apiURL+d3ProfilePath+"/"+account+profileHeroPartialPath+"/"+strconv.Itoa(heroID)+profileFollowerItemsPartialPath+"?"+localeQuery+c.locale.String(), "")
	if err != nil {
		return &dat, b, err
	}

	err = json.Unmarshal(b, &dat)
	if err != nil {
		return &dat, b, err
	}

	return &dat, b, nil
}
