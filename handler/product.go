package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
