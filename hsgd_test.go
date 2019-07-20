package blizzard

import (
	"fmt"
	"testing"
)

func TestHSCardsAll(t *testing.T) {
	dat, _, err := c.HSCardsAll()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCards(t *testing.T) {
	dat, _, err := c.HSCards("rise-of-shadows", "mage", "legendary", "minion", "dragon", "battlecry", "kalecgos",
		[]int{10}, []int{4}, []int{10}, 1, 5,
		CollectiblilityCollectible, SortName, OrderAsc)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestHSCardByIDOrSlug(t *testing.T) {
	dat, _, err := c.HSCardByIDOrSlug("52119-arch-villain-rafaam")
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
