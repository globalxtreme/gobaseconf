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

func ErrXtremeRabbitMQMessageGet(internalMsg string) {
	response.Error(http.StatusNotFound, "RabbitMQ message not found", internalMsg, nil)
}

func ErrXtremeRabbitMQMessageDeliveryGet(internalMsg string) {
	response.Error(http.StatusNotFound, "RabbitMQ message delivery not found", internalMsg, nil)
}

func ErrXtremeRabbitMQMessageDeliveryValidation(internalMsg string) {
	response.Error(http.StatusBadRequest, "RabbitMQ message delivery form invalid", internalMsg, nil)
}

func ErrXtremeAPI(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Calling external api is invalid!", internalMsg, nil)
}
