package router

import (
	"github.com/globalxtreme/gobaseconf/controller"
	"github.com/globalxtreme/gobaseconf/middleware"
	"github.com/gorilla/mux"
)

type CallbackRouter func(*mux.Router)

func RegisterRouter(router *mux.Router, callback CallbackRouter) {
	router.Use(middleware.PanicHandler)
	router.Use(middleware.PrepareRequestHandler)

	// Storage route
	stController := controller.BaseStorageController{}
	router.HandleFunc("/storages/{path:.*}", stController.ShowFile).Methods("GET")

	callback(router)
}
