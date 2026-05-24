package routes

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Home - {{ .AppName }}</h1><p>Welcome to the Full Starter.</p>"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>About {{ .AppName }}</h1>"))
}
