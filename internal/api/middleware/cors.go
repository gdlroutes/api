package middleware

import (
	"net/http"

	"github.com/justinas/alice"
)

// CORS is a function that returns a middleware that adds the CORS headers
func CORS(origin string) alice.Constructor {

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			h.ServeHTTP(w, r)
		})
	}
}
