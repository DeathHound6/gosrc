package v2

type Category struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	URL                 string          `json:"url"`
	Pos                 int             `json:"pos"`
	GameID              string          `json:"gameId"`
	IsMisc              bool            `json:"isMisc"`
	IsPerLevel          bool            `json:"isPerLevel"`
	NumPlayers          int             `json:"numPlayers"`
	ExactPlayers        bool            `json:"exactPlayers"`
	PlayerMatchMode     PlayerMatchMode `json:"playerMatchMode"`
	TimeDirection       TimeDirection   `json:"timeDirection"`
	EnforceMilliseconds bool            `json:"enforceMs"`
	Rules               *string         `json:"rules"`
	Archived            *bool           `json:"archived"`
}

type Level struct {
	ID       string  `json:"id"`
	GameID   string  `json:"gameId"`
	Name     string  `json:"name"`
	URL      string  `json:"url"`
	Pos      int     `json:"pos"`
	Rules    *string `json:"rules"`
	Archived bool    `json:"archived"`
}
