package v1

import (
	"encoding/json"

	"github.com/DeathHound6/gosrc"
)

type Notification struct {
	ID      string  `json:"id"`
	Created string  `json:"created"`
	Status  string  `json:"status"`
	Text    string  `json:"text"`
	Item    *Link   `json:"item"`
	Links   []*Link `json:"links"`
}

type NotificationsResponse struct {
	Data []*Notification `json:"data"`
}

// This endpoint requires Authentication
func (client *APIClient) GetNotifications() (*NotificationsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"X-API-Key":    client.token,
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, "notifications", gosrc.HTTPMethodGET, headers, nil)
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
	data := new(NotificationsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
