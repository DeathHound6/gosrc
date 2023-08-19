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
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Weblink       string           `json:"weblink"`
	Type          string           `json:"type"`
	Rules         string           `json:"rules"`
	Players       *CategoryPlayers `json:"players"`
	Miscellaneous bool             `json:"miscellaneous"`
	Links         []*Link          `json:"links"`
}

type CategoryResponse struct {
	Data *Category `json:"data"`
}

type CategoryVariableResponse struct {
	Data []*Variable `json:"data"`
}

type CategoryRecordsResponse struct {
	Data []*Leaderboard `json:"data"`
}

func GetCategory(categoryId string) (*CategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("categories/%s", categoryId), http.MethodGet, headers, reqBody)
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
	data := new(CategoryResponse)
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetCategoryVariables(categoryId string) (*CategoryVariableResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("categories/%s/variables", categoryId), http.MethodGet, headers, reqBody)
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
	data := new(CategoryVariableResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetCategoryRecords(categoryId string) (*CategoryRecordsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("categories/%s/variables", categoryId), http.MethodGet, headers, reqBody)
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
	data := new(CategoryRecordsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
