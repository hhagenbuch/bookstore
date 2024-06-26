package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the incoming HTTP requests with their method, URL, and processing time.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("method=%s url=%s time=%s", r.Method, r.URL, time.Since(start))
	})
}

// RecoverMiddleware recovers from any panics and writes a 500 internal server error response.
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
