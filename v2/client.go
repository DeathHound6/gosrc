package v2

import (
	"reflect"
)

type APIClient struct {
	phpsessid *string
	csrfToken *string
}

func NewAPIClient(phpSessionID *string, csrfToken *string) *APIClient {
	return &APIClient{
		phpsessid: phpSessionID,
		csrfToken: csrfToken,
	}
}

func (client *APIClient) ApplyAuthorization(headers map[string]string, body any) {
	if client.phpsessid != nil {
		headers["Cookie"] = "PHPSESSID=" + *client.phpsessid
		if client.csrfToken != nil && body != nil {
			// body may be any custom struct, so we need to check if it has a csrfToken field
			v := reflect.ValueOf(body)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				for i := 0; i < v.NumField(); i++ {
					if tag, ok := v.Type().Field(i).Tag.Lookup("json"); ok && tag == "csrfToken" {
						if v.Field(i).CanSet() && v.Field(i).Kind() == reflect.String {
							v.Field(i).SetString(*client.csrfToken)
						}
						break
					}
				}
			}
		}
	}
}
