package handlers

import (
	service "ecommerce-service/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(c *gin.Context) {
	// Extract userID from the URL path parameter
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)

	items, err := h.service.GetUserCart(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *CartHandler) AddItem(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddToCart(uint(userID), req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Item added to cart"})
}

func (h *CartHandler) RemoveItem(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("userID"), 10, 32)
	productID, _ := strconv.ParseUint(c.Param("productID"), 10, 32)

	if err := h.service.RemoveFromCart(uint(userID), uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
