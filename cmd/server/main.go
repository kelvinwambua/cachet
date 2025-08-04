package main

import (
	"cachet/internal/db"
	"cachet/internal/server"
	"log"
)

func main() {
	store := db.NewStore()
	s := server.New(":6380", store)
	log.Fatal(s.ListenAndServe())
}
