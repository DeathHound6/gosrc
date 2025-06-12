package v1

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/DeathHound6/gosrc"
)

type CategoryVariablesOrderBy string

const (
	CategoryVariablesOrderByName        CategoryVariablesOrderBy = "name"
	CategoryVariablesOrderByMandatory   CategoryVariablesOrderBy = "mandatory"
	CategoryVariablesOrderByUserDefined CategoryVariablesOrderBy = "user-defined"
	CategoryVariablesOrderByPos         CategoryVariablesOrderBy = "pos"
)

type CategoryPlayers struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Category struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Weblink       string           `json:"weblink"`
	Type          string           `json:"type"`
	Rules         string           `json:"rules"`
	Players       *CategoryPlayers `json:"players"`
	Miscellaneous bool             `json:"miscellaneous"`
	Links         []*Link          `json:"links"`
}

type CategoryResponse struct {
	Data *Category `json:"data"`
}

type CategoryQuery struct {
	Embed *[]Embed
}

func (client *APIClient) GetCategory(categoryId string, queryParams *CategoryQuery) (*CategoryResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	query := map[string]string{}
	if queryParams != nil {
		if queryParams.Embed != nil {
			allowedEmbeds := []Embed{EmbedGame, EmbedVariables}
			embeds := ValidateQueryEmbeds(*queryParams.Embed, allowedEmbeds)
			query["embed"] = strings.Join(embeds, ",")
		}
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("categories/%s?%s", categoryId, gosrc.MakeURLQuery(query)), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(CategoryResponse)
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return nil, err
	}
	return data, nil
}

type CategoryVariableResponse struct {
	Data       []*Variable         `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

type CategoryVariablesQuery struct {
	Embed     *[]Embed
	OrderBy   *CategoryVariablesOrderBy
	Direction *Direction
}

func (client *APIClient) GetCategoryVariables(categoryId string, queryParams *CategoryVariablesQuery) (*CategoryVariableResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	query := map[string]string{}
	if queryParams != nil {
		if queryParams.OrderBy != nil {
			query["orderby"] = string(*queryParams.OrderBy)
		}
		if queryParams.Direction != nil {
			query["direction"] = string(*queryParams.Direction)
		}
		if queryParams.Embed != nil {
			allowedEmbeds := []Embed{EmbedGame, EmbedVariables}
			embeds := ValidateQueryEmbeds(*queryParams.Embed, allowedEmbeds)
			query["embed"] = strings.Join(embeds, ",")
		}
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("categories/%s/variables?%s", categoryId, gosrc.MakeURLQuery(query)), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(CategoryVariableResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

type CategoryRecordsResponse struct {
	Data       []*Leaderboard      `json:"data"`
	Pagination *PaginationResponse `json:"pagination"`
}

type CategoryRecordsQuery struct {
	Embed     *[]Embed
	Top       *int
	SkipEmpty *bool
	Offset    *int
	Max       *int
}

func (client *APIClient) GetCategoryRecords(categoryId string, queryParams *CategoryRecordsQuery) (*CategoryRecordsResponse, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	query := map[string]string{}
	if queryParams != nil {
		if queryParams.Top != nil {
			query["top"] = fmt.Sprintf("%d", *queryParams.Top)
		}
		if queryParams.SkipEmpty != nil {
			query["skip-empty"] = fmt.Sprintf("%t", *queryParams.SkipEmpty)
		}
		if queryParams.Offset != nil {
			query["offset"] = fmt.Sprintf("%d", *queryParams.Offset)
		}
		if queryParams.Max != nil {
			query["max"] = fmt.Sprintf("%d", *queryParams.Max)
		}
		if queryParams.Embed != nil {
			allowedEmbeds := []Embed{EmbedGame, EmbedVariables, EmbedLevel, EmbedPlayers, EmbedRegions, EmbedPlatforms}
			embeds := ValidateQueryEmbeds(*queryParams.Embed, allowedEmbeds)
			query["embed"] = strings.Join(embeds, ",")
		}
	}
	resp, err := gosrc.MakeRequest(gosrc.APIVersionV1, fmt.Sprintf("categories/%s/records?%s", categoryId, gosrc.MakeURLQuery(query)), gosrc.HTTPMethodGET, headers, nil)
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
	data := new(CategoryRecordsResponse)
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}
