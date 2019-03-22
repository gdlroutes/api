package middleware

import (
	"context"
	"net/http"

	"github.com/justinas/alice"
)

type key int

const (
	// AccessTokenCookieName is the name of the access token cookie
	AccessTokenCookieName = "session"
	// AccessTokenCookieKey is the key with which the access token cookie value is placed in the context
	AccessTokenCookieKey = key(0)
)

// Token is a function that returns a middleware that extracts the access token cookie and adds it to the context.
// If the cookie is not present, the value is simply not added.
func Token() alice.Constructor {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			cookie, _ := r.Cookie(AccessTokenCookieName)
			if cookie != nil {
				ctx = context.WithValue(ctx, AccessTokenCookieKey, cookie.Value)
			}

			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
