package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DeathHound6/gosrc"
	"net/http"
)

type UserNames struct {
	International string `json:"international"`
	Japanese      string `json:"japanese"`
}

type UserNameStyleColor struct {
	Light string `json:"light"`
	Dark  string `json:"dark"`
}

type UserNameStyle struct {
	Style string              `json:"style"`
	Color *UserNameStyleColor `json:"color"`
}

type UserLocationLocation struct {
	Code  string     `json:"code"`
	Names *UserNames `json:"names"`
}

type UserLocation struct {
	Country *UserLocationLocation `json:"country"`
	Region  *UserLocationLocation `json:"region"`
}

type UserConnection struct {
	URI string `json:"uri"`
}

type User struct {
	ID            string          `json:"id"`
	Names         *UserNames      `json:"names"`
	Weblink       string          `json:"weblink"`
	NameStyle     *UserNameStyle  `json:"name-style"`
	Role          string          `json:"role"`
	Signup        string          `json:"signup"`
	Location      *UserLocation   `json:"location"`
	Twitch        *UserConnection `json:"twitch"`
	Hitbox        *UserConnection `json:"hitbox"`
	Youtube       *UserConnection `json:"youtube"`
	Twitter       *UserConnection `json:"twitter"`
	Speedrunslive *UserConnection `json:"speedrunslive"`
	Links         []*Link         `json:"links"`
}

type UsersResponse struct {
	Data []*User `json:"data"`
}

type UserResponse struct {
	Data *User `json:"data"`
}

func GetUsers() (*UsersResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "users", http.MethodGet, headers, reqBody)
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
	data := new(UsersResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetUser(userId string) (*UserResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("users/%s", userId), http.MethodGet, headers, reqBody)
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
	data := new(UserResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// This endpoint requires Authentication
func GetProfile() (*UserResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, "profile", http.MethodGet, headers, reqBody)
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
	data := new(UserResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
