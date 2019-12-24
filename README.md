# blizzard

[![GoDoc](https://godoc.org/github.com/FuzzyStatic/blizzard?status.svg)](http://godoc.org/github.com/FuzzyStatic/blizzard) [![Go Report Card](https://goreportcard.com/badge/github.com/FuzzyStatic/blizzard)](https://goreportcard.com/report/github.com/FuzzyStatic/blizzard) [![Build Status](https://travis-ci.org/FuzzyStatic/blizzard.svg?branch=master)](https://travis-ci.org/FuzzyStatic/blizzard) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/fa25319c93814ff4878ee049f04317d4)](https://www.codacy.com/app/FuzzyStatic/blizzard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=FuzzyStatic/blizzard&amp;utm_campaign=Badge_Grade)

> This is a Go client library for gathering [Blizzard API reference](https://develop.battle.net/documentation/api-reference)  data

## Table of Contents

- [blizzard](#blizzard)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Fetching OAuth Data](#fetching-oauth-data)
    - [Fetching Diablo 3 Data](#fetching-diablo-3-data)
    - [Fetching Hearthstone Data](#fetching-hearthstone-data)
    - [Fetching StarCraft 2 Data](#fetching-starcraft-2-data)
    - [Fetching World of Warcraft Data](#fetching-world-of-warcraft-data)
    - [Fetching World of Warcraft Classic Data](#fetching-world-of-warcraft-classic-data)
  - [Documentation](#documentation)
  - [Special Thanks](#special-thanks)

## Getting Started

First, download the Blizzard library:

```shell
go get github.com/FuzzyStatic/blizzard
```

Start using the library by initiating a new Blizzard config structure for desired region and locale (client_id and client_secret can be acquired through your developer account at [https://develop.battle.net/](https://develop.battle.net/)) and requesting an access token:

```go
blizz := blizzard.NewClient("client_id", "client_secret", blizzard.US, blizzard.EnUS)

err := blizz.AccessTokenReq()
if err != nil {
  fmt.Println(err)
}
```

### Fetching OAuth Data

Now you can fetch data from the Blizzard API. For example, you validate your token:

```go
dat, _, err := blizz.TokenValidation()
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching Diablo 3 Data

You can use the functions prefixed with "D3" to acquire Diablo 3 information. For example, you can get information about the current D3 hardcore necromancer leaderboards:

```go
dat, _, err := blizz.D3SeasonLeaderboardHardcoreNecromancer(15)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching Hearthstone Data

You can use the functions prefixed with "HS" to acquire Hearthstone information. For example, you can get information about all the Hearthstone cards:

```go
dat, _, err := blizz.HSCardsAll()
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching StarCraft 2 Data

You can use the functions prefixed with "SC2" to acquire StarCraft 2 information. For example, you can get information about the current SC2 grandmaster ladder:

```go
dat, _, err := blizz.SC2LadderGrandmaster(blizzard.EU)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching World of Warcraft Data

You can use the functions prefixed with "WoW" to acquire World of Warcraft information. For example, you can get information about your WoW character profile:

```go
dat, _, err := blizz.WoWCharacterProfileSummary("emerald-dream", "wildz")
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

or get information about specific spells:

```go
dat, _, err := blizz.WoWSpell(17086)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

or the PVP leaderboards:

```go
dat, _, err := blizz.WoWPVPLeaderboard(wowc.Bracket3v3)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching World of Warcraft Classic Data

You can use the functions prefixed with "ClassicWoW" to acquire World of Warcraft Classic information. For example, you can get information about WoW Classic creature data:

```go
dat, _, err := blizz.ClassicWoWCreature(30)
if err != nil {
  fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

## Documentation

See the [Blizzard API reference](https://develop.battle.net/documentation/api-reference) and the [godoc](http://godoc.org/github.com/FuzzyStatic/blizzard) for all the different datasets that can be acquired.

## Special Thanks

Thanks to [JSON-to-Go](https://mholt.github.io/json-to-go/) for making JSON to Go structure creation simple.
