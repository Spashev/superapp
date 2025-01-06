package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"superapp/cmd/app/router"
	"superapp/config"
	"superapp/internal/db"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

}

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
