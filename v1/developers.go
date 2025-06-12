package v1

import (
	"encoding/json"
	"fmt"

	"github.com/DeathHound6/gosrc"
)

type Developer struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Links []*Link `json:"links"`
}

type DevelopersResponse struct {
	Data []*Developer `json:"data"`
}

type DeveloperResponse struct {
	Data *Developer `json:"data"`
}

func (client *APIClient) GetDevelopers() (*DevelopersResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, "developers", gosrc.HTTPMethodGET, headers, nil)
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

func (client *APIClient) GetDeveloper(developerId string) (*DeveloperResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("developers/%s", developerId), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(DeveloperResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
