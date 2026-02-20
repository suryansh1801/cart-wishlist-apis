package routes

import (
	"ecommerce-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, cartHandler *handlers.CartHandler, wishlistHandler *handlers.WishlistHandler) {
	v1 := router.Group("/api/v1")
	{
		// Cart Routes -> /api/v1/carts/1/items
		carts := v1.Group("/carts/:userID")
		{
			carts.GET("/", cartHandler.GetCart)
			carts.POST("/items", cartHandler.AddItem)
		}

		// Wishlist Routes -> /api/v1/wishlists/1/items
		wishlists := v1.Group("/wishlists/:userID")
		{
			wishlists.GET("/", wishlistHandler.GetWishlist)
			wishlists.POST("/items", wishlistHandler.AddItem)
		}
	}
}
