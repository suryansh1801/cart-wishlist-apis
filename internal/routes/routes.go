package routes

import (
	"ecommerce-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, cartHandler *handlers.CartHandler, wishlistHandler *handlers.WishlistHandler) {
	v1 := router.Group("/api/v1")
	{
		carts := v1.Group("/carts/:userID")
		{
			carts.GET("/", cartHandler.GetCart)
			carts.POST("/items", cartHandler.AddItem)
			carts.DELETE("/items/:productID", cartHandler.RemoveItem)
		}

		wishlists := v1.Group("/wishlists/:userID")
		{
			wishlists.GET("/", wishlistHandler.GetWishlist)
			wishlists.POST("/items", wishlistHandler.AddItem)
			wishlists.DELETE("/items/:productID", wishlistHandler.RemoveItem)
			wishlists.POST("/items/:productID/move-to-cart", wishlistHandler.MoveToCart)
		}
	}
}
