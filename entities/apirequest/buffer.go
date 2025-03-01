package apirequest

import (
	"bytes"
	"encoding/json"

	"github.com/samber/mo"
)

func (apiRequestBody APIRequestBody) GetBuffer() mo.Result[*bytes.Buffer] {
	jsonData, err := json.Marshal(apiRequestBody.apiRequestBodyJSON)
	if err != nil {
		return mo.Err[*bytes.Buffer](err)
	}
	buf := bytes.NewBuffer(jsonData)
	return mo.Ok(buf)
}
