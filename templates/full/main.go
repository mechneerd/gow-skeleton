package main

import "{{ .ModulePath }}/bootstrap"

func main() {
	app := bootstrap.NewApplication()
	app.Serve()
}
