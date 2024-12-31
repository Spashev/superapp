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
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/products", handler.GetProductList)

	addr := ":8080"
	fmt.Printf("Server running on %s...\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
