package middleware

import (
	"api"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// UnAuthorizedError is a custom error used when the username or token is invalid.
var UnAuthorizedError = errors.New("Invalid username or token")

// Authorization is a middleware function that checks the validity of a username and authorization token
// passed through the request. If either the username or token is missing or invalid, it returns an unauthorized error.
// Otherwise, it allows the request to proceed to the next handler in the middleware chain.
func Authorization(next http.Handler) http.Handler {
	// Return an http.HandlerFunc that wraps the next handler in the middleware chain.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the 'username' query parameter from the request URL.
		var userName string = r.URL.Query().Get("username")
		// Extract the 'Authorization' header from the request.
		var token = r.Header.Get("Authorization")

		// If either the 'username' or 'token' is missing, log the error and return an unauthorized error response.
		if userName == "" || token == "" {
			log.Error(UnAuthorizedError)                  // Log the error for missing or invalid credentials.
			api.RequestErrorHandler(w, UnAuthorizedError) // Call the custom error handler to respond with an error.
			return
		}

		// If both username and token are provided, call the next handler in the middleware chain.
		next.ServeHTTP(w, r)
	})
}
