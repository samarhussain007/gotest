package handlers

import (
	"api/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Handler function sets up the routes and middleware for the API.
// It is responsible for configuring the HTTP request handling behavior for the application.
func Handler(r *chi.Mux) {
	// Use StripSlashes middleware to remove trailing slashes from the URL path
	r.Use(chimiddle.StripSlashes)

	// Define the "/account" route group
	r.Route("/account", func(router chi.Router) {

		// Apply the Authorization middleware to routes inside the "/account" group.
		// This ensures that the user is authenticated before accessing the routes.
		router.Use(middleware.Authorization)

		// Define the route for GET requests to "/account/coins".
		// This route calls the GetCoinBalance function to handle the request and return the user's coin balance.
		router.Get("/coins", GetCoinBalance)
	})
}
