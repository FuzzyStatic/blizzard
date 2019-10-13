package blizzard

import (
	"fmt"
	"log"
	"os"
)

var printOutput string

func init() {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
	}

	printOutput = os.Getenv("PRINT_OUTPUT")
	if printOutput == "" {
		log.Println("Output will not be printed for tests.")
	}

	c = NewClient(clientID, clientSecret, US, EnUS)

	err := c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		return
	}
}
