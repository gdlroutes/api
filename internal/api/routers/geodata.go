package routers

import (
	"net/http"

	"github.com/gdlroutes/api/internal/api/controllers/geodata"
	"github.com/gorilla/mux"
)

// GeodataPrefix is used to route all hostpot-related requests
const GeodataPrefix = "/geodata/"

// Hotspot controlls hotspot-related routes
type Hotspot struct {
	Controller geodata.Controller
	router     *mux.Router
}

var _ http.Handler = &Hotspot{}

func (h *Hotspot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.router == nil {
		// Initialize Mux
		h.router = mux.NewRouter()
		h.router.HandleFunc(GeodataPrefix+"categories/{id}", h.Controller.GetCategoryByID).
			Methods(http.MethodGet)
		h.router.HandleFunc(GeodataPrefix+"categories", h.Controller.GetCategories).
			Methods(http.MethodGet)
	}

	h.router.ServeHTTP(w, r)
}
