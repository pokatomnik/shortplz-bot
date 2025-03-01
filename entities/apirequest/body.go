package apirequest

type apiRequestBodyJSON struct {
	ArticleUrl string `json:"article_url"`
}

type APIRequestBody struct {
	apiRequestBodyJSON apiRequestBodyJSON
}

func New(articleURL string) APIRequestBody {
	return APIRequestBody{
		apiRequestBodyJSON: apiRequestBodyJSON{
			ArticleUrl: articleURL,
		},
	}
}
