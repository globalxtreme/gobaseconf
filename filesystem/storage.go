package filesystem

import (
	"github.com/globalxtreme/gobaseconf/helpers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Storage struct {
	IsPublic bool
}

func (repo Storage) GetFullPath(path string) string {
	if repo.IsPublic {
		path = "public/" + path
	}

	baseDir, _ := os.Getwd()
	storageDir := helpers.SetStorageDir(path)

	return baseDir + "/" + storageDir
}

func (repo Storage) GetFullPathURL(path string) string {
	return os.Getenv("API_GATEWAY_LINK_URL") + path
}

func (repo Storage) ShowFile(w http.ResponseWriter, r *http.Request, paths ...string) {
	var path string

	if len(paths) > 0 {
		path = paths[0]
	} else {
		vars := mux.Vars(r)
		path = vars["path"]
	}

	if repo.IsPublic {
		path = "public/" + path
	}

	realPath := checkAndSetDefaultFile(path)
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
