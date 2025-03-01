package apiresponse

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/samber/mo"
)

const successStatus = "success"

type apiResponseJSON struct {
	Status     string `json:"status"`
	SharingURL string `json:"sharing_url"`
}

type APIResponse struct {
	apiResponseJSON apiResponseJSON
}

func FromReader(reader io.Reader) mo.Result[APIResponse] {
	response := APIResponse{
		apiResponseJSON: apiResponseJSON{},
	}
	err := json.NewDecoder(reader).Decode(&response.apiResponseJSON)
	if err != nil {
		return mo.Err[APIResponse](errors.New(errorParseResponse))
	}
	return mo.Ok(response)
}
