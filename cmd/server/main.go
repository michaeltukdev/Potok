package main

import (
	"fmt"
	"log"
	"os" // <-- Import the "os" package

	"github.com/michaeltukdev/Potok/internal/api"
	"github.com/michaeltukdev/Potok/internal/database"
)

func main() {
	// Read DATABASE_PATH from environment, with a default if not set
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "potok.db" // Default database path
	}
	db, err := database.InitDB(dbPath)
	if err != nil {
		log.Fatal("Failed to init DB:", err)
	}
	fmt.Println("Database running...")

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	fmt.Println("Migrations completed...")

	// Read PORT from environment, with a default if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Println("Starting HTTP server on :" + port)
	// Pass the port to the StartServer function
	api.StartServer(port)
}
