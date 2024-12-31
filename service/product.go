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
