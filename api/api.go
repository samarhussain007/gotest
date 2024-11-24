package api

import (
	"encoding/json"
	"net/http"
)

// CoinBalanceParams struct is used to capture the parameters
// sent by the client to fetch the coin balance, in this case, the username.
type CoinBalanceParams struct {
	Username string
}

// CoinResponse struct represents the response structure
// for returning the user's coin balance along with a status code.
type CoinResponse struct {
	Code    int   // HTTP status code
	Balance int64 // The user's coin balance
}

// Error struct is used for returning error responses with a code and message.
type Error struct {
	Code    int    // HTTP status code for the error
	Message string // The error message to be sent to the client
}

// writeError writes a formatted error response with the given message and HTTP status code.
// The response is encoded as JSON and sent to the client.
func writeError(w http.ResponseWriter, message string, code int) {
	// Create an Error object with the provided code and message
	resp := Error{
		Code:    code,
		Message: message,
	}

	// Set the response header to indicate that the response is in JSON format
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code for the response
	w.WriteHeader(code)
	// Encode the error message as JSON and send it to the client
	json.NewEncoder(w).Encode(resp)
}

// RequestErrorHandler is a function that handles client errors (e.g., bad request),
// sending a detailed error message with HTTP status 400 (Bad Request).
var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusBadRequest)
}

// InternalErrorHandler is a function that handles internal server errors,
// sending a generic error message with HTTP status 500 (Internal Server Error).
var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
}
