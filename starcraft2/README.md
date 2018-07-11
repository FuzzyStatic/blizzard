# blizzard/starcraft2

[![GoDoc](https://godoc.org/github.com/FuzzyStatic/blizzard/starcraft2?status.svg)](http://godoc.org/github.com/FuzzyStatic/blizzard/starcraft2)

This is a Go client library for gathering Blizzard Starcraft 2 API data.

### Getting started

Start by initiating a new Starcraft2 structure (accessToken and apiKey can be acquired through your Mashery account at https://dev.battle.net/):

```go
var s *Starcraft2

s := New(
	blizzard.Auth{
		AccessToken: "myAccessToken",
		APIKey:      "myApiKey",
	},
	blizzard.US,
)
```

### Fetching data

Now you can fetch data from the Starcraft 2 API. For example, you can get all the data for a profile:

```go
profile, err := s.GetProfile(2537456, 1, "Neeb")
if err != nil {
	fmt.Println(err)
}

fmt.Println(*profile)
```

or all the ladders for a profile:

```go
profileLadders, err := s.GetProfileLadders(2537456, 1, "Neeb")
if err != nil {
	fmt.Println(err)
}

fmt.Println(*profileLadders)
```
