package blizzard

import (
	"fmt"
	"testing"
)

func TestClassicWoWCreatureFamiliesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamiliesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamily(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamily(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureTypesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureType(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureType(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreature(t *testing.T) {
	dat, _, err := c.ClassicWoWCreature(30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureDisplayMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureDisplayMedia(30)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWCreatureFamilyMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWCreatureFamilyMedia(1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestComponentsIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestComponentsIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestBorderMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestBorderMedia(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWGuildCrestEmblemMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWGuildCrestEmblemMedia(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
