package service

import (
	"superapp/internal/models"
	"superapp/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts(page, limit int) (*models.ProductsPaginate, error) {
	return s.repo.GetAllProducts(page, limit)
}

func (s *ProductService) GetProductBySlug(slug string) (models.Product, error) {
	productPtr, err := s.repo.GetProductBySlug(slug)
	if err != nil {
		return models.Product{}, err
	}
	return *productPtr, nil
}
