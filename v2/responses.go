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

type GetGameSummaryResponse struct {
	Game           Game            `json:"game"`
	GameBoosts     []GameBoost     `json:"gameBoosts"`
	GameModerators []GameModerator `json:"gameModerators"`
	Forum          Forum           `json:"forum"`
	NewsList       []News          `json:"newsList"`
	GameStats      []GameStats     `json:"gameStats"`
	Stats          GameStats       `json:"stats"`
	RelatedGames   []Game          `json:"relatedGames"`
	SeriesList     []Series        `json:"seriesList"`
	Theme          Theme           `json:"theme"`
	ThreadList     []Thread        `json:"threadList"`
	Users          []User          `json:"users"`
	ChallengeList  []Challenge     `json:"challengeList"`
	ChallengeCount int             `json:"challengeCount"`
	GuideCount     int             `json:"guideCount"`
	LevelCount     int             `json:"levelCount"`
	NewsCount      int             `json:"newsCount"`
	RelatedCount   int             `json:"relatedCount"`
	ResourceCount  int             `json:"resourceCount"`
	StreamCount    int             `json:"streamCount"`
	ThreadCount    int             `json:"threadCount"`
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
