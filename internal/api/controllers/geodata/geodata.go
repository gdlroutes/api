package geodata

import (
	"errors"
	"net/http"

	"github.com/gdlroutes/api/internal/api/usecases/geodata"
)

type controller struct {
	useCases geodata.UseCases
}

var _ Controller = &controller{}

// New returns a new, initialized, hotspot controller
func New(useCases geodata.UseCases) (Controller, error) {
	if useCases == nil {
		return nil, errors.New("nil useCases")
	}

	return &controller{}, nil
}

// TODO: implement
func (c *controller) GetCategories(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "GetCategories", http.StatusNotImplemented)
}

// TODO: implement
func (c *controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "GetCategoryByID", http.StatusNotImplemented)
}
