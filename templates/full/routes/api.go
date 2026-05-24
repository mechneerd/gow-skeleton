package routes

import (
	"encoding/json"
	"net/http"
)

func RegisterAPIRoutes() {
	http.HandleFunc("/api/health", apiHealthHandler)
	http.HandleFunc("/api/version", apiVersionHandler)
}

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"app":    "{{ .AppName }}",
	})
}

func apiVersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"version": "1.0.0",
		"app":     "{{ .AppName }}",
	})
}
