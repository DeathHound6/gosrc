package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type GameResponse struct {
	Data *Game `json:"data"`
}

type GameCategoriesResponse struct {
	Data []*Category `json:"data"`
}

type GameLevelsResponse struct {
	Data []*Level `json:"data"`
}

type GameVariablesResponse struct {
	Data []*Variable `json:"data"`
}

type GameRecordsResponse struct {
	Data []*Leaderboard `json:"data"`
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

func GetGame(gameId string) (*GameResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s", gameId), http.MethodGet, headers, reqBody)
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
	data := new(GameResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetGameCategories(gameId string) (*GameCategoriesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s/categories", gameId), http.MethodGet, headers, reqBody)
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
	data := new(GameCategoriesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetGameLevels(gameId string) (*GameLevelsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s/levels", gameId), http.MethodGet, headers, reqBody)
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
	data := new(GameLevelsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetGameVariables(gameId string) (*GameVariablesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s/variables", gameId), http.MethodGet, headers, reqBody)
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
	data := new(GameVariablesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetGameDerived(gameId string) (*GamesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s/derived-games", gameId), http.MethodGet, headers, reqBody)
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

func GetGameRecords(gameId string) (*GameRecordsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("games/%s/records", gameId), http.MethodGet, headers, reqBody)
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
	data := new(GameRecordsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
