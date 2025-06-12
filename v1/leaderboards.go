package v1

import (
	"encoding/json"
	"fmt"

	"github.com/DeathHound6/gosrc"
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

func (client *APIClient) GetLeaderboardCategory(gameId string, categoryId string) (*LeaderboardCategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("leaderboards/%s/categories/%s", gameId, categoryId), gosrc.HTTPMethodGET, headers, nil)
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

func (client *APIClient) GetLeaderboardLevelCategory(gameId string, levelId string, categoryId string) (*LeaderboardCategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("leaderboards/%s/level/%s/%s", gameId, levelId, categoryId), gosrc.HTTPMethodGET, headers, nil)
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
