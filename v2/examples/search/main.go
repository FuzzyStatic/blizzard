package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/FuzzyStatic/blizzard/v2"
	"github.com/FuzzyStatic/blizzard/v2/wowsearch"
)

var (
	bnetID     string
	bnetSECRET string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	bnetID = os.Getenv("BNET_ID")
	bnetSECRET = os.Getenv("BNET_SECRET")
	if bnetID == "" || bnetSECRET == "" {
		log.Fatal("missing BNET_ID or BNET_SECRET")
	}
}
func main() {
	blizz := blizzard.NewClient(bnetID, bnetSECRET, blizzard.EU, blizzard.EnUS)

	realmSearch, _, err := blizz.ClassicRealmSearch(
		context.TODO(),
		wowsearch.Page(1),
		wowsearch.PageSize(5),
		wowsearch.OrderBy("name.EN_US:asc"),
		wowsearch.Field().
			AND("timezone", "Europe/Paris").
			AND("data.locale", "enGB").
			NOT("type.type", "PVP").
			NOT("id", "4756||4757").
			OR("type.type", "NORMAL", "RP"),
	)
	if err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(realmSearch, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out[:]))
}
