package v2

type Pagination struct {
	Count   uint `json:"count"`
	Page    uint `json:"page"`
	Pages   uint `json:"pages"`
	PerPage uint `json:"per"`
}
