# GoSRC
A Go library for interacting with the Speedrun.com API

## Install
```sh
$ go get github.com/DeathHound6/gosrc
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/DeathHound6/gosrc/v1"
    "github.com/DeathHound6/gosrc/v2"
)

func main() {
    V1()
    V2()
}

func V1() {
    // Speedrun.com API v1
    client := v1.NewAPIClient("your speedrun.com token")
    game, err := client.GetGame("game id")
    // error handling here
    fmt.Printf("Game URL is %s\n", game.Data.Weblink)
}

func V2() {
    // Speedrun.com API v2
    client := v2.NewAPIClient("your php session id", "your csrf token")
    params := GetGameDataFilters{GameID: "game id"}
    game, err := client.GetGameData(params)
    // error handling here
    fmt.Printf("Game URL is %s", game.Game.URL)
}
```

## API Documentation
- [API v1](https://github.com/speedruncomorg/api)
- [API v2](https://github.com/ManicJamie/speedruncompy)
