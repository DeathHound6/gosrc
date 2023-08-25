package v1

import (
	"bytes"
	"encoding/json"
	"github.com/DeathHound6/gosrc"
	"net/http"
)

type GameNames struct {
	International string `json:"international"`
	Japanese      string `json:"japanese"`
	Twitch        string `json:"twitch"`
}

type GameRuleset struct {
	ShowMilliseconds    bool     `json:"show-milliseconds"`
	RequireVerification bool     `json:"require-verification"`
	RequireVideo        bool     `json:"require-video"`
	RunTimes            []string `json:"run-times"`
	DefaultTime         string   `json:"default-time"`
	EmulatorsAllowed    bool     `json:"emulators-allowed"`
}

type Game struct {
	ID           string            `json:"id"`
	Names        *GameNames        `json:"names"`
	Abbreviation string            `json:"abbreviation"`
	Weblink      string            `json:"weblink"`
	Released     int               `json:"released"`
	ReleaseDate  string            `json:"release-date"`
	Ruleset      *GameRuleset      `json:"ruleset"`
	Romhack      bool              `json:"romhack"`
	Gametypes    []string          `json:"gametypes"`
	Platforms    []string          `json:"platforms"`
	Regions      []string          `json:"regions"`
	Genres       []string          `json:"genres"`
	Engines      []string          `json:"engines"`
	Developers   []string          `json:"developers"`
	Publishers   []string          `json:"publishers"`
	Moderators   map[string]string `json:"moderators"`
	Created      string            `json:"created"`
	Assets       map[string]*Asset `json:"assets"`
	Links        []*Link           `json:"links"`
}

type GamesResponse struct {
	Data []*Game `json:"data"`
}

func GetGames() (*GamesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "games", http.MethodGet, headers, reqBody)
	if err != nil {
		return nil, err
	}

	bodyBytes := make([]byte, 0)
	if _, err := resp.Body.Read(bodyBytes); err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	data := new(GamesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
