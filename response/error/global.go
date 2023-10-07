package error

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

func ErrXtremeUnauthenticated(internalMsg string) {
	response.Error(http.StatusUnauthorized, "Unauthenticated.", internalMsg, nil)
}

func ErrXtremeBadRequest(internalMsg string) {
	response.Error(http.StatusBadRequest, "Bad request!", internalMsg, nil)
}

func ErrXtremePayloadVeryLarge(internalMsg string) {
	response.Error(http.StatusRequestEntityTooLarge, "Your payload very large!", internalMsg, nil)
}

func ErrXtremeValidation(attributes []interface{}) {
	response.Error(http.StatusBadRequest, "Missing Required Parameter", "", attributes)
}

func ErrXtremeNotFound(internalMsg string) {
	response.Error(http.StatusNotFound, "Data not found", internalMsg, nil)
}
