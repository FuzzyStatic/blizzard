# blizzard

[![GoDoc](https://godoc.org/github.com/FuzzyStatic/blizzard?status.svg)](http://godoc.org/github.com/FuzzyStatic/blizzard) [![Go Report Card](https://goreportcard.com/badge/github.com/FuzzyStatic/blizzard)](https://goreportcard.com/report/github.com/FuzzyStatic/blizzard) [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/FuzzyStatic)

This is a Go client library for gathering Blizzard API game data.

### Game Client Libraries

- [Diablo 3](https://github.com/FuzzyStatic/blizzard/tree/master/diablo3)

- [Starcraft 2](https://github.com/FuzzyStatic/blizzard/tree/master/starcraft2)

- [World of Warcraft](https://github.com/FuzzyStatic/blizzard/tree/master/worldofwarcraft)

### Getting started

Start by initiating a new Blizzard structure (accessToken and apiKey can be acquired through your Mashery account at https://dev.battle.net/):

```go
var c *Config

c = New(
	Auth{
		AccessToken: "myAccessToken",
		APIKey:      "myApiKey",
	},
	US,
)
```

### Fetching data

Now you can fetch data from the Blizzard API. For example, you can get all the data for your account:

```go
accountUser, err := c.GetAccountUser()
if err != nil {
	fmt.Println(err)
}

fmt.Println(accountUser)
```
