package v1

type APIClient struct {
	token string
}

func NewAPIClient(token string) *APIClient {
	return &APIClient{
		token: token,
	}
}
