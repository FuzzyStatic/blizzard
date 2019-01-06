package blizzard

import (
	"fmt"
	"log"
	"os"
)

var printStruct string

func init() {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
	}

	printStruct = os.Getenv("PRINT_STRUCT")
	if printStruct == "" {
		log.Println("Structures will not be printed.")
	}

	c = NewClient(clientID, clientSecret, US, enUS)

	err := c.AccessTokenReq()
	if err != nil {
		fmt.Println(err)
		return
	}
}
