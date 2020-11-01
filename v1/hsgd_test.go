package blizzard

import (
	"fmt"
	"testing"

	"github.com/FuzzyStatic/blizzard/v1/hsgd"
)

func TestHSCardsSearch(t *testing.T) {
	dat, _, err := c.HSCardsSearch()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSDetailedCardsSearch(t *testing.T) {
	dat, _, err := c.HSDetailedCardsSearch(
		"rise-of-shadows", "mage", "legendary", "minion", "dragon", "battlecry", "kalecgos",
		[]int{10}, []int{4}, []int{10}, 1, 5,
		hsgd.CollectibilityCollectible, hsgd.SortName, hsgd.OrderAsc,
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSBattlegroundsCardsSearch(t *testing.T) {
	dat, _, err := c.HSBattlegroundsCardsSearch(
		"", "", "", "", "",
		[]int{}, []int{}, []int{}, 0, 0,
		[]hsgd.Tier{hsgd.TierHero, hsgd.Tier3}, "", "", "",
	)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCardByIDOrSlug(t *testing.T) {
	dat, _, err := c.HSCardByIDOrSlug("52119-arch-villain-rafaam", hsgd.GameModeConstructed)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCardBackSearchAllLocales(t *testing.T) {
	dat, _, err := c.HSCardBackSearchAllLocales(hsgd.CardBackCategoryBase, "", hsgd.SortName, hsgd.OrderAsc)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCardBackByIDOrSlug(t *testing.T) {
	dat, _, err := c.HSCardBackByIDOrSlug("155-pizza-stone")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCardBackSearch(t *testing.T) {
	dat, _, err := c.HSCardBackSearch("", "", hsgd.SortName, hsgd.OrderAsc)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSDeck(t *testing.T) {
	dat, _, err := c.HSDeck("AAECAQcG+wyd8AKS+AKggAOblAPanQMMS6IE/web8wLR9QKD+wKe+wKz/AL1gAOXlAOalAOSnwMA")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadata(t *testing.T) {
	dat, _, err := c.HSMetadata()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataSets(t *testing.T) {
	dat, _, err := c.HSMetadataSets()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataSetGroups(t *testing.T) {
	dat, _, err := c.HSMetadataSetGroups()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataTypes(t *testing.T) {
	dat, _, err := c.HSMetadataTypes()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataRarities(t *testing.T) {
	dat, _, err := c.HSMetadataRarities()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataClasses(t *testing.T) {
	dat, _, err := c.HSMetadataClasses()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataMinionTypes(t *testing.T) {
	dat, _, err := c.HSMetadataMinionTypes()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSMetadataKeywords(t *testing.T) {
	dat, _, err := c.HSMetadataKeywords()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
