# go-blizzard/worldofwarcraft

This is a Go API for gathering World of Warcraft data.

### Getting started

Start by initiating a new WorldOfWarcraft structure (accessToken and apiKey can be acquired through your Mashery account at https://dev.battle.net/):

```go
var w *WorldOfWarcraft

w = New(
	blizzard.Auth{
		AccessToken: "myAccessToken",
		APIKey:      "myApiKey",
	},
	blizzard.US,
)
```

### Polling data

Now you can poll data for the World of Warcraft API. 

TODO
* Add all calls
* Improved GoDoc documentation
