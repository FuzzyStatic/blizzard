package blizzard

import (
	"context"
	"fmt"

	"github.com/FuzzyStatic/blizzard/v2/d3c"
)

// D3ActIndex returns an index of acts
func (c *Client) D3ActIndex(ctx context.Context) (*d3c.ActIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/act"),
		"",
		&d3c.ActIndex{},
	)
	return dat.(*d3c.ActIndex), b, err
}

// D3Act returns information about act based on ID
func (c *Client) D3Act(ctx context.Context, id int) (*d3c.Act, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/act/%d", id),
		"",
		&d3c.Act{},
	)
	return dat.(*d3c.Act), b, err
}

// D3Blacksmith returns blacksmith data
func (c *Client) D3Blacksmith(ctx context.Context) (*d3c.Artisan, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/artisan/blacksmith"),
		"",
		&d3c.Artisan{},
	)
	return dat.(*d3c.Artisan), b, err
}

// D3Jeweler returns jeweler data
func (c *Client) D3Jeweler(ctx context.Context) (*d3c.Artisan, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/artisan/jeweler"),
		"",
		&d3c.Artisan{},
	)
	return dat.(*d3c.Artisan), b, err
}

// D3Mystic returns mystic data
func (c *Client) D3Mystic(ctx context.Context) (*d3c.Artisan, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/artisan/mystic"),
		"",
		&d3c.Artisan{},
	)
	return dat.(*d3c.Artisan), b, err
}

// D3BlacksmithRecipe returns blacksmith recipe data
func (c *Client) D3BlacksmithRecipe(ctx context.Context, recipeSlug string) (*d3c.Recipe, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/artisan/blacksmith/recipe/%s", recipeSlug),
		"",
		&d3c.Recipe{},
	)
	return dat.(*d3c.Recipe), b, err
}

// D3JewelerRecipe returns jeweler recipe data
func (c *Client) D3JewelerRecipe(ctx context.Context, recipeSlug string) (*d3c.Recipe, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/artisan/jeweler/recipe/%s", recipeSlug),
		"",
		&d3c.Recipe{},
	)
	return dat.(*d3c.Recipe), b, err
}

// D3Enchantress returns enchantress data
func (c *Client) D3Enchantress(ctx context.Context) (*d3c.Follower, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/follower/enchantress"),
		"",
		&d3c.Follower{},
	)
	return dat.(*d3c.Follower), b, err
}

// D3Scoundrel returns scoundrel data
func (c *Client) D3Scoundrel(ctx context.Context) (*d3c.Follower, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/follower/scoundrel"),
		"",
		&d3c.Follower{},
	)
	return dat.(*d3c.Follower), b, err
}

// D3Templar returns templar data
func (c *Client) D3Templar(ctx context.Context) (*d3c.Follower, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/follower/templar"),
		"",
		&d3c.Follower{},
	)
	return dat.(*d3c.Follower), b, err
}

// D3Barbarian returns barbarian data
func (c *Client) D3Barbarian(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/barbarian"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3Crusader returns crusader data
func (c *Client) D3Crusader(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/crusader"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3DemonHunter returns demon hunter data
func (c *Client) D3DemonHunter(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/demon-hunter"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3Monk returns monk data
func (c *Client) D3Monk(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/monk"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3Necromancer returns necromancer data
func (c *Client) D3Necromancer(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/necromancer"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3Wizard returns wizard data
func (c *Client) D3Wizard(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/wizard"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3WitchDoctor returns witch doctor data
func (c *Client) D3WitchDoctor(ctx context.Context) (*d3c.Hero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/witch-doctor"),
		"",
		&d3c.Hero{},
	)
	return dat.(*d3c.Hero), b, err
}

// D3BarbarianSkill returns barbarian skill data
func (c *Client) D3BarbarianSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/barbarian/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3CrusaderSkill returns crusader skill data
func (c *Client) D3CrusaderSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/crusader/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3DemonHunterSkill returns demon hunter skill data
func (c *Client) D3DemonHunterSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/demon-hunter/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3MonkSkill returns monk skill data
func (c *Client) D3MonkSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/monk/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3NecromancerSkill returns necromancer skill data
func (c *Client) D3NecromancerSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/necromancer/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3WizardSkill returns wizard skill data
func (c *Client) D3WizardSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/wizard/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3WitchDoctorSkill returns witch doctor skill data
func (c *Client) D3WitchDoctorSkill(ctx context.Context, skillSlug string) (*d3c.Skill, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/hero/witch-doctor/skill/%s", skillSlug),
		"",
		&d3c.Skill{},
	)
	return dat.(*d3c.Skill), b, err
}

// D3ItemTypeIndex returns an index of item types
func (c *Client) D3ItemTypeIndex(ctx context.Context) (*d3c.ItemTypeIndex, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/item-type"),
		"",
		&d3c.ItemTypeIndex{},
	)
	return dat.(*d3c.ItemTypeIndex), b, err
}

// D3ItemType returns item type data
func (c *Client) D3ItemType(ctx context.Context, itemTypeSlug string) (*d3c.ItemType, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/item-type/%s", itemTypeSlug),
		"",
		&d3c.ItemType{},
	)
	return dat.(*d3c.ItemType), b, err
}

// D3Item returns item data
func (c *Client) D3Item(ctx context.Context, itemSlug, itemID string) (*d3c.Item, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/data/item/%s-%s", itemSlug, itemID),
		"",
		&d3c.Item{},
	)
	return dat.(*d3c.Item), b, err
}

// D3Profile returns profile data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3Profile(ctx context.Context, account string) (*d3c.Profile, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/profile/%s/", formatAccount(account)),
		"",
		&d3c.Profile{},
	)
	return dat.(*d3c.Profile), b, err
}

// D3ProfileHero returns profile's hero data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHero(ctx context.Context, account string, heroID int) (*d3c.ProfileHero, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/profile/%s/hero/%d", formatAccount(account), heroID),
		"",
		&d3c.ProfileHero{},
	)
	return dat.(*d3c.ProfileHero), b, err
}

// D3ProfileHeroItems returns profile's hero's item data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHeroItems(ctx context.Context, account string, heroID int) (*d3c.ProfileHeroItems, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/profile/%s/hero/%d/items", formatAccount(account), heroID),
		"",
		&d3c.ProfileHeroItems{},
	)
	return dat.(*d3c.ProfileHeroItems), b, err
}

// D3ProfileHeroFollowerItems returns profile's hero's follower item data
// Formats accepted: Player-1234 or Player#1234
func (c *Client) D3ProfileHeroFollowerItems(ctx context.Context, account string, heroID int) (*d3c.ProfileHeroFollowerItems, []byte, error) {
	dat, b, err := c.getStructData(ctx,
		fmt.Sprintf("/d3/profile/%s/hero/%d/follower-items", formatAccount(account), heroID),
		"",
		&d3c.ProfileHeroFollowerItems{},
	)
	return dat.(*d3c.ProfileHeroFollowerItems), b, err
}
