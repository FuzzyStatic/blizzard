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
