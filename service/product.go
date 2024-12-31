package service

import (
	"superapp/models"
	"superapp/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() ([]models.ProductPaginate, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) GetProductBySlug(slug string) (models.Product, error) {
	productPtr, err := s.repo.GetProductBySlug(slug)
	if err != nil {
		return models.Product{}, err
	}
	return *productPtr, nil
}
