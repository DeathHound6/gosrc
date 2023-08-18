package gosrc

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func MakeRequest(apiVersion string, endpoint string, method string, headers map[string]string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("https://speedrun.com/api/%s/%s", apiVersion, endpoint)

	req, err := http.NewRequest(method, url, body)
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
