package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWClassicCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureFamiliesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreatureFamily(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureFamily(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreatureTypesIndex(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreatureType(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureType(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreature(t *testing.T) {
	dat, _, err := c.WoWClassicCreature(30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreatureDisplayMedia(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureDisplayMedia(30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWClassicCreatureFamilyMedia(t *testing.T) {
	dat, _, err := c.WoWClassicCreatureFamilyMedia(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
