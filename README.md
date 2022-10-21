# blizzard

[![Go Reference](https://pkg.go.dev/badge/github.com/FuzzyStatic/blizzard/v3.svg)](https://pkg.go.dev/github.com/FuzzyStatic/blizzard/v3) [![Go Report Card](https://goreportcard.com/badge/github.com/FuzzyStatic/blizzard)](https://goreportcard.com/report/github.com/FuzzyStatic/blizzard)

> This is a Go client library for gathering [Blizzard API reference](https://develop.battle.net/documentation/api-reference) data

## Table of Contents

- [blizzard](#blizzard)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Fetching Diablo 3 Data](#fetching-diablo-3-data)
    - [Fetching Hearthstone Data](#fetching-hearthstone-data)
    - [Fetching StarCraft 2 Data](#fetching-starcraft-2-data)
    - [Fetching World of Warcraft Data](#fetching-world-of-warcraft-data)
    - [Fetching World of Warcraft Classic Data](#fetching-world-of-warcraft-classic-data)
  - [Authorization for User Data](#authorization-for-user-data)
    - [Fetching OAuth Data](#fetching-oauth-data)
  - [Header Information](#header-information)
  - [Documentation](#documentation)
  - [Special Thanks](#special-thanks)

## Getting Started

First, download the Blizzard library:

```shell
go get github.com/FuzzyStatic/blizzard/v3
```

Start using the library by initiating a new Blizzard config structure for your desired region and locale (client_id and client_secret can be acquired through your developer account at [https://develop.battle.net/](https://develop.battle.net/)) and requesting an access token:

```go
usBlizzClient := blizzard.NewClient(blizzard.Config{
  ClientID:     "my_client_id",
  ClientSecret: "my_client_secret",
  HTTPClient:   http.DefaultClient,
  Region:       blizzard.US,
  Locale:       blizzard.EnUS,
})

err := usBlizzClient.AccessTokenRequest(ctx)
if err != nil {
  fmt.Println(err)
}

euBlizzClient := blizzard.NewClient(blizzard.Config{
  ClientID:     "my_client_id",
  ClientSecret: "my_client_secret",
  HTTPClient:   http.DefaultClient,
  Region:       blizzard.EU,
  Locale:       blizzard.EnGB,
})

err := euBlizzClient.AccessTokenRequest(ctx)
if err != nil {
  fmt.Println(err)
}
```

### Fetching Diablo 3 Data

You can use the functions prefixed with "D3" to acquire Diablo 3 information. For example, you can get information about the current D3 hardcore necromancer leaderboards:

```go
dat, _, err := usBlizzClient.D3SeasonLeaderboardHardcoreNecromancer(ctx, 15)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching Hearthstone Data

You can use the functions prefixed with "HS" to acquire Hearthstone information. For example, you can get information about all the Hearthstone cards:

```go
dat, _, err := usBlizzClient.HSCardsAll(ctx)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching StarCraft 2 Data

You can use the functions prefixed with "SC2" to acquire StarCraft 2 information. For example, you can get information about the current SC2 grandmaster ladder:

```go
dat, _, err := usBlizzClient.SC2LadderGrandmaster(ctx, blizzard.EU)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching World of Warcraft Data

You can use the functions prefixed with "WoW" to acquire World of Warcraft information. For example, you can get information about your WoW character profile:

```go
dat, _, err := usBlizzClient.WoWCharacterProfileSummary(ctx, "illidan", "wildz")
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

or get information about specific spells:

```go
dat, _, err := usBlizzClient.WoWSpell(ctx, 17086)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

or the PvP leaderboards:

```go
dat, _, err := usBlizzClient.WoWCharacterPvPBracketStatistics(ctx, wowp.Bracket3v3)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching World of Warcraft Classic Data

You can use the functions prefixed with "ClassicWoW" to acquire World of Warcraft Classic information. For example, you can get information about WoW Classic creature data:

```go
dat, _, err := usBlizzClient.ClassicWoWCreature(ctx, 30)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

## Authorization for User Data

To use the `UserInfoHeader` or `WoWUserCharacters` functions to acquire data about other users (and not your own), you must use the OAuth2 redirect method to get an authorized token. This is useful for building websites that display more personal or individualized data. The following code snippet is an example on how to acquire authorized tokens for other users. A working example can be found in the [examples/authCodeFlow](https://github.com/FuzzyStatic/blizzard/tree/master/examples/authCodeFlow) directory. Before the redirect URI will work, you will have to add it to your client settings at <https://develop.battle.net/access>:

```go
package main

import (
  "context"
  "encoding/json"
  "fmt"
  "log"
  "net/http"

  "github.com/FuzzyStatic/blizzard"
  "github.com/FuzzyStatic/blizzard/oauth"
  "golang.org/x/oauth2"
)

var (
  cfg           oauth2.Config
  usBlizzClient *blizzard.Client
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Homepage Hit!")
  u := cfg.AuthCodeURL("my_random_state")
  http.Redirect(w, r, u, http.StatusFound)
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

  dat1, _, err := usBlizzClient.UserInfoHeader(token)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Printf("%+v\n", dat1)

  dat2, _, err := usBlizzClient.WoWUserCharacters(token)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  fmt.Printf("%+v\n", dat2)
}

func main() {
  blizz = blizzard.NewClient("client_id", "client_secret", blizzard.US, blizzard.EnUS)
  cfg = usBlizzClient.AuthorizeConfig("http://<mydomain>:9094/oauth2", oauth.ProfileD3, oauth.ProfileSC2, oauth.ProfileWoW)

  http.HandleFunc("/", HomePage)
  http.HandleFunc("/oauth2", Authorize)

  // We start up our Client on port 9094
  log.Println("Client is running at 9094 port.")
  log.Fatal(http.ListenAndServe(":9094", nil))
}
```

### Fetching OAuth Data

Now you can validate those tokens with the OAuth API:

```go
dat, _, err := usBlizzClient.TokenValidation(ctx, token)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

## Header Information

Each API call will return HTTP response header information, if any. Use the second return variable to get a structure containing the response header information.

```go
dat, header, err := usBlizzClient.WoWAuctions(context.Background(), 1138)
if err != nil {
  fmt.Println(err)
}

fmt.Println(header.BattlenetNamespace)
fmt.Println(header.LastModified)
...
```

## Documentation

See the [Blizzard API reference](https://develop.battle.net/documentation/guides) and the [Go reference](https://pkg.go.dev/github.com/FuzzyStatic/blizzard/v2) for all the different datasets that can be acquired.

## Special Thanks

Thanks to [JSON-to-Go](https://mholt.github.io/json-to-go/) for making JSON to Go structure creation simple.

Thanks to all who contribute to keep this package current.
