package main

import (
	"fmt"
	"github.com/mechneerd/gow/bootstrap"
)

func main() {
	app := bootstrap.NewApplication()
	fmt.Println("{{ .AppName }} is running!")
	app.Serve()
}
