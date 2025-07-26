package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{user}/vaults", handleVaults)

	r.HandleFunc("/users/{user}/vaults/{vault}", handlePostVault).Methods("POST")
	r.HandleFunc("/users/{user}/vaults/{vault}", handleDeleteVault).Methods("DELETE")

	r.HandleFunc("/users/{user}/vaults/{vault}/upload", handleUploadVault).Methods("POST")
	r.HandleFunc("/users/{user}/vaults/{vault}/download", handleDownloadVault).Methods("GET")

	r.HandleFunc("/users/{user}/vaults/{vault}/files/{filepath:.*}", handleDownloadFile).Methods("GET")
	r.HandleFunc("/users/{user}/vaults/{vault}/files/{filepath:.*}", handleUploadFile).Methods("POST")

	http.ListenAndServe(":8080", r)

	log.Println("Starting server on :8080")
}

func handleVaults(w http.ResponseWriter, r *http.Request)        {}
func handlePostVault(w http.ResponseWriter, r *http.Request)     {}
func handleDeleteVault(w http.ResponseWriter, r *http.Request)   {}
func handleUploadVault(w http.ResponseWriter, r *http.Request)   {}
func handleDownloadVault(w http.ResponseWriter, r *http.Request) {}
func handleDownloadFile(w http.ResponseWriter, r *http.Request)  {}
func handleUploadFile(w http.ResponseWriter, r *http.Request)    {}
