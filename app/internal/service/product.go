package service

import (
	"github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts(userId, limit, offset int) (*models.ProductsPaginate, error) {
	return s.repo.GetAllProducts(userId, limit, offset)
}

func (s *ProductService) GetProductBySlug(slug string) (models.Product, error) {
	productPtr, err := s.repo.GetProductBySlug(slug)
	if err != nil {
		return models.Product{}, err
	}
	return *productPtr, nil
}

func (s *ProductService) ToggleLike(userId, id int) error {
	if err := s.repo.ToggleLike(userId, id); err != nil {
		return err
	}
	return nil
}
