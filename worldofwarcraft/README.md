# go-blizzard/worldofwarcraft

This is a Go client library for gathering Blizzard World of Warcraft API data.

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

### Fetching data

Now you can fetch data from the World of Warcraft API. For example, you can get all the data for a character:

```go
character, err := w.GetCharacter("illidan", "flowbs")
if err != nil {
	t.Fail()
}

fmt.Println(*character)
```
