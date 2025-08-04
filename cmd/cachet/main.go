package main

import (
	"cachet/internal/server"
	"cachet/internal/store"
	"log"
)

func main() {

	dataStore := store.NewMemoryStore()

	srv := server.New(":6380", dataStore)

	log.Printf("Starting Cachet server...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
