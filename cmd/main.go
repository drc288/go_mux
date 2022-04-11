package main

import (
	"log"

	"github.com/drc288/go_mux/database"
	"github.com/drc288/go_mux/handlers"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Connection refuse in database")
		return
	}

	handlers.Mannage()
}
