package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"

	"superapp/handler"
)

func main() {
	// Create a new router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger) // Logs requests

	// Register routes
	r.Get("/products", handler.GetProductList)

	// Start the server
	addr := ":8080"
	fmt.Printf("Server running on %s...\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
