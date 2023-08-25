package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeathHound6/gosrc"
	"net/http"
)

type Level struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Weblink string  `json:"weblink"`
	Rules   string  `json:"rules"`
	Links   []*Link `json:"links"`
}

type LevelResponse struct {
	Data *Level `json:"data"`
}

func GetLevel(levelId string) (*LevelResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("levels/%s", levelId), http.MethodGet, headers, reqBody)
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
	data := new(LevelResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
