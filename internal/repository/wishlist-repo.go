package repository

import (
	"ecommerce-service/internal/models"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{db: db}
}

func (r *WishlistRepository) GetByUserID(userID uint) ([]models.WishlistItem, error) {
	var items []models.WishlistItem
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *WishlistRepository) AddItem(item *models.WishlistItem) error {
	// Ignore if already exists
	var existing models.WishlistItem
	if r.db.Where("user_id = ? AND product_id = ?", item.UserID, item.ProductID).First(&existing).Error == nil {
		return nil
	}
	return r.db.Create(item).Error
}

func (r *WishlistRepository) DeleteItem(userID, productID uint) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.WishlistItem{}).Error
}
