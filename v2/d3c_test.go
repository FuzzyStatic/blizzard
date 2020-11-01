package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestD3ActIndex(t *testing.T) {
	dat, _, err := c.D3ActIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Act(t *testing.T) {
	dat, _, err := c.D3Act(context.Background(), 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Blacksmith(t *testing.T) {
	dat, _, err := c.D3Blacksmith(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Jeweler(t *testing.T) {
	dat, _, err := c.D3Jeweler(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Mystic(t *testing.T) {
	dat, _, err := c.D3Mystic(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3BlacksmithRecipe(t *testing.T) {
	dat, _, err := c.D3BlacksmithRecipe(context.Background(), "apprentice-flamberge")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3JewelerRecipe(t *testing.T) {
	dat, _, err := c.D3JewelerRecipe(context.Background(), "flawless-amethyst")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Enchantress(t *testing.T) {
	dat, _, err := c.D3Enchantress(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Scoundrel(t *testing.T) {
	dat, _, err := c.D3Scoundrel(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Templar(t *testing.T) {
	dat, _, err := c.D3Templar(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Barbarian(t *testing.T) {
	dat, _, err := c.D3Barbarian(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Crusader(t *testing.T) {
	dat, _, err := c.D3Crusader(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3DemonHunter(t *testing.T) {
	dat, _, err := c.D3DemonHunter(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Monk(t *testing.T) {
	dat, _, err := c.D3Monk(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Necromancer(t *testing.T) {
	dat, _, err := c.D3Necromancer(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Wizard(t *testing.T) {
	dat, _, err := c.D3Wizard(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WitchDoctor(t *testing.T) {
	dat, _, err := c.D3WitchDoctor(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3BarbarianSkill(t *testing.T) {
	dat, _, err := c.D3BarbarianSkill(context.Background(), "bash")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3CrusaderSkill(t *testing.T) {
	dat, _, err := c.D3CrusaderSkill(context.Background(), "condemn")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3DemonHunterSkill(t *testing.T) {
	dat, _, err := c.D3DemonHunterSkill(context.Background(), "smoke-screen")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3MonkSkill(t *testing.T) {
	dat, _, err := c.D3MonkSkill(context.Background(), "dashing-strike")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3NecromancerSkill(t *testing.T) {
	dat, _, err := c.D3NecromancerSkill(context.Background(), "grim-scythe")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WizardSkill(t *testing.T) {
	dat, _, err := c.D3WizardSkill(context.Background(), "energy-twister")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3WitchDoctorSkill(t *testing.T) {
	dat, _, err := c.D3WitchDoctorSkill(context.Background(), "gargantuan")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ItemTypeIndex(t *testing.T) {
	dat, _, err := c.D3ItemTypeIndex(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ItemType(t *testing.T) {
	dat, _, err := c.D3ItemType(context.Background(), "chestarmorwizard")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Item(t *testing.T) {
	dat, _, err := c.D3Item(context.Background(), "firebirds-breast", "Unique_Chest_Set_06_x1")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3Profile(t *testing.T) {
	dat, _, err := c.D3Profile(context.Background(), "FuzzyStatic#1384")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHero(t *testing.T) {
	dat, _, err := c.D3ProfileHero(context.Background(), "FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHeroItems(t *testing.T) {
	dat, _, err := c.D3ProfileHeroItems(context.Background(), "FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestD3ProfileHeroFollowerItems(t *testing.T) {
	dat, _, err := c.D3ProfileHeroFollowerItems(context.Background(), "FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
