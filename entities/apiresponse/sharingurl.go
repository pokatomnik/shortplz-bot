package apiresponse

import (
	"errors"

	"github.com/samber/mo"
)

func (apiResponse APIResponse) GetSharingURL() mo.Result[string] {
	if apiResponse.apiResponseJSON.Status == successStatus {
		return mo.Ok(apiResponse.apiResponseJSON.SharingURL)
	}
	return mo.Err[string](errors.New(errorGetSharingURL))
}
