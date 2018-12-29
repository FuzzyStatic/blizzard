package blizzard

import (
	"fmt"
	"testing"
)

func TestWoWConnectedRealmIndex(t *testing.T) {
	dat, err := c.WoWConnectedRealmIndex()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}

func TestWoWConnectedRealm(t *testing.T) {
	dat, err := c.WoWConnectedRealm(11)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", dat)
}
