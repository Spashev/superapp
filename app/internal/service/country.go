package service

import (
	"github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
)

type CountryService struct {
	repo *repository.CountryRepository
}

func NewCountryService(repo *repository.CountryRepository) *CountryService {
	return &CountryService{repo: repo}
}

func (s *CountryService) GetAllCountries() ([]models.Country, error) {
	return s.repo.GetAllCountries()
}

func (s *CountryService) GetAllCities() ([]models.City, error) {
	return s.repo.GetAllCities()
}

func (s *CountryService) GetAllCitiesByCountryId(countryId int) ([]models.City, error) {
	return s.repo.GetAllCitiesByCountryId(countryId)
}
