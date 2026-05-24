package routes

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Welcome to {{ .AppName }} API"}`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy", "app": "{{ .AppName }}"}`)
}
