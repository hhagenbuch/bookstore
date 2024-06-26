package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hhagenbuch/bookstore/internal/database"
	"github.com/hhagenbuch/bookstore/internal/handlers"
	"github.com/hhagenbuch/bookstore/internal/middleware"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// Seed the database
	if err := database.SeedDB(); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	// Create a new router
	r := mux.NewRouter()

	// Apply middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoverMiddleware)

	// Register handlers
	handlers.RegisterHandlers(r)

	// Start the server
	port := ":8000"
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
