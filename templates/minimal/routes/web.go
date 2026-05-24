package routes

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Welcome to %s</h1><p>Minimal GoW starter.</p>", "{{ .AppName }}")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>About</h1>"))
}
