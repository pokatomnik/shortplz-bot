package apiclient

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/pokatomnik/shortplz-bot/entities/apirequest"
	"github.com/pokatomnik/shortplz-bot/entities/apiresponse"
	"github.com/samber/mo"
)

func (client *YandexAPIClient) GetShortResponseURL(token string, articleURL string) mo.Result[string] {
	_, err := url.Parse(articleURL)
	if err != nil {
		return mo.Err[string](errors.New(errorParseIncomingURL))
	}

	requestBody := apirequest.New(articleURL)
	requestBodyBuffer := requestBody.GetBuffer()

	if requestBodyBuffer.IsError() {
		return mo.Err[string](errors.New(errorSerializeRequest))
	}

	request, err := http.NewRequest(
		yandexShortAPIHTTPMethod,
		yandexShortAPIURL,
		requestBodyBuffer.MustGet(),
	)

	if err != nil {
		return mo.Err[string](errors.New(errorInitializeRequest))
	}

	headers := client.getHeaders(token)
	for hKey, hVal := range headers {
		request.Header.Add(hKey, hVal)
	}

	response, err := client.client.Do(request)
	if err != nil {
		return mo.Err[string](errors.New(errorCompleteRequest))
	}
	defer response.Body.Close()

	apiResponse := apiresponse.FromReader(response.Body)

	if apiResponse.IsError() {
		return mo.Err[string](apiResponse.Error())
	}

	return apiResponse.MustGet().GetSharingURL()
}
