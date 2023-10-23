package filesystem

import (
	"github.com/globalxtreme/gobaseconf/helpers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Storage struct{}

func (repo Storage) GetFullPath(path string) string {
	baseDir, _ := os.Getwd()
	storageDir := helpers.SetStorageDir(path)

	return baseDir + "/" + storageDir
}

func (repo Storage) GetFullPathURL(path string) string {
	return os.Getenv("API_GATEWAY_LINK_URL") + path
}

func (repo Storage) ShowFile(w http.ResponseWriter, path string) {
	var request http.Request

	realPath := checkAndSetDefaultFile(path)
	if realPath == nil {
		http.NotFound(w, &request)
		return
	}

	http.ServeFile(w, &request, realPath.(string))
}

func (repo Storage) ShowPublicFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	realPath := checkAndSetDefaultFile("public/" + vars["path"])
	if realPath == nil {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, realPath.(string))
}

func checkAndSetDefaultFile(path string) any {
	path = helpers.SetStorageAppDir(path)

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
	}

	if info.IsDir() {
		return nil
	}

	return path
}
