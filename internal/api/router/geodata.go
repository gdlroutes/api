package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gdlroutes/api/internal/api/middleware"
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
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *Router) createRoute(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value(middleware.AccessTokenCookieKey).(string)
	if !ok {
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in POST /routes:", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	route := &models.Route{}
	if err := json.Unmarshal(body, route); err != nil {
		http.Error(w, "Invalid body.", http.StatusBadRequest)
		return
	}

	if err := h.GeodataUseCases.CreateRoute(route); err != nil {
		log.Println("Error in POST /routes:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Router) getRoutes(w http.ResponseWriter, r *http.Request) {
	routes, err := h.GeodataUseCases.GetRoutes()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(routes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *Router) getRoutesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID, err := strconv.Atoi(vars["category_id"])
	if err != nil {
		http.Error(w, "ID must be numeric.", http.StatusBadRequest)
		return
	}

	routes, err := h.GeodataUseCases.GetRoutesByCategory(categoryID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	bytes, _ := json.Marshal(routes)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
