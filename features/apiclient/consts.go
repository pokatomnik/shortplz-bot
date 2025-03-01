package apiclient

const (
	yandexShortAPIHTTPMethod = "POST"
	yandexShortAPIURL        = "https://300.ya.ru/api/sharing-url"
)

const (
	errorParseIncomingURL  = "failed to parse incoming url"
	errorSerializeRequest  = "failed to serialize request"
	errorInitializeRequest = "failed to initialize http request"
	errorCompleteRequest   = "failed to make request"
)
