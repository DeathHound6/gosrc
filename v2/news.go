package v2

type News struct {
	ID            string  `json:"id"`
	GameID        string  `json:"gameId"`
	UserID        string  `json:"userId"`
	Title         string  `json:"title"`
	Body          *string `json:"body"`
	DateSubmitted int     `json:"dateSubmitted"`
}
