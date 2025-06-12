package v1

import (
	"encoding/json"
	"fmt"

	"github.com/DeathHound6/gosrc"
)

type Engine struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Links []*Link `json:"links"`
}

type EnginesResponse struct {
	Data []*Engine `json:"data"`
}

type EngineResponse struct {
	Data *Engine `json:"data"`
}

func (client *APIClient) GetEngines() (*EnginesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, "engines", gosrc.HTTPMethodGET, headers, nil)
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
	data := new(EnginesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (client *APIClient) GetEngine(engineId string) (*EngineResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("engines/%s", engineId), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(EngineResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
