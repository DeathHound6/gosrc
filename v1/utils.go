package v1

const (
	APIVersion = "v1"
)

type Link struct {
	Rel string `json:"rel"`
	URI string `json:"uri"`
}
