package repository

import (
	"ecommerce-service/internal/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetCartByUserID(userID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := r.db.Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func (r *CartRepository) AddOrUpdateItem(item *models.CartItem) error {
	var existing models.CartItem
	result := r.db.Where("user_id = ? AND product_id = ?", item.UserID, item.ProductID).First(&existing)

	if result.Error == nil {
		existing.Quantity += item.Quantity
		return r.db.Save(&existing).Error
	}
	return r.db.Create(item).Error
}

func (r *CartRepository) DeleteItem(userID, productID uint) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.CartItem{}).Error
}
