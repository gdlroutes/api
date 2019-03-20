package user

import "net/http"

// Controller is in charge of transport-level validation.
// Errors are returned as HTTP codes.
type Controller interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
}
