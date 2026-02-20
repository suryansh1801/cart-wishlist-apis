package service

import (
	"ecommerce-service/internal/models"
	"ecommerce-service/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) GetUserCart(userID uint) ([]models.CartItem, error) {
	return s.repo.GetCartByUserID(userID)
}

func (s *CartService) AddToCart(userID, productID uint, quantity int) error {
	item := &models.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return s.repo.AddOrUpdateItem(item)
}

func (s *CartService) RemoveFromCart(userID, productID uint) error {
	return s.repo.DeleteItem(userID, productID)
}
