package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gdlroutes/api/internal/api/middleware"

	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/user"
)

type controller struct {
	useCases     user.UseCases
	cookieDomain string
}

var _ Controller = &controller{}

// New returns a new, initialized, hotsgeodatapot controller
func New(useCases user.UseCases, cookieDomain string) (Controller, error) {
	if useCases == nil {
		return nil, errors.New("nil useCases")
	}

	return &controller{
		useCases:     useCases,
		cookieDomain: cookieDomain,
	}, nil
}

// SignUp creates a new user and returns a session for newly created user
func (c *controller) SignUp(w http.ResponseWriter, r *http.Request) {
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

	token, err := c.useCases.CreateUserAndToken(user)
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

	cookie := c.buildTokenCookie(token)
	http.SetCookie(w, cookie)
}

// LogIn returns a session for an existing user
func (c *controller) LogIn(w http.ResponseWriter, r *http.Request) {
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

	token, err := c.useCases.CreateToken(user)
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

	cookie := c.buildTokenCookie(token)
	http.SetCookie(w, cookie)
}

// LogOut closes a session
func (c *controller) LogOut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_, ok := ctx.Value(middleware.AccessTokenCookieKey).(string)
	if !ok {
		http.Error(w, "No session found", http.StatusForbidden)
		return
	}

	cookie := c.buildEmptyCookie()
	http.SetCookie(w, cookie)
}

func (c *controller) buildTokenCookie(token *models.Token) *http.Cookie {
	return &http.Cookie{
		Domain:  c.cookieDomain,
		Expires: token.Expires,
		Name:    middleware.AccessTokenCookieName,
		Value:   token.Token,
	}
}

func (c *controller) buildEmptyCookie() *http.Cookie {
	return &http.Cookie{
		Domain:  c.cookieDomain,
		Expires: time.Unix(0, 0),
		Name:    middleware.AccessTokenCookieName,
		Value:   "",
	}
}
