package apiclient

import "fmt"

const (
	keyAuthorization = "Authorization"
	keyContentType   = "Content-Type"
)

const (
	valueContentType = "application/json"
)

func (client YandexAPIClient) getHeaders(token string) map[string]string {
	headers := make(map[string]string)
	headers[keyAuthorization] = fmt.Sprintf("OAuth %s", token)
	headers[keyContentType] = valueContentType
	return headers
}
