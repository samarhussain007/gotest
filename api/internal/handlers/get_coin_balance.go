package handlers

import (
	"api"                // Importing the 'api' package for access to related methods and types.
	"api/internal/tools" // Importing tools package which contains utility functions and types.
	"encoding/json"      // Importing the JSON encoding/decoding package.
	"fmt"
	"net/http" // Importing the HTTP package for handling requests and responses.

	"github.com/gorilla/schema"      // Importing the schema package to decode URL query parameters.
	log "github.com/sirupsen/logrus" // Importing logrus for structured logging.
)

// GetCoinBalance is an HTTP handler function that handles the request to fetch the user's coin balance.
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the query parameters, typically used to capture request params.
	var params = api.CoinBalanceParams{}

	// Create a new schema decoder instance to decode the query parameters from the URL.
	var decoder *schema.Decoder = schema.NewDecoder()

	// Variable to hold potential errors during decoding.
	var err error

	// Decode the URL query parameters into the params struct.
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		// Log the error and handle it by sending an internal server error response.
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Initialize the database interface to interact with the database.
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase() // Create a new database connection.

	fmt.Println(database, err)
	if err != nil {
		// Handle any error in database connection.
		api.InternalErrorHandler(w)
		return
	}

	fmt.Println("Its here")

	// Fetch the coin details of the user using the provided username.
	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)

	// If the coin details are not found (nil), log the error and send an internal error response.
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Prepare the response containing the user's coin balance.
	var response = api.CoinResponse{
		Balance: (*tokenDetails).Coins, // Set the balance field from the retrieved token details.
		Code:    http.StatusOK,         // Set the HTTP status code to 200 (OK).
	}

	// Set the content type of the response to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode the response struct into JSON and send it to the client.
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// If encoding the response fails, log the error and send an internal error response.
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
