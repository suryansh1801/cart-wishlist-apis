package service

import (
	"ecommerce-service/internal/models"
	"ecommerce-service/internal/repository"
)

type WishlistService struct {
	repo *repository.WishlistRepository
}

func NewWishlistService(repo *repository.WishlistRepository) *WishlistService {
	return &WishlistService{repo: repo}
}

func (s *WishlistService) GetUserWishlist(userID uint) ([]models.WishlistItem, error) {
	return s.repo.GetByUserID(userID)
}

func (s *WishlistService) AddToWishlist(userID, productID uint) error {
	item := &models.WishlistItem{
		UserID:    userID,
		ProductID: productID,
	}
	return s.repo.AddItem(item)
}

func (s *WishlistService) RemoveFromWishlist(userID, productID uint) error {
	return s.repo.DeleteItem(userID, productID)
}
