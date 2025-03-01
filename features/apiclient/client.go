package apiclient

import (
	"net/http"
)

type YandexAPIClient struct {
	client *http.Client
}

func New() YandexAPIClient {
	return YandexAPIClient{
		client: &http.Client{},
	}
}
