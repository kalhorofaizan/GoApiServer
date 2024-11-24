package middleware

import (
	"net/http"
)

type Middleware func(handler http.Handler) http.Handler

func HandleChainMiddleware(xs ...Middleware) Middleware {
	return func(handler http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			handler = xs[i](handler)
		}
		return handler
	}
}

func EnableCors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost:2222")
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}
