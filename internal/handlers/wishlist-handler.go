package handlers

import (
	service "ecommerce-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	service *service.WishlistService
}

func NewWishlistHandler(service *service.WishlistService) *WishlistHandler {
	return &WishlistHandler{service: service}
}

func (h *WishlistHandler) GetWishlist(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	items, err := h.service.GetUserWishlist(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *WishlistHandler) AddItem(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddToWishlist(uint(userID), req.ProductID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Item added to wishlist"})
}

func (h *WishlistHandler) RemoveItem(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	productID, _ := strconv.ParseUint(c.Param("productID"), 10, 32)

	if err := h.service.RemoveFromWishlist(uint(userID), uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from wishlist"})
}

func (h *WishlistHandler) MoveToCart(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	productID, _ := strconv.ParseUint(c.Param("productID"), 10, 32)

	if err := h.service.MoveItemToCart(uint(userID), uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move item to cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item successfully moved to cart"})
}
