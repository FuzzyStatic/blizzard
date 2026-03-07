package blizzard

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"
)

func TestClient_SPSPlayableTitles(t *testing.T) {
	// This test is a placeholder; exercising the real Streaming
	// Provider Service requires a valid OAuth2 token obtained with
	// the streaming.titles and openid scopes.  We skip it by default.
	t.Skip("integration test requiring valid SPS token")

	// Example code (not executed) follows:
	//
	// token := &oauth2.Token{AccessToken: "<token>"}
	// pts, hdr, err := usClient.SPSPlayableTitles(context.Background(), token)
	// if err != nil {
	//     t.Fatalf("unexpected error: %v", err)
	// }
	// if pts == nil || len(pts.Titles) == 0 {
	//     t.Errorf("no titles returned")
	// }
	// _ = hdr
}

// TestClient_SPSPlayableTitles_Mock exercises the call path using an httptest
// server.  It verifies that the bearer token and locale query param are sent
// and that the JSON response is decoded correctly.
func TestClient_SPSPlayableTitles_Mock(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/streaming/titles", func(w http.ResponseWriter, r *http.Request) {
		// header checks
		auth := r.Header.Get("Authorization")
		if auth != "Bearer testtoken" {
			t.Errorf("unexpected auth header: %s", auth)
		}
		if r.URL.Query().Get("locale") != "en_US" {
			t.Errorf("locale query missing or wrong: %s", r.URL.Query().Get("locale"))
		}
		// send simple JSON
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"titles":[{"id":1,"name":"Foo"}]}`)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewClient(Config{
		ClientID:     "id",
		ClientSecret: "secret",
		HTTPClient:   http.DefaultClient,
		Region:       US,
		Locale:       EnUS,
	})
	if err != nil {
		t.Fatal(err)
	}
	// override host to mock
	client.apiHost = srv.URL

	token := &oauth2.Token{AccessToken: "testtoken"}
	pts, _, err := client.SPSPlayableTitles(context.Background(), token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pts == nil || len(pts.Titles) != 1 || pts.Titles[0].ID != 1 {
		t.Errorf("unexpected result: %+v", pts)
	}
}
