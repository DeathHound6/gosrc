package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeathHound6/gosrc"
	"net/http"
)

type LeaderboardRun struct {
	Place int  `json:"place"`
	Run   *Run `json:"run"`
}

type Leaderboard struct {
	Weblink   string            `json:"weblink"`
	Game      string            `json:"game"`
	Category  string            `json:"category"`
	Level     string            `json:"level"`
	Platform  string            `json:"platform"`
	Region    string            `json:"region"`
	Emulators bool              `json:"emulators"`
	VideoOnly bool              `json:"video-only"`
	Timing    string            `json:"timing"`
	Values    map[string]string `json:"values"`
	Runs      []*LeaderboardRun `json:"runs"`
	Links     []*Link           `json:"links"`
}

type LeaderboardCategoryResponse struct {
	Data *Leaderboard `json:"data"`
}

func GetLeaderboardCategory(gameId string, categoryId string) (*LeaderboardCategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("leaderboards/%s/categories/%s", gameId, categoryId), http.MethodGet, headers, reqBody)
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
	data := new(LeaderboardCategoryResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetLeaderboardLevelCategory(gameId string, levelId string, categoryId string) (*LeaderboardCategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("leaderboards/%s/level/%s/%s", gameId, levelId, categoryId), http.MethodGet, headers, reqBody)
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
	data := new(LeaderboardCategoryResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
