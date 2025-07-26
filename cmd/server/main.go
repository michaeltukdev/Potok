package main

import (
	"fmt"
	"log"

	"github.com/michaeltukdev/Potok/internal/api"
	"github.com/michaeltukdev/Potok/internal/database"
)

func main() {
	db, err := database.InitDB("potok.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database running...")

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Migrations completed...")

	fmt.Println("Starting HTTP server on :8080")
	api.StartServer()
}
