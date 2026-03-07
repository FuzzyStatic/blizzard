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
	http.Redirect(w, r, cfg.AuthCodeURL("state123"), http.StatusFound)
}

func authorize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := r.Form.Get("state")
	if state != "state123" {
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

	// request playable titles from SPS
	pts, _, err := usBlizzClient.SPSPlayableTitles(context.Background(), token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(pts)
}

func init() {
	clientID = os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("CLIENT_ID must be set")
	}
	clientSecret = os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("CLIENT_SECRET must be set")
	}
	myDomain = os.Getenv("MY_DOMAIN")
	if myDomain == "" {
		log.Fatal("MY_DOMAIN must be set")
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

	cfg = usBlizzClient.AuthorizeConfig(
		fmt.Sprintf("http://%s:9094/oauth2", myDomain),
		oauth.ProfileStreamingTitles,
		oauth.ProfileOpenID,
	)

	http.HandleFunc("/", homepage)
	http.HandleFunc("/oauth2", authorize)

	log.Println("SPS example running on :9094")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
