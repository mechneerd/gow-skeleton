package main

import (
	"{{ .ModulePath }}/bootstrap"
	"{{ .ModulePath }}/routes"
)

func main() {
	// Register routes (web + API)
	routes.RegisterWebRoutes()
	routes.RegisterAPIRoutes()

	app := bootstrap.NewApplication()
	app.Serve()
}
