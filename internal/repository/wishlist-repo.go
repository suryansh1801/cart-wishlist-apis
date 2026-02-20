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

func (r *WishlistRepository) MoveToCart(userID, productID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingCartItem models.CartItem
		result := tx.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingCartItem)

		if result.Error == nil {
			existingCartItem.Quantity += 1
			if err := tx.Save(&existingCartItem).Error; err != nil {
				return err // Returns error will rollback transaction
			}
		} else {
			newCartItem := &models.CartItem{
				UserID:    userID,
				ProductID: productID,
				Quantity:  1,
			}

			if err := tx.Create(newCartItem).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.WishlistItem{}).Error; err != nil {
			return err
		}

		// Return nil commits the transaction
		return nil
	})
}
