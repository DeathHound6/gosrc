package v1

type CategoryPlayers struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Category struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Weblink       string          `json:"weblink"`
	Type          string          `json:"type"`
	Rules         string          `json:"rules"`
	Players       CategoryPlayers `json:"players"`
	Miscellaneous bool            `json:"miscellaneous"`
	Links         []Link          `json:"links"`
}
