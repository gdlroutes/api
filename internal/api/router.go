package api

import (
	"net/http"

	"github.com/gdlroutes/api/internal/api/controllers/geodata"
	"github.com/gdlroutes/api/internal/api/controllers/user"
	"github.com/gorilla/mux"
)

const geodataPrefix = "/geodata/"
const userPrefix = "/users/"

// Router routes all requests to correct endpoints
type Router struct {
	GeodataController geodata.Controller
	UserController    user.Controller
	router            *mux.Router
}

var _ http.Handler = &Router{}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.router == nil {
		// Initialize Mux
		h.router = mux.NewRouter()

		// Geodata
		h.router.HandleFunc(geodataPrefix+"categories/{id}", h.GeodataController.GetCategoryByID).
			Methods(http.MethodGet)
		h.router.HandleFunc(geodataPrefix+"categories", h.GeodataController.GetCategories).
			Methods(http.MethodGet)

		// User
		h.router.HandleFunc(userPrefix+"login", h.UserController.LogIn).
			Methods(http.MethodPost)
		h.router.HandleFunc(userPrefix+"logout", h.UserController.LogOut).
			Methods(http.MethodGet)
		h.router.HandleFunc(userPrefix+"signup", h.UserController.SignUp).
			Methods(http.MethodPost)
	}

	h.router.ServeHTTP(w, r)
}
