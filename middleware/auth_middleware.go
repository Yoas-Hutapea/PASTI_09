package middleware

import (
	"net/http"
)

// NewAuthMiddleware creates a new authentication middleware.
// Modify this function to implement your user authentication logic.
func NewAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the user is authenticated
		if !isAuthenticated(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the user is authenticated, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// isAuthenticated is an example function to check if the user is authenticated.
// Modify this function to implement your actual user authentication logic.
func isAuthenticated(r *http.Request) bool {
	// Implement your user authentication logic here.
	// For example, you can check if the user has a valid session or token.
	// Return true if the user is authenticated, false otherwise.
	// Modify this function to fit your authentication implementation.
	return true
}
