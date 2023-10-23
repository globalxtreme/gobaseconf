package controller

import (
	"github.com/globalxtreme/gobaseconf/filesystem"
	"github.com/gorilla/mux"
	"net/http"
)

type BaseStorageController struct{}

func (ctr BaseStorageController) ShowFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	storage := filesystem.Storage{IsPublic: true}
	storage.ShowFile(w, r, vars["path"])
}
