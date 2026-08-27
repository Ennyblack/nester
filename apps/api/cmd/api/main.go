package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Version and Commit are injected via ldflags at build time.
var (
	Version = "dev"
	Commit  = "unknown"
)

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	log.Printf("Starting Nester API server | version=%s commit=%s", Version, Commit)

	mux := http.NewServeMux()

	// Health status endpoint exposing build version and commit
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK">
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":    "ok",
			"version":   Version,
			"commit":    Commit,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	})

	addr := fmt.Sprintf(":%s", *port)
	log.Printf("Listening on %s...", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
		os.Exit(1)
	}
}
