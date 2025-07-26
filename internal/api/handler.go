package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")

		http.Error(w, "Not Found", http.StatusNotFound)
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
