package main

import (
	"log"
	"net/http"

	"go-tasks/config"
	"go-tasks/database"
	_ "go-tasks/docs"
	"go-tasks/handlers"
	"go-tasks/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db := database.InitializeDB()
	defer db.Close()

	// Pass the database connection to the handlers
	handlers.SetDB(db)

	// Load configuration from YAML file
	cfg, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Create a new router
	r := mux.NewRouter()

	// Apply middleware (logging, error handling, etc.)
	r.Use(routes.Middleware)

	// Initialize routes from config
	routes.InitializeRoutes(r, cfg)

	// Start the server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
