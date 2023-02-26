package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FuzzyStatic/blizzard/v3"
	"github.com/FuzzyStatic/blizzard/v3/oauth"
	"golang.org/x/oauth2"
)

var (
	clientID      string
	clientSecret  string
	myDomain      string
	cfg           oauth2.Config
	usBlizzClient *blizzard.Client
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Hit!")
	http.Redirect(w, r, cfg.AuthCodeURL("my_random_state"), http.StatusFound)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := r.Form.Get("state")
	if state != "my_random_state" {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(*token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validation, _, err := usBlizzClient.TokenValidation(context.Background(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", validation)

	userInfo, _, err := usBlizzClient.UserInfoHeader(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", userInfo)
}

func init() {
	clientID = os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("Set the environment variable CLIENT_ID before retrying.")
	}

	clientSecret = os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("Set the environment variable CLIENT_SECRET before retrying.")
	}

	myDomain = os.Getenv("MY_DOMAIN")
	if myDomain == "" {
		log.Fatal("Set the environment variable MY_DOMAIN before retrying.")
	}
}

func main() {
	var err error

	usBlizzClient, err = blizzard.NewClient(blizzard.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
		Region:       blizzard.US,
		Locale:       blizzard.EnUS,
	})
	if err != nil {
		panic(err)
	}

	cfg = usBlizzClient.AuthorizeConfig(fmt.Sprintf("http://%s:9094/oauth2", myDomain), oauth.ProfileD3, oauth.ProfileSC2, oauth.ProfileWoW)

	http.HandleFunc("/", homepage)
	http.HandleFunc("/oauth2", authorize)

	// We start up our Client on port 9094
	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
