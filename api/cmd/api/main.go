package main

import (
	"api/internal/handlers"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// Enable reporting the caller in log entries for better debugging.
	log.SetReportCaller(true)

	// Initialize a new router from the chi package to handle HTTP requests.
	var r *chi.Mux = chi.NewRouter()
	// Call the Handler function from the 'handlers' package to set up routes and middleware.
	handlers.Handler(r)

	// Print a message indicating that the API service is starting.
	fmt.Println("Starting GO API service ...")

	// Start the HTTP server and listen on localhost at port 8000 with the chi router to handle requests.
	err := http.ListenAndServe("localhost:8000", r)

	// If there is an error while starting the server, log the error.
	if err != nil {
		log.Error(err)
	}
}
