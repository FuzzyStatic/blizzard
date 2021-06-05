package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FuzzyStatic/blizzard/v2"
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
		// This is the client default rate config, we allow 100 requests per second, and a burst budget of 10 requests
		blizzard.NewRateOpt(1*time.Second/100, 10),
		// This is the client default retry config
		blizzard.NewRetryOpt(
			retry.Attempts(3),
			retry.Delay(100*time.Millisecond),
			retry.DelayType(retry.BackOffDelay),
			retry.MaxJitter(0),
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

	mount, _, err := blizz.WoWMountIndex(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(mount, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out[:]))
}
