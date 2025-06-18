package v2

type GameRunner struct {
	GameID   string `json:"gameId"`
	UserID   string `json:"userId"`
	RunCount int    `json:"runCount"`
}
