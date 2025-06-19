package v2

type Challenge struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	Announcement        string               `json:"announcement"`
	URL                 string               `json:"url"`
	GameID              string               `json:"gameId"`
	CreateDate          int                  `json:"createDate"`
	UpdateDate          int                  `json:"updateDate"`
	StartDate           int                  `json:"startDate"`
	EndDate             int                  `json:"endDate"`
	State               ChallengeState       `json:"state"`
	Description         string               `json:"description"`
	Rules               string               `json:"rules"`
	NumPlayers          int                  `json:"numPlayers"`
	ExactPlayers        bool                 `json:"exactPlayers"`
	PlayerMatchMode     PlayerMatchMode      `json:"playerMatchMode"`
	TimeDirection       TimeDirection        `json:"timeDirection"`
	EnforceMilliseconds bool                 `json:"enforceMs"`
	CoverImagePath      string               `json:"coverImagePath"`
	Contest             bool                 `json:"contest"`
	ContestRules        string               `json:"contestRules"`
	RunCommentsMode     PermissionType       `json:"runCommentsMode"`
	PrizeConfig         ChallengePrizeConfig `json:"prizeConfig"`
}

type ChallengePrizeConfig struct {
	PrizePool int              `json:"prizePool"`
	Currency  string           `json:"currency"`
	Prizes    []ChallengePrize `json:"prizes"`
}

type ChallengePrize struct {
	Place  int `json:"place"`
	Amount int `json:"amount"`
}

type ChallengeModerator struct {
	ChallengeID string             `json:"challengeId"`
	UserID      string             `json:"userId"`
	Level       GameModeratorLevel `json:"level"`
}
