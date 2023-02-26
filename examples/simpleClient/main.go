package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/FuzzyStatic/blizzard/v3"
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

	err = usBlizzClient.AccessTokenRequest(context.Background())
	if err != nil {
		panic(err)
	}
}
