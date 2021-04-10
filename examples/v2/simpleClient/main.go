package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/FuzzyStatic/blizzard/v2"
)

var (
	clientID     string
	clientSecret string
	blizz        *blizzard.Client
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
	blizz = blizzard.NewClient(
		clientID,
		clientSecret,
		blizzard.US,
		blizzard.EnUS,
	)

	err := blizz.AccessTokenRequest(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
