package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"{{ .ModulePath }}/config"
	"{{ .ModulePath }}/routes"
)

// Application represents the full starter application.
type Application struct {
	Config config.AppConfig
}

// NewApplication initializes config and prepares routes.
func NewApplication() *Application {
	cfg := config.Load()
	return &Application{Config: cfg}
}

// Serve starts the HTTP server with both web and API routes registered.
func (a *Application) Serve() {
	routes.RegisterRoutes()
	routes.RegisterAPIRoutes()

	// Recommended one-time setup for RBAC:
	// import "yourmodule/auth/rbac"
	// rbac.SetDefaultDB(yourDB)

	// Example middleware usage:
	// import "yourmodule/http/middleware"
	// http.Handle("/admin", middleware.RoleMiddleware("Super Admin")(adminHandler))

	addr := ":" + a.Config.Port
	fmt.Printf("%s (Full Starter) is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
