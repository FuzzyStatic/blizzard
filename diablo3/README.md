
# blizzard/diablo3

This is a Go client library for gathering Blizzard Diablo 3 API data.

### Getting started

Start by initiating a new Diablo3 structure (accessToken and apiKey can be acquired through your Mashery account at https://dev.battle.net/):

```go
var d *Diablo3

d = New(
	blizzard.Auth{
		AccessToken: "myAccessToken",
		APIKey:      "myApiKey",
	},
	blizzard.US,
)
```

### Fetching data

Now you can fetch data from the Diablo 3 API. For example, you can get all the data for a specific season:

```go
seasonIndex, err := d.GetSeasonIndex()
if err != nil {
  fmt.Println(err)
}


season, err := d.GetSeason(seasonIndex.CurrentSeason)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*season)
```

or the top players for the hardcore Necromancer leaderboard:

```go
seasonRift, err := d.GetCurrentSeasonNecromancerRift(true)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*seasonRift)
```

or a player's profile:

```go
battleTag := "FuzzyStatic#1384"

profile, err := d.GetProfile(battleTag)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*profile)
```

TODO
* Verify all class AttributesRaw exist
* Add more tests
