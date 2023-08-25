package v1

import (
	"bytes"
	"encoding/json"
	"github.com/DeathHound6/gosrc"
	"net/http"
)

type Developer struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Links []*Link `json:"links"`
}

type DevelopersResponse struct {
	Data []*Developer `json:"data"`
}

func GetDevelopers() (*DevelopersResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "developers", http.MethodGet, headers, reqBody)
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
	data := new(DevelopersResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
