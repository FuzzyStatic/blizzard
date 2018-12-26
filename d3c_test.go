package blizzard

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("credentials") // replace with TOML file similar to sample.toml
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	clientID := viper.GetString("authentication.client_id")
	clientSecret := viper.GetString("authentication.client_secret")

	c = New(clientID, clientSecret, US)
}

func TestD3GetActIndex(t *testing.T) {
	dat, err := c.D3GetActIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetAct(t *testing.T) {
	dat, err := c.D3GetAct(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetBlacksmith(t *testing.T) {
	dat, err := c.D3GetBlacksmith()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetJeweler(t *testing.T) {
	dat, err := c.D3GetJeweler()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetMystic(t *testing.T) {
	dat, err := c.D3GetMystic()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetBlacksmithRecipe(t *testing.T) {
	dat, err := c.D3GetBlacksmithRecipe("apprentice-flamberge")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetJewelerRecipe(t *testing.T) {
	dat, err := c.D3GetJewelerRecipe("flawless-amethyst")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetEnchantress(t *testing.T) {
	dat, err := c.D3GetEnchantress()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetScoundrel(t *testing.T) {
	dat, err := c.D3GetScoundrel()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetTemplar(t *testing.T) {
	dat, err := c.D3GetTemplar()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetBarbarian(t *testing.T) {
	dat, err := c.D3GetBarbarian()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetCrusader(t *testing.T) {
	dat, err := c.D3GetCrusader()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetDemonHunter(t *testing.T) {
	dat, err := c.D3GetDemonHunter()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetMonk(t *testing.T) {
	dat, err := c.D3GetMonk()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetNecromancer(t *testing.T) {
	dat, err := c.D3GetNecromancer()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetWizard(t *testing.T) {
	dat, err := c.D3GetWizard()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetWitchDoctor(t *testing.T) {
	dat, err := c.D3GetWitchDoctor()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetBarbarianSkill(t *testing.T) {
	dat, err := c.D3GetBarbarianSkill("bash")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetCrusaderSkill(t *testing.T) {
	dat, err := c.D3GetCrusaderSkill("condemn")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetDemonHunterSkill(t *testing.T) {
	dat, err := c.D3GetDemonHunterSkill("smoke-screen")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetMonkSkill(t *testing.T) {
	dat, err := c.D3GetMonkSkill("dashing-strike")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetNecromancerSkill(t *testing.T) {
	dat, err := c.D3GetNecromancerSkill("grim-scythe")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetWizardSkill(t *testing.T) {
	dat, err := c.D3GetWizardSkill("energy-twister")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetWitchDoctorSkill(t *testing.T) {
	dat, err := c.D3GetWitchDoctorSkill("gargantuan")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetItemTypeIndex(t *testing.T) {
	dat, err := c.D3GetItemTypeIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetItemType(t *testing.T) {
	dat, err := c.D3GetItemType("chestarmorwizard")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetItem(t *testing.T) {
	dat, err := c.D3GetItem("firebirds-breast", "Unique_Chest_Set_06_x1")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetProfile(t *testing.T) {
	dat, err := c.D3GetProfile("FuzzyStatic#1384")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetProfileHero(t *testing.T) {
	dat, err := c.D3GetProfileHero("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetProfileHeroItems(t *testing.T) {
	dat, err := c.D3GetProfileHeroItems("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestD3GetProfileHeroFollowerItems(t *testing.T) {
	dat, err := c.D3GetProfileHeroFollowerItems("FuzzyStatic#1384", 98663224)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
