package v2

type GetGameDataResponse struct {
	Game       Game            `json:"game"`
	Categories []Category      `json:"categories"`
	Levels     []Level         `json:"levels"`
	Moderators []GameModerator `json:"moderators"`
	Platforms  []Platform      `json:"platforms"`
	Regions    []Region        `json:"regions"`
	Theme      *Theme          `json:"theme"`
	Users      []User          `json:"users"`
	Values     []Value         `json:"values"`
	Variables  []Variable      `json:"variables"`
}

type GetSessionResponse struct {
	Session *Session `json:"session"`
}

type GetArticleResponse struct {
	Article         *Article   `json:"article"`
	RelatedArticles []*Article `json:"relatedArticleList"`
	GameList        []string   `json:"gameList"`
	UserList        []string   `json:"userList"`
}

type GetArticleListResponse struct {
	Articles   []*Article  `json:"articleList"`
	Pagination *Pagination `json:"pagination"`
	GameList   []string    `json:"gameList"`
	UserList   []string    `json:"userList"`
}
