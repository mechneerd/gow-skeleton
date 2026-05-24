package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from {{ .AppName }} Full Starter!"))
}
