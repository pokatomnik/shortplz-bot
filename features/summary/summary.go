package summary

import (
	"github.com/pokatomnik/shortplz-bot/features/apiclient"
	"github.com/pokatomnik/shortplz-bot/features/shortclient"
)

type SummaryClient struct {
	apiClient   apiclient.YandexAPIClient
	shortClient shortclient.HTMLClient
}

func New() SummaryClient {
	return SummaryClient{
		apiClient:   apiclient.New(),
		shortClient: shortclient.New(),
	}
}
