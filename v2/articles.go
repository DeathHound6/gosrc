package v2

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DeathHound6/gosrc"
)

type Article struct {
	ID            int      `json:"id"`
	Slug          string   `json:"slug"`
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	Body          string   `json:"body"`
	CreateDate    string   `json:"createDate"`
	UpdateDate    string   `json:"updateDate"`
	PublishDate   string   `json:"publishDate"`
	PublishTarget string   `json:"publishTarget"`
	PublishTags   []string `json:"publishTags"`
	CommentsCount int      `json:"commentsCount"`
}

type GetArticleResponse struct {
	Article         *Article   `json:"article"`
	RelatedArticles []*Article `json:"relatedArticleList"`
	GameList        []string   `json:"gameList"`
	UserList        []string   `json:"userList"`
}

func GetArticle(id *int, slug *string) (*GetArticleResponse, error) {
	if id == nil && slug == nil {
		return nil, errors.New("either id or slug must be provided")
	}

	query := ""
	if id != nil {
		query = fmt.Sprintf("int=%d", *id)
	} else if slug != nil {
		query = fmt.Sprintf("slug=%s", *slug)
	}

	headers := map[string]string{
		"Accept": "application/json",
	}

	resp, err := gosrc.MakeRequest(APIVersion, fmt.Sprintf("GetArticle?%s", query), "GET", headers, nil)
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
	data := new(GetArticleResponse)
	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return nil, err
	}

	return data, nil
}
