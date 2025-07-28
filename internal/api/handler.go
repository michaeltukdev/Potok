package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/michaeltukdev/Potok/internal/database"
	"github.com/michaeltukdev/Potok/internal/middleware"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{user}/vaults", handleVaults).Handler(middleware.ApiMiddleware(http.HandlerFunc(handleVaults)))

	r.HandleFunc("/users/{user}/vaults/{vault}", handlePostVault).Methods("POST")
	r.HandleFunc("/users/{user}/vaults/{vault}", handleDeleteVault).Methods("DELETE")

	r.HandleFunc("/users/{user}/vaults/{vault}/upload", handleUploadVault).Methods("POST")
	r.HandleFunc("/users/{user}/vaults/{vault}/download", handleDownloadVault).Methods("GET")

	r.HandleFunc("/users/{user}/vaults/{vault}/files/{filepath:.*}", handleDownloadFile).Methods("GET")
	r.HandleFunc("/users/{user}/vaults/{vault}/files/{filepath:.*}", handleUploadFile).Methods("POST").Handler(middleware.ApiMiddleware(http.HandlerFunc(handleUploadFile)))

	r.HandleFunc("/me", handleMe).Methods("GET").Handler(middleware.ApiMiddleware(http.HandlerFunc(handleMe)))

	http.ListenAndServe(":8080", r)

	log.Println("Starting server on :8080")
}

func handleVaults(w http.ResponseWriter, r *http.Request) {
	vaults, err := database.FetchUserVaults(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Error fetching vaults", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(vaults); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func handlePostVault(w http.ResponseWriter, r *http.Request)     {}
func handleDeleteVault(w http.ResponseWriter, r *http.Request)   {}
func handleUploadVault(w http.ResponseWriter, r *http.Request)   {}
func handleDownloadVault(w http.ResponseWriter, r *http.Request) {}
func handleDownloadFile(w http.ResponseWriter, r *http.Request)  {}

func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	user, err := database.FindByAPIKey(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Authentication failed!", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	urlUser := vars["user"]
	vault := vars["vault"]
	filepathInVault := vars["filepath"]

	// TODO: See how this will work on the client (thinking of just removing this)
	if user.Username != urlUser {
		http.Error(w, "Authentication failed!", http.StatusUnauthorized)
		return
	}

	fetchedVault, err := database.FetchUserVaultByName(user.Api_key, vault)
	if err != nil {
		http.Error(w, "Unauthorized Access!", http.StatusUnauthorized)
		return
	}

	baseDir := "../../data"
	fullPath := path.Join(baseDir, user.Username, fetchedVault.Name, filepathInVault)

	if err := os.MkdirAll(path.Dir(fullPath), 0700); err != nil {
		http.Error(w, "Failed to create directories", http.StatusInternalServerError)
		return
	}

	f, err := os.Create(fullPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if _, err := io.Copy(f, r.Body); err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func handleMe(w http.ResponseWriter, r *http.Request) {
	user, err := database.FindByAPIKey(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Authentication failed!", http.StatusUnauthorized)
		return
	}

	resp := struct {
		Username string `json:"username"`
		ID       int    `json:"id"`
	}{
		Username: user.Username,
		ID:       user.Id,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
