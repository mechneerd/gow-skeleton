package routes

import (
	"net/http"

	"{{ .ModulePath }}/app/Livewire"
	"github.com/mechneerd/gow/http/livewire"
	"github.com/mechneerd/gow/routing"
	"github.com/mechneerd/gow/view"
)

var router *routing.Router
var viewEngine *view.Engine

func init() {
	router = routing.NewRouter()

	// Setup view engine (GoBlade)
	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	// Register Livewire
	livewire.RegisterRoutes(router)

	// Main landing page using the beautiful GoBlade template + Livewire demo
	router.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("welcome", map[string]any{
			"AppName": "{{ .AppName }}",
			"Year":    "2026",
		})
		if err != nil {
			// Fallback simple response
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Welcome to {{ .AppName }}</h1>"))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})

	// About page (simple example)
	router.Get("/about", func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<h1>About {{ .AppName }}</h1><p>This is a GoW starter kit.</p>"))
		return nil
	})
}

func RegisterRoutes() {
	// Mount the router
	http.Handle("/", router)
}
