package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gosrc"
	"net/http"
)

type VariableScope struct {
	Type string `json:"type"`
}

type VariableValuesValueFlags struct {
	Miscellaneous bool `json:"miscellaneous"`
}

type VariableValuesValue struct {
	Label string                    `json:"label"`
	Rules string                    `json:"rules"`
	Flags *VariableValuesValueFlags `json:"flags"`
}

type VariableValues struct {
	Values  map[string]*VariableValuesValue `json:"values"`
	Default string                          `json:"default"`
}

type Variable struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	Scope         *VariableScope  `json:"scope"`
	Mandatory     bool            `json:"mandatory"`
	UserDefined   bool            `json:"user-defined"`
	Obsoletes     bool            `json:"obsoletes"`
	Values        *VariableValues `json:"values"`
	IsSubcategory bool            `json:"is-subcategory"`
	Links         []*Link         `json:"links"`
}

type VariableResponse struct {
	Data *Variable `json:"data"`
}

func GetVariable(variableId string) (*VariableResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	jsonBody, err := json.Marshal(map[string]string{})
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewBuffer(jsonBody)
	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("variables/%s", variableId), http.MethodGet, headers, reqBody)
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
	data := new(VariableResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
