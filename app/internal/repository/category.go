package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/internal/models"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	imageBaseUrl := os.Getenv("IMAGE_BASE_URL")
	if imageBaseUrl == "" {
		return nil, fmt.Errorf("IMAGE_BASE_URL not set")
	}

	var categories []models.Category
	query := `
		SELECT 
			id,
			name_ru AS name,
			slug,
			icon
		FROM
			categories
	`
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	for i := range categories {
		categories[i].Icon = imageBaseUrl + "/" + categories[i].Icon
	}

	return categories, nil
}
