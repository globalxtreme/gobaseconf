package filesystem

import (
	"github.com/globalxtreme/gobaseconf/helpers"
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

func (repo Storage) ShowFile(w http.ResponseWriter, path string) {
	var request http.Request

	if repo.IsPublic {
		path = "public/" + path
	}

	realPath := checkAndSetDefaultFile(path)
	if realPath == nil {
		http.NotFound(w, &request)
		return
	}

	http.ServeFile(w, &request, realPath.(string))
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
