package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/internal/models"
)

type CountryRepository struct {
	db *sqlx.DB
}

func NewCountryRepository(db *sqlx.DB) *CountryRepository {
	return &CountryRepository{db: db}
}

func (r *CountryRepository) GetAllCountries() ([]models.Country, error) {
	var countries []models.Country
	query := `
		SELECT 
			id,
			name_ru AS name,
			code
		FROM
			country
	`
	err := r.db.Select(&countries, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return countries, nil
}

func (r *CountryRepository) GetAllCities() ([]models.City, error) {
	var cities []models.City
	query := `
		SELECT 
			id,
			name_ru AS name,
			postall_code
		FROM
			city
	`
	err := r.db.Select(&cities, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return cities, nil
}

func (r *CountryRepository) GetAllCitiesByCountryId(countryId int) ([]models.City, error) {
	var cities []models.City
	query := `
		SELECT 
			id,
			name_ru AS name,
			postall_code
		FROM
			city
		WHERE 
			country_id = $1
	`

	err := r.db.Select(&cities, query, countryId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return cities, nil
}
