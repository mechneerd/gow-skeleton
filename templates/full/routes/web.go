package routes

import (
	"net/http"

	"gow/routing"
	"gow/view"
)

var router *routing.Router
var viewEngine *view.Engine

func init() {
	router = routing.NewRouter()
	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	// Beautiful landing page + Livewire support
	router.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("welcome", map[string]any{
			"AppName": "{{ .AppName }}",
			"Year":    "2026",
		})
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Welcome to {{ .AppName }}</h1>"))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})
}

func RegisterRoutes() {
	http.Handle("/", router)

	// Keep legacy simple handlers if needed
	http.HandleFunc("/about", aboutHandler)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>About {{ .AppName }}</h1>"))
}
