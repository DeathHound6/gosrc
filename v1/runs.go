package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gosrc"
	"net/http"
)

type RunVideosLink struct {
	URI string `json:"uri"`
}

type RunVideos struct {
	Text  string           `json:"text"`
	Links []*RunVideosLink `json:"links"`
}

type RunStatus struct {
	Status     string `json:"status"`
	Examiner   string `json:"examiner"`
	VerifyDate string `json:"verify-date"`
}

type RunPlayers struct {
	Rel  string `json:"rel"`
	Name string `json:"name"`
	ID   string `json:"id"`
	URI  string `json:"uri"`
}

type RunTimes struct {
	Primary          string `json:"primary"`
	PrimaryT         int    `json:"primary_t"`
	Realtime         string `json:"realtime"`
	RealtimeT        int    `json:"realtime_t"`
	RealtimeNoloads  string `json:"realtime_noloads"`
	RealtimeNoloadsT int    `json:"realtime_noloads_t"`
	Ingame           string `json:"ingame"`
	IngameT          int    `json:"ingame_t"`
}

type RunSystem struct {
	Platform string `json:"platform"`
	Emulated bool   `json:"emulated"`
	Region   string `json:"region"`
}

type Run struct {
	ID        string            `json:"id"`
	Weblink   string            `json:"weblink"`
	Game      string            `json:"game"`
	Level     string            `json:"level"`
	Category  string            `json:"category"`
	Videos    *RunVideos        `json:"videos"`
	Comment   string            `json:"comment"`
	Status    *RunStatus        `json:"status"`
	Players   []*RunPlayers     `json:"players"`
	Date      string            `json:"date"`
	Submitted string            `json:"submitted"`
	Times     *RunTimes         `json:"times"`
	System    *RunSystem        `json:"system"`
	Splits    *Link             `json:"splits"`
	Values    map[string]string `json:"values"`
	Links     []*Link           `json:"links"`
}

type PostRunBodyRunTimes struct {
	Realtime        int `json:"realtime"`
	RealtimeNoloads int `json:"realtime_noloads"`
	Ingame          int `json:"ingame"`
}

type PostRunBodyPlayers struct {
	Rel  string `json:"rel"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostRunBodyVariable struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PostRunBodyRun struct {
	Category string `json:"category"`
	Level    string `json:"level"`
	Date     string `json:"date"`
	Region   string `json:"region"`
	Platform string `json:"platform"`
	// This can only be set if the Authenticated user is a Moderator for the game
	Verified  bool                            `json:"verified"`
	Times     *PostRunBodyRunTimes            `json:"times"`
	Players   []*PostRunBodyPlayers           `json:"players"`
	Emulated  bool                            `json:"emulated"`
	Video     string                          `json:"video"`
	Comment   string                          `json:"comment"`
	Splitsio  string                          `json:"splitsio"`
	Variables map[string]*PostRunBodyVariable `json:"variables"`
}

type PostRunBody struct {
	Run *PostRunBodyRun `json:"run"`
}

type PutRunStatusBodyStatus struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type PutRunStatusBody struct {
	Status *PutRunStatusBodyStatus `json:"status"`
}

type PutRunPlayersBodyPlayer struct {
	Rel  string `json:"rel"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PutRunPlayersBody struct {
	Players []*PutRunPlayersBodyPlayer `json:"players"`
}

type RunsResponse struct {
	Data []*Run `json:"data"`
}

type RunResponse struct {
	Data *Run `json:"data"`
}

type PostRunErrorResponse struct {
	error
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
	Links   []*Link  `json:"links"`
}

func GetRuns() (*RunsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("runs%s", ""), http.MethodGet, headers, reqBody)
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
	data := new(RunsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetRun(runId string) (*RunResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("runs/%s", runId), http.MethodGet, headers, reqBody)
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
	data := new(RunResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// This endpoint requires Authentication
func PostRun(body *PostRunBody) (*RunResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "runs", http.MethodPost, headers, reqBody)
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
	if resp.StatusCode == 201 {
		data := new(RunResponse)
		if err := json.Unmarshal(bodyBytes, &data); err != nil {
			return nil, err
		}
		return data, nil
	} else {
		data := new(PostRunErrorResponse)
		if err := json.Unmarshal(bodyBytes, &data); err != nil {
			return nil, err
		}
		data.error = errors.New(data.Message)
		return nil, data
	}
}

// This endpoint requires Authentication
func PutRunStatus(runId string, body *PutRunStatusBody) (*RunResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("runs/%s/status", runId), http.MethodPut, headers, reqBody)
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
	data := new(RunResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// This endpoint requires Authentication
func PutRunPlayers(runId string, body *PutRunPlayersBody) (*RunResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("runs/%s/players", runId), http.MethodPut, headers, reqBody)
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
	data := new(RunResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
