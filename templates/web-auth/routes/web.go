package routes

import (
	"net/http"

	"{{ .ModulePath }}/app/Http/Controllers/Auth"
	"gow/routing"
	"gow/view"
)

var router *routing.Router
var viewEngine *view.Engine

func init() {
	router = routing.NewRouter()

	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	// GoW Livewire + beautiful landing page
	// Note: livewire.RegisterRoutes(router) would be ideal here once imported
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

func RegisterWebRoutes() {
	// Mount router for new style routes (landing page + future Livewire)
	http.Handle("/", router)

	// Existing auth routes (raw for now)
	http.HandleFunc("/login", Auth.LoginHandler)
	http.HandleFunc("/register", Auth.RegisterHandler)
	http.HandleFunc("/logout", Auth.LogoutHandler)
	http.HandleFunc("/dashboard", Auth.DashboardHandler)
}

func RegisterAPIRoutes() {
	http.HandleFunc("/api/user", Auth.MeHandler)
}
