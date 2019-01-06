package blizzard

import (
	"fmt"
	"testing"
)

func TestD3ActIndex(t *testing.T) {
	dat, _, err := c.D3ActIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Act(t *testing.T) {
	dat, _, err := c.D3Act(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Blacksmith(t *testing.T) {
	dat, _, err := c.D3Blacksmith()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Jeweler(t *testing.T) {
	dat, _, err := c.D3Jeweler()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Mystic(t *testing.T) {
	dat, _, err := c.D3Mystic()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3BlacksmithRecipe(t *testing.T) {
	dat, _, err := c.D3BlacksmithRecipe("apprentice-flamberge")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3JewelerRecipe(t *testing.T) {
	dat, _, err := c.D3JewelerRecipe("flawless-amethyst")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Enchantress(t *testing.T) {
	dat, _, err := c.D3Enchantress()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Scoundrel(t *testing.T) {
	dat, _, err := c.D3Scoundrel()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Templar(t *testing.T) {
	dat, _, err := c.D3Templar()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Barbarian(t *testing.T) {
	dat, _, err := c.D3Barbarian()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Crusader(t *testing.T) {
	dat, _, err := c.D3Crusader()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3DemonHunter(t *testing.T) {
	dat, _, err := c.D3DemonHunter()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Monk(t *testing.T) {
	dat, _, err := c.D3Monk()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Necromancer(t *testing.T) {
	dat, _, err := c.D3Necromancer()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Wizard(t *testing.T) {
	dat, _, err := c.D3Wizard()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WitchDoctor(t *testing.T) {
	dat, _, err := c.D3WitchDoctor()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3BarbarianSkill(t *testing.T) {
	dat, _, err := c.D3BarbarianSkill("bash")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3CrusaderSkill(t *testing.T) {
	dat, _, err := c.D3CrusaderSkill("condemn")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3DemonHunterSkill(t *testing.T) {
	dat, _, err := c.D3DemonHunterSkill("smoke-screen")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3MonkSkill(t *testing.T) {
	dat, _, err := c.D3MonkSkill("dashing-strike")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3NecromancerSkill(t *testing.T) {
	dat, _, err := c.D3NecromancerSkill("grim-scythe")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WizardSkill(t *testing.T) {
	dat, _, err := c.D3WizardSkill("energy-twister")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WitchDoctorSkill(t *testing.T) {
	dat, _, err := c.D3WitchDoctorSkill("gargantuan")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ItemTypeIndex(t *testing.T) {
	dat, _, err := c.D3ItemTypeIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ItemType(t *testing.T) {
	dat, _, err := c.D3ItemType("chestarmorwizard")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Item(t *testing.T) {
	dat, _, err := c.D3Item("firebirds-breast", "Unique_Chest_Set_06_x1")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Profile(t *testing.T) {
	dat, _, err := c.D3Profile("FuzzyStatic#1384")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHero(t *testing.T) {
	dat, _, err := c.D3ProfileHero("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHeroItems(t *testing.T) {
	dat, _, err := c.D3ProfileHeroItems("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHeroFollowerItems(t *testing.T) {
	dat, _, err := c.D3ProfileHeroFollowerItems("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}
