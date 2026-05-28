package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"test-project/app/Http/Controllers/Auth"
	"github.com/mechneerd/gow/routing"
	"github.com/mechneerd/gow/view"
)

var router *routing.Router
var viewEngine *view.Engine
var counterValue int

func init() {
	router = routing.NewRouter()

	viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

	// Serve static files from public/
	http.HandleFunc("/js/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public"+r.URL.Path)
	})
	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public"+r.URL.Path)
	})

	// Landing page
	router.Get("/", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("welcome", map[string]any{
			"AppName": "{{ .AppName }}",
			"Year":    "2026",
		})
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Welcome to test-project</h1>"))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})

	// Documentation page
	router.Get("/docs", func(w http.ResponseWriter, r *http.Request) error {
		html, err := viewEngine.Make("docs", map[string]any{
			"AppName": "{{ .AppName }}",
			"Year":    "2026",
		})
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Documentation</h1><p>Docs coming soon.</p>"))
			return nil
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return nil
	})

	// Livewire counter endpoint
	http.HandleFunc("/livewire/counter", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodPost {
			var req struct {
				Count int `json:"count"`
				Delta int `json:"delta"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err == nil {
				counterValue = req.Count + req.Delta
			}
		}

		json.NewEncoder(w).Encode(map[string]any{
			"count": counterValue,
			"html":  fmt.Sprintf(`<div class="flex flex-col items-center gap-4"><div class="text-6xl font-bold text-white">%s</div></div>`, strconv.Itoa(counterValue)),
		})
	})
}

func RegisterWebRoutes() {
	http.Handle("/", router)

	http.HandleFunc("/login", Auth.LoginHandler)
	http.HandleFunc("/register", Auth.RegisterHandler)
	http.HandleFunc("/logout", Auth.LogoutHandler)
	http.HandleFunc("/dashboard", Auth.DashboardHandler)
}

func RegisterAPIRoutes() {
	http.HandleFunc("/api/user", Auth.MeHandler)
}
