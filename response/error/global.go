package error

import (
	"github.com/globalxtreme/gomodule-baseconf/response"
	"net/http"
)

func ErrUnauthenticated(internalMsg string) {
	response.Error(http.StatusUnauthorized, "Unauthenticated.", internalMsg, nil)
}

func ErrBadRequest(internalMsg string) {
	response.Error(http.StatusBadRequest, "Bad request!", internalMsg, nil)
}

func ErrPayloadVeryLarge(internalMsg string) {
	response.Error(http.StatusRequestEntityTooLarge, "Your payload very large!", internalMsg, nil)
}

func ErrValidation(attributes []interface{}) {
	response.Error(http.StatusBadRequest, "Missing Required Parameter", "", attributes)
}

func ErrNotFound(internalMsg string) {
	response.Error(http.StatusNotFound, "Data not found", internalMsg, nil)
}
