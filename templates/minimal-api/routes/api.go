package routes

import (
	"net/http"

	"github.com/mechneerd/gow/routing"
	"github.com/mechneerd/gow/view"
)

var router *routing.Router
var viewEngine *view.Engine

func init() {
	router = routing.NewRouter()
	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	// Beautiful landing page (even for API kits, it's a nice welcome)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("welcome", map[string]any{
			"AppName": "{{ .AppName }}",
			"Year":    "2026",
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Welcome to {{ .AppName }} API"}`))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})
}

func RegisterRoutes() {
	http.Handle("/", router)

	// Keep health endpoint as JSON
	http.HandleFunc("/health", healthHandler)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy", "app": "{{ .AppName }}"}`))
}
