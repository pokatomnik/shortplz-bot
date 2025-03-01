package summary

import (
	"github.com/samber/mo"
)

func (client SummaryClient) Get(token string, url string) mo.Result[[]string] {
	shortUrl := client.apiClient.GetShortResponseURL(token, url)
	if shortUrl.IsError() {
		return mo.Err[[]string](shortUrl.Error())
	}

	summary := client.shortClient.Get(shortUrl.MustGet())
	if summary.IsError() {
		return mo.Err[[]string](summary.Error())
	}

	return mo.Ok(summary.MustGet())
}
