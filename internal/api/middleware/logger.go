package middleware

import (
	"log"
	"net/http"

	"github.com/justinas/alice"
)

// RequestLogger is a function that returns a middleware that logs all requests
func RequestLogger() alice.Constructor {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path += "?" + r.URL.RawQuery
			}
			log.Printf("%s - %s\n", r.Method, path)
			h.ServeHTTP(w, r)
		})
	}
}
