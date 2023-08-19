package gosrc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	DeprecatedAPIVersions = []string{}
	logger                = log.Default()
)

func MakeRequest(apiVersion string, endpoint string, method string, headers map[string]string, body io.Reader) (*http.Response, error) {
	for index := range DeprecatedAPIVersions {
		if DeprecatedAPIVersions[index] == apiVersion {
			logger.Print(fmt.Sprintf("WARN: API Version %s is deprecated", apiVersion))
			break
		}
	}

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
