package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gdlroutes/api/internal/api/middleware"

	"github.com/gdlroutes/api/internal/api/models"
)

func (h *Router) signUp(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in /signup:", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	user := &models.User{}
	if err := json.Unmarshal(body, user); err != nil {
		http.Error(w, "Invalid body.", http.StatusBadRequest)
		return
	}

	token, err := h.UserUseCases.CreateUserAndToken(user)
	switch err.(type) {
	case nil:
		break
	case models.ConflictError:
		http.Error(w, err.Error(), http.StatusConflict)
		return
	default:
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie := h.buildTokenCookie(token)
	http.SetCookie(w, cookie)
}

func (h *Router) logIn(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in /login:", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}

	user := &models.User{}
	if err := json.Unmarshal(body, user); err != nil {
		http.Error(w, "Invalid body.", http.StatusBadRequest)
		return
	}

	token, err := h.UserUseCases.CreateToken(user)
	switch err.(type) {
	case nil:
		break
	case models.InvalidCredentialsError:
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	default:
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	cookie := h.buildTokenCookie(token)
	http.SetCookie(w, cookie)
}

func (h *Router) logOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value(middleware.AccessTokenCookieKey).(string)
	if !ok {
		http.Error(w, "No session found", http.StatusForbidden)
		return
	}

	cookie := h.buildEmptyCookie()
	http.SetCookie(w, cookie)
}

func (h *Router) buildTokenCookie(token *models.Token) *http.Cookie {
	return &http.Cookie{
		Domain:  h.CookieDomain,
		Expires: token.Expires,
		Name:    middleware.AccessTokenCookieName,
		Value:   token.Token,
	}
}

func (h *Router) buildEmptyCookie() *http.Cookie {
	return &http.Cookie{
		Domain:  h.CookieDomain,
		Expires: time.Unix(0, 0),
		Name:    middleware.AccessTokenCookieName,
		Value:   "",
	}
}
