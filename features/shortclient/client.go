package shortclient

import "net/http"

type HTMLClient struct {
	httpClient   *http.Client
	replacements map[string]string
}

func New() HTMLClient {
	replacements := map[string]string{
		"•": "",
	}
	return HTMLClient{
		httpClient:   &http.Client{},
		replacements: replacements,
	}
}
