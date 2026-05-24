package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"{{ .ModulePath }}/config"
	"{{ .ModulePath }}/routes"
)

// Application holds the minimal app state.
type Application struct {
	Config config.AppConfig
}

// NewApplication initializes the minimal API application.
func NewApplication() *Application {
	cfg := config.Load()
	return &Application{Config: cfg}
}

// Serve starts the HTTP server.
func (a *Application) Serve() {
	routes.RegisterRoutes()

	addr := ":" + a.Config.Port
	fmt.Printf("%s (Minimal API) is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
