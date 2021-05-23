package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FuzzyStatic/blizzard/v2"
	"github.com/FuzzyStatic/blizzard/v2/wowsearch"
	"github.com/avast/retry-go"
)

var (
	bnetID     string
	bnetSECRET string
)

func init() {
	bnetID = os.Getenv("BNET_ID")
	bnetSECRET = os.Getenv("BNET_SECRET")
	if bnetID == "" || bnetSECRET == "" {
		log.Fatal("missing BNET_ID or BNET_SECRET")
	}
}
func main() {
	blizz := blizzard.NewClient(
		bnetID,
		bnetSECRET,
		blizzard.EU,
		blizzard.EnUS,
		// This is the client default rate config, we allow 100 requests per second, and a burst of the
		blizzard.NewRateOpt(1*time.Second/100, 10),
		// This is the client default retry config
		blizzard.NewRetryOpt(
			retry.Attempts(3),
			retry.Delay(100*time.Millisecond),
			retry.DelayType(retry.BackOffDelay),
			retry.RetryIf(func(err error) bool {
				switch {
				case err.Error() == "429 Too Many Requests":
					return true // recoverable error, retry
				case err.Error() == "403 Forbidden":
					return false
				case err.Error() == "404 Not Found":
					return false
				default:
					return false // We cannot retry this away
				}
			}),
		),
	)

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
