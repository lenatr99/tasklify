package middlewares

import (
	"net/http"
)

func CSPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "base-uri 'self'; default-src 'self'; style-src 'unsafe-inline' 'self'; object-src 'none';")
		next.ServeHTTP(w, r)
	})
}

func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
