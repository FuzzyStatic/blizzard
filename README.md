# blizzard

[![GoDoc](https://godoc.org/github.com/FuzzyStatic/blizzard?status.svg)](http://godoc.org/github.com/FuzzyStatic/blizzard) [![Go Report Card](https://goreportcard.com/badge/github.com/FuzzyStatic/blizzard)](https://goreportcard.com/report/github.com/FuzzyStatic/blizzard) [![Build Status](https://travis-ci.org/FuzzyStatic/blizzard.svg?branch=master)](https://travis-ci.org/FuzzyStatic/blizzard) [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/FuzzyStatic)

This is a Go client library for gathering Blizzard API game data.

### Getting started

Start by initiating a new Blizzard config structure for desired region (client_id and client_secret can be acquired through your developer account at https://develop.battle.net/) and requesting an access token:

```go
blizz := blizzard.NewClient("client_id", "client_secret", blizzard.US, blizzard.enUS)

err := blizz.AccessTokenReq()
if err != nil {
	fmt.Println(err)
}
```

### Fetching basic data

Now you can fetch data from the Blizzard API. For example, you validate your token:

```go
dat, _, err := blizz.TokenValidation()
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching Diablo 3 data

You can use the functions prefixed with "D3" to acquire Diablo 3 information. For example, you can get information about the current D3 hardcore necromancer leaderboards:

```go
dat, _, err := blizz.D3SeasonLeaderboardHardcoreNecromancer(15)
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching StarCraft 2 data

You can use the functions prefixed with "SC2" to acquire StarCraft 2 information. For example, you can get information about the current SC2 grandmaster ladder:

```go
dat, _, err := blizz.SC2LadderGrandmaster(EU)
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

### Fetching World of Warcraft data

You can use the functions prefixed with "WoW" to acquire World of Warcraft information. For example, you can get information about your WoW character profile:

```go
dat, _, err := blizz.WoWCharacterProfile("emerald-dream", "Limejelly",
	[]string{
		wowc.FieldCharacterAchievements,
		wowc.FieldCharacterAppearance,
		wowc.FieldCharacterAudit,
		wowc.FieldCharacterFeed,
		wowc.FieldCharacterGuild,
		wowc.FieldCharacterItems,
		wowc.FieldCharacterMounts,
		wowc.FieldCharacterPVP,
		wowc.FieldCharacterPetSlots,
		wowc.FieldCharacterPets,
		wowc.FieldCharacterProfessions,
		wowc.FieldCharacterProgression,
		wowc.FieldCharacterQuests,
		wowc.FieldCharacterReputation,
		wowc.FieldCharacterStatistics,
		wowc.FieldCharacterStats,
		wowc.FieldCharacterTalents,
		wowc.FieldCharacterTitle,
	},
)
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
	t.Fail()
}

fmt.Printf("%+v\n", dat)
```

### Documentation

See the [Blizzard API reference](https://develop.battle.net/documentation/api-reference) and the godoc for all the different datasets that can be acquired.

### Thanks

Thanks to [JSON-to-Go](https://mholt.github.io/json-to-go/) for making JSON to Go structure creation simple.