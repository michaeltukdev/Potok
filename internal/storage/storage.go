package storage

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

func HandleUploadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	vault := vars["vault"]
	filepathInVault := vars["filepath"]

	dir := path.Join("data", user, vault, path.Dir(filepathInVault))
	os.MkdirAll(dir, 0700)

	f, err := os.Create(path.Join("data", user, vault, filepathInVault))
	if err != nil {
		http.Error(w, "Failed to save file", 500)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
	w.WriteHeader(http.StatusCreated)
}
