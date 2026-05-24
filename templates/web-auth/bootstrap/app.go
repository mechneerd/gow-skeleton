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

	// Recommended: Set up RBAC DB connection here (after opening your DB)
	// import "yourmodule/auth/rbac"
	// rbac.SetDefaultDB(yourDB)

	// Example protected route registration (uncomment and adjust as needed):
	// import "yourmodule/http/middleware"
	// http.Handle("/admin", middleware.RoleMiddleware("Super Admin")(yourHandler))

	addr := ":" + a.Config.Port
	fmt.Printf("%s (with Auth + RBAC) is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
