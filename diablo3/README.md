
# go-blizzard/diablo3

This is a Go API for gathering Diablo 3 data.

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

### Polling data

Now you can poll data for the Diablo 3 API. For example, you can get all the data for a specific season:

```go
var (
  seasonIndex *SeasonIndex
  season      *Season
  err         error
)

seasonIndex, err = d.GetSeasonIndex()
if err != nil {
  fmt.Println(err)
}


season, err = d.GetSeason(seasonIndex.CurrentSeason)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*season)
```

or the top players for the hardcore Necromancer leaderboard:

```go
var (
  seasonIndex    *SeasonIndex
  seasonRift     *SeasonRift
  err            error
)

seasonIndex, err = d.GetSeasonIndex()
if err != nil {
  fmt.Println(err)
}

seasonRift, err = d.GetSeasonRift(seasonIndex.CurrentSeason, NecromancerPath, true)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*seasonRift)
```

or with better abstraction:

```go
var (
  seasonRift     *SeasonRift
  err            error
)

seasonRift, err = d.GetCurrentSeasonNecromancerRift(true)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*seasonRift)
```

or a player's profile:

```go
var (
  battleTag   = "FuzzyStatic#1384"
  profile     *Profile
  err         error
)

profile, err = d.GetProfile(battleTag)
if err != nil {
  fmt.Println(err)
}

fmt.Println(*profile)
```

TODO
* ~~Add Necromancer AttributesRaw~~
* Verify all class AttributesRaw exist
* Add more tests
* Improved GoDoc documentation
