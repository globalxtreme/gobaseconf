package privateapi

import (
	"errors"
	xtremeapi "github.com/globalxtreme/gobaseconf/api"
	"github.com/globalxtreme/gobaseconf/response"
	"os"
)

type BusinessAPI interface {
	NotificationPush(payload interface{}) response.ResponseSuccessWithPagination
}

func NewBusinessAPI() (BusinessAPI, error) {
	host := os.Getenv("CLIENT_PRIVATE_API_ASA_HOST")
	clientId := os.Getenv("CLIENT_PRIVATE_API_ASA_ID")
	clientName := os.Getenv("CLIENT_PRIVATE_API_ASA_NAME")
	clientSecret := os.Getenv("CLIENT_PRIVATE_API_ASA_SECRET")

	if host == "" || clientId == "" || clientName == "" || clientSecret == "" {
		return nil, errors.New("Please set private api ASA environment variables")
	}

	client := xtremeapi.NewXtremeAPI(xtremeapi.XtremeAPIOption{
		Headers: map[string]string{
			"Client-ID":     clientId,
			"Client-Name":   clientName,
			"Client-Secret": clientSecret,
		},
	})

	api := businessAPI{
		baseURL: host,
		client:  client,
	}

	return &api, nil
}

type businessAPI struct {
	baseURL string
	client  xtremeapi.XtremeAPI
}

func (api *businessAPI) NotificationPush(payload interface{}) response.ResponseSuccessWithPagination {
	return api.client.Post(api.baseURL+"/notifications", payload)
}
