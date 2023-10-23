package controller

import (
	"github.com/globalxtreme/gobaseconf/filesystem"
	"net/http"
)

type BaseStorageController struct{}

func (ctr BaseStorageController) ShowFile(w http.ResponseWriter, r *http.Request) {
	storage := filesystem.Storage{}
	storage.ShowPublicFile(w, r)
}
