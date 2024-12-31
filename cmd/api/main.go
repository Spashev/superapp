package main

import (
	"fmt"
	"log"
	"net/http"

	"superapp/cmd/router"
	"superapp/config"
	"superapp/db"
)

func main() {
	cfg := config.NewConfig()

	database, err := db.NewDatabase(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	r := router.RegisterRoutes(database.Conn)

	addr := ":8080"
	fmt.Printf("Server running on %s...\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
