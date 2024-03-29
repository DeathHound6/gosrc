package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeathHound6/gosrc"
	"net/http"
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

type SeriessResponse struct {
	Data []*Series `json:"data"`
}

type SeriesResponse struct {
	Data *Series `json:"data"`
}

type SeriesGamesResponse struct {
	Data []*Game `json:"data"`
}

func GetSeriess() (*SeriessResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "series", http.MethodGet, headers, reqBody)
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
	data := new(SeriessResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetSeries(seriesId string) (*SeriesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("series/%s", seriesId), http.MethodGet, headers, reqBody)
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

func GetSeriesGames(seriesId string) (*SeriesGamesResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("series/%s/games", seriesId), http.MethodGet, headers, reqBody)
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
