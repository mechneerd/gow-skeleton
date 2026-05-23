package bootstrap

import "fmt"

func NewApplication() *Application {
	return &Application{}
}

type Application struct{}

func (a *Application) Serve() {
	fmt.Println("Application serving on :8080")
}
