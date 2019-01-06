package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWCharacterMythicKeystoneProfile(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfile("illidan", "norilockxo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestWoWCharacterMythicKeystoneProfileSeason(t *testing.T) {
	dat, _, err := c.WoWCharacterMythicKeystoneProfileSeason("illidan", "norilockxo", 1)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printStruct != "" {
		fmt.Printf("%+v\n", dat)
	}
}
