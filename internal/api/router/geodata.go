package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gdlroutes/api/internal/api/models"
)

func (h *Router) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.GeodataUseCases.GetCategories()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(categories)
	w.Write(bytes)
}

func (h *Router) getCategoryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID must be numeric.", http.StatusBadRequest)
		return
	}

	category, err := h.GeodataUseCases.GetCategoryByID(categoryID)
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
