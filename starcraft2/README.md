# go-blizzard/worldofwarcraft

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

Now you can fetch data from the Starcraft 2 API. 

TODO
* Add all calls
