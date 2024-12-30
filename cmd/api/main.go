package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// Replace with your actual connection details
	dsn := "postgres://user:password@localhost:6432/mydatabase?sslmode=disable"

	// Initialize the database
	db, err := NewDatabase(dsn)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	fmt.Println("Successfully connected to the database!")

	// Example: Query the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var currentTime time.Time
	err = db.Conn.QueryRowContext(ctx, "SELECT NOW()").Scan(&currentTime)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Printf("Current time in database: %v\n", currentTime)
}
