package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"superapp/internal/repository"
	"superapp/internal/service"
)

func GetCategories(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewCategoryRepository(db)
		categoryService := service.NewCategoryService(repo)

		categories, err := categoryService.GetAllCategories()
		if err != nil {
			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
