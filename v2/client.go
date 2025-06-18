package v2

type APIClient struct {
	phpsessid string
	csrfToken string
}

func NewAPIClient(phpSessionID string) *APIClient {
	return &APIClient{
		phpsessid: phpSessionID,
		csrfToken: "",
	}
}
