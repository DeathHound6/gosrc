package gosrc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type APIVersion string
type HTTPMethod string

const (
	HTTPMethodGET    HTTPMethod = "GET"
	HTTPMethodPOST   HTTPMethod = "POST"
	HTTPMethodPUT    HTTPMethod = "PUT"
	HTTPMethodDELETE HTTPMethod = "DELETE"

	APIVersionV1 APIVersion = "v1"
	APIVersionV2 APIVersion = "v2"
)

var (
	DeprecatedAPIVersions = []APIVersion{APIVersionV1}
	logger                = log.Default()
)

func MakeRequest(apiVersion APIVersion, endpoint string, method HTTPMethod, headers map[string]string, body io.Reader) (*http.Response, error) {
	for index := range DeprecatedAPIVersions {
		if DeprecatedAPIVersions[index] == apiVersion {
			logger.Printf("WARN: API Version %s is deprecated", apiVersion)
			break
		}
	}

	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	url := fmt.Sprintf("https://speedrun.com/api/%s/%s", apiVersion, endpoint)

	req, err := http.NewRequest(string(method), url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := http.Client{
		Transport: http.DefaultTransport,
		Timeout:   time.Second * 30,
	}
	return client.Do(req)
}

func MakeURLQuery(query map[string]string) string {
	if len(query) > 0 {
		q := make([]string, 0)
		for key, value := range query {
			q = append(q, fmt.Sprintf("%s=%s", key, value))
		}
		return fmt.Sprintf("?%s", strings.Join(q, "&"))
	}
	return ""
}

func SliceContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ReadBody[T any](body io.ReadCloser) (*T, error) {
	bodyBytes := make([]byte, 0)
	if _, err := body.Read(bodyBytes); err != nil {
		return nil, err
	}
	if err := body.Close(); err != nil {
		return nil, err
	}
	data := new(T)
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}
