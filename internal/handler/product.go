package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"

	"superapp/internal/repository"
	"superapp/internal/service"
)

func GetProductList(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := r.Context().Value("page").(int)
		limit := r.Context().Value("limit").(int)

		repo := repository.NewProductRepository(db)
		productService := service.NewProductService(repo)

		products, err := productService.GetAllProducts(page, limit)
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

func GetProductBySlug(db *sqlx.DB) http.HandlerFunc {
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
