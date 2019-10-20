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

func TestClassicWoWItemClassesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWItemClassesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemClass(t *testing.T) {
	dat, _, err := c.ClassicWoWItemClass(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemSubclass(t *testing.T) {
	dat, _, err := c.ClassicWoWItemSubclass(0, 0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItem(t *testing.T) {
	dat, _, err := c.ClassicWoWItem(19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWItemMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWItemMedia(19019)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClassesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClass(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClass(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableClassMedia(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableClassMedia(7)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRacesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableRacesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPlayableRace(t *testing.T) {
	dat, _, err := c.ClassicWoWPlayableRace(2)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerTypesIndex(t *testing.T) {
	dat, _, err := c.ClassicWoWPowerTypesIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestClassicWoWPowerType(t *testing.T) {
	dat, _, err := c.ClassicWoWPowerType(0)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
