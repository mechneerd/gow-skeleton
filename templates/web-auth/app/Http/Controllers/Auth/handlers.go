package Auth

import (
	"encoding/json"
	"net/http"

	"github.com/mechneerd/gow/view"
)

var viewEngine = view.NewEngine([]string{"resources/views"}, "storage/framework/views")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html, err := viewEngine.Make("auth.login", nil)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("<h1>Login</h1><form method='POST'><input name='email' placeholder='Email'><input name='password' type='password' placeholder='Password'><button type='submit'>Login</button></form>"))
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return
	}

	http.Error(w, "Login not yet implemented", http.StatusNotImplemented)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html, err := viewEngine.Make("auth.register", nil)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte("<h1>Register</h1><form method='POST'><input name='name' placeholder='Name'><input name='email' placeholder='Email'><input name='password' type='password' placeholder='Password'><button type='submit'>Register</button></form>"))
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
		return
	}

	http.Error(w, "Registration not yet implemented", http.StatusNotImplemented)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	html, err := viewEngine.Make("dashboard", nil)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Dashboard</h1><p>Welcome! Set up your dashboard view.</p>"))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Not authenticated",
	})
}
