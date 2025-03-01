package shortclient

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/samber/mo"
)

func (client HTMLClient) Get(url string) mo.Result[[]string] {
	response, err := client.httpClient.Get(url)
	if err != nil {
		return mo.Err[[]string](errors.New(errorNetworkGet))
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return mo.Err[[]string](errors.New(errorBuildDocument))
	}

	content := getShortContent(document)

	if content.IsError() {
		return mo.Err[[]string](content.Error())
	}

	lines := client.parse(content.MustGet())

	if len(lines) == 0 {
		return mo.Err[[]string](errors.New(errorEmptySummarization))
	}

	return mo.Ok(lines)
}
