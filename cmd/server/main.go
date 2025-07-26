package main

import (
	"log"

	"github.com/michaeltukdev/Potok/internal/api"
	"github.com/michaeltukdev/Potok/internal/database"
)

func main() {
	err := database.InitDB("potok.db")
	if err != nil {
		log.Fatal(err)
	}

	api.StartServer()
}
