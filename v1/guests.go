package v1

type Guest struct {
	Name  string  `json:"name"`
	Links []*Link `json:"links"`
}
