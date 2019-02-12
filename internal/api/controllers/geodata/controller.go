package geodata

import "net/http"

// Controller is in charge of transport-level validation.
// Errors are returned as HTTP codes.
type Controller interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetCategoryByID(w http.ResponseWriter, r *http.Request)
}
