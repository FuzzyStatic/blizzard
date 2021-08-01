package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FuzzyStatic/blizzard/v3"
	"github.com/FuzzyStatic/blizzard/wowsearch"
)

var (
	clientID     string
	clientSecret string
)

func init() {
	clientID = os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
	}

	clientSecret = os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
	}
}

func main() {
	usBlizzClient, err := blizzard.NewClient(blizzard.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
		Region:       blizzard.US,
		Locale:       blizzard.EnUS,
	})
	if err != nil {
		panic(err)
	}

	realmSearch, _, err := usBlizzClient.ClassicWoWRealmSearch(
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
