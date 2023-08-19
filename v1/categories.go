package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gosrc"
	"net/http"
)

type CategoryPlayers struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Category struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Weblink       string          `json:"weblink"`
	Type          string          `json:"type"`
	Rules         string          `json:"rules"`
	Players       CategoryPlayers `json:"players"`
	Miscellaneous bool            `json:"miscellaneous"`
	Links         []Link          `json:"links"`
}

type CategoryResponse struct {
	Data Category `json:"data"`
}

func GetCategory(id string) (*CategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("categories/%s", id), http.MethodGet, headers, reqBody)
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
	data := &CategoryResponse{}
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetCategoryVariable(id string) []*Variable {

}

func GetCategoryRecords(id string) []*Leaderboard {

}
