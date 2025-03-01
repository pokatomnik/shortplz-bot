package apirequest

func (apiRequestBody APIRequestBody) GetArticleUrl() string {
	return apiRequestBody.apiRequestBodyJSON.ArticleUrl
}
