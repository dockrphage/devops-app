package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0.1"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Version: %s\n", version)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s\n", version)
	})

	fmt.Printf("Starting app version %s on :8080\n", version)
	http.ListenAndServe(":8080", nil)
}
// Pipeline test Wed Jul 22 12:15:09 PM BST 2026
