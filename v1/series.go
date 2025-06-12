package v1

import (
	"encoding/json"
	"fmt"

	"github.com/DeathHound6/gosrc"
)

type SeriesNames struct {
	International string `json:"international"`
	Japanese      string `json:"japanese"`
}

type Series struct {
	ID           string            `json:"id"`
	Names        *SeriesNames      `json:"names"`
	Abbreviation string            `json:"abbreviation"`
	Weblink      string            `json:"weblink"`
	Moderators   map[string]string `json:"moderators"`
	Created      string            `json:"created"`
	Assets       map[string]*Asset `json:"assets"`
	Links        []*Link           `json:"links"`
}

type SeriesesResponse struct {
	Data []*Series `json:"data"`
}

type SeriesResponse struct {
	Data *Series `json:"data"`
}

type SeriesGamesResponse struct {
	Data []*Game `json:"data"`
}

func (client *APIClient) GetSerieses() (*SeriesesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, "series", gosrc.HTTPMethodGET, headers, nil)
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
	data := new(SeriesesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (client *APIClient) GetSeries(seriesId string) (*SeriesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("series/%s", seriesId), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(SeriesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (client *APIClient) GetSeriesGames(seriesId string) (*SeriesGamesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("series/%s/games", seriesId), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(SeriesGamesResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
