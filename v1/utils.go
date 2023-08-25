package v1

const (
	APIVersion = "v1"
)

type Link struct {
	Rel string `json:"rel"`
	URI string `json:"uri"`
}

type Asset struct {
	URI    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
