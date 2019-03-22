package routers

import (
	"net/http"

	"github.com/gdlroutes/api/internal/api/controllers/user"
	"github.com/gorilla/mux"
)

// UserPrefix is used to route all user-related requests
const UserPrefix = "/users/"

var authenticationRoutes = map[string]struct{}{}

// User controlls user-related routes
type User struct {
	Controller user.Controller
	router     *mux.Router
}

var _ http.Handler = &Hotspot{}

func (h *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.router == nil {
		// Initialize Mux
		h.router = mux.NewRouter()
		h.router.HandleFunc(UserPrefix+"login", h.Controller.LogIn).
			Methods(http.MethodPost)
		h.router.HandleFunc(UserPrefix+"logout", h.Controller.LogOut).
			Methods(http.MethodGet)
		h.router.HandleFunc(UserPrefix+"signup", h.Controller.SignUp).
			Methods(http.MethodPost)
	}

	h.router.ServeHTTP(w, r)
}
