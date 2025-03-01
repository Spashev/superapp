package service

import (
	"github.com/spashev/superapp/internal/models"
	"github.com/spashev/superapp/internal/repository"
)

type UserProductService struct {
	repo *repository.UserProductRepository
}

func NewUserProductService(repo *repository.UserProductRepository) *UserProductService {
	return &UserProductService{repo: repo}
}

func (s *UserProductService) GetUserFavoriteProducts(userId, limit, offset int) (*models.ProductsPaginate, error) {
	return s.repo.GetUserFavoriteProducts(userId, limit, offset)
}
