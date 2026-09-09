package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	addr := os.Getenv("APP_ADDR")
	if addr == "" { addr = ":8080" }
	if err := http.ListenAndServe(addr, mux); err != nil { panic(err) }
}
