package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FuzzyStatic/blizzard"
	"github.com/FuzzyStatic/blizzard/oauth"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	myDomain     string
	cfg          oauth2.Config
	blizz        *blizzard.Client
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Hit!")
	http.Redirect(w, r, cfg.AuthCodeURL("my_random_state"), http.StatusFound)
}

// Authorize
func Authorize(w http.ResponseWriter, r *http.Request) {
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

	validation, _, err := blizz.TokenValidation(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", validation)

	userInfo, _, err := blizz.UserInfoHeader(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", userInfo)

	wowProfile, _, err := blizz.WoWUserCharacters(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", wowProfile)
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
	if clientSecret == "" {
		log.Fatal("Set the environment variable MY_DOMAIN before retrying.")
	}
}

func main() {
	blizz = blizzard.NewClient(clientID, clientSecret, blizzard.US, blizzard.EnUS)
	cfg = blizz.AuthorizeConfig(fmt.Sprintf("http://%s:9094/oauth2", myDomain), oauth.ProfileD3, oauth.ProfileSC2, oauth.ProfileWoW)

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/oauth2", Authorize)

	// We start up our Client on port 9094
	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
