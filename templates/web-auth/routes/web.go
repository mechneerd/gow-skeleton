package routes

import (
	"net/http"

	"{{ .ModulePath }}/app/Http/Controllers/Auth"
)

func RegisterWebRoutes() {
	// Public routes
	http.HandleFunc("/login", Auth.LoginHandler)
	http.HandleFunc("/register", Auth.RegisterHandler)
	http.HandleFunc("/logout", Auth.LogoutHandler)

	// Protected routes example (middleware will be added later)
	http.HandleFunc("/dashboard", Auth.DashboardHandler)
}

func RegisterAPIRoutes() {
	// Sanctum protected API routes
	http.HandleFunc("/api/user", Auth.MeHandler)
}
