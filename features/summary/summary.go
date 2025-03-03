package summary

import "github.com/samber/mo"

type ApiClient interface {
	GetShortResponseURL(token string, articleURL string) mo.Result[string]
}

type ShortClient interface {
	Get(url string) mo.Result[[]string]
}

type SummaryClient struct {
	apiClient   ApiClient
	shortClient ShortClient
}

func New(apiClient ApiClient, shortClient ShortClient) SummaryClient {
	return SummaryClient{
		apiClient:   apiClient,
		shortClient: shortClient,
	}
}
