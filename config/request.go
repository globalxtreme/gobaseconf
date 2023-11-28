package config

import (
	"encoding/json"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

var (
	RequestBody map[string]interface{}
)

func SetRequestBody(r *http.Request) {
	RequestBody = make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&RequestBody)
	if err != nil {
		error.ErrXtremeBadRequest(err.Error())
	}
}
