package router

import (
	"net/http"

	"github.com/gdlroutes/api/internal/api/usecases/geodata"
	"github.com/gdlroutes/api/internal/api/usecases/user"
	"github.com/gorilla/mux"
)

const geodataPrefix = "/geodata/"
const userPrefix = "/users/"

// Router routes all requests to correct endpoints
type Router struct {
	GeodataUseCases geodata.UseCases
	UserUseCases    user.UseCases
	CookieDomain    string

	router *mux.Router
}

var _ http.Handler = &Router{}

func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.router == nil {
		// Initialize Mux
		h.router = mux.NewRouter()

		// Geodata
		h.router.HandleFunc(geodataPrefix+"categories/{id}", h.getCategoryByID).
			Methods(http.MethodGet)
		h.router.HandleFunc(geodataPrefix+"categories", h.getCategories).
			Methods(http.MethodGet)
		h.router.HandleFunc(geodataPrefix+"routeCategories", h.getRouteCategories).
			Methods(http.MethodGet)
		h.router.HandleFunc(geodataPrefix+"routes", h.createRoute).
			Methods(http.MethodPost)
		h.router.HandleFunc(geodataPrefix+"routes", h.getRoutesByCategory).
			Methods(http.MethodGet).
			Queries("category_id", "{category_id}")
		h.router.HandleFunc(geodataPrefix+"routes", h.getRoutes).
			Methods(http.MethodGet)

		// User
		h.router.HandleFunc(userPrefix+"login", h.logIn).
			Methods(http.MethodPost)
		h.router.HandleFunc(userPrefix+"logout", h.logOut).
			Methods(http.MethodGet)
		h.router.HandleFunc(userPrefix+"signup", h.signUp).
			Methods(http.MethodPost)
	}

	h.router.ServeHTTP(w, r)
}
