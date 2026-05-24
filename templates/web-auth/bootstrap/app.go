package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"{{ .ModulePath }}/config"
	"{{ .ModulePath }}/routes"
)

// Application holds the web-auth application state.
type Application struct {
	Config config.AppConfig
}

// NewApplication initializes the authenticated web application.
func NewApplication() *Application {
	cfg := config.Load()
	return &Application{Config: cfg}
}

// Serve starts the server with auth routes already registered.
func (a *Application) Serve() {
	// Routes are registered in main.go before calling Serve
	addr := ":" + a.Config.Port
	fmt.Printf("%s (with Auth + RBAC) is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
