package geodata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gdlroutes/api/internal/api/models"
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

	return &controller{useCases: useCases}, nil
}

func (c *controller) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.useCases.GetCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(categories)
	w.Write(bytes)
}

func (c *controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID must be numeric.", http.StatusBadRequest)
		return
	}

	category, err := c.useCases.GetCategoryByID(categoryID)
	switch err.(type) {
	case nil:
		break
	case models.NotFoundError:
		http.Error(w, fmt.Sprintf("Category with id %d not found.", categoryID), http.StatusNotFound)
		return
	default:
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(category)
	w.Write(bytes)
}
