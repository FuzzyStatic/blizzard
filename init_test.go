package blizzard

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	usClient    *Client
	printOutput string
)

func init() {
	var err error
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
	}

	printOutput = os.Getenv("PRINT_OUTPUT")
	if printOutput == "" {
		log.Println("Output will not be printed for tests.")
	}

	usClient, err = NewClient(Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
		Region:       US,
		Locale:       EnUS,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = usClient.AccessTokenRequest(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}
