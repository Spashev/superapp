package service

import (
	"github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *CategoryService) GetAllTypes() ([]models.ProductType, error) {
	return s.repo.GetAllTypes()
}
