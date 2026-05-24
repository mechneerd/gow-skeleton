package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"{{ .ModulePath }}/config"
	"{{ .ModulePath }}/routes"
)

// Application is the minimal starter app.
type Application struct {
	Config config.AppConfig
}

// NewApplication creates the minimal app.
func NewApplication() *Application {
	return &Application{Config: config.Load()}
}

// Serve starts the server.
func (a *Application) Serve() {
	routes.RegisterRoutes()
	addr := ":" + a.Config.Port
	fmt.Printf("%s is running on http://localhost%s\n", a.Config.AppName, addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
