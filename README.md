# blizzard

[![GoDoc](https://godoc.org/github.com/FuzzyStatic/blizzard?status.svg)](http://godoc.org/github.com/FuzzyStatic/blizzard) [![Go Report Card](https://goreportcard.com/badge/github.com/FuzzyStatic/blizzard)](https://goreportcard.com/report/github.com/FuzzyStatic/blizzard) [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/FuzzyStatic)

This is a Go client library for gathering Blizzard API game data.

### Getting started

Start by initiating a new Blizzard config structure for desired region (client_id and client_secret can be acquired through your developer account at https://develop.battle.net/):

```go
blizz = New("client_id", "client_secret", US)
```

### Fetching data

Now you can fetch data from the Blizzard API. For example, you validate your token:

```go
dat, err := blizz.TokenValidation()
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

You can get information about the current D3 hardcore necromancer leaderboards:

```go
dat, err := blizz.D3GetSeasonLeaderboardHardcoreNecromancer(15)
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

You can get information about the current SC2 grandmaster ladder:

```go
dat, err := blizz.SC2GetLadderGrandmaster(EU)
if err != nil {
	fmt.Println(err)
}

fmt.Printf("%+v\n", dat)
```

See godoc for all the different datasets that can be acquired.

### Thanks

Thanks to [JSON-to-Go](https://mholt.github.io/json-to-go/) for making JSON to Go structure creation simple.