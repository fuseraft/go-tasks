package routes

import (
	"log"
	"net/http"
	"time"
)

// Middleware to log incoming requests and their response time
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		log.Printf("Started %s %s", r.Method, r.RequestURI)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the completion of the request
		log.Printf("Completed in %v", time.Since(start))
	})
}
