package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"superapp/repository"
	"superapp/service"
)

func GetProductList(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		products, err := productService.GetAllProducts()
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println(err)
		}
	}
}

func GetProductBySlug(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "slug")
		if slug == "" {
			http.Error(w, "Slug is required", http.StatusBadRequest)
			return
		}

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		product, err := productService.GetProductBySlug(slug)
		if err != nil {
			http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
