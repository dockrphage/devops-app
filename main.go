package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0.3"
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
// Trigger CI pipeline: Tue Jul 21 05:07:31 PM BST 2026
// Test fix Tue Jul 21 05:10:21 PM BST 2026
// Trigger pipeline with fixed auth Tue Jul 21 05:17:55 PM BST 2026
// Clean pipeline trigger: Tue Jul 21 05:22:50 PM BST 2026
