package routes

import (
	"encoding/json"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/mechneerd/gow/routing"
	"github.com/mechneerd/gow/view"
)

var router *routing.Router
var viewEngine *view.Engine
var counterValue int

// Monitor stats
var (
	monitorStartTime = time.Now()
	monitorTotalReqs int64
	monitorMu        sync.RWMutex
)

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

	// Monitor page
	router.Get("/monitor", func(w http.ResponseWriter, r *http.Request) error {
		if r.URL.Query().Get("format") == "json" {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			monitorMu.RLock()
			totalReqs := monitorTotalReqs
			monitorMu.RUnlock()
			uptime := time.Since(monitorStartTime).Round(time.Second)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"memory_alloc":     m.Alloc,
				"memory_sys":       m.Sys,
				"go_routines":      runtime.NumGoroutine(),
				"open_connections": 0,
				"total_requests":   totalReqs,
				"uptime":           uptime.String(),
				"response_time":    0,
				"timestamp":        time.Now().Unix(),
			})
			return nil
		}

		html, err := viewEngine.Make("monitor", map[string]any{
			"AppName": "{{ .AppName }}",
		})
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("<h1>Monitor</h1><p>Monitor page not found.</p>"))
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
